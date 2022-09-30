package storage

import (
	"auth/internal/domain/device"
	"context"
	"os/user"

	"github.com/google/uuid"
)

type Storage interface {
	User
}

type User interface {
	CreateUser(ctx context.Context, user *user.User) (err error)
	UpdateUser(ctx context.Context, ID uuid.UUID, updateFn func(user *user.User) (*user.User, error)) (user *user.User, err error)
	DeleteUser(ctx context.Context, ID uuid.UUID) (err error)
	ReadUserByID(ctx context.Context, ID uuid.UUID) (user *user.User, err error)
}

type Device interface {
	CreateDevice(ctx context.Context, device *device.Device) (err error)
	UpdateDevice(ctx context.Context, ID uuid.UUID, updateFn func(device *device.Device) (*device.Device, error)) (device *device.Device, err error)
	DeleteDevice(ctx context.Context, ID uuid.UUID) (err error)
	ReadDeviceByID(ctx context.Context, ID uuid.UUID) (device *device.Device, err error)
}
