package storage

import (
	"auth/internal/domain/device"
	"auth/internal/domain/device/deviceID"
	"auth/internal/domain/device/refreshToken"
	"auth/internal/domain/user"
	"auth/internal/domain/user/email"
	"auth/internal/domain/user/pass"
	"auth/internal/domain/user/phone"
	"context"

	"github.com/google/uuid"
)

type User interface {
	CreateUser(ctx context.Context, user *user.User) (err error)
	UpdateUser(ctx context.Context, ID uuid.UUID, updateFn func(user *user.User) (*user.User, error)) (user *user.User, err error)
	DeleteUser(ctx context.Context, ID uuid.UUID) (err error)
	ReadUserByID(ctx context.Context, ID uuid.UUID) (user *user.User, err error)
	ReadUserByCredetinals(ctx context.Context, phone phone.Phone, pass pass.Pass) (*user.User, error)
	ReadUserByEmail(ctx context.Context, email email.Email) (*user.User, error)
}

type Device interface {
	CreateDevice(ctx context.Context, device *device.Device) (err error)
	UpdateDevice(ctx context.Context, ID uuid.UUID, updateFn func(*device.Device) (*device.Device, error)) (device *device.Device, err error)
	DeleteDevice(ctx context.Context, ID uuid.UUID) (err error)
	ReadDeviceByID(ctx context.Context, ID uuid.UUID) (device *device.Device, err error)
	ReadDeviceByUserIDAndRefresh(ctx context.Context, userID uuid.UUID, refreshToken refreshToken.RefreshToken) (device *device.Device, err error)
	ReadDevicesByDeviceID(ctx context.Context, deviceID deviceID.DeviceID) (device *device.Device, err error)
}
