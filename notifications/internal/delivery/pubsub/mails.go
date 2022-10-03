package pubsub

import (
	"context"
	"fmt"

	notification "notifications/internal/delivery/pubsub/interface"
	"notifications/internal/domain/mail"
	"notifications/internal/domain/mail/email"
	"notifications/internal/domain/mail/message"
	"notifications/internal/domain/mail/subject"

	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
)

func (d *Delivery) SubscribeToMessages(ctx context.Context) (messages chan *kafka.Message, err error) {
	messages = make(chan *kafka.Message)
	go func(messages chan<- *kafka.Message) {
		for {
			msg, err := d.kServer.ReadMessage(ctx)
			if err != nil {
				d.logger.Error(fmt.Errorf("error reading message: %w", err))
			}
			messages <- msg
		}
	}(messages)
	return
}
func (d *Delivery) ProcessMessage(ctx context.Context, messages <-chan *kafka.Message) {
	for {
		select {
		case msg := <-messages:
			err := d.doProcessMessage(ctx, msg)
			if err != nil {
				d.logger.Error(err)
			}
		case <-ctx.Done():
			fmt.Println("closing reading")
			return
		}
	}
}

func (d *Delivery) doProcessMessage(ctx context.Context, msg *kafka.Message) (err error) {
	mailPB := &notification.Mail{}
	err = proto.Unmarshal(msg.Value, mailPB)

	if err != nil {
		return fmt.Errorf("error unmarshaling message: %w", err)
	}

	recipient, err := email.NewEmail(mailPB.Recipient)
	if err != nil {
		return err
	}
	subject, err := subject.NewSubject(mailPB.Subject)
	if err != nil {
		return err
	}
	message, err := message.NewMessage(mailPB.Message)
	if err != nil {
		return err
	}
	mail, err := mail.New(*recipient, *subject, *message)
	if err != nil {
		return err
	}

	err = d.mailsUseCase.StoreMail(ctx, mail)
	if err != nil {
		return err
	}
	return
}
