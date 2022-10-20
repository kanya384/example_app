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
	"auth/pkg/helpers"
	"context"
	"errors"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
)

func (d *Delivery) SignUp(ctx context.Context, request *auth.SignUpRequest) (response *auth.SignUpResponse, err error) {
	phone, err := phone.NewPhone(request.Phone)
	if err != nil {
		helpers.NewGrpcError(codes.InvalidArgument, err)
		return
	}
	name, err := name.NewName(request.Name)
	if err != nil {
		helpers.NewGrpcError(codes.InvalidArgument, err)
		return
	}
	surname, err := surname.NewSurname(request.Surname)
	if err != nil {
		helpers.NewGrpcError(codes.InvalidArgument, err)
		return
	}
	pass, err := pass.NewPass(request.Pass, d.salt)
	if err != nil {
		helpers.NewGrpcError(codes.InvalidArgument, err)
		return
	}
	email, err := email.NewEmail(request.Email)
	if err != nil {
		helpers.NewGrpcError(codes.InvalidArgument, err)
		return
	}
	role := user.UserRole(request.Role)

	user, err := user.New(*name, *surname, *phone, *pass, *email, role)
	if err != nil {
		helpers.NewGrpcError(codes.InvalidArgument, err)
		return
	}

	err = d.ucAuth.SignUp(ctx, user)
	if err != nil {
		helpers.NewGrpcError(codes.InvalidArgument, err)
		return
	}
	return &auth.SignUpResponse{Message: "registration success, please confirm email"}, nil
}

func (d *Delivery) SignIn(ctx context.Context, request *auth.SignInRequest) (response *auth.SignInResponse, err error) {
	phone, err := phone.NewPhone(request.Phone)
	if err != nil {
		helpers.NewGrpcError(codes.InvalidArgument, err)
		return
	}
	pass, err := pass.NewPass(request.Pass, d.salt)
	if err != nil {
		helpers.NewGrpcError(codes.InvalidArgument, err)
		return
	}

	agent, err := agent.NewAgent(request.Device.Agent)
	if err != nil {
		helpers.NewGrpcError(codes.InvalidArgument, err)
		return
	}
	deviceID, err := deviceID.NewDeviceID(request.Device.DeviceID)
	if err != nil {
		helpers.NewGrpcError(codes.InvalidArgument, err)
		return
	}
	ip, err := ip.NewIp(request.Device.Ip)
	if err != nil {
		helpers.NewGrpcError(codes.InvalidArgument, err)
		return
	}
	userID, err := uuid.Parse(request.Device.UserID)
	if err != nil {
		helpers.NewGrpcError(codes.InvalidArgument, err)
		return
	}
	dtype := device.DeviceType(request.Device.Dtype)

	refreshToken, err := refreshToken.New()
	if err != nil {
		helpers.NewGrpcError(codes.InvalidArgument, err)
		return
	}

	device, err := device.New(userID, *deviceID, *ip, *agent, dtype, *refreshToken)
	if err != nil {
		helpers.NewGrpcError(codes.InvalidArgument, err)
		return
	}

	result, err := d.ucAuth.SignIn(ctx, *phone, *pass, device)
	if err != nil {
		helpers.NewGrpcError(codes.NotFound, err)
		return
	}
	return &auth.SignInResponse{Token: result.Token, RefreshToken: result.Refresh}, nil
}

func (d *Delivery) VerifyEmail(ctx context.Context, request *auth.VerifyEmailRequest) (response *auth.VerifyEmailResponse, err error) {
	requestID, err := uuid.Parse(request.VerificationCode)
	if err != nil {
		helpers.NewGrpcError(codes.InvalidArgument, err)
		return
	}
	err = d.ucAuth.VerifyEmail(ctx, requestID)
	if err != nil {
		helpers.NewGrpcError(codes.InvalidArgument, err)
		return
	}
	return &auth.VerifyEmailResponse{}, nil
}

func (d *Delivery) ResetPassword(ctx context.Context, request *auth.ResetPasswordRequest) (response *auth.ResetPasswordResponse, err error) {
	return nil, errors.New("not implemented")
}

func (d *Delivery) RefreshToken(ctx context.Context, request *auth.RefreshTokenRequest) (response *auth.RefreshTokenResponse, err error) {
	return nil, errors.New("not implemented")
}
