package useCase

import (
	"auth/internal/domain/device"
	"auth/internal/domain/device/agent"
	"auth/internal/domain/device/ip"
	"auth/internal/domain/user"
	"auth/internal/domain/user/email"
	"auth/internal/domain/user/pass"
	"auth/internal/domain/user/phone"
	"context"

	"github.com/google/uuid"
)

type UseCase interface {
	SignUp(ctx context.Context, user *user.User) (err error)
	SignIn(ctx context.Context, phone phone.Phone, pass pass.Pass, device *device.Device) (err error)
	VerifyEmail(ctx context.Context, verificationCode string) (err error)
	ResetPassword(ctx context.Context, email email.Email) (err error)
	RefreshToken(ctx context.Context, deviceID uuid.UUID, ip ip.Ip, agent agent.Agent) (err error)
}
