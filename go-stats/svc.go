/**
* Generated by go-doudou v2.0.3.
* You can edit it as your need.
 */
package service

import (
	"context"
	"go-doudou-tutorials/go-stats/vo"
)

//go:generate go-doudou svc http -c
//go:generate go-doudou svc grpc

// GoStats is a demo gRPC service developed by go-doudou
type GoStats interface {
	// LargestRemainder implements Largest Remainder Method https://en.wikipedia.org/wiki/Largest_remainder_method
	LargestRemainder(ctx context.Context, payload vo.PercentageReqVo) (data []vo.PercentageRespVo, err error)
}
