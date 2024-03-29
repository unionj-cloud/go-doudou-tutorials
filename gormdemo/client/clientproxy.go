/**
* Generated by go-doudou v2.0.5.
* You can edit it as your need.
 */
package client

import (
	"context"
	"gormdemo/dto"
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

type GormdemoClientProxy struct {
	client *GormdemoClient
	logger zerolog.Logger
	runner goresilience.Runner
}

func (receiver *GormdemoClientProxy) PageUsers(ctx context.Context, _headers map[string]string, query dto.PageQuery, options Options) (_resp *resty.Response, data dto.PageRet, err error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_resp, data, err = receiver.client.PageUsers(
			ctx,
			_headers,
			query,
			options,
		)
		if err != nil {
			return errors.Wrap(err, "call PageUsers fail")
		}
		return nil
	}); _err != nil {
		// you can implement your fallback logic here
		if errors.Is(_err, rerrors.ErrCircuitOpen) {
			receiver.logger.Error().Err(_err).Msg("")
		}
		err = errors.Wrap(_err, "call PageUsers fail")
	}
	return
}

type ProxyOption func(*GormdemoClientProxy)

func WithRunner(runner goresilience.Runner) ProxyOption {
	return func(proxy *GormdemoClientProxy) {
		proxy.runner = runner
	}
}

func WithLogger(logger zerolog.Logger) ProxyOption {
	return func(proxy *GormdemoClientProxy) {
		proxy.logger = logger
	}
}

func NewGormdemoClientProxy(client *GormdemoClient, opts ...ProxyOption) *GormdemoClientProxy {
	cp := &GormdemoClientProxy{
		client: client,
		logger: zlogger.Logger,
	}

	for _, opt := range opts {
		opt(cp)
	}

	if cp.runner == nil {
		var mid []goresilience.Middleware
		mid = append(mid, metrics.NewMiddleware("gormdemo_client", metrics.NewPrometheusRecorder(prometheus.DefaultRegisterer)))
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

func (receiver *GormdemoClientProxy) GetTUser_Id(ctx context.Context, _headers map[string]string, id int32, options Options) (_resp *resty.Response, data dto.TUser, err error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_resp, data, err = receiver.client.GetTUser_Id(
			ctx,
			_headers,
			id,
			options,
		)
		if err != nil {
			return errors.Wrap(err, "call GetTUser_Id fail")
		}
		return nil
	}); _err != nil {
		// you can implement your fallback logic here
		if errors.Is(_err, rerrors.ErrCircuitOpen) {
			receiver.logger.Error().Err(_err).Msg("")
		}
		err = errors.Wrap(_err, "call GetTUser_Id fail")
	}
	return
}

func (receiver *GormdemoClientProxy) PostTUser(ctx context.Context, _headers map[string]string, tUser dto.TUser, options Options) (_resp *resty.Response, data int32, err error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_resp, data, err = receiver.client.PostTUser(
			ctx,
			_headers,
			tUser,
			options,
		)
		if err != nil {
			return errors.Wrap(err, "call PostTUser fail")
		}
		return nil
	}); _err != nil {
		// you can implement your fallback logic here
		if errors.Is(_err, rerrors.ErrCircuitOpen) {
			receiver.logger.Error().Err(_err).Msg("")
		}
		err = errors.Wrap(_err, "call PostTUser fail")
	}
	return
}
func (receiver *GormdemoClientProxy) PutTUser(ctx context.Context, _headers map[string]string, tUser dto.TUser, options Options) (_resp *resty.Response, re error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_resp, re = receiver.client.PutTUser(
			ctx,
			_headers,
			tUser,
			options,
		)
		if re != nil {
			return errors.Wrap(re, "call PutTUser fail")
		}
		return nil
	}); _err != nil {
		// you can implement your fallback logic here
		if errors.Is(_err, rerrors.ErrCircuitOpen) {
			receiver.logger.Error().Err(_err).Msg("")
		}
		re = errors.Wrap(_err, "call PutTUser fail")
	}
	return
}
func (receiver *GormdemoClientProxy) DeleteTUser_Id(ctx context.Context, _headers map[string]string, id int32, options Options) (_resp *resty.Response, re error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_resp, re = receiver.client.DeleteTUser_Id(
			ctx,
			_headers,
			id,
			options,
		)
		if re != nil {
			return errors.Wrap(re, "call DeleteTUser_Id fail")
		}
		return nil
	}); _err != nil {
		// you can implement your fallback logic here
		if errors.Is(_err, rerrors.ErrCircuitOpen) {
			receiver.logger.Error().Err(_err).Msg("")
		}
		re = errors.Wrap(_err, "call DeleteTUser_Id fail")
	}
	return
}
func (receiver *GormdemoClientProxy) GetTUsers(ctx context.Context, _headers map[string]string, parameter dto.Parameter, options Options) (_resp *resty.Response, data dto.Page, err error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_resp, data, err = receiver.client.GetTUsers(
			ctx,
			_headers,
			parameter,
			options,
		)
		if err != nil {
			return errors.Wrap(err, "call GetTUsers fail")
		}
		return nil
	}); _err != nil {
		// you can implement your fallback logic here
		if errors.Is(_err, rerrors.ErrCircuitOpen) {
			receiver.logger.Error().Err(_err).Msg("")
		}
		err = errors.Wrap(_err, "call GetTUsers fail")
	}
	return
}

func (receiver *GormdemoClientProxy) PostGenTUser(ctx context.Context, _headers map[string]string, tUser dto.TUser, options Options) (_resp *resty.Response, data int32, err error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_resp, data, err = receiver.client.PostGenTUser(
			ctx,
			_headers,
			tUser,
			options,
		)
		if err != nil {
			return errors.Wrap(err, "call PostGenTUser fail")
		}
		return nil
	}); _err != nil {
		// you can implement your fallback logic here
		if errors.Is(_err, rerrors.ErrCircuitOpen) {
			receiver.logger.Error().Err(_err).Msg("")
		}
		err = errors.Wrap(_err, "call PostGenTUser fail")
	}
	return
}
func (receiver *GormdemoClientProxy) GetGenTUser_Id(ctx context.Context, _headers map[string]string, id int32, options Options) (_resp *resty.Response, data dto.TUser, err error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_resp, data, err = receiver.client.GetGenTUser_Id(
			ctx,
			_headers,
			id,
			options,
		)
		if err != nil {
			return errors.Wrap(err, "call GetGenTUser_Id fail")
		}
		return nil
	}); _err != nil {
		// you can implement your fallback logic here
		if errors.Is(_err, rerrors.ErrCircuitOpen) {
			receiver.logger.Error().Err(_err).Msg("")
		}
		err = errors.Wrap(_err, "call GetGenTUser_Id fail")
	}
	return
}
func (receiver *GormdemoClientProxy) PutGenTUser(ctx context.Context, _headers map[string]string, tUser dto.TUser, options Options) (_resp *resty.Response, re error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_resp, re = receiver.client.PutGenTUser(
			ctx,
			_headers,
			tUser,
			options,
		)
		if re != nil {
			return errors.Wrap(re, "call PutGenTUser fail")
		}
		return nil
	}); _err != nil {
		// you can implement your fallback logic here
		if errors.Is(_err, rerrors.ErrCircuitOpen) {
			receiver.logger.Error().Err(_err).Msg("")
		}
		re = errors.Wrap(_err, "call PutGenTUser fail")
	}
	return
}
func (receiver *GormdemoClientProxy) DeleteGenTUser_Id(ctx context.Context, _headers map[string]string, id int32, options Options) (_resp *resty.Response, re error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_resp, re = receiver.client.DeleteGenTUser_Id(
			ctx,
			_headers,
			id,
			options,
		)
		if re != nil {
			return errors.Wrap(re, "call DeleteGenTUser_Id fail")
		}
		return nil
	}); _err != nil {
		// you can implement your fallback logic here
		if errors.Is(_err, rerrors.ErrCircuitOpen) {
			receiver.logger.Error().Err(_err).Msg("")
		}
		re = errors.Wrap(_err, "call DeleteGenTUser_Id fail")
	}
	return
}
func (receiver *GormdemoClientProxy) GetGenTUsers(ctx context.Context, _headers map[string]string, parameter dto.Parameter, options Options) (_resp *resty.Response, data dto.Page, err error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_resp, data, err = receiver.client.GetGenTUsers(
			ctx,
			_headers,
			parameter,
			options,
		)
		if err != nil {
			return errors.Wrap(err, "call GetGenTUsers fail")
		}
		return nil
	}); _err != nil {
		// you can implement your fallback logic here
		if errors.Is(_err, rerrors.ErrCircuitOpen) {
			receiver.logger.Error().Err(_err).Msg("")
		}
		err = errors.Wrap(_err, "call GetGenTUsers fail")
	}
	return
}
