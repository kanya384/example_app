package dao

import (
	"time"

	"github.com/google/uuid"
)

type Device struct {
	ID           uuid.UUID `db:"id"`
	UserID       uuid.UUID `db:"user_id"`
	DeviceID     string    `db:"device_id"`
	Ip           string    `db:"ip"`
	Agent        string    `db:"agent"`
	Dtype        int       `db:"dtype"`
	RefreshToken string    `db:"refresh_token"`
	RefreshExp   time.Time `db:"refresh_exp"`
	LastSeen     time.Time `db:"last_seen"`
	CreatedAt    time.Time `db:"created_at"`
	ModifiedAt   time.Time `db:"modified_at"`
}

var ColumnsDevice = []string{
	"id",
	"user_id",
	"device_id",
	"ip",
	"agent",
	"dtype",
	"refresh_token",
	"refresh_exp",
	"last_seen",
	"created_at",
	"modified_at",
}
