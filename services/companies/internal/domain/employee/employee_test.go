package employee

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestNewWithID(t *testing.T) {
	req := require.New(t)
	timeNow := time.Now()
	role := Administrator
	employeeID := uuid.New()
	userID := uuid.New()
	projectID := uuid.New()

	t.Run("create employee with id success", func(t *testing.T) {
		employee, err := NewWithID(employeeID, timeNow, timeNow, role, userID, projectID)
		req.Equal(err, nil)
		req.Equal(employee.ID(), employeeID)
		req.Equal(employee.UserID(), userID)
		req.Equal(employee.ProjectID(), projectID)
		req.Equal(employee.Role(), role)
		req.Equal(employee.CreatedAt(), timeNow)
		req.Equal(employee.ModifiedAt(), timeNow)
	})
}

func TestNew(t *testing.T) {
	req := require.New(t)
	role := Administrator
	projectID := uuid.New()
	userID := uuid.New()

	t.Run("create employee success", func(t *testing.T) {
		employee, err := New(role, userID, projectID)
		req.Equal(err, nil)
		req.NotNil(employee.ID())
		req.Equal(employee.UserID(), userID)
		req.Equal(employee.ProjectID(), projectID)
		req.Equal(employee.Role(), role)
		req.NotNil(employee.CreatedAt())
		req.NotNil(employee.ModifiedAt())
	})
}
