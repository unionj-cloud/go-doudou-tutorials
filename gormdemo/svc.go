/**
* Generated by go-doudou v2.0.5.
* You can edit it as your need.
 */
package service

import (
	"context"
	"gormdemo/dto"
)

//go:generate go-doudou svc http -c
//go:generate go-doudou svc grpc

type Gormdemo interface {
	PostTUser(ctx context.Context, tUser dto.TUser) (data int32, err error)
	GetTUser_Id(ctx context.Context, id int32) (data dto.TUser, err error)
	PutTUser(ctx context.Context, tUser dto.TUser) error
	DeleteTUser_Id(ctx context.Context, id int32) error
	GetTUsers(ctx context.Context, parameter dto.Parameter) (data dto.Page, err error)
}