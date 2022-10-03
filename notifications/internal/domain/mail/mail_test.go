package mail

import (
	"notifications/internal/domain"
	"notifications/internal/domain/mail/email"
	"notifications/internal/domain/mail/message"
	"notifications/internal/domain/mail/subject"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestNewWithID(t *testing.T) {
	req := require.New(t)
	id := uuid.New()
	timeNow := time.Now()
	recipient, _ := email.NewEmail("test01@mail.ru")
	subject, _ := subject.NewSubject("test subject")
	message, _ := message.NewMessage("test message body")
	status := domain.NotSended
	t.Run("create message with id success", func(t *testing.T) {
		device, err := NewWithID(id, timeNow, timeNow, *recipient, *subject, *message, status)
		req.Equal(err, nil)
		req.Equal(device.ID(), id)
		req.Equal(device.CreatedAt(), timeNow)
		req.Equal(device.ModifiedAt(), timeNow)
		req.Equal(device.Recipient(), *recipient)
		req.Equal(device.Message(), *message)
		req.Equal(device.Status(), status)
	})
}

func TestNew(t *testing.T) {
	req := require.New(t)
	recipient, _ := email.NewEmail("test01@mail.ru")
	subject, _ := subject.NewSubject("test subject")
	message, _ := message.NewMessage("test message body")
	t.Run("create message", func(t *testing.T) {
		device, err := New(*recipient, *subject, *message)
		req.Equal(err, nil)
		req.NotEmpty(device.ID())
		req.NotEmpty(device.CreatedAt())
		req.NotEmpty(device.ModifiedAt())
		req.Equal(device.Recipient(), *recipient)
		req.Equal(device.Message(), *message)
		req.Equal(device.Status(), domain.NotSended)
	})
}
