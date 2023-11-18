/**
* Generated by go-doudou v2.1.8.
* You can edit it as your need.
 */
package service

import (
	"context"
	"testhttp2grpc/config"
	pb "testhttp2grpc/transport/grpc"

	"github.com/brianvoe/gofakeit/v6"

	"google.golang.org/protobuf/types/known/emptypb"
)

var _ pb.Testhttp2GrpcServiceServer = (*Testhttp2GrpcImpl)(nil)

type Testhttp2GrpcImpl struct {
	pb.UnimplementedTesthttp2GrpcServiceServer
	conf *config.Config
}

func (receiver *Testhttp2GrpcImpl) PutUserRpc(ctx context.Context, request *pb.GddUser) (*emptypb.Empty, error) {
	return nil, nil
}
func (receiver *Testhttp2GrpcImpl) GetUsersRpc(ctx context.Context, request *pb.Parameter) (*pb.Page, error) {
	var _result struct {
		Reply *pb.Page
	}
	_ = gofakeit.Struct(&_result)
	return _result.Reply, nil
}
func (receiver *Testhttp2GrpcImpl) GetDictRpc(ctx context.Context, request *emptypb.Empty) (*pb.DictResp, error) {
	var _result struct {
		Reply *pb.DictResp
	}
	_ = gofakeit.Struct(&_result)
	return _result.Reply, nil
}

func NewTesthttp2Grpc(conf *config.Config) *Testhttp2GrpcImpl {
	return &Testhttp2GrpcImpl{
		conf: conf,
	}
}