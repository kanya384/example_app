package project

import (
	"companies/internal/domain/project/description"
	"companies/internal/domain/project/location"
	"companies/internal/domain/project/name"
	"companies/pkg/types/address"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestNewWithID(t *testing.T) {
	req := require.New(t)
	projectID := uuid.New()
	testTime := time.Now()
	projectName, _ := name.NewName("Эксельсиор")
	projectDescripiton, _ := description.NewDescription("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua")
	projectAddress, _ := address.NewAdress("ул. Депутатская, 94")
	projectLocation, _ := location.NewLocation(55.44, 44.55)
	projectCompanyID := uuid.New()

	t.Run("create company with id success", func(t *testing.T) {
		project, err := NewWithID(
			projectID, testTime, testTime, *projectName, *projectDescripiton, *projectAddress, *projectLocation, projectCompanyID)
		req.Equal(project.Name(), *projectName)
		req.Equal(project.Address(), *projectAddress)
		req.Equal(project.Description(), *projectDescripiton)
		req.Equal(project.Location(), *projectLocation)
		req.Equal(project.CompanyID(), projectCompanyID)
		req.Equal(err, nil)
	})
}

func TestNew(t *testing.T) {
	req := require.New(t)
	projectName, _ := name.NewName("Эксельсиор")
	projectDescripiton, _ := description.NewDescription("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua")
	projectAddress, _ := address.NewAdress("ул. Депутатская, 94")
	projectLocation, _ := location.NewLocation(55.44, 44.55)
	projectCompanyID := uuid.New()

	t.Run("create company with id success", func(t *testing.T) {
		project, err := New(
			*projectName, *projectDescripiton, *projectAddress, *projectLocation, projectCompanyID)
		req.NotEmpty(project.ID())
		req.Equal(project.Name(), *projectName)
		req.Equal(project.Address(), *projectAddress)
		req.Equal(project.Description(), *projectDescripiton)
		req.Equal(project.Location(), *projectLocation)
		req.Equal(project.CompanyID(), projectCompanyID)
		req.Equal(err, nil)
	})
}
