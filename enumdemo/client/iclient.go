package client

import (
	"abc2/vo"
	"context"

	"github.com/go-resty/resty/v2"
)

type IEnumDemoClient interface {
	GetKeyboard(ctx context.Context, _headers map[string]string, layout vo.KeyboardLayout) (_resp *resty.Response, data string, err error)
	GetKeyboard2(ctx context.Context, _headers map[string]string, layout *vo.KeyboardLayout) (_resp *resty.Response, data string, err error)
	GetKeyboards(ctx context.Context, _headers map[string]string, layout []vo.KeyboardLayout) (_resp *resty.Response, data []string, err error)
	GetKeyboards2(ctx context.Context, _headers map[string]string, layout *[]vo.KeyboardLayout) (_resp *resty.Response, data []string, err error)
	GetKeyboards5(ctx context.Context, _headers map[string]string, layout ...vo.KeyboardLayout) (_resp *resty.Response, data []string, err error)
	Keyboard(ctx context.Context, _headers map[string]string, keyboard vo.Keyboard) (_resp *resty.Response, data string, err error)
}
