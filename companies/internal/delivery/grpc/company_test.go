package grpc

import (
	company "companies/internal/delivery/grpc/interface"
	dCompany "companies/internal/domain/company"
	"companies/internal/domain/company/inn"
	"companies/internal/domain/company/name"
	"companies/pkg/types/address"
	"companies/test/mocks/useCase"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

var (
	testName, _    = name.NewName("ООО Тесла")
	testInn, _     = inn.NewInn(1234567890)
	testAddress, _ = address.NewAdress("ул. Красная 494")
	testId         = uuid.New()
)

func TestCreateCompany(t *testing.T) {
	req := require.New(t)

	any := gomock.Any()

	delivery, useCaseMock := initDelivery(t)

	t.Run("create company success", func(t *testing.T) {
		useCaseMock.EXPECT().CreateCompany(any, any).Return(nil).Times(1)
		resp, err := delivery.CreateCompany(context.Background(), &company.CreateCompanyRequest{
			Name:    testName.String(),
			Inn:     testInn.Inn(),
			Address: testAddress.String(),
		})
		req.Empty(err)
		req.NotEmpty(resp.GetResponse().Id)
		req.Equal(resp.GetResponse().Name, testName.String())
		req.Equal(resp.GetResponse().Inn, testInn.Inn())
		req.Equal(resp.GetResponse().Address, testAddress.String())
		req.NotEmpty(resp.GetResponse().CreatedAt, resp.GetResponse().ModifiedAt)
	})

	t.Run("invalid name value", func(t *testing.T) {
		_, err := delivery.CreateCompany(context.Background(), &company.CreateCompanyRequest{
			Name:    "ООО",
			Inn:     testInn.Inn(),
			Address: testAddress.String(),
		})
		req.ErrorContains(err, name.ErrWrongLength.Error())
	})

	t.Run("invalid inn value", func(t *testing.T) {
		_, err := delivery.CreateCompany(context.Background(), &company.CreateCompanyRequest{
			Name:    testName.String(),
			Inn:     123456,
			Address: testAddress.String(),
		})
		req.ErrorContains(err, inn.ErrWrongLength.Error())
	})

	t.Run("invalid address value", func(t *testing.T) {
		_, err := delivery.CreateCompany(context.Background(), &company.CreateCompanyRequest{
			Name:    testName.String(),
			Inn:     testInn.Inn(),
			Address: "s",
		})
		req.ErrorContains(err, address.ErrWrongLength.Error())
	})

	t.Run("useCase error", func(t *testing.T) {
		errR := errors.New("ERROR: duplicate key value violates unique constraint \"company_pkey\" (SQLSTATE 23505)")
		useCaseMock.EXPECT().CreateCompany(any, any).Return(errR).Times(1)
		_, err := delivery.CreateCompany(context.Background(), &company.CreateCompanyRequest{
			Name:    testName.String(),
			Inn:     testInn.Inn(),
			Address: testAddress.String(),
		})
		req.NotEmpty(err)
	})
}

func TestUpdateCompany(t *testing.T) {
	req := require.New(t)

	any := gomock.Any()

	delivery, useCaseMock := initDelivery(t)

	t.Run("update company success", func(t *testing.T) {
		useCaseMock.EXPECT().UpdateCompany(any, any).Return(nil).Times(1)
		resp, err := delivery.UpdateCompany(context.Background(), &company.UpdateCompanyRequest{
			Id:      testId.String(),
			Name:    testName.String(),
			Inn:     testInn.Inn(),
			Address: testAddress.String(),
		})
		req.Empty(err)
		req.Equal(resp.GetResponse().Id, testId.String())
		req.Equal(resp.GetResponse().Name, testName.String())
		req.Equal(resp.GetResponse().Inn, testInn.Inn())
		req.Equal(resp.GetResponse().Address, testAddress.String())
		req.NotEmpty(resp.GetResponse().CreatedAt, resp.GetResponse().ModifiedAt)
	})

	t.Run("invalid id value", func(t *testing.T) {
		_, err := delivery.UpdateCompany(context.Background(), &company.UpdateCompanyRequest{
			Id:      "sdsd",
			Name:    testName.String(),
			Inn:     testInn.Inn(),
			Address: testAddress.String(),
		})
		req.ErrorContains(err, "InvalidArgument")
	})

	t.Run("invalid name value", func(t *testing.T) {
		_, err := delivery.UpdateCompany(context.Background(), &company.UpdateCompanyRequest{
			Id:      testId.String(),
			Name:    "ООО",
			Inn:     testInn.Inn(),
			Address: testAddress.String(),
		})
		req.ErrorContains(err, name.ErrWrongLength.Error())
	})

	t.Run("invalid inn value", func(t *testing.T) {
		_, err := delivery.UpdateCompany(context.Background(), &company.UpdateCompanyRequest{
			Id:      testId.String(),
			Name:    testName.String(),
			Inn:     123456,
			Address: testAddress.String(),
		})
		req.ErrorContains(err, inn.ErrWrongLength.Error())
	})

	t.Run("invalid address value", func(t *testing.T) {
		_, err := delivery.UpdateCompany(context.Background(), &company.UpdateCompanyRequest{
			Id:      testId.String(),
			Name:    testName.String(),
			Inn:     testInn.Inn(),
			Address: "s",
		})
		req.ErrorContains(err, address.ErrWrongLength.Error())
	})

	t.Run("useCase error", func(t *testing.T) {
		errR := errors.New("not found")
		useCaseMock.EXPECT().UpdateCompany(any, any).Return(errR).Times(1)
		_, err := delivery.UpdateCompany(context.Background(), &company.UpdateCompanyRequest{
			Id:      testId.String(),
			Name:    testName.String(),
			Inn:     testInn.Inn(),
			Address: testAddress.String(),
		})
		req.ErrorContains(err, "Internal")
	})
}

func TestDeleteCompany(t *testing.T) {
	req := require.New(t)

	any := gomock.Any()

	delivery, useCaseMock := initDelivery(t)

	t.Run("delete company success", func(t *testing.T) {
		useCaseMock.EXPECT().DeleteCompany(any, any).Return(nil).Times(1)
		_, err := delivery.DeleteCompany(context.Background(), &company.DeleteCompanyRequest{Id: uuid.NewString()})
		req.Empty(err)
	})

	t.Run("wrong ig", func(t *testing.T) {
		useCaseMock.EXPECT().DeleteCompany(any, any).Return(errors.New("not foun")).Times(1)
		_, err := delivery.DeleteCompany(context.Background(), &company.DeleteCompanyRequest{Id: "sd"})
		req.ErrorContains(err, "InvalidArgument")
	})

	t.Run("delete company error", func(t *testing.T) {
		useCaseMock.EXPECT().DeleteCompany(any, any).Return(errors.New("not foun")).Times(1)
		_, err := delivery.DeleteCompany(context.Background(), &company.DeleteCompanyRequest{Id: uuid.NewString()})
		req.ErrorContains(err, "Internal")
	})
}

func TestReadCompanyByID(t *testing.T) {
	req := require.New(t)

	any := gomock.Any()

	delivery, useCaseMock := initDelivery(t)

	testCompany, _ := dCompany.New(*testName, *testInn, *testAddress)

	t.Run("read company success", func(t *testing.T) {

		useCaseMock.EXPECT().ReadCompanyByID(any, any).Return(testCompany, nil).Times(1)
		_, err := delivery.ReadCompanyByID(context.Background(), &company.ReadCompanyByIdRequest{Id: testCompany.ID().String()})
		req.Empty(err)
	})

	t.Run("wrong ig", func(t *testing.T) {
		useCaseMock.EXPECT().ReadCompanyByID(any, any).Return(nil, errors.New("not found")).Times(1)
		_, err := delivery.ReadCompanyByID(context.Background(), &company.ReadCompanyByIdRequest{Id: "sd"})
		req.ErrorContains(err, "InvalidArgument")
	})

	t.Run("read company error", func(t *testing.T) {
		errN := errors.New("not found")

		useCaseMock.EXPECT().ReadCompanyByID(any, any).Return(nil, errN).Times(1)
		_, err := delivery.ReadCompanyByID(context.Background(), &company.ReadCompanyByIdRequest{Id: testCompany.ID().String()})
		req.ErrorContains(err, "Internal")
	})
}

func initDelivery(t *testing.T) (*Delivery, *useCase.MockStorage) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	useCaseMock := useCase.NewMockStorage(mockCtrl)

	delivery := New(useCaseMock, Options{DefaultTimeout: time.Duration(time.Second * 5)})

	return delivery, useCaseMock
}
