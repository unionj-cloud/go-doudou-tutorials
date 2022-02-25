package client

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
	"github.com/slok/goresilience"
	"github.com/slok/goresilience/circuitbreaker"
	rerrors "github.com/slok/goresilience/errors"
	"github.com/slok/goresilience/metrics"
	"github.com/slok/goresilience/retry"
	"github.com/slok/goresilience/timeout"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-bff/vo"
	v3 "github.com/unionj-cloud/go-doudou/toolkit/openapi/v3"
)

type WordcloudBffClientProxy struct {
	client *WordcloudBffClient
	logger *logrus.Logger
	runner goresilience.Runner
}

func (receiver *WordcloudBffClientProxy) Upload(ctx context.Context, file v3.FileModel, lang string, top *int) (data vo.UploadResult, err error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_, data, err = receiver.client.Upload(
			ctx,
			file,
			lang,
			top,
		)
		if err != nil {
			return errors.Wrap(err, "call Upload fail")
		}
		return nil
	}); _err != nil {
		// you can implement your fallback logic here
		if errors.Is(_err, rerrors.ErrCircuitOpen) {
			receiver.logger.Error(_err)
		}
		err = errors.Wrap(_err, "call Upload fail")
	}
	return
}

type ProxyOption func(*WordcloudBffClientProxy)

func WithRunner(runner goresilience.Runner) ProxyOption {
	return func(proxy *WordcloudBffClientProxy) {
		proxy.runner = runner
	}
}

func WithLogger(logger *logrus.Logger) ProxyOption {
	return func(proxy *WordcloudBffClientProxy) {
		proxy.logger = logger
	}
}

func NewWordcloudBffClientProxy(client *WordcloudBffClient, opts ...ProxyOption) *WordcloudBffClientProxy {
	cp := &WordcloudBffClientProxy{
		client: client,
		logger: logrus.StandardLogger(),
	}

	for _, opt := range opts {
		opt(cp)
	}

	if cp.runner == nil {
		var mid []goresilience.Middleware
		mid = append(mid, metrics.NewMiddleware("github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-bff_client", metrics.NewPrometheusRecorder(prometheus.DefaultRegisterer)))
		mid = append(mid, circuitbreaker.NewMiddleware(circuitbreaker.Config{
			ErrorPercentThresholdToOpen:        50,
			MinimumRequestToOpen:               6,
			SuccessfulRequiredOnHalfOpen:       1,
			WaitDurationInOpenState:            5 * time.Second,
			MetricsSlidingWindowBucketQuantity: 10,
			MetricsBucketDuration:              1 * time.Second,
		}),
			timeout.NewMiddleware(timeout.Config{
				Timeout: 3 * time.Minute,
			}),
			retry.NewMiddleware(retry.Config{
				Times: 3,
			}))

		cp.runner = goresilience.RunnerChain(mid...)
	}

	return cp
}
