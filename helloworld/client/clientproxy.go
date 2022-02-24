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
)

type HelloworldClientProxy struct {
	client *HelloworldClient
	logger *logrus.Logger
	runner goresilience.Runner
}

func (receiver *HelloworldClientProxy) Greeting(ctx context.Context, greeting string) (data string, err error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_, data, err = receiver.client.Greeting(
			ctx,
			greeting,
		)
		if err != nil {
			return errors.Wrap(err, "call Greeting fail")
		}
		return nil
	}); _err != nil {
		// you can implement your fallback logic here
		if errors.Is(_err, rerrors.ErrCircuitOpen) {
			receiver.logger.Error(_err)
		}
		err = errors.Wrap(_err, "call Greeting fail")
	}
	return
}
func (receiver *HelloworldClientProxy) Bye(ctx context.Context, greeting string) (data string, err error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_, data, err = receiver.client.Bye(
			ctx,
			greeting,
		)
		if err != nil {
			return errors.Wrap(err, "call Bye fail")
		}
		return nil
	}); _err != nil {
		// you can implement your fallback logic here
		if errors.Is(_err, rerrors.ErrCircuitOpen) {
			receiver.logger.Error(_err)
		}
		err = errors.Wrap(_err, "call Bye fail")
	}
	return
}

type ProxyOption func(*HelloworldClientProxy)

func WithRunner(runner goresilience.Runner) ProxyOption {
	return func(proxy *HelloworldClientProxy) {
		proxy.runner = runner
	}
}

func WithLogger(logger *logrus.Logger) ProxyOption {
	return func(proxy *HelloworldClientProxy) {
		proxy.logger = logger
	}
}

func NewHelloworldClientProxy(client *HelloworldClient, opts ...ProxyOption) *HelloworldClientProxy {
	cp := &HelloworldClientProxy{
		client: client,
		logger: logrus.StandardLogger(),
	}

	for _, opt := range opts {
		opt(cp)
	}

	if cp.runner == nil {
		var mid []goresilience.Middleware
		mid = append(mid, metrics.NewMiddleware("github.com/unionj-cloud/helloworld_client", metrics.NewPrometheusRecorder(prometheus.DefaultRegisterer)))
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
