package nacosservicej

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/opentracing-contrib/go-stdlib/nethttp"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	ddhttp "github.com/unionj-cloud/go-doudou/framework/http"
	"github.com/unionj-cloud/go-doudou/framework/registry"
)

type EchoClient struct {
	provider registry.IServiceProvider
	client   *resty.Client
	rootPath string
}

func (receiver *EchoClient) SetRootPath(rootPath string) {
	receiver.rootPath = rootPath
}

func (receiver *EchoClient) SetProvider(provider registry.IServiceProvider) {
	receiver.provider = provider
}

func (receiver *EchoClient) SetClient(client *resty.Client) {
	receiver.client = client
}
func (receiver *EchoClient) GetEchoString(ctx context.Context, _headers map[string]string,
	// required
	string string) (ret string, _resp *resty.Response, err error) {
	var _err error

	_req := receiver.client.R()
	_req.SetContext(ctx)
	if len(_headers) > 0 {
		_req.SetHeaders(_headers)
	}
	_req.SetPathParam("string", fmt.Sprintf("%v", string))

	_resp, _err = _req.Get("/echo/{string}")
	if _err != nil {
		err = errors.Wrap(_err, "")
		return
	}
	if _resp.IsError() {
		err = errors.New(_resp.String())
		return
	}
	ret = _resp.String()
	return
}

func NewEcho(opts ...ddhttp.DdClientOption) *EchoClient {
	defaultProvider := ddhttp.NewServiceProvider("NACOS_SERVICE_J")
	defaultClient := ddhttp.NewClient()

	svcClient := &EchoClient{
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
