package grpc

import (
	auth "auth/internal/delivery/grpc/interface"
	"auth/internal/domain/device"
	"auth/internal/domain/device/agent"
	"auth/internal/domain/device/deviceID"
	"auth/internal/domain/device/ip"
	"auth/internal/domain/device/refreshToken"
	"auth/internal/domain/user"
	"auth/internal/domain/user/email"
	"auth/internal/domain/user/name"
	"auth/internal/domain/user/pass"
	"auth/internal/domain/user/phone"
	"auth/internal/domain/user/surname"
	"context"
	"errors"

	"github.com/google/uuid"
)

func (d *Delivery) SignUp(ctx context.Context, request *auth.SignUpRequest) (response *auth.SignUpResponse, err error) {
	phone, err := phone.NewPhone(request.Phone)
	if err != nil {
		return
	}
	name, err := name.NewName(request.Name)
	if err != nil {
		return
	}
	surname, err := surname.NewSurname(request.Surname)
	if err != nil {
		return
	}
	pass, err := pass.NewPass(request.Pass, d.salt)
	if err != nil {
		return
	}
	email, err := email.NewEmail(request.Email)
	if err != nil {
		return
	}
	role := user.UserRole(request.Role)

	user, err := user.New(*name, *surname, *phone, *pass, *email, role)
	if err != nil {
		return
	}

	err = d.ucAuth.SignUp(ctx, user)
	if err != nil {
		return
	}
	return &auth.SignUpResponse{Message: "registration success, please confirm email"}, nil
}

func (d *Delivery) SignIn(ctx context.Context, request *auth.SignInRequest) (response *auth.SignInResponse, err error) {
	phone, err := phone.NewPhone(request.Phone)
	if err != nil {
		return
	}
	pass, err := pass.NewPass(request.Pass, d.salt)
	if err != nil {
		return
	}

	agent, err := agent.NewAgent(request.Device.Agent)
	if err != nil {
		return
	}
	deviceID, err := deviceID.NewDeviceID(request.Device.DeviceID)
	if err != nil {
		return
	}
	ip, err := ip.NewIp(request.Device.Ip)
	if err != nil {
		return
	}
	userID, err := uuid.Parse(request.Device.UserID)
	if err != nil {
		return
	}
	dtype := device.DeviceType(request.Device.Dtype)

	refreshToken, err := refreshToken.New()
	if err != nil {
		return
	}

	device, err := device.New(userID, *deviceID, *ip, *agent, dtype, *refreshToken)
	if err != nil {
		return
	}

	d.ucAuth.SignIn(ctx, *phone, *pass, device)
	return &auth.SignInResponse{Message: "successfully registred"}, nil
}

func (d *Delivery) VerifyEmail(ctx context.Context, request *auth.VerifyEmailRequest) (response *auth.VerifyEmailResponse, err error) {
	return nil, errors.New("not implemented")
}

func (d *Delivery) ResetPassword(ctx context.Context, request *auth.ResetPasswordRequest) (response *auth.ResetPasswordResponse, err error) {
	return nil, errors.New("not implemented")
}

func (d *Delivery) RefreshToken(ctx context.Context, request *auth.RefreshTokenRequest) (response *auth.RefreshTokenResponse, err error) {
	return nil, errors.New("not implemented")
}
