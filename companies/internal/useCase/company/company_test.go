package company

import (
	"companies/internal/domain/company"
	"companies/internal/domain/company/inn"
	"companies/internal/domain/company/name"
	"companies/internal/repository/company/postgres"
	"companies/pkg/types/address"
	"companies/test/mocks/storage"
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestCreateCompany(t *testing.T) {
	req := require.New(t)
	any := gomock.Any()
	uc, storageMock := initUseCase(t)

	company := createTestCompany()

	t.Run("create company success", func(t *testing.T) {
		storageMock.EXPECT().CreateCompany(any, any).Return(nil).Times(1)
		err := uc.CreateCompany(context.Background(), company)
		req.Empty(err)
	})

	t.Run("duplicate key error", func(t *testing.T) {
		storageMock.EXPECT().CreateCompany(any, any).Return(postgres.ErrDuplicateKey).Times(1)
		err := uc.CreateCompany(context.Background(), company)
		req.ErrorContains(err, postgres.ErrDuplicateKey.Error())
	})
}

func TestUpdateCompany(t *testing.T) {
	req := require.New(t)
	any := gomock.Any()
	uc, storageMock := initUseCase(t)

	company := createTestCompany()

	t.Run("update company success", func(t *testing.T) {
		storageMock.EXPECT().UpdateCompany(any, any).Return(nil).Times(1)
		err := uc.UpdateCompany(context.Background(), company)
		req.Empty(err)
	})

	t.Run("not found", func(t *testing.T) {
		storageMock.EXPECT().UpdateCompany(any, any).Return(postgres.ErrNotFound).Times(1)
		err := uc.UpdateCompany(context.Background(), company)
		req.ErrorContains(err, postgres.ErrNotFound.Error())
	})
}

func TestDeleteCompany(t *testing.T) {
	req := require.New(t)
	any := gomock.Any()
	uc, storageMock := initUseCase(t)

	company := createTestCompany()

	t.Run("delete company success", func(t *testing.T) {
		storageMock.EXPECT().DeleteCompany(any, any).Return(nil).Times(1)
		err := uc.DeleteCompany(context.Background(), company.ID())
		req.Empty(err)
	})

	t.Run("not found", func(t *testing.T) {
		storageMock.EXPECT().DeleteCompany(any, any).Return(postgres.ErrNotFound).Times(1)
		err := uc.DeleteCompany(context.Background(), company.ID())
		req.ErrorContains(err, postgres.ErrNotFound.Error())
	})
}

func TestReadCompanyByID(t *testing.T) {
	req := require.New(t)
	any := gomock.Any()
	uc, storageMock := initUseCase(t)

	companyT := createTestCompany()

	t.Run("read company success", func(t *testing.T) {
		storageMock.EXPECT().ReadCompanyByID(any, any).Return(companyT, nil).Times(1)
		_, err := uc.ReadCompanyByID(context.Background(), companyT.ID())
		req.Empty(err)
	})

	t.Run("not found", func(t *testing.T) {
		storageMock.EXPECT().ReadCompanyByID(any, any).Return(&company.Company{}, postgres.ErrNotFound).Times(1)
		_, err := uc.ReadCompanyByID(context.Background(), companyT.ID())
		req.ErrorContains(err, postgres.ErrNotFound.Error())
	})
}

func createTestCompany() *company.Company {
	companyID := uuid.New()
	companyName, _ := name.NewName("ООО \"Лидген\"")
	companyInn, _ := inn.NewInn(1234567890)
	companyAddress, _ := address.NewAdress("г. Майкоп, ул. Краснооктябрьская 10А")
	testTime := time.Now()
	testCompany, _ := company.NewWithID(companyID, *companyName, *companyInn, *companyAddress, testTime, testTime)
	return testCompany
}

func initUseCase(t *testing.T) (*UseCase, *storage.MockStorage) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	storageMock := storage.NewMockStorage(mockCtrl)

	uc := New(storageMock, Options{})
	return uc, storageMock
}
