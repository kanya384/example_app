package dao

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID `db:"id"`
	Name       string    `db:"name"`
	Surname    string    `db:"surname"`
	Phone      string    `db:"phone"`
	Pass       string    `db:"pass"`
	Email      string    `db:"email"`
	Verified   bool      `db:"verified"`
	Role       string    `db:"int"`
	CreatedAt  time.Time `db:"created_at"`
	ModifiedAt time.Time `db:"modified_at"`
}

var ColumnsUser = []string{
	"id",
	"name",
	"surname",
	"phone",
	"pass",
	"email",
	"role",
	"created_at",
	"modified_at",
}
