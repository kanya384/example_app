package company

import (
	company "companies/internal/delivery/grpc/interface"
	dCompany "companies/internal/domain/company"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToCompanyResponse(response *dCompany.Company) *company.CompanyResponse {
	return &company.CompanyResponse{
		Id:         response.ID().String(),
		Name:       response.Name().String(),
		Inn:        uint32(response.Inn()),
		Address:    response.Address().String(),
		CreatedAt:  timestamppb.New(response.ModifiedAt()),
		ModifiedAt: timestamppb.New(response.ModifiedAt()),
	}
}
