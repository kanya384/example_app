package grpc

import (
	company "companies/internal/delivery/grpc/interface"
	dCompany "companies/internal/domain/company"
	"companies/internal/domain/company/inn"
	"companies/internal/domain/company/name"
	contextI "companies/pkg/context"
	"companies/pkg/helpers"
	"companies/pkg/types/address"
	"context"
	"fmt"
	"time"

	grpcCompany "companies/internal/delivery/grpc/company"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
)

func (d *Delivery) CreateCompany(c context.Context, request *company.CreateCompanyRequest) (*company.CreateCompanyResponse, error) {
	var ctx = contextI.New(c)

	name, err := name.NewName(request.Name)
	if err != nil {
		return nil, helpers.NewGrpcError(codes.InvalidArgument, err)
	}

	inn, err := inn.NewInn(request.Inn)
	if err != nil {
		return nil, helpers.NewGrpcError(codes.InvalidArgument, err)
	}

	address, err := address.NewAdress(request.Address)
	if err != nil {
		return nil, helpers.NewGrpcError(codes.InvalidArgument, err)
	}

	response, _ := dCompany.New(*name, *inn, *address)

	err = d.ucCompany.CreateCompany(ctx, response)
	if err != nil {
		return nil, helpers.NewGrpcError(codes.Internal, err)
	}

	return &company.CreateCompanyResponse{Response: grpcCompany.ToCompanyResponse(response)}, nil
}

func (d *Delivery) UpdateCompany(c context.Context, request *company.UpdateCompanyRequest) (*company.UpdateCompanyResponse, error) {
	var ctx = contextI.New(c)

	id, err := uuid.Parse(request.Id)
	if err != nil {
		return nil, helpers.NewGrpcError(codes.InvalidArgument, fmt.Errorf("invalid ID format %s", err))
	}

	name, err := name.NewName(request.Name)
	if err != nil {
		return nil, helpers.NewGrpcError(codes.InvalidArgument, err)
	}

	inn, err := inn.NewInn(request.Inn)
	if err != nil {
		return nil, helpers.NewGrpcError(codes.InvalidArgument, err)
	}

	address, err := address.NewAdress(request.Address)
	if err != nil {
		return nil, helpers.NewGrpcError(codes.InvalidArgument, err)
	}

	response, _ := dCompany.NewWithID(id, *name, *inn, *address, time.Now(), time.Now())

	err = d.ucCompany.UpdateCompany(ctx, response)
	if err != nil {
		return nil, helpers.NewGrpcError(codes.Internal, err)
	}

	return &company.UpdateCompanyResponse{Response: grpcCompany.ToCompanyResponse(response)}, nil
}

func (d *Delivery) DeleteCompany(c context.Context, request *company.DeleteCompanyRequest) (*company.DeleteCompanyResponse, error) {
	var ctx = contextI.New(c)

	id, err := uuid.Parse(request.Id)
	if err != nil {
		return nil, helpers.NewGrpcError(codes.InvalidArgument, fmt.Errorf("invalid ID format %s", err))
	}

	err = d.ucCompany.DeleteCompany(ctx, id)
	if err != nil {
		return &company.DeleteCompanyResponse{Success: false}, helpers.NewGrpcError(codes.Internal, err)
	}

	return &company.DeleteCompanyResponse{Success: true}, nil
}

func (d *Delivery) ReadCompanyByID(c context.Context, request *company.ReadCompanyByIdRequest) (*company.ReadCompanyByIdResponse, error) {
	var ctx = contextI.New(c)

	id, err := uuid.Parse(request.Id)
	if err != nil {
		return nil, helpers.NewGrpcError(codes.InvalidArgument, fmt.Errorf("invalid ID format %s", err))
	}

	response, err := d.ucCompany.ReadCompanyByID(ctx, id)
	if err != nil {
		return nil, helpers.NewGrpcError(codes.Internal, err)
	}

	return &company.ReadCompanyByIdResponse{Response: grpcCompany.ToCompanyResponse(response)}, nil
}
