package project

import (
	"companies/internal/domain/project/description"
	"companies/internal/domain/project/location"
	"companies/internal/domain/project/name"
	"companies/pkg/types/address"
	"time"

	"github.com/google/uuid"
)

type Project struct {
	id         uuid.UUID
	createdAt  time.Time
	modifiedAt time.Time

	name        name.Name
	description description.Description
	address     address.Address
	location    location.Location
	companyID   uuid.UUID
}

func NewWithID(
	id uuid.UUID,
	createdAt time.Time,
	modifiedAt time.Time,
	name name.Name,
	description description.Description,
	address address.Address,
	location location.Location,
	companyID uuid.UUID,
) (project *Project, err error) {
	project = &Project{
		id:          id,
		createdAt:   createdAt,
		modifiedAt:  modifiedAt,
		name:        name,
		description: description,
		address:     address,
		location:    location,
		companyID:   companyID,
	}
	return
}

func New(
	name name.Name,
	description description.Description,
	address address.Address,
	location location.Location,
	companyID uuid.UUID,
) (project *Project, err error) {
	timeNow := time.Now().UTC()
	project = &Project{
		id:          uuid.New(),
		createdAt:   timeNow,
		modifiedAt:  timeNow,
		name:        name,
		description: description,
		address:     address,
		location:    location,
		companyID:   companyID,
	}
	return
}

func (p Project) ID() uuid.UUID {
	return p.id
}

func (p Project) Name() name.Name {
	return p.name
}

func (p Project) Description() description.Description {
	return p.description
}

func (p Project) Address() address.Address {
	return p.address
}

func (p Project) Location() location.Location {
	return p.location
}

func (p Project) CompanyID() uuid.UUID {
	return p.companyID
}
