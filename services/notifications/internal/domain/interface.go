package domain

import (
	"notifications/internal/domain/mail/email"
	"notifications/internal/domain/mail/message"
	"time"

	"github.com/google/uuid"
)

type Notification interface {
	ID() uuid.UUID
	Recipient() email.Email
	Message() message.Message
	Status() SendStatus
	CreatedAt() time.Time
	ModifiedAt() time.Time
}
