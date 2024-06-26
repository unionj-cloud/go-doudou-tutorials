// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ComponentaServiceClient is the client API for ComponentaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ComponentaServiceClient interface {
	// You can define your service methods as your need. Below is an example.
	// You can also add annotations here like @role(admin) to add meta data to routes for
	// implementing your own middlewares
	PostUserRpc(ctx context.Context, in *GddUser, opts ...grpc.CallOption) (*PostUserRpcResponse, error)
	GetUserIdRpc(ctx context.Context, in *GetUserIdRpcRequest, opts ...grpc.CallOption) (*GddUser, error)
	PutUserRpc(ctx context.Context, in *GddUser, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteUserIdRpc(ctx context.Context, in *DeleteUserIdRpcRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetUsersRpc(ctx context.Context, in *Parameter, opts ...grpc.CallOption) (*Page, error)
}

type componentaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewComponentaServiceClient(cc grpc.ClientConnInterface) ComponentaServiceClient {
	return &componentaServiceClient{cc}
}

func (c *componentaServiceClient) PostUserRpc(ctx context.Context, in *GddUser, opts ...grpc.CallOption) (*PostUserRpcResponse, error) {
	out := new(PostUserRpcResponse)
	err := c.cc.Invoke(ctx, "/componenta.ComponentaService/PostUserRpc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *componentaServiceClient) GetUserIdRpc(ctx context.Context, in *GetUserIdRpcRequest, opts ...grpc.CallOption) (*GddUser, error) {
	out := new(GddUser)
	err := c.cc.Invoke(ctx, "/componenta.ComponentaService/GetUserIdRpc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *componentaServiceClient) PutUserRpc(ctx context.Context, in *GddUser, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/componenta.ComponentaService/PutUserRpc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *componentaServiceClient) DeleteUserIdRpc(ctx context.Context, in *DeleteUserIdRpcRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/componenta.ComponentaService/DeleteUserIdRpc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *componentaServiceClient) GetUsersRpc(ctx context.Context, in *Parameter, opts ...grpc.CallOption) (*Page, error) {
	out := new(Page)
	err := c.cc.Invoke(ctx, "/componenta.ComponentaService/GetUsersRpc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ComponentaServiceServer is the server API for ComponentaService service.
// All implementations must embed UnimplementedComponentaServiceServer
// for forward compatibility
type ComponentaServiceServer interface {
	// You can define your service methods as your need. Below is an example.
	// You can also add annotations here like @role(admin) to add meta data to routes for
	// implementing your own middlewares
	PostUserRpc(context.Context, *GddUser) (*PostUserRpcResponse, error)
	GetUserIdRpc(context.Context, *GetUserIdRpcRequest) (*GddUser, error)
	PutUserRpc(context.Context, *GddUser) (*emptypb.Empty, error)
	DeleteUserIdRpc(context.Context, *DeleteUserIdRpcRequest) (*emptypb.Empty, error)
	GetUsersRpc(context.Context, *Parameter) (*Page, error)
	mustEmbedUnimplementedComponentaServiceServer()
}

// UnimplementedComponentaServiceServer must be embedded to have forward compatible implementations.
type UnimplementedComponentaServiceServer struct {
}

func (UnimplementedComponentaServiceServer) PostUserRpc(context.Context, *GddUser) (*PostUserRpcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostUserRpc not implemented")
}
func (UnimplementedComponentaServiceServer) GetUserIdRpc(context.Context, *GetUserIdRpcRequest) (*GddUser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserIdRpc not implemented")
}
func (UnimplementedComponentaServiceServer) PutUserRpc(context.Context, *GddUser) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutUserRpc not implemented")
}
func (UnimplementedComponentaServiceServer) DeleteUserIdRpc(context.Context, *DeleteUserIdRpcRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserIdRpc not implemented")
}
func (UnimplementedComponentaServiceServer) GetUsersRpc(context.Context, *Parameter) (*Page, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsersRpc not implemented")
}
func (UnimplementedComponentaServiceServer) mustEmbedUnimplementedComponentaServiceServer() {}

// UnsafeComponentaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ComponentaServiceServer will
// result in compilation errors.
type UnsafeComponentaServiceServer interface {
	mustEmbedUnimplementedComponentaServiceServer()
}

func RegisterComponentaServiceServer(s grpc.ServiceRegistrar, srv ComponentaServiceServer) {
	s.RegisterService(&ComponentaService_ServiceDesc, srv)
}

func _ComponentaService_PostUserRpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GddUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComponentaServiceServer).PostUserRpc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/componenta.ComponentaService/PostUserRpc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComponentaServiceServer).PostUserRpc(ctx, req.(*GddUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComponentaService_GetUserIdRpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserIdRpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComponentaServiceServer).GetUserIdRpc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/componenta.ComponentaService/GetUserIdRpc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComponentaServiceServer).GetUserIdRpc(ctx, req.(*GetUserIdRpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComponentaService_PutUserRpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GddUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComponentaServiceServer).PutUserRpc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/componenta.ComponentaService/PutUserRpc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComponentaServiceServer).PutUserRpc(ctx, req.(*GddUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComponentaService_DeleteUserIdRpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserIdRpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComponentaServiceServer).DeleteUserIdRpc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/componenta.ComponentaService/DeleteUserIdRpc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComponentaServiceServer).DeleteUserIdRpc(ctx, req.(*DeleteUserIdRpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComponentaService_GetUsersRpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Parameter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComponentaServiceServer).GetUsersRpc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/componenta.ComponentaService/GetUsersRpc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComponentaServiceServer).GetUsersRpc(ctx, req.(*Parameter))
	}
	return interceptor(ctx, in, info, handler)
}

// ComponentaService_ServiceDesc is the grpc.ServiceDesc for ComponentaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ComponentaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "componenta.ComponentaService",
	HandlerType: (*ComponentaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostUserRpc",
			Handler:    _ComponentaService_PostUserRpc_Handler,
		},
		{
			MethodName: "GetUserIdRpc",
			Handler:    _ComponentaService_GetUserIdRpc_Handler,
		},
		{
			MethodName: "PutUserRpc",
			Handler:    _ComponentaService_PutUserRpc_Handler,
		},
		{
			MethodName: "DeleteUserIdRpc",
			Handler:    _ComponentaService_DeleteUserIdRpc_Handler,
		},
		{
			MethodName: "GetUsersRpc",
			Handler:    _ComponentaService_GetUsersRpc_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transport/grpc/componenta.proto",
}
