package useCase

import (
	"context"
	"fmt"
	"notifications/internal/domain"
	"notifications/internal/domain/mail"
	"notifications/internal/useCase/adapters/storage"
)

type useCase struct {
	storage storage.Storage
	options Options
}

type Options struct{}

func New(storage storage.Storage, options Options) *useCase {
	var uc = &useCase{
		storage: storage,
	}
	uc.SetOptions(options)
	return uc
}

func (uc *useCase) SetOptions(options Options) {
	if uc.options != options {
		uc.options = options
		//log.Info("set new options", zap.Any("options", uc.options))
	}
}

func (u *useCase) StoreMail(ctx context.Context, mail *mail.Mail) (err error) {
	return u.storage.CreateMail(ctx, mail)
}

func (u *useCase) ProcessEmails(ctx context.Context) (err error) {
	mails, err := u.storage.ReadMailFiltredList(ctx, map[string]interface{}{"sended": domain.NotSended})

	for _, mail := range mails {
		fmt.Println(mail)
	}
	return
}
