/**
* Generated by go-doudou v2.0.6.
* Don't edit!
 */
package client

import (
	"context"
	"service-b/dto"

	"github.com/go-resty/resty/v2"
)

type Options struct {
	GzipReqBody bool
}

type IServiceBClient interface {
	GetDeptById(ctx context.Context, _headers map[string]string, deptId int, options Options) (_resp *resty.Response, dept dto.DeptDto, err error)
}