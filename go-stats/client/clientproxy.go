/**
* Generated by go-doudou v2.0.4.
* You can edit it as your need.
 */
package client

import (
	"context"
	"go-doudou-tutorials/go-stats/vo"
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

type GoStatsClientProxy struct {
	client *GoStatsClient
	logger zerolog.Logger
	runner goresilience.Runner
}

func (receiver *GoStatsClientProxy) LargestRemainder(ctx context.Context, _headers map[string]string, payload vo.PercentageReqVo, options Options) (_resp *resty.Response, data []vo.PercentageRespVo, err error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_resp, data, err = receiver.client.LargestRemainder(
			ctx,
			_headers,
			payload,
			options,
		)
		if err != nil {
			return errors.Wrap(err, "call LargestRemainder fail")
		}
		return nil
	}); _err != nil {
		// you can implement your fallback logic here
		if errors.Is(_err, rerrors.ErrCircuitOpen) {
			receiver.logger.Error().Err(_err).Msg("")
		}
		err = errors.Wrap(_err, "call LargestRemainder fail")
	}
	return
}
func (receiver *GoStatsClientProxy) GetShelves_ShelfBooks_Book(ctx context.Context, _headers map[string]string, shelf int, book string, options Options) (_resp *resty.Response, data string, err error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_resp, data, err = receiver.client.GetShelves_ShelfBooks_Book(
			ctx,
			_headers,
			shelf,
			book,
			options,
		)
		if err != nil {
			return errors.Wrap(err, "call GetShelves_ShelfBooks_Book fail")
		}
		return nil
	}); _err != nil {
		// you can implement your fallback logic here
		if errors.Is(_err, rerrors.ErrCircuitOpen) {
			receiver.logger.Error().Err(_err).Msg("")
		}
		err = errors.Wrap(_err, "call GetShelves_ShelfBooks_Book fail")
	}
	return
}

type ProxyOption func(*GoStatsClientProxy)

func WithRunner(runner goresilience.Runner) ProxyOption {
	return func(proxy *GoStatsClientProxy) {
		proxy.runner = runner
	}
}

func WithLogger(logger zerolog.Logger) ProxyOption {
	return func(proxy *GoStatsClientProxy) {
		proxy.logger = logger
	}
}

func NewGoStatsClientProxy(client *GoStatsClient, opts ...ProxyOption) *GoStatsClientProxy {
	cp := &GoStatsClientProxy{
		client: client,
		logger: zlogger.Logger,
	}

	for _, opt := range opts {
		opt(cp)
	}

	if cp.runner == nil {
		var mid []goresilience.Middleware
		mid = append(mid, metrics.NewMiddleware("go-doudou-tutorials/go-stats_client", metrics.NewPrometheusRecorder(prometheus.DefaultRegisterer)))
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
