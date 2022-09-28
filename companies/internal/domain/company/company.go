package company

import (
	"time"

	"companies/internal/domain/company/inn"
	"companies/internal/domain/company/name"
	"companies/pkg/types/address"

	"github.com/google/uuid"
)

type Company struct {
	id         uuid.UUID
	createdAt  time.Time
	modifiedAt time.Time

	name    name.Name
	inn     inn.Inn
	address address.Address
}

func NewWithID(
	id uuid.UUID,
	name name.Name,
	inn inn.Inn,
	address address.Address,
	createdAt time.Time,
	modifiedAt time.Time,
) (company *Company, err error) {
	return &Company{
		id:         id,
		name:       name,
		inn:        inn,
		address:    address,
		createdAt:  createdAt,
		modifiedAt: modifiedAt,
	}, nil
}

func New(
	name name.Name,
	inn inn.Inn,
	address address.Address,
) (company *Company, err error) {
	timeNow := time.Now().UTC()
	return &Company{
		id:         uuid.New(),
		createdAt:  timeNow,
		modifiedAt: timeNow,

		name:    name,
		inn:     inn,
		address: address,
	}, nil
}

func (c Company) ID() uuid.UUID {
	return c.id
}

func (c Company) CreatedAt() time.Time {
	return c.createdAt
}

func (c Company) ModifiedAt() time.Time {
	return c.modifiedAt
}

func (c Company) Inn() inn.Inn {
	return c.inn
}

func (c Company) Address() address.Address {
	return c.address
}

func (c Company) Name() name.Name {
	return c.name
}
