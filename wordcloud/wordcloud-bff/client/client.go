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
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-bff/vo"
	ddhttp "github.com/unionj-cloud/go-doudou/framework/http"
	"github.com/unionj-cloud/go-doudou/framework/registry"
	v3 "github.com/unionj-cloud/go-doudou/toolkit/openapi/v3"
)

type WordcloudBffClient struct {
	provider registry.IServiceProvider
	client   *resty.Client
}

func (receiver *WordcloudBffClient) SetProvider(provider registry.IServiceProvider) {
	receiver.provider = provider
}

func (receiver *WordcloudBffClient) SetClient(client *resty.Client) {
	receiver.client = client
}
func (receiver *WordcloudBffClient) Upload(ctx context.Context, file v3.FileModel, lang string, top *int) (_resp *resty.Response, data vo.UploadResult, err error) {
	var _err error
	_urlValues := url.Values{}
	_req := receiver.client.R()
	_req.SetContext(ctx)
	_req.SetFileReader("file", file.Filename, file.Reader)
	_urlValues.Set("lang", fmt.Sprintf("%v", lang))
	if top != nil {
		_urlValues.Set("top", fmt.Sprintf("%v", *top))
	}
	_path := "/upload"
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
		Data vo.UploadResult `json:"data"`
	}
	if _err = json.Unmarshal(_resp.Body(), &_result); _err != nil {
		err = errors.Wrap(_err, "error")
		return
	}
	return _resp, _result.Data, nil
}

func NewWordcloudBffClient(opts ...ddhttp.DdClientOption) *WordcloudBffClient {
	defaultProvider := ddhttp.NewServiceProvider("WORDCLOUDBFF")
	defaultClient := ddhttp.NewClient()

	svcClient := &WordcloudBffClient{
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
