// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: service.proto

package carwash

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ServicesService_CreateService_FullMethodName     = "/car_wash.ServicesService/CreateService"
	ServicesService_GetService_FullMethodName        = "/car_wash.ServicesService/GetService"
	ServicesService_UpdateService_FullMethodName     = "/car_wash.ServicesService/UpdateService"
	ServicesService_DeleteService_FullMethodName     = "/car_wash.ServicesService/DeleteService"
	ServicesService_ListServices_FullMethodName      = "/car_wash.ServicesService/ListServices"
	ServicesService_SearchServices_FullMethodName    = "/car_wash.ServicesService/SearchServices"
	ServicesService_GetPopularService_FullMethodName = "/car_wash.ServicesService/GetPopularService"
)

// ServicesServiceClient is the client API for ServicesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServicesServiceClient interface {
	CreateService(ctx context.Context, in *CreateServiceRequest, opts ...grpc.CallOption) (*CreateServiceResponse, error)
	GetService(ctx context.Context, in *GetServiceRequest, opts ...grpc.CallOption) (*GetServiceResponse, error)
	UpdateService(ctx context.Context, in *UpdateServiceRequest, opts ...grpc.CallOption) (*UpdateServiceResponse, error)
	DeleteService(ctx context.Context, in *DeleteServiceRequest, opts ...grpc.CallOption) (*DeleteServiceResponse, error)
	ListServices(ctx context.Context, in *ListServicesRequest, opts ...grpc.CallOption) (*ListServicesResponse, error)
	SearchServices(ctx context.Context, in *SearchServicesRequest, opts ...grpc.CallOption) (*SearchServicesResponse, error)
	GetPopularService(ctx context.Context, in *PopularServiceRequest, opts ...grpc.CallOption) (*PopularServicesResponse, error)
}

type servicesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewServicesServiceClient(cc grpc.ClientConnInterface) ServicesServiceClient {
	return &servicesServiceClient{cc}
}

func (c *servicesServiceClient) CreateService(ctx context.Context, in *CreateServiceRequest, opts ...grpc.CallOption) (*CreateServiceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateServiceResponse)
	err := c.cc.Invoke(ctx, ServicesService_CreateService_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesServiceClient) GetService(ctx context.Context, in *GetServiceRequest, opts ...grpc.CallOption) (*GetServiceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetServiceResponse)
	err := c.cc.Invoke(ctx, ServicesService_GetService_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesServiceClient) UpdateService(ctx context.Context, in *UpdateServiceRequest, opts ...grpc.CallOption) (*UpdateServiceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateServiceResponse)
	err := c.cc.Invoke(ctx, ServicesService_UpdateService_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesServiceClient) DeleteService(ctx context.Context, in *DeleteServiceRequest, opts ...grpc.CallOption) (*DeleteServiceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteServiceResponse)
	err := c.cc.Invoke(ctx, ServicesService_DeleteService_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesServiceClient) ListServices(ctx context.Context, in *ListServicesRequest, opts ...grpc.CallOption) (*ListServicesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListServicesResponse)
	err := c.cc.Invoke(ctx, ServicesService_ListServices_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesServiceClient) SearchServices(ctx context.Context, in *SearchServicesRequest, opts ...grpc.CallOption) (*SearchServicesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchServicesResponse)
	err := c.cc.Invoke(ctx, ServicesService_SearchServices_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesServiceClient) GetPopularService(ctx context.Context, in *PopularServiceRequest, opts ...grpc.CallOption) (*PopularServicesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PopularServicesResponse)
	err := c.cc.Invoke(ctx, ServicesService_GetPopularService_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServicesServiceServer is the server API for ServicesService service.
// All implementations must embed UnimplementedServicesServiceServer
// for forward compatibility.
type ServicesServiceServer interface {
	CreateService(context.Context, *CreateServiceRequest) (*CreateServiceResponse, error)
	GetService(context.Context, *GetServiceRequest) (*GetServiceResponse, error)
	UpdateService(context.Context, *UpdateServiceRequest) (*UpdateServiceResponse, error)
	DeleteService(context.Context, *DeleteServiceRequest) (*DeleteServiceResponse, error)
	ListServices(context.Context, *ListServicesRequest) (*ListServicesResponse, error)
	SearchServices(context.Context, *SearchServicesRequest) (*SearchServicesResponse, error)
	GetPopularService(context.Context, *PopularServiceRequest) (*PopularServicesResponse, error)
	mustEmbedUnimplementedServicesServiceServer()
}

// UnimplementedServicesServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedServicesServiceServer struct{}

func (UnimplementedServicesServiceServer) CreateService(context.Context, *CreateServiceRequest) (*CreateServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateService not implemented")
}
func (UnimplementedServicesServiceServer) GetService(context.Context, *GetServiceRequest) (*GetServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetService not implemented")
}
func (UnimplementedServicesServiceServer) UpdateService(context.Context, *UpdateServiceRequest) (*UpdateServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateService not implemented")
}
func (UnimplementedServicesServiceServer) DeleteService(context.Context, *DeleteServiceRequest) (*DeleteServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteService not implemented")
}
func (UnimplementedServicesServiceServer) ListServices(context.Context, *ListServicesRequest) (*ListServicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListServices not implemented")
}
func (UnimplementedServicesServiceServer) SearchServices(context.Context, *SearchServicesRequest) (*SearchServicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchServices not implemented")
}
func (UnimplementedServicesServiceServer) GetPopularService(context.Context, *PopularServiceRequest) (*PopularServicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPopularService not implemented")
}
func (UnimplementedServicesServiceServer) mustEmbedUnimplementedServicesServiceServer() {}
func (UnimplementedServicesServiceServer) testEmbeddedByValue()                         {}

// UnsafeServicesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServicesServiceServer will
// result in compilation errors.
type UnsafeServicesServiceServer interface {
	mustEmbedUnimplementedServicesServiceServer()
}

func RegisterServicesServiceServer(s grpc.ServiceRegistrar, srv ServicesServiceServer) {
	// If the following call pancis, it indicates UnimplementedServicesServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ServicesService_ServiceDesc, srv)
}

func _ServicesService_CreateService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServiceServer).CreateService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServicesService_CreateService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServiceServer).CreateService(ctx, req.(*CreateServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServicesService_GetService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServiceServer).GetService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServicesService_GetService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServiceServer).GetService(ctx, req.(*GetServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServicesService_UpdateService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServiceServer).UpdateService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServicesService_UpdateService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServiceServer).UpdateService(ctx, req.(*UpdateServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServicesService_DeleteService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServiceServer).DeleteService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServicesService_DeleteService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServiceServer).DeleteService(ctx, req.(*DeleteServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServicesService_ListServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListServicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServiceServer).ListServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServicesService_ListServices_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServiceServer).ListServices(ctx, req.(*ListServicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServicesService_SearchServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchServicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServiceServer).SearchServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServicesService_SearchServices_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServiceServer).SearchServices(ctx, req.(*SearchServicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServicesService_GetPopularService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PopularServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServiceServer).GetPopularService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServicesService_GetPopularService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServiceServer).GetPopularService(ctx, req.(*PopularServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ServicesService_ServiceDesc is the grpc.ServiceDesc for ServicesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServicesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "car_wash.ServicesService",
	HandlerType: (*ServicesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateService",
			Handler:    _ServicesService_CreateService_Handler,
		},
		{
			MethodName: "GetService",
			Handler:    _ServicesService_GetService_Handler,
		},
		{
			MethodName: "UpdateService",
			Handler:    _ServicesService_UpdateService_Handler,
		},
		{
			MethodName: "DeleteService",
			Handler:    _ServicesService_DeleteService_Handler,
		},
		{
			MethodName: "ListServices",
			Handler:    _ServicesService_ListServices_Handler,
		},
		{
			MethodName: "SearchServices",
			Handler:    _ServicesService_SearchServices_Handler,
		},
		{
			MethodName: "GetPopularService",
			Handler:    _ServicesService_GetPopularService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
