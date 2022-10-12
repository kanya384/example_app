package storage

import (
	"context"
	"notifications/internal/domain/mail"

	"github.com/google/uuid"
)

type Storage interface {
	Mail
	Push
}

type Mail interface {
	CreateMail(ctx context.Context, mail *mail.Mail) (err error)
	UpdateMail(ctx context.Context, ID uuid.UUID, updateFn func(mail *mail.Mail) (*mail.Mail, error)) (mail *mail.Mail, err error)
	DeleteMail(ctx context.Context, ID uuid.UUID) (err error)
	ReadMailByID(ctx context.Context, ID uuid.UUID) (mail *mail.Mail, err error)
	ReadMailFiltredList(ctx context.Context, filter map[string]interface{}) (mailsList []*mail.Mail, err error)
}

type Push interface {
}
