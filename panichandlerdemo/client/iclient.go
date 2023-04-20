/**
* Generated by go-doudou v2.0.5.
* Don't edit!
 */
package client

import (
	"context"
	"testsvc/dto"

	"github.com/go-resty/resty/v2"
)

type Options struct {
	GzipReqBody bool
}

type ITestsvcClient interface {
	GetBookNotFoundException(ctx context.Context, _headers map[string]string, options Options) (_resp *resty.Response, re error)
	GetConversionFailedException(ctx context.Context, _headers map[string]string, options Options) (_resp *resty.Response, re error)
	GetBookPage(ctx context.Context, _headers map[string]string, name string, author string, page dto.Page, options Options) (_resp *resty.Response, re error)
	PostBookPage(ctx context.Context, _headers map[string]string, name string, author string, options Options) (_resp *resty.Response, re error)
}
