package useCase

import (
	"context"
	"errors"
	"testing"
	"time"

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
	userRepo "auth/internal/repository/postgres/user"
	"auth/test/mocks/auth"
	"auth/test/mocks/cache"
	"auth/test/mocks/logger"
	"auth/test/mocks/pubsub"
	"auth/test/mocks/storage"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func setupUseCase(t *testing.T) UseCaseSuite {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	usersMock := storage.NewMockUser(mockCtrl)
	devicesMock := storage.NewMockDevice(mockCtrl)
	cacheMock := cache.NewMockCache(mockCtrl)
	pubsubMock := pubsub.NewMockNotification(mockCtrl)
	tokenManager := auth.NewMockTokenManager(mockCtrl)
	loggerMock := logger.NewMockInterface(mockCtrl)
	useCase := New(usersMock, devicesMock, cacheMock, pubsubMock, tokenManager, loggerMock, Options{TokenTTL: time.Second})
	return UseCaseSuite{
		any:          gomock.Any(),
		usersMock:    usersMock,
		devicesMock:  devicesMock,
		cacheMock:    cacheMock,
		pubsubMock:   pubsubMock,
		tokenManager: tokenManager,
		loggerMock:   loggerMock,
		useCase:      useCase,
	}
}

type UseCaseSuite struct {
	any          gomock.Matcher
	cacheMock    *cache.MockCache
	pubsubMock   *pubsub.MockNotification
	usersMock    *storage.MockUser
	devicesMock  *storage.MockDevice
	tokenManager *auth.MockTokenManager
	loggerMock   *logger.MockInterface
	useCase      UseCase
}

func TestSignUp(t *testing.T) {
	req := require.New(t)
	useCaseSuite := setupUseCase(t)
	testUser := createTestUser()
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		useCaseSuite.usersMock.EXPECT().CreateUser(ctx, testUser).Return(nil).Times(1)
		useCaseSuite.cacheMock.EXPECT().Set(useCaseSuite.any, useCaseSuite.any, useCaseSuite.any).Times(1)
		useCaseSuite.pubsubMock.EXPECT().SendMessage(useCaseSuite.any, useCaseSuite.any, useCaseSuite.any).Times(1)
		err := useCaseSuite.useCase.SignUp(ctx, testUser)
		req.Empty(err)
	})

	t.Run("error", func(t *testing.T) {
		useCaseSuite.usersMock.EXPECT().CreateUser(ctx, testUser).Return(userRepo.ErrDuplicateKey).Times(1)
		err := useCaseSuite.useCase.SignUp(context.Background(), testUser)
		req.Equal(userRepo.ErrDuplicateKey, err)
	})

}

func TestSignIn(t *testing.T) {
	req := require.New(t)
	useCaseSuite := setupUseCase(t)
	testUser := createTestUser()
	inputDevice := createTestDevice()
	ctx := context.Background()
	token := "test"

	t.Run("success old device", func(t *testing.T) {
		useCaseSuite.usersMock.EXPECT().ReadUserByCredetinals(ctx, testUser.Phone(), testUser.Pass()).Return(testUser, nil).Times(1)
		useCaseSuite.devicesMock.EXPECT().ReadDevicesByDeviceID(ctx, inputDevice.DeviceID()).Return(inputDevice, nil).Times(1)
		useCaseSuite.devicesMock.EXPECT().UpdateDevice(ctx, useCaseSuite.any, useCaseSuite.any).Return(inputDevice, nil).Times(1)
		useCaseSuite.tokenManager.EXPECT().NewJWT(useCaseSuite.any, useCaseSuite.any).Return(token, nil).Times(1)

		authParams, err := useCaseSuite.useCase.SignIn(context.Background(), testUser.Phone(), testUser.Pass(), inputDevice)
		req.Empty(err)
		req.Equal(authParams.Token, token)
		req.NotEmpty(authParams.Refresh)
	})

	t.Run("success new device", func(t *testing.T) {
		useCaseSuite.usersMock.EXPECT().ReadUserByCredetinals(ctx, testUser.Phone(), testUser.Pass()).Return(testUser, nil).Times(1)
		useCaseSuite.devicesMock.EXPECT().ReadDevicesByDeviceID(ctx, inputDevice.DeviceID()).Return(nil, errors.New("not found")).Times(1)
		useCaseSuite.devicesMock.EXPECT().CreateDevice(ctx, inputDevice).Return(nil).Times(1)
		useCaseSuite.tokenManager.EXPECT().NewJWT(useCaseSuite.any, useCaseSuite.any).Return(token, nil).Times(1)

		authParams, err := useCaseSuite.useCase.SignIn(context.Background(), testUser.Phone(), testUser.Pass(), inputDevice)
		req.Empty(err)
		req.Equal(authParams.Token, token)
		req.NotEmpty(authParams.Refresh)
	})

	t.Run("no such user", func(t *testing.T) {
		useCaseSuite.usersMock.EXPECT().ReadUserByCredetinals(ctx, testUser.Phone(), testUser.Pass()).Return(nil, errors.New("not found")).Times(1)
		authParams, err := useCaseSuite.useCase.SignIn(context.Background(), testUser.Phone(), testUser.Pass(), inputDevice)
		req.Equal(err, errors.New("not found"))
		req.Empty(authParams)
	})
}

func TestVerifyEmail(t *testing.T) {
	req := require.New(t)
	useCaseSuite := setupUseCase(t)
	ctx := context.Background()
	verificationCode := uuid.New()
	testUser := createTestUser()

	t.Run("success", func(t *testing.T) {
		useCaseSuite.cacheMock.EXPECT().Get(useCaseSuite.any).Return(testUser.ID(), nil).Times(1)
		useCaseSuite.usersMock.EXPECT().ReadUserByID(ctx, testUser.ID()).Return(testUser, nil).Times(1)
		useCaseSuite.usersMock.EXPECT().UpdateUser(useCaseSuite.any, testUser.ID(), useCaseSuite.any).Return(testUser, nil).Times(1)
		err := useCaseSuite.useCase.VerifyEmail(ctx, verificationCode)
		req.Empty(err)
	})

	t.Run("error", func(t *testing.T) {
		useCaseSuite.cacheMock.EXPECT().Get(useCaseSuite.any).Return(nil, errors.New("not found")).Times(1)
		err := useCaseSuite.useCase.VerifyEmail(ctx, verificationCode)
		req.Equal(err, errors.New("not found"))
	})
}

func TestResendVerificationCode(t *testing.T) {
	req := require.New(t)
	useCaseSuite := setupUseCase(t)
	ctx := context.Background()
	testUser := createTestUser()

	t.Run("success", func(t *testing.T) {
		useCaseSuite.usersMock.EXPECT().ReadUserByID(ctx, testUser.ID()).Return(testUser, nil).Times(1)
		useCaseSuite.cacheMock.EXPECT().Set(useCaseSuite.any, useCaseSuite.any, useCaseSuite.any).Times(1)
		useCaseSuite.pubsubMock.EXPECT().SendMessage(useCaseSuite.any, useCaseSuite.any, useCaseSuite.any).Times(1)
		err := useCaseSuite.useCase.ResendVerificationCode(ctx, testUser.ID())
		req.Empty(err)
	})

	t.Run("not found", func(t *testing.T) {
		useCaseSuite.usersMock.EXPECT().ReadUserByID(ctx, testUser.ID()).Return(nil, errors.New("not found")).Times(1)
		err := useCaseSuite.useCase.ResendVerificationCode(ctx, testUser.ID())
		req.Equal(err, errors.New("not found"))
	})
}

func createTestDevice() (deviceT *device.Device) {

	userID := uuid.New()
	deviceID, _ := deviceID.NewDeviceID("00000-00000-4445654")
	ip, _ := ip.NewIp("10.21.3.56")
	agent, _ := agent.NewAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/52.0.2743.116 Safari/537.36 Edge/15.15063")
	dType := device.Web
	refreshToken, _ := refreshToken.New()
	deviceT, _ = device.New(userID, *deviceID, *ip, *agent, dType, *refreshToken)
	return
}

func createTestUser() (userT *user.User) {
	name, _ := name.NewName("test")
	surname, _ := surname.NewSurname("testSurname")
	phone, _ := phone.NewPhone("+7 (962) 885-41-32")
	pass, _ := pass.NewPass("Qwerty123!", "asdaqeq214234")
	email, _ := email.NewEmail("test01@mail.ru")
	role := user.Administrator

	userT, _ = user.New(*name, *surname, *phone, *pass, *email, role)
	return
}
