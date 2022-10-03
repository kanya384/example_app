package main

import (
	"auth/internal/config"
	"auth/pkg/kafka/client"
	"auth/pkg/types/notification"
	"context"
	"fmt"

	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
)

func main() {

	cfg, err := config.InitConfig("auth")
	if err != nil {
		panic(fmt.Sprintf("error initializing config %s", err))
	}

	kafkaClent := client.NewKafkaClient(cfg.Kafka.Brokers, cfg.Kafka.Topic, cfg.Kafka.GroupID)
	mail := &notification.Mail{
		Recipient: "test01@mail.ru",
		Subject:   "test subject",
		Message:   "test message",
	}
	protoMail, err := proto.Marshal(mail)
	if err != nil {
		panic(fmt.Sprintf("error initializing config %s", err))
	}
	fmt.Println(string(protoMail))
	err = kafkaClent.SendMessage(context.Background(), uuid.NewString(), protoMail)
	if err != nil {
		panic(fmt.Sprintf("error sending message to kafka %s", err))
	}

}
