/**
* Generated by go-doudou v2.0.1.
* You can edit it as your need.
 */
package service

import (
	"context"
	"grpc3/config"
	pb "grpc3/transport/grpc"
)

var _ pb.Grpc3ServiceServer = (*Grpc3Impl)(nil)

type Grpc3Impl struct {
	pb.UnimplementedGrpc3ServiceServer
	conf *config.Config
}

func (receiver *Grpc3Impl) SegRpc(ctx context.Context, request *pb.SegPayload) (*pb.SegResult, error) {
	//TODO implement me
	panic("implement me")
}

func NewGrpc3(conf *config.Config) *Grpc3Impl {
	return &Grpc3Impl{
		conf: conf,
	}
}
