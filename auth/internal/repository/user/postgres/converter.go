package postgres

import (
	"auth/internal/domain/user"
	"auth/internal/domain/user/email"
	"auth/internal/domain/user/name"
	"auth/internal/domain/user/phone"
	"auth/internal/domain/user/surname"
	"auth/internal/repository/user/postgres/dao"
)

func (r Repository) toDomainUser(dao *dao.User) (result *user.User, err error) {
	name, err := name.NewName(dao.Name)
	if err != nil {
		return
	}
	surname, err := surname.NewSurname(dao.Surname)
	if err != nil {
		return
	}
	phone, err := phone.NewPhone(dao.Phone)
	if err != nil {
		return
	}
	email, err := email.NewEmail(dao.Email)
	if err != nil {
		return
	}

	result, err = user.NewWithID(
		dao.ID,
		dao.CreatedAt,
		dao.ModifiedAt,
		*name,
		*surname,
		*phone,
		*email,
		user.UserRole(dao.Role),
	)
	if err != nil {
		return nil, err
	}
	return result, nil
}
