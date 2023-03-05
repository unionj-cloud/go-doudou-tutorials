/**
* Generated by go-doudou v2.0.5.
* You can edit it as your need.
 */
package client

import (
	"context"
	"testsvc/dto"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/rs/zerolog"
	"github.com/slok/goresilience"
	"github.com/slok/goresilience/circuitbreaker"
	rerrors "github.com/slok/goresilience/errors"
	"github.com/slok/goresilience/metrics"
	"github.com/slok/goresilience/retry"
	"github.com/slok/goresilience/timeout"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/zlogger"
)

type TestsvcClientProxy struct {
	client *TestsvcClient
	logger zerolog.Logger
	runner goresilience.Runner
}

func (receiver *TestsvcClientProxy) GetBookNotFoundException(ctx context.Context, _headers map[string]string, options Options) (_resp *resty.Response, re error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_resp, re = receiver.client.GetBookNotFoundException(
			ctx,
			_headers,
			options,
		)
		if re != nil {
			return errors.Wrap(re, "call GetBookNotFoundException fail")
		}
		return nil
	}); _err != nil {
		// you can implement your fallback logic here
		if errors.Is(_err, rerrors.ErrCircuitOpen) {
			receiver.logger.Error().Err(_err).Msg("")
		}
		re = errors.Wrap(_err, "call GetBookNotFoundException fail")
	}
	return
}
func (receiver *TestsvcClientProxy) GetConversionFailedException(ctx context.Context, _headers map[string]string, options Options) (_resp *resty.Response, re error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_resp, re = receiver.client.GetConversionFailedException(
			ctx,
			_headers,
			options,
		)
		if re != nil {
			return errors.Wrap(re, "call GetConversionFailedException fail")
		}
		return nil
	}); _err != nil {
		// you can implement your fallback logic here
		if errors.Is(_err, rerrors.ErrCircuitOpen) {
			receiver.logger.Error().Err(_err).Msg("")
		}
		re = errors.Wrap(_err, "call GetConversionFailedException fail")
	}
	return
}
func (receiver *TestsvcClientProxy) GetBookPage(ctx context.Context, _headers map[string]string, name string, author string, page dto.Page, options Options) (_resp *resty.Response, re error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_resp, re = receiver.client.GetBookPage(
			ctx,
			_headers,
			name,
			author,
			page,
			options,
		)
		if re != nil {
			return errors.Wrap(re, "call GetBookPage fail")
		}
		return nil
	}); _err != nil {
		// you can implement your fallback logic here
		if errors.Is(_err, rerrors.ErrCircuitOpen) {
			receiver.logger.Error().Err(_err).Msg("")
		}
		re = errors.Wrap(_err, "call GetBookPage fail")
	}
	return
}

type ProxyOption func(*TestsvcClientProxy)

func WithRunner(runner goresilience.Runner) ProxyOption {
	return func(proxy *TestsvcClientProxy) {
		proxy.runner = runner
	}
}

func WithLogger(logger zerolog.Logger) ProxyOption {
	return func(proxy *TestsvcClientProxy) {
		proxy.logger = logger
	}
}

func NewTestsvcClientProxy(client *TestsvcClient, opts ...ProxyOption) *TestsvcClientProxy {
	cp := &TestsvcClientProxy{
		client: client,
		logger: zlogger.Logger,
	}

	for _, opt := range opts {
		opt(cp)
	}

	if cp.runner == nil {
		var mid []goresilience.Middleware
		mid = append(mid, metrics.NewMiddleware("testsvc_client", metrics.NewPrometheusRecorder(prometheus.DefaultRegisterer)))
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