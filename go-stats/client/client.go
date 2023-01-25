/**
* Generated by go-doudou v2.0.4.
* Don't edit!
 */
package client

import (
	"context"
	"encoding/json"
	"fmt"
	"go-doudou-tutorials/go-stats/vo"
	"io"
	"net/http"
	"net/url"

	"github.com/go-resty/resty/v2"
	"github.com/klauspost/compress/gzip"
	"github.com/opentracing-contrib/go-stdlib/nethttp"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/unionj-cloud/go-doudou/v2/framework/registry"
	"github.com/unionj-cloud/go-doudou/v2/framework/restclient"
)

type GoStatsClient struct {
	provider registry.IServiceProvider
	client   *resty.Client
	rootPath string
}

func (receiver *GoStatsClient) SetRootPath(rootPath string) {
	receiver.rootPath = rootPath
}

func (receiver *GoStatsClient) SetProvider(provider registry.IServiceProvider) {
	receiver.provider = provider
}

func (receiver *GoStatsClient) SetClient(client *resty.Client) {
	receiver.client = client
}
func (receiver *GoStatsClient) LargestRemainder(ctx context.Context, _headers map[string]string, payload vo.PercentageReqVo, options Options) (_resp *resty.Response, data []vo.PercentageRespVo, err error) {
	var _err error
	_urlValues := url.Values{}
	_req := receiver.client.R()
	if len(_headers) > 0 {
		_req.SetHeaders(_headers)
	}
	_req.SetContext(ctx)
	if options.GzipReqBody {
		pr, pw := io.Pipe()
		go func() {
			gw := gzip.NewWriter(pw)
			_err = json.NewEncoder(gw).Encode(payload)
			if _err != nil {
				err = errors.Wrap(_err, "error")
				return
			}
			_err = gw.Close()
			if _err != nil {
				err = errors.Wrap(_err, "error")
				return
			}
			defer pw.CloseWithError(err)
		}()
		_req.SetHeader("Content-Type", "application/json")
		_req.SetHeader("Content-Encoding", "gzip")
		_req.SetBody(pr)
	} else {
		_req.SetBody(payload)
	}
	_path := "/largest/remainder"
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
		Data []vo.PercentageRespVo `json:"data"`
	}
	if _err = json.Unmarshal(_resp.Body(), &_result); _err != nil {
		err = errors.Wrap(_err, "error")
		return
	}
	return _resp, _result.Data, nil
}
func (receiver *GoStatsClient) GetShelves_ShelfBooks_Book(ctx context.Context, _headers map[string]string, shelf int, book string, options Options) (_resp *resty.Response, data string, err error) {
	var _err error
	_urlValues := url.Values{}
	_req := receiver.client.R()
	if len(_headers) > 0 {
		_req.SetHeaders(_headers)
	}
	_req.SetContext(ctx)
	_req.SetPathParam("shelf", fmt.Sprintf("%v", shelf))
	_req.SetPathParam("book", fmt.Sprintf("%v", book))
	_path := "/shelves/{shelf}/books/{book}"
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
		Data string `json:"data"`
	}
	if _err = json.Unmarshal(_resp.Body(), &_result); _err != nil {
		err = errors.Wrap(_err, "error")
		return
	}
	return _resp, _result.Data, nil
}

func NewGoStatsClient(opts ...restclient.RestClientOption) *GoStatsClient {
	defaultProvider := restclient.NewServiceProvider("GOSTATS")
	defaultClient := restclient.NewClient()

	svcClient := &GoStatsClient{
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
