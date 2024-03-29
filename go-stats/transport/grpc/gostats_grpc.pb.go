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

// GoStatsServiceClient is the client API for GoStatsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GoStatsServiceClient interface {
	// LargestRemainder implements Largest Remainder Method https://en.wikipedia.org/wiki/Largest_remainder_method
	LargestRemainderRpc(ctx context.Context, in *PercentageReqVo, opts ...grpc.CallOption) (*LargestRemainderRpcResponse, error)
	// /shelves/:shelf/books/:book
	// shelves_shelf_books_book
	GetShelvesShelfBooksBookRpc(ctx context.Context, in *GetShelvesShelfBooksBookRpcRequest, opts ...grpc.CallOption) (*GetShelvesShelfBooksBookRpcResponse, error)
}

type goStatsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGoStatsServiceClient(cc grpc.ClientConnInterface) GoStatsServiceClient {
	return &goStatsServiceClient{cc}
}

func (c *goStatsServiceClient) LargestRemainderRpc(ctx context.Context, in *PercentageReqVo, opts ...grpc.CallOption) (*LargestRemainderRpcResponse, error) {
	out := new(LargestRemainderRpcResponse)
	err := c.cc.Invoke(ctx, "/go_stats.GoStatsService/LargestRemainderRpc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goStatsServiceClient) GetShelvesShelfBooksBookRpc(ctx context.Context, in *GetShelvesShelfBooksBookRpcRequest, opts ...grpc.CallOption) (*GetShelvesShelfBooksBookRpcResponse, error) {
	out := new(GetShelvesShelfBooksBookRpcResponse)
	err := c.cc.Invoke(ctx, "/go_stats.GoStatsService/GetShelvesShelfBooksBookRpc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoStatsServiceServer is the server API for GoStatsService service.
// All implementations must embed UnimplementedGoStatsServiceServer
// for forward compatibility
type GoStatsServiceServer interface {
	// LargestRemainder implements Largest Remainder Method https://en.wikipedia.org/wiki/Largest_remainder_method
	LargestRemainderRpc(context.Context, *PercentageReqVo) (*LargestRemainderRpcResponse, error)
	// /shelves/:shelf/books/:book
	// shelves_shelf_books_book
	GetShelvesShelfBooksBookRpc(context.Context, *GetShelvesShelfBooksBookRpcRequest) (*GetShelvesShelfBooksBookRpcResponse, error)
	mustEmbedUnimplementedGoStatsServiceServer()
}

// UnimplementedGoStatsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGoStatsServiceServer struct {
}

func (UnimplementedGoStatsServiceServer) LargestRemainderRpc(context.Context, *PercentageReqVo) (*LargestRemainderRpcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LargestRemainderRpc not implemented")
}
func (UnimplementedGoStatsServiceServer) GetShelvesShelfBooksBookRpc(context.Context, *GetShelvesShelfBooksBookRpcRequest) (*GetShelvesShelfBooksBookRpcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShelvesShelfBooksBookRpc not implemented")
}
func (UnimplementedGoStatsServiceServer) mustEmbedUnimplementedGoStatsServiceServer() {}

// UnsafeGoStatsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GoStatsServiceServer will
// result in compilation errors.
type UnsafeGoStatsServiceServer interface {
	mustEmbedUnimplementedGoStatsServiceServer()
}

func RegisterGoStatsServiceServer(s grpc.ServiceRegistrar, srv GoStatsServiceServer) {
	s.RegisterService(&GoStatsService_ServiceDesc, srv)
}

func _GoStatsService_LargestRemainderRpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PercentageReqVo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoStatsServiceServer).LargestRemainderRpc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go_stats.GoStatsService/LargestRemainderRpc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoStatsServiceServer).LargestRemainderRpc(ctx, req.(*PercentageReqVo))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoStatsService_GetShelvesShelfBooksBookRpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShelvesShelfBooksBookRpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoStatsServiceServer).GetShelvesShelfBooksBookRpc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go_stats.GoStatsService/GetShelvesShelfBooksBookRpc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoStatsServiceServer).GetShelvesShelfBooksBookRpc(ctx, req.(*GetShelvesShelfBooksBookRpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GoStatsService_ServiceDesc is the grpc.ServiceDesc for GoStatsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GoStatsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "go_stats.GoStatsService",
	HandlerType: (*GoStatsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LargestRemainderRpc",
			Handler:    _GoStatsService_LargestRemainderRpc_Handler,
		},
		{
			MethodName: "GetShelvesShelfBooksBookRpc",
			Handler:    _GoStatsService_GetShelvesShelfBooksBookRpc_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transport/grpc/gostats.proto",
}
