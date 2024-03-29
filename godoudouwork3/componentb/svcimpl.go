/**
* Generated by go-doudou v2.1.5.
* You can edit it as your need.
 */
package service

import (
	"context"
	"godoudouwork1/componentb/config"
	"godoudouwork1/componentb/dto"

	"github.com/brianvoe/gofakeit/v6"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "godoudouwork1/componentb/transport/grpc"
)

var _ Componentb = (*ComponentbImpl)(nil)

var _ pb.ComponentbServiceServer = (*ComponentbImpl)(nil)

type ComponentbImpl struct {
	pb.UnimplementedComponentbServiceServer

	conf *config.Config
}

func (receiver *ComponentbImpl) PostUser(ctx context.Context, user dto.GddUser) (data int32, err error) {
	var _result struct {
		Data int32
	}
	_ = gofakeit.Struct(&_result)
	return _result.Data, nil
}
func (receiver *ComponentbImpl) GetUser_Id(ctx context.Context, id int32) (data dto.GddUser, err error) {
	return dto.GddUser{
		Id:    id,
		Name:  "bbbbbb",
		Phone: "",
		Dept:  "",
	}, nil
}
func (receiver *ComponentbImpl) PutUser(ctx context.Context, user dto.GddUser) (re error) {
	var _result struct {
	}
	_ = gofakeit.Struct(&_result)
	return nil
}
func (receiver *ComponentbImpl) DeleteUser_Id(ctx context.Context, id int32) (re error) {
	var _result struct {
	}
	_ = gofakeit.Struct(&_result)
	return nil
}
func (receiver *ComponentbImpl) GetUsers(ctx context.Context, parameter dto.Parameter) (data dto.Page, err error) {
	var _result struct {
		Data dto.Page
	}
	_ = gofakeit.Struct(&_result)
	return _result.Data, nil
}

func NewComponentb(conf *config.Config) *ComponentbImpl {
	return &ComponentbImpl{
		conf: conf,
	}
}

func (receiver *ComponentbImpl) PostUserRpc(ctx context.Context, request *pb.GddUser) (*pb.PostUserRpcResponse, error) {
	//TODO implement me
	panic("implement me")
}
func (receiver *ComponentbImpl) GetUserIdRpc(ctx context.Context, request *pb.GetUserIdRpcRequest) (*pb.GddUser, error) {
	//TODO implement me
	panic("implement me")
}
func (receiver *ComponentbImpl) PutUserRpc(ctx context.Context, request *pb.GddUser) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}
func (receiver *ComponentbImpl) DeleteUserIdRpc(ctx context.Context, request *pb.DeleteUserIdRpcRequest) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}
func (receiver *ComponentbImpl) GetUsersRpc(ctx context.Context, request *pb.Parameter) (*pb.Page, error) {
	//TODO implement me
	panic("implement me")
}
