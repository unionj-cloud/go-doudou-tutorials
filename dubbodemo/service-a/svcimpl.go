/**
* Generated by go-doudou v2.0.6.
* You can edit it as your need.
 */
package service

import (
	"context"
	"service-a/config"
	protobuf "service-a/deps"
	"service-a/dto"
	"service-b/client"
	pb "service-b/transport/grpc"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/errorx"
)

var _ ServiceA = (*ServiceAImpl)(nil)

type ServiceAImpl struct {
	conf          *config.Config
	bClient       client.IServiceBClient
	grpcClient    pb.ServiceBServiceClient
	greeterClient protobuf.GreeterClient
}

func (receiver *ServiceAImpl) GetUserById(ctx context.Context, userId int) (user dto.UserDto, err error) {
	var _result struct {
		User dto.UserDto
	}
	_ = gofakeit.Struct(&_result)
	_result.User.Id = userId
	_, deptDto, err := receiver.bClient.GetDeptById(ctx, nil, 1, client.Options{})
	if err != nil {
		errorx.Panic(err.Error())
	}
	_result.User.Dept = deptDto.Name
	return _result.User, nil
}

func NewServiceA(conf *config.Config, bClient client.IServiceBClient, grpcClient pb.ServiceBServiceClient, greeterClient protobuf.GreeterClient) *ServiceAImpl {
	return &ServiceAImpl{
		conf:          conf,
		bClient:       bClient,
		grpcClient:    grpcClient,
		greeterClient: greeterClient,
	}
}

func (receiver *ServiceAImpl) GetRpcUserById(ctx context.Context, userId int) (user dto.UserDto, err error) {
	var _result struct {
		User dto.UserDto
	}
	_ = gofakeit.Struct(&_result)
	_result.User.Id = userId
	deptDto, err := receiver.grpcClient.GetDeptByIdRpc(ctx, &pb.GetDeptByIdRpcRequest{
		DeptId: 2,
	})
	if err != nil {
		errorx.Panic(err.Error())
	}
	_result.User.Dept = deptDto.Name
	return _result.User, nil
}

func (receiver *ServiceAImpl) GetRpcSayHello(ctx context.Context, name string) (reply string, err error) {
	hr, err := receiver.greeterClient.SayHello(ctx, &protobuf.HelloRequest{
		Name: name,
	})
	if err != nil {
		errorx.Panic(err.Error())
	}
	return hr.Message, nil
}