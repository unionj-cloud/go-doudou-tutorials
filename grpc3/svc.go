/**
* Generated by go-doudou v1.2.4.
* You can edit it as your need.
 */
package service

import (
	"context"
	"grpc3/vo"
)

//go:generate go-doudou svc http --handler -c --doc
//go:generate go-doudou svc grpc

type Grpc3 interface {
	Seg(ctx context.Context, payload vo.SegPayload) (rs vo.SegResult, err error)
}
