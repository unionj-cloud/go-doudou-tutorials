package client

import (
	"context"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
	"github.com/slok/goresilience"
	"github.com/slok/goresilience/circuitbreaker"
	rerrors "github.com/slok/goresilience/errors"
	"github.com/slok/goresilience/metrics"
	"github.com/slok/goresilience/retry"
	"github.com/slok/goresilience/timeout"
)

type StatsvcClientProxy struct {
	client *StatsvcClient
	logger *logrus.Logger
	runner goresilience.Runner
}

func (receiver *StatsvcClientProxy) Add(ctx context.Context, _headers map[string]string, x int, y int) (_resp *resty.Response, data int, err error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_resp, data, err = receiver.client.Add(
			ctx,
			_headers,
			x,
			y,
		)
		if err != nil {
			return errors.Wrap(err, "call Add fail")
		}
		return nil
	}); _err != nil {
		// you can implement your fallback logic here
		if errors.Is(_err, rerrors.ErrCircuitOpen) {
			receiver.logger.Error(_err)
		}
		err = errors.Wrap(_err, "call Add fail")
	}
	return
}

type ProxyOption func(*StatsvcClientProxy)

func WithRunner(runner goresilience.Runner) ProxyOption {
	return func(proxy *StatsvcClientProxy) {
		proxy.runner = runner
	}
}

func WithLogger(logger *logrus.Logger) ProxyOption {
	return func(proxy *StatsvcClientProxy) {
		proxy.logger = logger
	}
}

func NewStatsvcClientProxy(client *StatsvcClient, opts ...ProxyOption) *StatsvcClientProxy {
	cp := &StatsvcClientProxy{
		client: client,
		logger: logrus.StandardLogger(),
	}

	for _, opt := range opts {
		opt(cp)
	}

	if cp.runner == nil {
		var mid []goresilience.Middleware
		mid = append(mid, metrics.NewMiddleware("statsvc_client", metrics.NewPrometheusRecorder(prometheus.DefaultRegisterer)))
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

func (receiver *StatsvcClientProxy) GetEcho(ctx context.Context, _headers map[string]string, s string) (_resp *resty.Response, data string, err error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_resp, data, err = receiver.client.GetEcho(
			ctx,
			_headers,
			s,
		)
		if err != nil {
			return errors.Wrap(err, "call GetEcho fail")
		}
		return nil
	}); _err != nil {
		// you can implement your fallback logic here
		if errors.Is(_err, rerrors.ErrCircuitOpen) {
			receiver.logger.Error(_err)
		}
		err = errors.Wrap(_err, "call GetEcho fail")
	}
	return
}
