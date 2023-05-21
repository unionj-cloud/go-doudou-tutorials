/**
* Generated by go-doudou v2.0.6.
* Don't edit!
 */
package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"service-a/dto"

	"github.com/go-resty/resty/v2"
	"github.com/opentracing-contrib/go-stdlib/nethttp"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/unionj-cloud/go-doudou/v2/framework/registry"
	"github.com/unionj-cloud/go-doudou/v2/framework/restclient"
)

type ServiceAClient struct {
	provider registry.IServiceProvider
	client   *resty.Client
	rootPath string
}

func (receiver *ServiceAClient) SetRootPath(rootPath string) {
	receiver.rootPath = rootPath
}

func (receiver *ServiceAClient) SetProvider(provider registry.IServiceProvider) {
	receiver.provider = provider
}

func (receiver *ServiceAClient) SetClient(client *resty.Client) {
	receiver.client = client
}
func (receiver *ServiceAClient) GetUserById(ctx context.Context, _headers map[string]string, userId int, options Options) (_resp *resty.Response, user dto.UserDto, err error) {
	var _err error
	_urlValues := url.Values{}
	_req := receiver.client.R()
	if len(_headers) > 0 {
		_req.SetHeaders(_headers)
	}
	_req.SetContext(ctx)
	_urlValues.Set("userId", fmt.Sprintf("%v", userId))
	_path := "/user/by/id"
	_req.SetQueryParamsFromValues(_urlValues)
	_resp, _err = _req.Get(_path)
	if _err != nil {
		err = errors.Wrap(_err, "error")
		return
	}
	if _resp.IsError() {
		err = errors.New(_resp.String())
		return
	}
	var _result struct {
		User dto.UserDto `json:"user"`
	}
	if _err = json.Unmarshal(_resp.Body(), &_result); _err != nil {
		err = errors.Wrap(_err, "error")
		return
	}
	return _resp, _result.User, nil
}
func (receiver *ServiceAClient) GetRpcUserById(ctx context.Context, _headers map[string]string, userId int, options Options) (_resp *resty.Response, user dto.UserDto, err error) {
	var _err error
	_urlValues := url.Values{}
	_req := receiver.client.R()
	if len(_headers) > 0 {
		_req.SetHeaders(_headers)
	}
	_req.SetContext(ctx)
	_urlValues.Set("userId", fmt.Sprintf("%v", userId))
	_path := "/rpc/user/by/id"
	_req.SetQueryParamsFromValues(_urlValues)
	_resp, _err = _req.Get(_path)
	if _err != nil {
		err = errors.Wrap(_err, "error")
		return
	}
	if _resp.IsError() {
		err = errors.New(_resp.String())
		return
	}
	var _result struct {
		User dto.UserDto `json:"user"`
	}
	if _err = json.Unmarshal(_resp.Body(), &_result); _err != nil {
		err = errors.Wrap(_err, "error")
		return
	}
	return _resp, _result.User, nil
}

func NewServiceAClient(opts ...restclient.RestClientOption) *ServiceAClient {
	defaultProvider := restclient.NewServiceProvider("SERVICEA")
	defaultClient := restclient.NewClient()

	svcClient := &ServiceAClient{
		provider: defaultProvider,
		client:   defaultClient,
	}

	for _, opt := range opts {
		opt(svcClient)
	}

	svcClient.client.OnBeforeRequest(func(_ *resty.Client, request *resty.Request) error {
		request.URL = svcClient.provider.SelectServer() + svcClient.rootPath + request.URL
		return nil
	})

	svcClient.client.SetPreRequestHook(func(_ *resty.Client, request *http.Request) error {
		traceReq, _ := nethttp.TraceRequest(opentracing.GlobalTracer(), request,
			nethttp.OperationName(fmt.Sprintf("HTTP %s: %s", request.Method, request.URL.Path)))
		*request = *traceReq
		return nil
	})

	svcClient.client.OnAfterResponse(func(_ *resty.Client, response *resty.Response) error {
		nethttp.TracerFromRequest(response.Request.RawRequest).Finish()
		return nil
	})

	return svcClient
}