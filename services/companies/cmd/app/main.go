package main

import (
	"companies/internal/config"
	deliveryGrpc "companies/internal/delivery/grpc"
	companyGrpc "companies/internal/delivery/grpc/interface"
	repositoryStorage "companies/internal/repository/company/postgres"
	useCaseCompany "companies/internal/useCase/company"
	"companies/pkg/helpers"
	lg "companies/pkg/logger"
	"companies/pkg/psql"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.InitConfig("COMPANIES")
	if err != nil {
		panic(fmt.Sprintf("error initializing config %s", err))
	}

	logger := lg.New(cfg.Log.Level)

	dsn := helpers.PostgresConnectionString(cfg.PG.User, cfg.PG.Pass, cfg.PG.Host, cfg.PG.Port, cfg.PG.DbName)

	err = helpers.MigrationsUP(dsn, "file://migrations")
	if err != nil {
		logger.Fatal(fmt.Errorf("migrations error: %w", err))
	}

	pg, err := psql.New(dsn, psql.MaxPoolSize(cfg.PG.PoolMax), psql.ConnTimeout(time.Duration(cfg.PG.Timeout)*time.Second))
	if err != nil {
		logger.Fatal(fmt.Errorf("postgres connection error: %w", err))
	}

	stCompany, err := repositoryStorage.New(pg, repositoryStorage.Options{})
	if err != nil {
		logger.Fatal(fmt.Errorf("storage initialization error: %w", err))
	}

	ucCompany := useCaseCompany.New(stCompany, useCaseCompany.Options{})

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		logger.Fatal(fmt.Errorf("listen to port error: %w", err))
	}

	server := grpc.NewServer()

	handlers := deliveryGrpc.New(ucCompany, deliveryGrpc.Options{})

	companyGrpc.RegisterCompanyServer(server, handlers)

	go func(server *grpc.Server, logger *lg.Logger) {
		if err := server.Serve(listen); err != nil {
			logger.Fatal(fmt.Errorf("error staring grpc server: %w", err))
		}
	}(server, logger)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	<-c

	if err := shutdown(server, logger, pg); err != nil {
		logger.Fatal(fmt.Errorf("failed shutdown with error: %w", err))
	}
}

func shutdown(server *grpc.Server, logger *lg.Logger, psql *psql.Postgres) error {
	logger.Info("Gracefully stopping...")
	server.GracefulStop()
	psql.Pool.Close()
	logger.Info("Gracefully shutdown done!")
	return nil
}
