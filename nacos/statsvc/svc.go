package service

import (
	"context"
)

//go:generate go-doudou svc http --doc
//go:generate go-doudou svc grpc

type Statsvc interface {
	Add(ctx context.Context, x, y int) (data int, err error)
	GetEcho(ctx context.Context, s string) (data string, err error)
}
