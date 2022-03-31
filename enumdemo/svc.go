package service

import (
	"abc2/vo"
	"context"
)

type EnumDemo interface {
	GetKeyboard(ctx context.Context, layout vo.KeyboardLayout) (data string, err error)
	GetKeyboard2(ctx context.Context, layout *vo.KeyboardLayout) (data string, err error)
	GetKeyboards(ctx context.Context, layout []vo.KeyboardLayout) (data []string, err error)
	GetKeyboards2(ctx context.Context, layout *[]vo.KeyboardLayout) (data []string, err error)
	GetKeyboards5(ctx context.Context, layout ...vo.KeyboardLayout) (data []string, err error)
	Keyboard(ctx context.Context, keyboard vo.Keyboard) (data string, err error)

	// Not support fully, don't define parameters like []*vo.KeyboardLayout
	//GetKeyboards3(ctx context.Context, layout []*vo.KeyboardLayout) error
	// Not support fully, don't define parameters like *[]*vo.KeyboardLayout
	//GetKeyboards4(ctx context.Context, layout *[]*vo.KeyboardLayout) error
}
