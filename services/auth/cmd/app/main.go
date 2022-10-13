package main

import (
	"auth/internal/config"
	delivery "auth/internal/delivery/grpc"
	authGrpc "auth/internal/delivery/grpc/interface"
	repository "auth/internal/repository/postgres"
	"auth/internal/useCase"
	"auth/pkg/helpers"
	"auth/pkg/kafka/client"
	lg "auth/pkg/logger"
	"auth/pkg/memcache"
	"auth/pkg/psql"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
)

const (
	serviceName = "auth"
)

func main() {
	cfg, err := config.InitConfig("auth")
	if err != nil {
		panic(fmt.Sprintf("error initializing config %s", err))
	}

	logger := lg.New(cfg.Log.Level, serviceName, cfg.Graylog.Host)

	kafkaClent := client.NewKafkaClient(cfg.Kafka.Brokers, cfg.Kafka.Topic, cfg.Kafka.GroupID)

	dsn := helpers.PostgresConnectionString(cfg.PG.User, cfg.PG.Pass, cfg.PG.Host, cfg.PG.Port, cfg.PG.DbName)
	err = helpers.MigrationsUP(dsn, "file://migrations")
	if err != nil {
		logger.Fatal(fmt.Errorf("migrations error: %w", err))
	}

	pg, err := psql.New(dsn, psql.MaxPoolSize(cfg.PG.PoolMax), psql.ConnTimeout(time.Duration(cfg.PG.Timeout)*time.Second))
	if err != nil {
		logger.Fatal(fmt.Errorf("postgres connection error: %w", err))
	}

	storage, err := repository.NewRepository(pg)
	if err != nil {
		logger.Fatal("storage initialization error: %w", err)
	}

	memcache := memcache.New(cfg.Cache.TimeToLive, cfg.Cache.CleanupInterval)

	uc := useCase.New(storage.Users, storage.Devices, memcache, kafkaClent, logger, useCase.Options{})

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		logger.Fatal(fmt.Errorf("listen to port error: %w", err))
	}

	server := grpc.NewServer()

	handlers := delivery.New(uc, cfg.Token.Salt, delivery.Options{})

	authGrpc.RegisterAuthServer(server, handlers)

	go func(server *grpc.Server, logger *lg.Logger) {
		if err := server.Serve(listen); err != nil {
			logger.Fatal(fmt.Errorf("error staring grpc server: %w", err))
		}
	}(server, logger)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	<-c

	if err := shutdown(server, pg, memcache, logger); err != nil {
		logger.Fatal(fmt.Errorf("failed shutdown with error: %w", err))
	}
}

func shutdown(server *grpc.Server, psql *psql.Postgres, memcache *memcache.Cache, logger *lg.Logger) error {
	logger.Info("Gracefully stopping...")
	server.GracefulStop()
	psql.Pool.Close()
	memcache.Stop()
	logger.Info("Gracefully shutdown done!")
	return nil
}
