package dao

import (
	"time"

	"github.com/google/uuid"
)

type Company struct {
	ID         uuid.UUID `db:"id"`
	Name       string    `db:"name"`
	Inn        uint32    `db:"inn"`
	Address    string    `db:"address"`
	CreatedAt  time.Time `db:"created_at"`
	ModifiedAt time.Time `db:"modified_at"`
}

var ColumnsCompany = []string{
	"id",
	"name",
	"inn",
	"address",
	"created_at",
	"modified_at",
}
