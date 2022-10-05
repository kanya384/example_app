package postgres

import (
	"auth/internal/repository/postgres/device"
	"auth/internal/repository/postgres/user"
	"auth/pkg/psql"
)

type Repository struct {
	Users   *user.Repository
	Devices *device.Repository
}

func NewRepository(pg *psql.Postgres) (*Repository, error) {
	users, err := user.New(pg, user.Options{})
	if err != nil {
		return nil, err
	}
	devices, err := device.New(pg, device.Options{})
	if err != nil {
		return nil, err
	}

	return &Repository{
		users,
		devices,
	}, nil

}
