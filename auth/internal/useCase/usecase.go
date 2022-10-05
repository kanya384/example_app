package useCase

import (
	"auth/internal/domain/device"
	"auth/internal/domain/device/agent"
	"auth/internal/domain/device/ip"
	"auth/internal/domain/user"
	"auth/internal/domain/user/email"
	"auth/internal/domain/user/pass"
	"auth/internal/domain/user/phone"
	"auth/internal/useCase/adapters/pubsub"
	"auth/internal/useCase/adapters/storage"
	"auth/pkg/helpers"
	"auth/pkg/logger"
	"context"
	"fmt"

	"github.com/google/uuid"
)

const (
	CONFIRMATION_MESSAGE_SUBJECT  = "Подтверждение электронной почты"
	CONFIRMATION_MESSAGE_TEMPLATE = "Пожалуйста перейдите по ссылке, чтобы подтвердить вашу почту:\r\n %s"
)

type useCase struct {
	usersStorage  storage.User
	deviceStorage storage.Device
	notification  pubsub.Notification
	logger        *logger.Logger
	options       Options
}

type Options struct{}

func New(usersStorage storage.User, deviceStorage storage.Device, notification pubsub.Notification, logger *logger.Logger, options Options) *useCase {
	var uc = &useCase{
		usersStorage:  usersStorage,
		deviceStorage: deviceStorage,
		notification:  notification,
		logger:        logger,
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
	err = uc.usersStorage.CreateUser(ctx, user)
	if err != nil {
		return
	}

	msg, err := helpers.GenerateEmail(user.Email().String(), CONFIRMATION_MESSAGE_SUBJECT, fmt.Sprintf(CONFIRMATION_MESSAGE_TEMPLATE, "http://url"))
	if err != nil {
		return
	}

	errM := uc.notification.SendMessage(ctx, uuid.New().String(), msg)
	if errM != nil {
		uc.logger.Error("error sending confirrmation message to kafka: ", errM.Error())
	}
	return
}

func (uc *useCase) SignIn(ctx context.Context, phone phone.Phone, pass pass.Pass, inputDevice *device.Device) (err error) {
	user, err := uc.usersStorage.ReadUserByCredetinals(ctx, phone, pass)
	if err != nil {
		return
	}

	inputDevice.SetUserID(user.ID())

	//TODO - optimize this
	_, err = uc.deviceStorage.ReadDevicesByDeviceID(ctx, inputDevice.UserID())
	switch err {
	case nil:
		_, err = uc.deviceStorage.UpdateDevice(ctx, inputDevice.ID(), func(oldDevice *device.Device) (*device.Device, error) {
			return device.NewWithID(oldDevice.ID(), oldDevice.CreatedAt(), inputDevice.ModifiedAt(), oldDevice.UserID(), inputDevice.DeviceID(), inputDevice.Ip(), inputDevice.Agent(), inputDevice.Type(), inputDevice.RefreshToken(), inputDevice.RefreshExpiration(), inputDevice.LastSeen())
		})
	default:
		err = uc.deviceStorage.CreateDevice(ctx, inputDevice)
	}
	if err != nil {
		return
	}

	//TODO - create refresh and access token

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
