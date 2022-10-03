package main

import (
	"context"
	"fmt"
	"notifications/internal/config"
	"notifications/internal/delivery/pubsub"
	repositoryStorage "notifications/internal/repository/mail/postgres"
	"notifications/internal/useCase"
	"notifications/pkg/emailClient"
	"notifications/pkg/helpers"
	"notifications/pkg/kafka/server"
	"notifications/pkg/psql"
	"os"
	"os/signal"
	"syscall"
	"time"

	lg "notifications/pkg/logger"
)

func main() {
	cfg, err := config.InitConfig("notifications")
	if err != nil {
		panic(fmt.Sprintf("error initializing config %s", err))
	}

	logger := lg.New(cfg.Log.Level)

	kServer := server.NewKafkaServer(cfg.Kafka.Brokers, cfg.Kafka.Topic, cfg.Kafka.Group)
	defer kServer.Close()

	dsn := helpers.PostgresConnectionString(cfg.PG.User, cfg.PG.Pass, cfg.PG.Host, cfg.PG.Port, cfg.PG.DbName)
	err = helpers.MigrationsUP(dsn, "file://migrations")
	if err != nil {
		logger.Fatal(fmt.Errorf("migrations error: %w", err))
	}

	pg, err := psql.New(dsn, psql.MaxPoolSize(cfg.PG.PoolMax), psql.ConnTimeout(time.Duration(cfg.PG.Timeout)*time.Second))
	if err != nil {
		logger.Fatal(fmt.Errorf("postgres connection error: %w", err))
	}

	stMail, err := repositoryStorage.New(pg, repositoryStorage.Options{})
	if err != nil {
		logger.Fatal(fmt.Errorf("storage initialization error: %w", err))
	}

	contextWithCancel, cancel := context.WithCancel(context.Background())
	defer cancel()

	emailClient := emailClient.NewEmailClient(cfg.Email.Host, cfg.Email.Port, cfg.Email.Login, cfg.Email.Pass)
	uc := useCase.New(stMail, emailClient, useCase.Options{})

	delivery := pubsub.New(uc, kServer, logger, pubsub.Options{})
	messagesChan, err := delivery.SubscribeToMessages(contextWithCancel)
	if err != nil {
		logger.Fatal(fmt.Errorf("kafka subscribe error: %w", err))
	}
	go delivery.ProcessMessage(contextWithCancel, messagesChan)

	go func(context context.Context, useCase *useCase.UseCase) {
		for {
			err = useCase.ProcessEmails(context)
			if err != nil {
				logger.Error(err)
			}
		}
	}(contextWithCancel, uc)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	<-c

	if err := shutdown(kServer); err != nil {
		panic(fmt.Errorf("failed shutdown with error: %w", err))
	}
}

func shutdown(kServer *server.KafkaServer) (err error) {
	fmt.Println("Gracefull shut down ....")
	err = kServer.Close()
	if err != nil {
		return
	}
	return nil
}
