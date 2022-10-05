package helpers

import (
	"auth/pkg/types/notification"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

func NewGrpcError(grpcErrorCode codes.Code, err error) error {
	errorResponse := status.Newf(
		grpcErrorCode,
		err.Error(),
	)
	return errorResponse.Err()
}

func PostgresConnectionString(user, pass, host, port, dbName string) string {
	return fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable",
		"postgres",
		user,
		pass,
		host,
		port,
		dbName)
}

func MigrationsUP(dsn, path string) (err error) {
	m, err := migrate.New(path, dsn)
	if err != nil {
		return err
	}
	m.Up()
	return nil
}

func GenerateEmail(recipient, subject, msg string) (message []byte, err error) {
	mail := &notification.Mail{
		Recipient: "kanya384@mail.ru",
		Subject:   "test subject",
		Message:   "test message",
	}
	message, err = proto.Marshal(mail)
	if err != nil {
		return
	}
	return
}
