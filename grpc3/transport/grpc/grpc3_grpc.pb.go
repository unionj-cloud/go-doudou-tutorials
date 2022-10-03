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

// Grpc3ServiceClient is the client API for Grpc3Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Grpc3ServiceClient interface {
	// You can define your service methods as your need. Below is an example.
	// You can also add annotations here like @role(admin) to add meta data to routes for
	// implementing your own middlewares
	PageUsersRpc(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*PageRet, error)
}

type grpc3ServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGrpc3ServiceClient(cc grpc.ClientConnInterface) Grpc3ServiceClient {
	return &grpc3ServiceClient{cc}
}

func (c *grpc3ServiceClient) PageUsersRpc(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*PageRet, error) {
	out := new(PageRet)
	err := c.cc.Invoke(ctx, "/grpc_3.Grpc3Service/PageUsersRpc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Grpc3ServiceServer is the server API for Grpc3Service service.
// All implementations must embed UnimplementedGrpc3ServiceServer
// for forward compatibility
type Grpc3ServiceServer interface {
	// You can define your service methods as your need. Below is an example.
	// You can also add annotations here like @role(admin) to add meta data to routes for
	// implementing your own middlewares
	PageUsersRpc(context.Context, *PageQuery) (*PageRet, error)
	mustEmbedUnimplementedGrpc3ServiceServer()
}

// UnimplementedGrpc3ServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGrpc3ServiceServer struct {
}

func (UnimplementedGrpc3ServiceServer) PageUsersRpc(context.Context, *PageQuery) (*PageRet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PageUsersRpc not implemented")
}
func (UnimplementedGrpc3ServiceServer) mustEmbedUnimplementedGrpc3ServiceServer() {}

// UnsafeGrpc3ServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Grpc3ServiceServer will
// result in compilation errors.
type UnsafeGrpc3ServiceServer interface {
	mustEmbedUnimplementedGrpc3ServiceServer()
}

func RegisterGrpc3ServiceServer(s grpc.ServiceRegistrar, srv Grpc3ServiceServer) {
	s.RegisterService(&Grpc3Service_ServiceDesc, srv)
}

func _Grpc3Service_PageUsersRpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PageQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Grpc3ServiceServer).PageUsersRpc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_3.Grpc3Service/PageUsersRpc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Grpc3ServiceServer).PageUsersRpc(ctx, req.(*PageQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// Grpc3Service_ServiceDesc is the grpc.ServiceDesc for Grpc3Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Grpc3Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc_3.Grpc3Service",
	HandlerType: (*Grpc3ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PageUsersRpc",
			Handler:    _Grpc3Service_PageUsersRpc_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transport/grpc/grpc3.proto",
}