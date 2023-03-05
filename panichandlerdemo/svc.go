/**
* Generated by go-doudou v2.0.4.
* You can edit it as your need.
 */
package service

import (
	"context"
	"testsvc/dto"
)

//go:generate go-doudou svc http -c
//go:generate go-doudou svc grpc

type Testsvc interface {
	// You can define your service methods as your need. Below is an example.
	// You can also add annotations here like @role(admin) to add metadata to routes for
	// implementing your own middlewares
	GetBookNotFoundException(ctx context.Context) error
	GetConversionFailedException(ctx context.Context) error
	GetBookPage(ctx context.Context, name string, author string, page dto.Page) error
}
