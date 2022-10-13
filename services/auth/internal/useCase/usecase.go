package useCase

import (
	"auth/internal/domain/device"
	"auth/internal/domain/device/agent"
	"auth/internal/domain/device/ip"
	"auth/internal/domain/user"
	"auth/internal/domain/user/email"
	"auth/internal/domain/user/pass"
	"auth/internal/domain/user/phone"
	"auth/internal/useCase/adapters/cache"
	"auth/internal/useCase/adapters/pubsub"
	"auth/internal/useCase/adapters/storage"
	"auth/pkg/helpers"
	"auth/pkg/logger"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

const (
	confirmationSubject    = "Подтверждение электронной почты"
	confirmationTemplate   = "Пожалуйста перейдите по ссылке, чтобы подтвердить вашу почту:\r\n %s"
	emailVerification      = ""
	verificationExpiration = time.Duration(time.Minute * 60 * 24)
)

type useCase struct {
	usersStorage  storage.User
	deviceStorage storage.Device
	cache         cache.Cache
	notification  pubsub.Notification
	logger        *logger.Logger
	options       Options
}

type Options struct{}

func New(usersStorage storage.User, deviceStorage storage.Device, cache cache.Cache, notification pubsub.Notification, logger *logger.Logger, options Options) *useCase {
	var uc = &useCase{
		usersStorage:  usersStorage,
		deviceStorage: deviceStorage,
		cache:         cache,
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

	go uc.sendVerificationCodeToEmail(ctx, user.ID().String(), user.Email().String())

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

func (uc *useCase) VerifyEmail(ctx context.Context, verificationCode uuid.UUID) (err error) {
	id, err := uc.cache.Get(verificationCode.String())
	if err != nil {
		return
	}

	userU, err := uc.usersStorage.ReadUserByID(ctx, id.(uuid.UUID))
	if err != nil {
		return
	}

	userU.SetVerified(true)

	_, err = uc.usersStorage.UpdateUser(ctx, userU.ID(), func(oldUser *user.User) (*user.User, error) {
		return user.NewWithID(oldUser.ID(), oldUser.CreatedAt(), time.Now(), oldUser.Name(), oldUser.Surname(), oldUser.Phone(), oldUser.Pass(), oldUser.Email(), true, oldUser.Role())
	})
	return
}

func (uc *useCase) ResendVerificationCode(ctx context.Context, userID uuid.UUID) (err error) {
	user, err := uc.usersStorage.ReadUserByID(ctx, userID)
	if err != nil {
		return
	}
	return uc.sendVerificationCodeToEmail(ctx, user.ID().String(), user.Email().String())
}

func (uc *useCase) ResetPassword(ctx context.Context, email email.Email) (err error) {
	return
}

func (uc *useCase) RefreshToken(ctx context.Context, deviceID uuid.UUID, ip ip.Ip, agent agent.Agent) (err error) {
	return
}

func (uc *useCase) sendVerificationCodeToEmail(ctx context.Context, userID string, email string) (err error) {
	verificationCode := uuid.NewString()
	uc.cache.Set(verificationCode, userID, verificationExpiration)
	msg, err := helpers.GenerateEmail(email, confirmationSubject, fmt.Sprintf(confirmationTemplate, fmt.Sprintf(`http://url/%s`, verificationCode)))
	if err != nil {
		return
	}

	err = uc.notification.SendMessage(ctx, uuid.New().String(), msg)
	if err != nil {
		uc.logger.Error("error sending confirmation message to kafka: ", err.Error())
	}
	return
}
