package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/go-resty/resty/v2"
	"github.com/opentracing-contrib/go-stdlib/nethttp"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	ddhttp "github.com/unionj-cloud/go-doudou/framework/http"
	"github.com/unionj-cloud/go-doudou/framework/registry"
)

type HelloworldClient struct {
	provider registry.IServiceProvider
	client   *resty.Client
}

func (receiver *HelloworldClient) SetProvider(provider registry.IServiceProvider) {
	receiver.provider = provider
}

func (receiver *HelloworldClient) SetClient(client *resty.Client) {
	receiver.client = client
}
func (receiver *HelloworldClient) Greeting(ctx context.Context, greeting string) (_resp *resty.Response, data string, err error) {
	var _err error
	_urlValues := url.Values{}
	_req := receiver.client.R()
	_req.SetContext(ctx)
	_urlValues.Set("greeting", fmt.Sprintf("%v", greeting))
	_path := "/greeting"
	if _req.Body != nil {
		_req.SetQueryParamsFromValues(_urlValues)
	} else {
		_req.SetFormDataFromValues(_urlValues)
	}
	_resp, _err = _req.Post(_path)
	if _err != nil {
		err = errors.Wrap(_err, "error")
		return
	}
	if _resp.IsError() {
		err = errors.New(_resp.String())
		return
	}
	var _result struct {
		Data string `json:"data"`
	}
	if _err = json.Unmarshal(_resp.Body(), &_result); _err != nil {
		err = errors.Wrap(_err, "error")
		return
	}
	return _resp, _result.Data, nil
}
func (receiver *HelloworldClient) Bye(ctx context.Context, greeting string) (_resp *resty.Response, data string, err error) {
	var _err error
	_urlValues := url.Values{}
	_req := receiver.client.R()
	_req.SetContext(ctx)
	_urlValues.Set("greeting", fmt.Sprintf("%v", greeting))
	_path := "/bye"
	if _req.Body != nil {
		_req.SetQueryParamsFromValues(_urlValues)
	} else {
		_req.SetFormDataFromValues(_urlValues)
	}
	_resp, _err = _req.Post(_path)
	if _err != nil {
		err = errors.Wrap(_err, "error")
		return
	}
	if _resp.IsError() {
		err = errors.New(_resp.String())
		return
	}
	var _result struct {
		Data string `json:"data"`
	}
	if _err = json.Unmarshal(_resp.Body(), &_result); _err != nil {
		err = errors.Wrap(_err, "error")
		return
	}
	return _resp, _result.Data, nil
}

func NewHelloworldClient(opts ...ddhttp.DdClientOption) *HelloworldClient {
	defaultProvider := ddhttp.NewServiceProvider("HELLOWORLD")
	defaultClient := ddhttp.NewClient()

	svcClient := &HelloworldClient{
		provider: defaultProvider,
		client:   defaultClient,
	}

	for _, opt := range opts {
		opt(svcClient)
	}

	svcClient.client.OnBeforeRequest(func(_ *resty.Client, request *resty.Request) error {
		request.URL = svcClient.provider.SelectServer() + request.URL
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
