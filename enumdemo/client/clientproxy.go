package client

import (
	"abc2/vo"
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

type EnumDemoClientProxy struct {
	client *EnumDemoClient
	logger *logrus.Logger
	runner goresilience.Runner
}

func (receiver *EnumDemoClientProxy) GetKeyboard(ctx context.Context, _headers map[string]string, layout vo.KeyboardLayout) (_resp *resty.Response, data string, err error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_resp, data, err = receiver.client.GetKeyboard(
			ctx,
			_headers,
			layout,
		)
		if err != nil {
			return errors.Wrap(err, "call GetKeyboard fail")
		}
		return nil
	}); _err != nil {
		// you can implement your fallback logic here
		if errors.Is(_err, rerrors.ErrCircuitOpen) {
			receiver.logger.Error(_err)
		}
		err = errors.Wrap(_err, "call GetKeyboard fail")
	}
	return
}
func (receiver *EnumDemoClientProxy) GetKeyboard2(ctx context.Context, _headers map[string]string, layout *vo.KeyboardLayout) (_resp *resty.Response, data string, err error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_resp, data, err = receiver.client.GetKeyboard2(
			ctx,
			_headers,
			layout,
		)
		if err != nil {
			return errors.Wrap(err, "call GetKeyboard2 fail")
		}
		return nil
	}); _err != nil {
		// you can implement your fallback logic here
		if errors.Is(_err, rerrors.ErrCircuitOpen) {
			receiver.logger.Error(_err)
		}
		err = errors.Wrap(_err, "call GetKeyboard2 fail")
	}
	return
}
func (receiver *EnumDemoClientProxy) GetKeyboards(ctx context.Context, _headers map[string]string, layout []vo.KeyboardLayout) (_resp *resty.Response, data []string, err error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_resp, data, err = receiver.client.GetKeyboards(
			ctx,
			_headers,
			layout,
		)
		if err != nil {
			return errors.Wrap(err, "call GetKeyboards fail")
		}
		return nil
	}); _err != nil {
		// you can implement your fallback logic here
		if errors.Is(_err, rerrors.ErrCircuitOpen) {
			receiver.logger.Error(_err)
		}
		err = errors.Wrap(_err, "call GetKeyboards fail")
	}
	return
}
func (receiver *EnumDemoClientProxy) GetKeyboards2(ctx context.Context, _headers map[string]string, layout *[]vo.KeyboardLayout) (_resp *resty.Response, data []string, err error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_resp, data, err = receiver.client.GetKeyboards2(
			ctx,
			_headers,
			layout,
		)
		if err != nil {
			return errors.Wrap(err, "call GetKeyboards2 fail")
		}
		return nil
	}); _err != nil {
		// you can implement your fallback logic here
		if errors.Is(_err, rerrors.ErrCircuitOpen) {
			receiver.logger.Error(_err)
		}
		err = errors.Wrap(_err, "call GetKeyboards2 fail")
	}
	return
}
func (receiver *EnumDemoClientProxy) GetKeyboards5(ctx context.Context, _headers map[string]string, layout ...vo.KeyboardLayout) (_resp *resty.Response, data []string, err error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_resp, data, err = receiver.client.GetKeyboards5(
			ctx,
			_headers,
			layout...,
		)
		if err != nil {
			return errors.Wrap(err, "call GetKeyboards5 fail")
		}
		return nil
	}); _err != nil {
		// you can implement your fallback logic here
		if errors.Is(_err, rerrors.ErrCircuitOpen) {
			receiver.logger.Error(_err)
		}
		err = errors.Wrap(_err, "call GetKeyboards5 fail")
	}
	return
}

type ProxyOption func(*EnumDemoClientProxy)

func WithRunner(runner goresilience.Runner) ProxyOption {
	return func(proxy *EnumDemoClientProxy) {
		proxy.runner = runner
	}
}

func WithLogger(logger *logrus.Logger) ProxyOption {
	return func(proxy *EnumDemoClientProxy) {
		proxy.logger = logger
	}
}

func NewEnumDemoClientProxy(client *EnumDemoClient, opts ...ProxyOption) *EnumDemoClientProxy {
	cp := &EnumDemoClientProxy{
		client: client,
		logger: logrus.StandardLogger(),
	}

	for _, opt := range opts {
		opt(cp)
	}

	if cp.runner == nil {
		var mid []goresilience.Middleware
		mid = append(mid, metrics.NewMiddleware("abc2_client", metrics.NewPrometheusRecorder(prometheus.DefaultRegisterer)))
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

func (receiver *EnumDemoClientProxy) Keyboard(ctx context.Context, _headers map[string]string, keyboard vo.Keyboard) (_resp *resty.Response, data string, err error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_resp, data, err = receiver.client.Keyboard(
			ctx,
			_headers,
			keyboard,
		)
		if err != nil {
			return errors.Wrap(err, "call Keyboard fail")
		}
		return nil
	}); _err != nil {
		// you can implement your fallback logic here
		if errors.Is(_err, rerrors.ErrCircuitOpen) {
			receiver.logger.Error(_err)
		}
		err = errors.Wrap(_err, "call Keyboard fail")
	}
	return
}
