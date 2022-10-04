package service

import (
	"context"
	"github.com/unionj-cloud/go-doudou-tutorials/grpcdemo/server/vo"
)

//go:generate go-doudou svc http --handler --doc -c
//go:generate go-doudou svc grpc

// Helloworld 单体服务
type Helloworld interface {
	// Greeting 问候接口
	Greeting(ctx context.Context,
	// 入参
		greeting string) (
	// 出参
		data string,
	// 错误信息
		err error)

	// Bye 再见接口
	Bye(ctx context.Context,
	// 入参
		greeting string) (
	// 出参
		data string,
	// 错误信息
		err error)

	BiStream(ctx context.Context, stream vo.Order) (stream1 vo.Page, err error)
	ClientStream(ctx context.Context, stream vo.Order) (data vo.Page, err error)
	ServerStream(ctx context.Context, payload vo.Order) (stream vo.Page, err error)
}
