package service

import (
	"context"
	"fmt"

	pb "github.com/unionj-cloud/helloworld/transport/grpc"

	"github.com/unionj-cloud/helloworld/config"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/unionj-cloud/helloworld/vo"
)

var _ pb.HelloworldRpcServer = (*HelloworldImpl)(nil)
var _ Helloworld = (*HelloworldImpl)(nil)

type HelloworldImpl struct {
	conf *config.Config
	pb.UnimplementedHelloworldRpcServer
}

func (receiver *HelloworldImpl) ByeRpc(ctx context.Context, request *pb.ByeRpcRequest) (*pb.ByeRpcResponse, error) {
	data, err := receiver.Bye(ctx, request.Greeting)
	if err != nil {
		return nil, err
	}
	return &pb.ByeRpcResponse{
		Data: data,
	}, nil
}

func (receiver *HelloworldImpl) GreetingRpc(ctx context.Context, request *pb.GreetingRpcRequest) (*pb.GreetingRpcResponse, error) {
	data, err := receiver.Greeting(ctx, request.Greeting)
	if err != nil {
		return nil, err
	}
	return &pb.GreetingRpcResponse{
		Data: data,
	}, nil
}

func (receiver *HelloworldImpl) Greeting(ctx context.Context, greeting string) (data string, err error) {
	return fmt.Sprintf("Hello %s", greeting), nil
}

func NewHelloworld(conf *config.Config) *HelloworldImpl {
	return &HelloworldImpl{
		conf: conf,
	}
}

func (receiver *HelloworldImpl) Bye(ctx context.Context, greeting string) (data string, err error) {
	return fmt.Sprintf("Bye %s", greeting), nil
}

func (receiver *HelloworldImpl) BiStream(ctx context.Context, stream vo.Order) (stream1 vo.Page, err error) {
	var _result struct {
		Stream1 vo.Page
	}
	_ = gofakeit.Struct(&_result)
	return _result.Stream1, nil
}
func (receiver *HelloworldImpl) ClientStream(ctx context.Context, stream vo.Order) (data vo.Page, err error) {
	var _result struct {
		Data vo.Page
	}
	_ = gofakeit.Struct(&_result)
	return _result.Data, nil
}
func (receiver *HelloworldImpl) ServerStream(ctx context.Context, payload vo.Order) (stream vo.Page, err error) {
	var _result struct {
		Stream vo.Page
	}
	_ = gofakeit.Struct(&_result)
	return _result.Stream, nil
}

func (receiver *HelloworldImpl) BiStreamRpc(server pb.HelloworldRpc_BiStreamRpcServer) error {
	//TODO implement me
	panic("implement me")
}
func (receiver *HelloworldImpl) ClientStreamRpc(server pb.HelloworldRpc_ClientStreamRpcServer) error {
	//TODO implement me
	panic("implement me")
}
func (receiver *HelloworldImpl) ServerStreamRpc(request *pb.Order, server pb.HelloworldRpc_ServerStreamRpcServer) error {
	//TODO implement me
	panic("implement me")
}
