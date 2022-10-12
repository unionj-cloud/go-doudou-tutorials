package service

import (
	"context"
	"statsvc/config"
	"statsvc/internal/nacosservicej"

	pb "statsvc/transport/grpc"
)

var _ Statsvc = (*StatsvcImpl)(nil)

var _ pb.StatsvcServiceServer = (*StatsvcImpl)(nil)

type StatsvcImpl struct {
	pb.UnimplementedStatsvcServiceServer

	conf       *config.Config
	echoClient *nacosservicej.EchoClient
}

func (receiver *StatsvcImpl) Add(ctx context.Context, x int, y int) (data int, err error) {
	return x + y, nil
}

func NewStatsvc(conf *config.Config, echoClient *nacosservicej.EchoClient) *StatsvcImpl {
	return &StatsvcImpl{
		conf:       conf,
		echoClient: echoClient,
	}
}

func (receiver *StatsvcImpl) GetEcho(ctx context.Context, s string) (data string, err error) {
	data, _, err = receiver.echoClient.GetEchoString(ctx, nil, s)
	return
}

func (receiver *StatsvcImpl) AddRpc(ctx context.Context, request *pb.AddRpcRequest) (*pb.AddRpcResponse, error) {
	data, err := receiver.Add(ctx, int(request.X), int(request.Y))
	if err != nil {
		return nil, err
	}
	return &pb.AddRpcResponse{
		Data: int32(data),
	}, nil
}
func (receiver *StatsvcImpl) GetEchoRpc(ctx context.Context, request *pb.GetEchoRpcRequest) (*pb.GetEchoRpcResponse, error) {
	//TODO implement me
	panic("implement me")
}
