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
	"auth/internal/useCase/adapters/cache"
	"auth/internal/useCase/adapters/pubsub"
	"auth/internal/useCase/adapters/storage"
	"auth/pkg/auth"
	"auth/pkg/helpers"
	"auth/pkg/logger"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

const (
	confirmationSubject     = "Подтверждение электронной почты"
	confirmationTemplate    = "Пожалуйста перейдите по ссылке, чтобы подтвердить вашу почту:\r\n %s"
	emailVerification       = ""
	emailVerificationTTL    = time.Duration(time.Minute * 60 * 24)
	resetPasswordSessionTTL = time.Duration(time.Minute * 20)
	resetSubject            = "Смена пароля"
	resetTemplate           = "Пожалуйста перейдите по ссылке, для сброса пароля:\r\n %s"
	ErrNoResetRequest       = "reset request not founded"
)

var (
	ErrInvalidRefresh = errors.New("refresh is invalid")
	ErrUnknownDevice  = errors.New("unknown device, please relogin")
	updateFunction    func(*device.Device) (*device.Device, error)
)

type useCase struct {
	usersStorage  storage.User
	deviceStorage storage.Device
	cache         cache.Cache
	notification  pubsub.Notification
	logger        logger.Interface
	tokenManager  auth.TokenManager
	options       Options
}

type Options struct {
	TokenTTL time.Duration
}

func New(usersStorage storage.User, deviceStorage storage.Device, cache cache.Cache, notification pubsub.Notification, tokenManager auth.TokenManager, logger logger.Interface, options Options) *useCase {
	var uc = &useCase{
		usersStorage:  usersStorage,
		deviceStorage: deviceStorage,
		cache:         cache,
		notification:  notification,
		tokenManager:  tokenManager,
		logger:        logger,
	}
	uc.SetOptions(options)
	return uc
}

func (uc *useCase) SetOptions(options Options) {
	if uc.options != options {
		uc.options = options
	}
}

func (uc *useCase) SignUp(ctx context.Context, user *user.User) (err error) {
	err = uc.usersStorage.CreateUser(ctx, user)
	if err != nil {
		return
	}

	uc.sendVerificationCodeToEmail(ctx, user.ID().String(), user.Email().String())

	return
}

func (uc *useCase) SignIn(ctx context.Context, phone phone.Phone, pass pass.Pass, inputDevice *device.Device) (res SignInResult, err error) {
	user, err := uc.usersStorage.ReadUserByCredetinals(ctx, phone, pass)
	if err != nil {
		return
	}

	inputDevice.SetUserID(user.ID())

	isDeviceExists := true
	if _, err := uc.deviceStorage.ReadDevicesByDeviceID(ctx, inputDevice.DeviceID()); err != nil {
		isDeviceExists = false
	}
	if isDeviceExists {
		_, err = uc.deviceStorage.UpdateDevice(ctx, inputDevice.ID(), func(oldDevice *device.Device) (*device.Device, error) {
			return device.NewWithID(oldDevice.ID(), oldDevice.CreatedAt(), inputDevice.ModifiedAt(), oldDevice.UserID(), inputDevice.DeviceID(), inputDevice.Ip(), inputDevice.Agent(), inputDevice.Type(), inputDevice.RefreshToken(), inputDevice.RefreshExpiration(), inputDevice.LastSeen())
		})
	} else {
		err = uc.deviceStorage.CreateDevice(ctx, inputDevice)
	}
	if err != nil {
		return
	}

	token, err := uc.tokenManager.NewJWT(auth.JwtClaims{UserID: user.ID().String(), UserName: user.Name().String() + " " + user.Surname().String()}, uc.options.TokenTTL)
	if err != nil {
		return
	}
	res.Token = token
	res.Refresh = inputDevice.RefreshToken().String()

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

func (uc *useCase) RefreshToken(ctx context.Context, userID uuid.UUID, refreshToken refreshToken.RefreshToken, ip ip.Ip, agent agent.Agent) (result RefreshResult, err error) {
	deviceUpd, err := uc.deviceStorage.ReadDeviceByUserIDAndRefresh(ctx, userID, refreshToken)
	if err != nil {
		err = ErrInvalidRefresh
		return
	}

	if deviceUpd.RefreshExpiration().UnixNano() < time.Now().UnixNano() {
		err = ErrInvalidRefresh
		return
	}

	if deviceUpd.Ip().String() != ip.String() || deviceUpd.Agent().String() != agent.String() {
		err = ErrUnknownDevice
		return
	}

	user, err := uc.usersStorage.ReadUserByID(ctx, userID)
	if err != nil {
		return
	}

	token, err := uc.tokenManager.NewJWT(auth.JwtClaims{UserID: user.ID().String(), UserName: user.Name().String() + " " + user.Surname().String()}, uc.options.TokenTTL)
	if err != nil {
		return
	}

	err = deviceUpd.UpdateRefreshToken()
	if err != nil {
		return
	}

	device, err := uc.deviceStorage.UpdateDevice(ctx, deviceUpd.ID(), func(oldDevice *device.Device) (*device.Device, error) {
		return device.NewWithID(oldDevice.ID(), oldDevice.CreatedAt(), time.Now(), oldDevice.UserID(), oldDevice.DeviceID(), oldDevice.Ip(), oldDevice.Agent(), oldDevice.Type(), deviceUpd.RefreshToken(), deviceUpd.RefreshExpiration(), time.Now())
	})
	if err != nil {
		return
	}

	result.Token = token
	result.Refresh = device.RefreshToken().String()

	return
}

func (uc *useCase) sendVerificationCodeToEmail(ctx context.Context, userID string, email string) (err error) {
	verificationID := uuid.NewString()
	uc.cache.Set(verificationID, userID, emailVerificationTTL)

	msg, err := helpers.GenerateEmail(email, confirmationSubject, fmt.Sprintf(confirmationTemplate, fmt.Sprintf(`http://url/%s`, verificationID)))
	if err != nil {
		return
	}

	err = uc.notification.SendMessage(ctx, uuid.New().String(), msg)
	if err != nil {
		uc.logger.Error("error sending confirmation message to kafka: %s", err.Error())
	}
	return
}

func (uc *useCase) ResetPasswordRequest(ctx context.Context, email email.Email) (err error) {
	user, err := uc.usersStorage.ReadUserByEmail(ctx, email)
	if err != nil {
		return
	}
	resetRequestID := uuid.NewString()
	msg, err := helpers.GenerateEmail(email.String(), resetSubject, fmt.Sprintf(resetTemplate, fmt.Sprintf(`http://url/%s`, resetRequestID)))
	if err != nil {
		return
	}

	err = uc.notification.SendMessage(ctx, uuid.New().String(), msg)
	if err != nil {
		uc.logger.Error("error sending reset password message to kafka: ", err.Error())
	}
	uc.cache.Set(resetRequestID, user.ID(), resetPasswordSessionTTL)
	return
}

func (uc *useCase) ResetPasswordProcess(ctx context.Context, resetRequestID uuid.UUID, newPass pass.Pass) (err error) {

	id, err := uc.cache.Get(resetRequestID.String())
	if err != nil {
		err = errors.New(ErrNoResetRequest)
		return
	}

	userID, ok := id.(uuid.UUID)
	if !ok {
		err = errors.New(ErrNoResetRequest)
		return
	}

	userU, err := uc.usersStorage.ReadUserByID(ctx, userID)
	if err != nil {
		return
	}

	userU.SetPassword(newPass)

	_, err = uc.usersStorage.UpdateUser(ctx, userU.ID(), func(oldUser *user.User) (*user.User, error) {
		return user.NewWithID(oldUser.ID(), oldUser.CreatedAt(), time.Now(), oldUser.Name(), oldUser.Surname(), oldUser.Phone(), userU.Pass(), oldUser.Email(), oldUser.Verified(), oldUser.Role())
	})

	if err != nil {
		return
	}
	return
}
