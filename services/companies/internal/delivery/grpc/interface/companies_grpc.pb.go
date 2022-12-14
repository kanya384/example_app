// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package company

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CompanyClient is the client API for Company service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CompanyClient interface {
	CreateCompany(ctx context.Context, in *CreateCompanyRequest, opts ...grpc.CallOption) (*CreateCompanyResponse, error)
	UpdateCompany(ctx context.Context, in *UpdateCompanyRequest, opts ...grpc.CallOption) (*UpdateCompanyResponse, error)
	DeleteCompany(ctx context.Context, in *DeleteCompanyRequest, opts ...grpc.CallOption) (*DeleteCompanyResponse, error)
	ReadCompanyByID(ctx context.Context, in *ReadCompanyByIdRequest, opts ...grpc.CallOption) (*ReadCompanyByIdResponse, error)
}

type companyClient struct {
	cc grpc.ClientConnInterface
}

func NewCompanyClient(cc grpc.ClientConnInterface) CompanyClient {
	return &companyClient{cc}
}

func (c *companyClient) CreateCompany(ctx context.Context, in *CreateCompanyRequest, opts ...grpc.CallOption) (*CreateCompanyResponse, error) {
	out := new(CreateCompanyResponse)
	err := c.cc.Invoke(ctx, "/companies.Company/CreateCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyClient) UpdateCompany(ctx context.Context, in *UpdateCompanyRequest, opts ...grpc.CallOption) (*UpdateCompanyResponse, error) {
	out := new(UpdateCompanyResponse)
	err := c.cc.Invoke(ctx, "/companies.Company/UpdateCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyClient) DeleteCompany(ctx context.Context, in *DeleteCompanyRequest, opts ...grpc.CallOption) (*DeleteCompanyResponse, error) {
	out := new(DeleteCompanyResponse)
	err := c.cc.Invoke(ctx, "/companies.Company/DeleteCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyClient) ReadCompanyByID(ctx context.Context, in *ReadCompanyByIdRequest, opts ...grpc.CallOption) (*ReadCompanyByIdResponse, error) {
	out := new(ReadCompanyByIdResponse)
	err := c.cc.Invoke(ctx, "/companies.Company/ReadCompanyByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CompanyServer is the server API for Company service.
// All implementations must embed UnimplementedCompanyServer
// for forward compatibility
type CompanyServer interface {
	CreateCompany(context.Context, *CreateCompanyRequest) (*CreateCompanyResponse, error)
	UpdateCompany(context.Context, *UpdateCompanyRequest) (*UpdateCompanyResponse, error)
	DeleteCompany(context.Context, *DeleteCompanyRequest) (*DeleteCompanyResponse, error)
	ReadCompanyByID(context.Context, *ReadCompanyByIdRequest) (*ReadCompanyByIdResponse, error)
	mustEmbedUnimplementedCompanyServer()
}

// UnimplementedCompanyServer must be embedded to have forward compatible implementations.
type UnimplementedCompanyServer struct {
}

func (UnimplementedCompanyServer) CreateCompany(context.Context, *CreateCompanyRequest) (*CreateCompanyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCompany not implemented")
}
func (UnimplementedCompanyServer) UpdateCompany(context.Context, *UpdateCompanyRequest) (*UpdateCompanyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCompany not implemented")
}
func (UnimplementedCompanyServer) DeleteCompany(context.Context, *DeleteCompanyRequest) (*DeleteCompanyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCompany not implemented")
}
func (UnimplementedCompanyServer) ReadCompanyByID(context.Context, *ReadCompanyByIdRequest) (*ReadCompanyByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadCompanyByID not implemented")
}
func (UnimplementedCompanyServer) mustEmbedUnimplementedCompanyServer() {}

// UnsafeCompanyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CompanyServer will
// result in compilation errors.
type UnsafeCompanyServer interface {
	mustEmbedUnimplementedCompanyServer()
}

func RegisterCompanyServer(s grpc.ServiceRegistrar, srv CompanyServer) {
	s.RegisterService(&Company_ServiceDesc, srv)
}

func _Company_CreateCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServer).CreateCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/companies.Company/CreateCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServer).CreateCompany(ctx, req.(*CreateCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Company_UpdateCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServer).UpdateCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/companies.Company/UpdateCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServer).UpdateCompany(ctx, req.(*UpdateCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Company_DeleteCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServer).DeleteCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/companies.Company/DeleteCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServer).DeleteCompany(ctx, req.(*DeleteCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Company_ReadCompanyByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadCompanyByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServer).ReadCompanyByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/companies.Company/ReadCompanyByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServer).ReadCompanyByID(ctx, req.(*ReadCompanyByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Company_ServiceDesc is the grpc.ServiceDesc for Company service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Company_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "companies.Company",
	HandlerType: (*CompanyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCompany",
			Handler:    _Company_CreateCompany_Handler,
		},
		{
			MethodName: "UpdateCompany",
			Handler:    _Company_UpdateCompany_Handler,
		},
		{
			MethodName: "DeleteCompany",
			Handler:    _Company_DeleteCompany_Handler,
		},
		{
			MethodName: "ReadCompanyByID",
			Handler:    _Company_ReadCompanyByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/companies.proto",
}
