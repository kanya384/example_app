package dao

import (
	"time"

	"github.com/google/uuid"
)

type Mail struct {
	ID         uuid.UUID `db:"id"`
	Recipient  string    `db:"recipient"`
	Subject    string    `db:"subject"`
	Message    string    `db:"message"`
	Status     uint8     `db:"status"`
	CreatedAt  time.Time `db:"created_at"`
	ModifiedAt time.Time `db:"modified_at"`
}

var ColumnsMail = []string{
	"id",
	"recipient",
	"subject",
	"message",
	"status",
	"created_at",
	"modified_at",
}
