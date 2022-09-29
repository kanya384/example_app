package user

import (
	"auth/internal/domain/user/email"
	"auth/internal/domain/user/name"
	"auth/internal/domain/user/pass"
	"auth/internal/domain/user/phone"
	"auth/internal/domain/user/surname"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

var (
	salt = "WcrMaOZhpYrpk79vFk"
)

func TestNewWithID(t *testing.T) {
	req := require.New(t)
	userID := uuid.New()
	timeNow := time.Now()
	name, _ := name.NewName("Василий")
	surname, _ := surname.NewSurname("Пирогов")
	phone, _ := phone.NewPhone("+7 900 862-15-54")
	email, _ := email.NewEmail("test01@mail.ru")
	role := Administrator

	t.Run("create company with id success", func(t *testing.T) {
		user, err := NewWithID(userID, timeNow, timeNow, *name, *surname, *phone, *email, role)
		req.Equal(err, nil)
		req.Equal(user.ID(), userID)
		req.Equal(user.Name(), *name)
		req.Equal(user.Surname(), *surname)
		req.Equal(user.Phone(), *phone)
		req.Equal(user.Email(), *email)
		req.Equal(user.Role(), role)
		req.Equal(user.ModifiedAt(), timeNow)
		req.Equal(user.CreatedAt(), timeNow)
	})
}

func TestNew(t *testing.T) {
	req := require.New(t)
	name, _ := name.NewName("Василий")
	surname, _ := surname.NewSurname("Пирогов")
	phone, _ := phone.NewPhone("+7 900 862-15-54")
	pass, _ := pass.NewPass("Password12!", salt)
	email, _ := email.NewEmail("test01@mail.ru")
	role := Administrator

	t.Run("create company with id success", func(t *testing.T) {
		user, err := New(*name, *surname, *phone, *pass, *email, role)
		req.Equal(err, nil)
		req.NotEmpty(user.ID())
		req.NotEmpty(user.CreatedAt())
		req.NotEmpty(user.ModifiedAt())
		req.Equal(user.Name(), *name)
		req.Equal(user.Surname(), *surname)
		req.Equal(user.Phone(), *phone)
		req.Equal(user.Email(), *email)
		req.Equal(user.Pass(), *pass)
		req.Equal(user.Role(), role)
	})
}
