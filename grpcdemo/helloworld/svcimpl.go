package service

import (
	"context"
	"fmt"

	pb "github.com/unionj-cloud/helloworld/transport/grpc"

	"github.com/unionj-cloud/helloworld/config"
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
