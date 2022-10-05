package useCase

import (
	"context"
	"notifications/internal/domain"
	"notifications/internal/domain/mail"
	"notifications/internal/useCase/adapters/emailClient"
	"notifications/internal/useCase/adapters/storage"
	"notifications/pkg/logger"
	"time"
)

type UseCase struct {
	storage    storage.Storage
	emailClent emailClient.EmailClient
	logger     *logger.Logger
	options    Options
}

type Options struct{}

func New(storage storage.Storage, emailClent emailClient.EmailClient, logger *logger.Logger, options Options) *UseCase {
	var uc = &UseCase{
		storage:    storage,
		emailClent: emailClent,
		logger:     logger,
	}
	uc.SetOptions(options)
	return uc
}

func (uc *UseCase) SetOptions(options Options) {
	if uc.options != options {
		uc.options = options
		//log.Info("set new options", zap.Any("options", uc.options))
	}
}

func (u *UseCase) StoreMail(ctx context.Context, mail *mail.Mail) (err error) {
	return u.storage.CreateMail(ctx, mail)
}

func (u *UseCase) ProcessEmails(ctx context.Context, done <-chan struct{}, interval time.Duration) {
L:
	for {
		select {
		case <-done:
			break L
		default:
			err := u.doProcessEmails(ctx)
			if err != nil {
				u.logger.Error(err)
			}
			time.Sleep(interval)
		}

	}
}

func (u *UseCase) doProcessEmails(ctx context.Context) (err error) {
	mails, err := u.storage.ReadMailFiltredList(ctx, map[string]interface{}{"status": domain.NotSended})
	if err != nil {
		return
	}
	for _, email := range mails {
		err = u.emailClent.SendEmail(ctx, email.Recipient().String(), email.Subject().String(), email.Message().String())
		if err != nil {
			continue
		}
		u.storage.UpdateMail(ctx, email.ID(), func(oldMessage *mail.Mail) (*mail.Mail, error) {
			return mail.NewWithID(oldMessage.ID(), oldMessage.CreatedAt(), time.Now(), oldMessage.Recipient(), oldMessage.Subject(), oldMessage.Message(), domain.Sended)
		})
	}
	return
}
