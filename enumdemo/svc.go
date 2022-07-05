package service

import (
	"abc2/vo"
	"context"
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
}
