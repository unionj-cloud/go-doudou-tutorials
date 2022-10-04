package service

import (
	"context"
	"github.com/unionj-cloud/go-doudou-tutorials/grpcdemo/client/vo"
)

//go:generate go-doudou svc http --handler --doc

type EnumDemo interface {
	GetKeyboard(ctx context.Context, layout vo.KeyboardLayout) (data string, err error)
	GetKeyboard2(ctx context.Context, layout *vo.KeyboardLayout) (data string, err error)
	GetKeyboards(ctx context.Context, layout []vo.KeyboardLayout) (data []string, err error)
	GetKeyboards2(ctx context.Context, layout *[]vo.KeyboardLayout) (data []string, err error)
	GetKeyboards5(ctx context.Context, layout ...vo.KeyboardLayout) (data []string, err error)
	Keyboard(ctx context.Context, keyboard vo.Keyboard) (data string, err error)

	// don't define parameters like []*vo.KeyboardLayout
	//GetKeyboards3(ctx context.Context, layout []*vo.KeyboardLayout) error
	// don't define parameters like *[]*vo.KeyboardLayout
	//GetKeyboards4(ctx context.Context, layout *[]*vo.KeyboardLayout) error

	// Greeting 问候接口
	Greeting(ctx context.Context,
	// 入参
		greeting string) (
	// 出参
		data string,
	// 错误信息
		err error)
	// Greeting 问候接口
	Greeting1(ctx context.Context,
	// 入参
		greeting string) (
	// 出参
		data string,
	// 错误信息
		err error)
}
