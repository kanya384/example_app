package useCase

import (
	"context"
	"notifications/internal/domain/mail"
)

type Mails interface {
	ProcessEmails(ctx context.Context) (err error)
	StoreMail(ctx context.Context, mail *mail.Mail) (err error)
}
