package employee

import (
	"time"

	"github.com/google/uuid"
)

type Employee struct {
	id         uuid.UUID
	createdAt  time.Time
	modifiedAt time.Time

	role      EmployeeRole
	userID    uuid.UUID
	projectID uuid.UUID
}

func NewWithID(
	id uuid.UUID,
	createdAt time.Time,
	modifiedAt time.Time,

	role EmployeeRole,
	userID uuid.UUID,
	projectID uuid.UUID,
) (*Employee, error) {
	return &Employee{
		id:         id,
		createdAt:  createdAt,
		modifiedAt: modifiedAt,
		role:       role,
		userID:     userID,
		projectID:  projectID,
	}, nil
}

func New(
	role EmployeeRole,
	userID uuid.UUID,
	projectID uuid.UUID,
) (*Employee, error) {
	id := uuid.New()
	timeNow := time.Now()
	return &Employee{
		id:         id,
		role:       role,
		userID:     userID,
		projectID:  projectID,
		createdAt:  timeNow,
		modifiedAt: timeNow,
	}, nil
}

func (e Employee) ID() uuid.UUID {
	return e.id
}

func (e Employee) Role() EmployeeRole {
	return e.role
}

func (e Employee) UserID() uuid.UUID {
	return e.userID
}

func (e Employee) ProjectID() uuid.UUID {
	return e.projectID
}

func (e Employee) CreatedAt() time.Time {
	return e.createdAt
}

func (e Employee) ModifiedAt() time.Time {
	return e.modifiedAt
}
