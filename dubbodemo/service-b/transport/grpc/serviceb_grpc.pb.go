// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpc

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

// ServiceBServiceClient is the client API for ServiceBService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceBServiceClient interface {
	GetDeptByIdRpc(ctx context.Context, in *GetDeptByIdRpcRequest, opts ...grpc.CallOption) (*DeptDto, error)
}

type serviceBServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceBServiceClient(cc grpc.ClientConnInterface) ServiceBServiceClient {
	return &serviceBServiceClient{cc}
}

func (c *serviceBServiceClient) GetDeptByIdRpc(ctx context.Context, in *GetDeptByIdRpcRequest, opts ...grpc.CallOption) (*DeptDto, error) {
	out := new(DeptDto)
	err := c.cc.Invoke(ctx, "/service_b.ServiceBService/GetDeptByIdRpc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceBServiceServer is the server API for ServiceBService service.
// All implementations must embed UnimplementedServiceBServiceServer
// for forward compatibility
type ServiceBServiceServer interface {
	GetDeptByIdRpc(context.Context, *GetDeptByIdRpcRequest) (*DeptDto, error)
	mustEmbedUnimplementedServiceBServiceServer()
}

// UnimplementedServiceBServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceBServiceServer struct {
}

func (UnimplementedServiceBServiceServer) GetDeptByIdRpc(context.Context, *GetDeptByIdRpcRequest) (*DeptDto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeptByIdRpc not implemented")
}
func (UnimplementedServiceBServiceServer) mustEmbedUnimplementedServiceBServiceServer() {}

// UnsafeServiceBServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceBServiceServer will
// result in compilation errors.
type UnsafeServiceBServiceServer interface {
	mustEmbedUnimplementedServiceBServiceServer()
}

func RegisterServiceBServiceServer(s grpc.ServiceRegistrar, srv ServiceBServiceServer) {
	s.RegisterService(&ServiceBService_ServiceDesc, srv)
}

func _ServiceBService_GetDeptByIdRpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeptByIdRpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceBServiceServer).GetDeptByIdRpc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service_b.ServiceBService/GetDeptByIdRpc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceBServiceServer).GetDeptByIdRpc(ctx, req.(*GetDeptByIdRpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ServiceBService_ServiceDesc is the grpc.ServiceDesc for ServiceBService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServiceBService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service_b.ServiceBService",
	HandlerType: (*ServiceBServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDeptByIdRpc",
			Handler:    _ServiceBService_GetDeptByIdRpc_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transport/grpc/serviceb.proto",
}