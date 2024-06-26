/**
* Generated by go-doudou v2.1.5.
* You can edit it as your need.
*/
package service

import (
	"context"
	"godoudouwork1/componentb/dto"
)

//go:generate go-doudou svc http -c
//go:generate go-doudou svc grpc

type Componentb interface {
	// You can define your service methods as your need. Below is an example.
	// You can also add annotations here like @role(admin) to add meta data to routes for 
	// implementing your own middlewares
	PostUser(ctx context.Context, user dto.GddUser) (data int32, err error)
	GetUser_Id(ctx context.Context, id int32) (data dto.GddUser, err error)
	PutUser(ctx context.Context, user dto.GddUser) error
	DeleteUser_Id(ctx context.Context, id int32) error
	GetUsers(ctx context.Context, parameter dto.Parameter) (data dto.Page, err error)
}
