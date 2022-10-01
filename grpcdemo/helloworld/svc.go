package service

import (
	"context"
)

//go:generate go-doudou svc http --handler --doc -c

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
}
