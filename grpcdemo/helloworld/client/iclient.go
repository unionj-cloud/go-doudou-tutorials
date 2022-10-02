/**
* Generated by go-doudou v1.2.4.
* Don't edit!
 */
package client

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/unionj-cloud/helloworld/vo"
)

type IHelloworldClient interface {
	Greeting(ctx context.Context, _headers map[string]string, greeting string) (_resp *resty.Response, data string, err error)
	Bye(ctx context.Context, _headers map[string]string, greeting string) (_resp *resty.Response, data string, err error)
	BiStream(ctx context.Context, _headers map[string]string, stream vo.Order) (_resp *resty.Response, stream1 vo.Page, err error)
	ClientStream(ctx context.Context, _headers map[string]string, stream vo.Order) (_resp *resty.Response, data vo.Page, err error)
	ServerStream(ctx context.Context, _headers map[string]string, payload vo.Order) (_resp *resty.Response, stream vo.Page, err error)
}
