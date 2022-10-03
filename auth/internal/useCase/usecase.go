package useCase

import (
	"auth/internal/domain/device"
	"auth/internal/domain/device/agent"
	"auth/internal/domain/device/ip"
	"auth/internal/domain/user"
	"auth/internal/domain/user/email"
	"auth/internal/domain/user/pass"
	"auth/internal/domain/user/phone"
	"auth/internal/useCase/adapters/storage"
	"context"

	"github.com/google/uuid"
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

func (uc *useCase) SignUp(ctx context.Context, user *user.User) (err error) {
	err = uc.storage.CreateUser(ctx, user)
	if err != nil {
		return
	}
	//send email confirmation to kafka
	return
}

func (uc *useCase) SignIn(ctx context.Context, phone phone.Phone, pass pass.Pass, device *device.Device) (err error) {
	return
}

func (uc *useCase) VerifyEmail(ctx context.Context, verificationCode string) (err error) {
	return
}

func (uc *useCase) ResetPassword(ctx context.Context, email email.Email) (err error) {
	return
}

func (uc *useCase) RefreshToken(ctx context.Context, deviceID uuid.UUID, ip ip.Ip, agent agent.Agent) (err error) {
	return
}
