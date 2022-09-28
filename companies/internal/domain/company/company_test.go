package company

import (
	"companies/internal/domain/company/inn"
	"companies/internal/domain/company/name"
	"companies/pkg/types/address"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestNewWithID(t *testing.T) {
	req := require.New(t)
	companyID := uuid.New()
	testTime := time.Now()
	companyName := name.Name("ООО \"Лидген\"")
	companyInn := inn.Inn(1234567890)
	companyAddress := address.Address("г. Майкоп, ул. Краснооктябрьская 10А")

	t.Run("create company with id success", func(t *testing.T) {
		company, err := NewWithID(companyID, companyName, companyInn, companyAddress, testTime, testTime)
		req.Equal(err, nil)
		req.Equal(company.ID(), companyID)
		req.Equal(company.Name(), companyName)
		req.Equal(company.Inn(), companyInn)
		req.Equal(company.Address(), companyAddress)
		req.Equal(company.CreatedAt(), testTime)
		req.Equal(company.ModifiedAt(), testTime)
	})
}

func TestNew(t *testing.T) {
	req := require.New(t)
	companyName := name.Name("ООО \"Лидген\"")
	companyInn := inn.Inn(1234567890)
	companyAddress := address.Address("г. Майкоп, ул. Краснооктябрьская 10А")

	t.Run("create company success", func(t *testing.T) {
		company, err := New(companyName, companyInn, companyAddress)
		req.Equal(err, nil)
		req.NotEmpty(company.ID())
		req.Equal(company.Name(), companyName)
		req.Equal(company.Inn(), companyInn)
		req.Equal(company.Address(), companyAddress)
		req.NotEmpty(company.CreatedAt())
		req.NotEmpty(company.ModifiedAt())
	})
}
