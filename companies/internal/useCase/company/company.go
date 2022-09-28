package company

import (
	"companies/internal/domain/company"
	"context"

	"github.com/google/uuid"
)

func (uc *UseCase) CreateCompany(ctx context.Context, company *company.Company) (err error) {
	return uc.adaptersStorage.CreateCompany(ctx, company)
}

func (uc *UseCase) UpdateCompany(ctx context.Context, company *company.Company) (err error) {
	return uc.adaptersStorage.UpdateCompany(ctx, company)
}

func (uc *UseCase) DeleteCompany(ctx context.Context, ID uuid.UUID) (err error) {
	return uc.adaptersStorage.DeleteCompany(ctx, ID)
}

func (uc *UseCase) ReadCompanyByID(ctx context.Context, ID uuid.UUID) (company *company.Company, err error) {
	return uc.adaptersStorage.ReadCompanyByID(ctx, ID)
}
