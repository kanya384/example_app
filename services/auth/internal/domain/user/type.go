package user

import (
	"auth/internal/domain/user/email"
	"auth/internal/domain/user/name"
	"auth/internal/domain/user/pass"
	"auth/internal/domain/user/phone"
	"auth/internal/domain/user/surname"
	"time"

	"github.com/google/uuid"
)

type User struct {
	id         uuid.UUID
	createdAt  time.Time
	modifiedAt time.Time

	name     name.Name
	surname  surname.Surname
	phone    phone.Phone
	pass     pass.Pass
	email    email.Email
	verified bool
	role     UserRole
}

func NewWithID(
	id uuid.UUID,
	createdAt time.Time,
	modifiedAt time.Time,

	name name.Name,
	surname surname.Surname,
	phone phone.Phone,
	pass pass.Pass,
	email email.Email,
	verified bool,
	role UserRole,
) (*User, error) {
	return &User{
		id:         id,
		createdAt:  createdAt,
		modifiedAt: modifiedAt,

		name:    name,
		surname: surname,
		phone:   phone,
		pass:    pass,
		email:   email,
		role:    role,
	}, nil
}

func New(
	name name.Name,
	surname surname.Surname,
	phone phone.Phone,
	pass pass.Pass,
	email email.Email,
	role UserRole,
) (*User, error) {
	var timeNow = time.Now().UTC()
	return &User{
		id:         uuid.New(),
		createdAt:  timeNow,
		modifiedAt: timeNow,

		name:     name,
		surname:  surname,
		phone:    phone,
		pass:     pass,
		email:    email,
		verified: false,
		role:     role,
	}, nil
}

func (c User) ID() uuid.UUID {
	return c.id
}

func (c User) CreatedAt() time.Time {
	return c.createdAt
}

func (c User) ModifiedAt() time.Time {
	return c.modifiedAt
}

func (c User) Name() name.Name {
	return c.name
}

func (c User) Surname() surname.Surname {
	return c.surname
}

func (c User) Phone() phone.Phone {
	return c.phone
}

func (c User) Pass() pass.Pass {
	return c.pass
}

func (c User) Email() email.Email {
	return c.email
}

func (c User) Verified() bool {
	return c.verified
}

func (c *User) SetVerified(value bool) {
	c.verified = value
}

func (c User) Role() UserRole {
	return c.role
}
