/**
* Generated by go-doudou v2.0.5.
* Don't edit!
 */
package client

import (
	"context"
	"gormdemo/dto"

	"github.com/go-resty/resty/v2"
)

type Options struct {
	GzipReqBody bool
}

type IGormdemoClient interface {
	PageUsers(ctx context.Context, _headers map[string]string, query dto.PageQuery, options Options) (_resp *resty.Response, data dto.PageRet, err error)
}
