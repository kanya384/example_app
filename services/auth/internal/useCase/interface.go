package useCase

import (
	"auth/internal/domain/device"
	"auth/internal/domain/device/agent"
	"auth/internal/domain/device/ip"
	"auth/internal/domain/device/refreshToken"
	"auth/internal/domain/user"
	"auth/internal/domain/user/email"
	"auth/internal/domain/user/pass"
	"auth/internal/domain/user/phone"
	"context"

	"github.com/google/uuid"
)

type UseCase interface {
	SignUp(ctx context.Context, user *user.User) (err error)
	SignIn(ctx context.Context, phone phone.Phone, pass pass.Pass, inputDevice *device.Device) (res SignInResult, err error)
	ResendVerificationCode(ctx context.Context, userID uuid.UUID) (err error)
	VerifyEmail(ctx context.Context, verificationCode uuid.UUID) (err error)
	RefreshToken(ctx context.Context, userID uuid.UUID, refreshToken refreshToken.RefreshToken, ip ip.Ip, agent agent.Agent) (result RefreshResult, err error)
	ResetPasswordRequest(ctx context.Context, email email.Email) (err error)
	ResetPasswordProcess(ctx context.Context, resetRequestID uuid.UUID, newPass pass.Pass) (err error)
}
