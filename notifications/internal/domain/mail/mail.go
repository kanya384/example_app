package mail

import (
	"notifications/internal/domain"
	"notifications/internal/domain/mail/email"
	"notifications/internal/domain/mail/message"
	"notifications/internal/domain/mail/subject"
	"time"

	"github.com/google/uuid"
)

const (
	MaxSendTries = 8
)

type Mail struct {
	id         uuid.UUID
	createdAt  time.Time
	modifiedAt time.Time

	recipient email.Email
	subject   subject.Subject
	message   message.Message
	status    domain.SendStatus
}

func NewWithID(
	id uuid.UUID,
	createdAt time.Time,
	modifiedAt time.Time,

	recipient email.Email,
	subject subject.Subject,
	message message.Message,
	status domain.SendStatus,
) (mail *Mail, err error) {
	return &Mail{
		id:         id,
		createdAt:  createdAt,
		modifiedAt: modifiedAt,

		recipient: recipient,
		subject:   subject,
		message:   message,
		status:    status,
	}, nil
}

func New(
	recipient email.Email,
	subject subject.Subject,
	message message.Message,
) (mail *Mail, err error) {
	timeNew := time.Now()
	return &Mail{
		id:         uuid.New(),
		createdAt:  timeNew,
		modifiedAt: timeNew,

		recipient: recipient,
		subject:   subject,
		message:   message,
		status:    domain.NotSended,
	}, nil
}

func (m *Mail) ID() uuid.UUID {
	return m.id
}

func (m *Mail) CreatedAt() time.Time {
	return m.createdAt
}

func (m *Mail) ModifiedAt() time.Time {
	return m.modifiedAt
}

func (m *Mail) Recipient() email.Email {
	return m.recipient
}

func (m *Mail) Subject() subject.Subject {
	return m.subject
}

func (m *Mail) Message() message.Message {
	return m.message
}

func (m *Mail) Status() domain.SendStatus {
	return m.status
}
