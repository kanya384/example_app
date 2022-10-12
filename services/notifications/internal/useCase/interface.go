package useCase

import (
	"context"
	"notifications/internal/domain/mail"
	"time"
)

type Mails interface {
	ProcessEmails(ctx context.Context, done <-chan struct{}, interval time.Duration)
	StoreMail(ctx context.Context, mail *mail.Mail) (err error)
}
