/**
* Generated by go-doudou v2.0.5.
* You can edit it as your need.
 */
package service

import (
	"context"
	"gormdemo/config"
	"gormdemo/dto"
	"gormdemo/query"

	"github.com/brianvoe/gofakeit/v6"

	pb "gormdemo/transport/grpc"
)

var _ Gormdemo = (*GormdemoImpl)(nil)

var _ pb.GormdemoServiceServer = (*GormdemoImpl)(nil)

type GormdemoImpl struct {
	pb.UnimplementedGormdemoServiceServer

	conf *config.Config
}

func (receiver *GormdemoImpl) PageUsers(ctx context.Context, q dto.PageQuery) (data dto.PageRet, err error) {
	var _result struct {
		Data dto.PageRet
	}
	_ = gofakeit.Struct(&_result)
	query.TUser.WithContext(ctx).First()
	return _result.Data, nil
}

func NewGormdemo(conf *config.Config) *GormdemoImpl {
	return &GormdemoImpl{
		conf: conf,
	}
}

func (receiver *GormdemoImpl) GetTUser_Id(ctx context.Context, id int32) (data dto.TUser, err error) {
	var _result struct {
		Data dto.TUser
	}
	_ = gofakeit.Struct(&_result)
	return _result.Data, nil
}

func (receiver *GormdemoImpl) PageUsersRpc(ctx context.Context, request *pb.PageQuery) (*pb.PageRet, error) {
	//TODO implement me
	panic("implement me")
}
func (receiver *GormdemoImpl) GetTUserIdRpc(ctx context.Context, request *pb.GetTUserIdRpcRequest) (*pb.TUser, error) {
	//TODO implement me
	panic("implement me")
}

func (receiver *GormdemoImpl) PostTUser(ctx context.Context, tUser dto.TUser) (data int32, err error) {
	var _result struct {
		Data int32
	}
	_ = gofakeit.Struct(&_result)
	return _result.Data, nil
}
func (receiver *GormdemoImpl) PutTUser(ctx context.Context, tUser dto.TUser) (re error) {
	var _result struct {
	}
	_ = gofakeit.Struct(&_result)
	return nil
}
func (receiver *GormdemoImpl) DeleteTUser_Id(ctx context.Context, id int32) (re error) {
	var _result struct {
	}
	_ = gofakeit.Struct(&_result)
	return nil
}
func (receiver *GormdemoImpl) GetTUsers(ctx context.Context, parameter dto.Parameter) (data dto.Page, err error) {
	var _result struct {
		Data dto.Page
	}
	_ = gofakeit.Struct(&_result)
	return _result.Data, nil
}

func (receiver *GormdemoImpl) PostGenTUser(ctx context.Context, tUser dto.TUser) (data int32, err error) {
	var _result struct {
		Data int32
	}
	_ = gofakeit.Struct(&_result)
	return _result.Data, nil
}
func (receiver *GormdemoImpl) GetGenTUser_Id(ctx context.Context, id int32) (data dto.TUser, err error) {
	var _result struct {
		Data dto.TUser
	}
	_ = gofakeit.Struct(&_result)
	return _result.Data, nil
}
func (receiver *GormdemoImpl) PutGenTUser(ctx context.Context, tUser dto.TUser) (re error) {
	var _result struct {
	}
	_ = gofakeit.Struct(&_result)
	return nil
}
func (receiver *GormdemoImpl) DeleteGenTUser_Id(ctx context.Context, id int32) (re error) {
	var _result struct {
	}
	_ = gofakeit.Struct(&_result)
	return nil
}
func (receiver *GormdemoImpl) GetGenTUsers(ctx context.Context, parameter dto.Parameter) (data dto.Page, err error) {
	var _result struct {
		Data dto.Page
	}
	_ = gofakeit.Struct(&_result)
	return _result.Data, nil
}
