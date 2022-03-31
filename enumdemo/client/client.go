package client

import (
	"abc2/vo"
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

type EnumDemoClient struct {
	provider registry.IServiceProvider
	client   *resty.Client
	rootPath string
}

func (receiver *EnumDemoClient) SetRootPath(rootPath string) {
	receiver.rootPath = rootPath
}

func (receiver *EnumDemoClient) SetProvider(provider registry.IServiceProvider) {
	receiver.provider = provider
}

func (receiver *EnumDemoClient) SetClient(client *resty.Client) {
	receiver.client = client
}
func (receiver *EnumDemoClient) GetKeyboard(ctx context.Context, _headers map[string]string, layout vo.KeyboardLayout) (_resp *resty.Response, data string, err error) {
	var _err error
	_urlValues := url.Values{}
	_req := receiver.client.R()
	if len(_headers) > 0 {
		_req.SetHeaders(_headers)
	}
	_req.SetContext(ctx)
	_urlValues.Set("layout", layout.StringGetter())
	_path := "/keyboard"
	_resp, _err = _req.SetQueryParamsFromValues(_urlValues).
		Get(_path)
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
func (receiver *EnumDemoClient) GetKeyboard2(ctx context.Context, _headers map[string]string, layout *vo.KeyboardLayout) (_resp *resty.Response, data string, err error) {
	var _err error
	_urlValues := url.Values{}
	_req := receiver.client.R()
	if len(_headers) > 0 {
		_req.SetHeaders(_headers)
	}
	_req.SetContext(ctx)
	if layout != nil {
		_urlValues.Set("layout", layout.StringGetter())
	}
	_path := "/keyboard/2"
	_resp, _err = _req.SetQueryParamsFromValues(_urlValues).
		Get(_path)
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
func (receiver *EnumDemoClient) GetKeyboards(ctx context.Context, _headers map[string]string, layout []vo.KeyboardLayout) (_resp *resty.Response, data []string, err error) {
	var _err error
	_urlValues := url.Values{}
	_req := receiver.client.R()
	if len(_headers) > 0 {
		_req.SetHeaders(_headers)
	}
	_req.SetContext(ctx)
	if len(layout) == 0 {
		err = errors.New("size of parameter layout should be greater than zero")
		return
	}
	for _, _item := range layout {
		_urlValues.Add("layout", _item.StringGetter())
	}
	_path := "/keyboards"
	_resp, _err = _req.SetQueryParamsFromValues(_urlValues).
		Get(_path)
	if _err != nil {
		err = errors.Wrap(_err, "error")
		return
	}
	if _resp.IsError() {
		err = errors.New(_resp.String())
		return
	}
	var _result struct {
		Data []string `json:"data"`
	}
	if _err = json.Unmarshal(_resp.Body(), &_result); _err != nil {
		err = errors.Wrap(_err, "error")
		return
	}
	return _resp, _result.Data, nil
}
func (receiver *EnumDemoClient) GetKeyboards2(ctx context.Context, _headers map[string]string, layout *[]vo.KeyboardLayout) (_resp *resty.Response, data []string, err error) {
	var _err error
	_urlValues := url.Values{}
	_req := receiver.client.R()
	if len(_headers) > 0 {
		_req.SetHeaders(_headers)
	}
	_req.SetContext(ctx)
	if layout != nil {
		for _, _item := range *layout {
			_urlValues.Add("layout", _item.StringGetter())
		}
	}
	_path := "/keyboards/2"
	_resp, _err = _req.SetQueryParamsFromValues(_urlValues).
		Get(_path)
	if _err != nil {
		err = errors.Wrap(_err, "error")
		return
	}
	if _resp.IsError() {
		err = errors.New(_resp.String())
		return
	}
	var _result struct {
		Data []string `json:"data"`
	}
	if _err = json.Unmarshal(_resp.Body(), &_result); _err != nil {
		err = errors.Wrap(_err, "error")
		return
	}
	return _resp, _result.Data, nil
}
func (receiver *EnumDemoClient) GetKeyboards5(ctx context.Context, _headers map[string]string, layout ...vo.KeyboardLayout) (_resp *resty.Response, data []string, err error) {
	var _err error
	_urlValues := url.Values{}
	_req := receiver.client.R()
	if len(_headers) > 0 {
		_req.SetHeaders(_headers)
	}
	_req.SetContext(ctx)
	if layout != nil {
		for _, _item := range layout {
			_urlValues.Add("layout", _item.StringGetter())
		}
	}
	_path := "/keyboards/5"
	_resp, _err = _req.SetQueryParamsFromValues(_urlValues).
		Get(_path)
	if _err != nil {
		err = errors.Wrap(_err, "error")
		return
	}
	if _resp.IsError() {
		err = errors.New(_resp.String())
		return
	}
	var _result struct {
		Data []string `json:"data"`
	}
	if _err = json.Unmarshal(_resp.Body(), &_result); _err != nil {
		err = errors.Wrap(_err, "error")
		return
	}
	return _resp, _result.Data, nil
}
func (receiver *EnumDemoClient) Keyboard(ctx context.Context, _headers map[string]string, keyboard vo.Keyboard) (_resp *resty.Response, data string, err error) {
	var _err error
	_urlValues := url.Values{}
	_req := receiver.client.R()
	if len(_headers) > 0 {
		_req.SetHeaders(_headers)
	}
	_req.SetContext(ctx)
	_req.SetBody(keyboard)
	_path := "/keyboard"
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

func NewEnumDemoClient(opts ...ddhttp.DdClientOption) *EnumDemoClient {
	defaultProvider := ddhttp.NewServiceProvider("ENUMDEMO")
	defaultClient := ddhttp.NewClient()

	svcClient := &EnumDemoClient{
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
