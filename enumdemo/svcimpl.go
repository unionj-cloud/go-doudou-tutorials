package service

import (
	"abc2/config"
	"abc2/vo"
	"context"
)

type EnumDemoImpl struct {
	conf *config.Config
}

func (receiver *EnumDemoImpl) GetKeyboard(ctx context.Context, layout vo.KeyboardLayout) (data string, err error) {
	return layout.StringGetter(), nil
}
func (receiver *EnumDemoImpl) GetKeyboard2(ctx context.Context, layout *vo.KeyboardLayout) (data string, err error) {
	return layout.StringGetter(), nil
}
func (receiver *EnumDemoImpl) GetKeyboards(ctx context.Context, layout []vo.KeyboardLayout) (data []string, err error) {
	for _, item := range layout {
		data = append(data, item.StringGetter())
	}
	return data, nil
}
func (receiver *EnumDemoImpl) GetKeyboards2(ctx context.Context, layout *[]vo.KeyboardLayout) (data []string, err error) {
	for _, item := range *layout {
		data = append(data, item.StringGetter())
	}
	return data, nil
}
func (receiver *EnumDemoImpl) GetKeyboards5(ctx context.Context, layout ...vo.KeyboardLayout) (data []string, err error) {
	for _, item := range layout {
		data = append(data, item.StringGetter())
	}
	return data, nil
}

func NewEnumDemo(conf *config.Config) EnumDemo {
	return &EnumDemoImpl{
		conf,
	}
}

func (receiver *EnumDemoImpl) Keyboard(ctx context.Context, keyboard vo.Keyboard) (data string, err error) {
	return keyboard.Layout.StringGetter(), nil
}
