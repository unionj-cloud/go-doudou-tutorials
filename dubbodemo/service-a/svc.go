/**
* Generated by go-doudou v2.0.6.
* You can edit it as your need.
 */
package service

import (
	"context"
	"service-a/dto"
)

//go:generate go-doudou svc http -c
//go:generate go-doudou svc grpc

type ServiceA interface {
	GetUserById(ctx context.Context, userId int) (user dto.UserDto, err error)
	GetRpcUserById(ctx context.Context, userId int) (user dto.UserDto, err error)
	GetRpcSayHello(ctx context.Context, name string) (reply string, err error)
}
