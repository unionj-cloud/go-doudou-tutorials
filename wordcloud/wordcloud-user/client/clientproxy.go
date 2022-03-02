package client

import (
	"context"
	"os"
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
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-user/vo"
	v3 "github.com/unionj-cloud/go-doudou/toolkit/openapi/v3"
)

type UsersvcClientProxy struct {
	client *UsersvcClient
	logger *logrus.Logger
	runner goresilience.Runner
}

func (receiver *UsersvcClientProxy) PageUsers(ctx context.Context, _headers map[string]string, query vo.PageQuery) (_resp *resty.Response, data vo.PageRet, err error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_resp, data, err = receiver.client.PageUsers(
			ctx,
			_headers,
			query,
		)
		if err != nil {
			return errors.Wrap(err, "call PageUsers fail")
		}
		return nil
	}); _err != nil {
		// you can implement your fallback logic here
		if errors.Is(_err, rerrors.ErrCircuitOpen) {
			receiver.logger.Error(_err)
		}
		err = errors.Wrap(_err, "call PageUsers fail")
	}
	return
}
func (receiver *UsersvcClientProxy) GetUser(ctx context.Context, _headers map[string]string, userId int) (_resp *resty.Response, data vo.UserVo, err error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_resp, data, err = receiver.client.GetUser(
			ctx,
			_headers,
			userId,
		)
		if err != nil {
			return errors.Wrap(err, "call GetUser fail")
		}
		return nil
	}); _err != nil {
		// you can implement your fallback logic here
		if errors.Is(_err, rerrors.ErrCircuitOpen) {
			receiver.logger.Error(_err)
		}
		err = errors.Wrap(_err, "call GetUser fail")
	}
	return
}
func (receiver *UsersvcClientProxy) GetMe(ctx context.Context, _headers map[string]string) (_resp *resty.Response, data vo.UserVo, err error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_resp, data, err = receiver.client.GetMe(
			ctx,
			_headers,
		)
		if err != nil {
			return errors.Wrap(err, "call GetMe fail")
		}
		return nil
	}); _err != nil {
		// you can implement your fallback logic here
		if errors.Is(_err, rerrors.ErrCircuitOpen) {
			receiver.logger.Error(_err)
		}
		err = errors.Wrap(_err, "call GetMe fail")
	}
	return
}
func (receiver *UsersvcClientProxy) PublicSignUp(ctx context.Context, _headers map[string]string, username string, password string, code *string) (_resp *resty.Response, data int, err error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_resp, data, err = receiver.client.PublicSignUp(
			ctx,
			_headers,
			username,
			password,
			code,
		)
		if err != nil {
			return errors.Wrap(err, "call PublicSignUp fail")
		}
		return nil
	}); _err != nil {
		// you can implement your fallback logic here
		if errors.Is(_err, rerrors.ErrCircuitOpen) {
			receiver.logger.Error(_err)
		}
		err = errors.Wrap(_err, "call PublicSignUp fail")
	}
	return
}
func (receiver *UsersvcClientProxy) PublicLogIn(ctx context.Context, _headers map[string]string, username string, password string) (_resp *resty.Response, data string, err error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_resp, data, err = receiver.client.PublicLogIn(
			ctx,
			_headers,
			username,
			password,
		)
		if err != nil {
			return errors.Wrap(err, "call PublicLogIn fail")
		}
		return nil
	}); _err != nil {
		// you can implement your fallback logic here
		if errors.Is(_err, rerrors.ErrCircuitOpen) {
			receiver.logger.Error(_err)
		}
		err = errors.Wrap(_err, "call PublicLogIn fail")
	}
	return
}
func (receiver *UsersvcClientProxy) UploadAvatar(ctx context.Context, _headers map[string]string, avatar v3.FileModel, id int) (_resp *resty.Response, data string, err error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_resp, data, err = receiver.client.UploadAvatar(
			ctx,
			_headers,
			avatar,
			id,
		)
		if err != nil {
			return errors.Wrap(err, "call UploadAvatar fail")
		}
		return nil
	}); _err != nil {
		// you can implement your fallback logic here
		if errors.Is(_err, rerrors.ErrCircuitOpen) {
			receiver.logger.Error(_err)
		}
		err = errors.Wrap(_err, "call UploadAvatar fail")
	}
	return
}
func (receiver *UsersvcClientProxy) GetPublicDownloadAvatar(ctx context.Context, _headers map[string]string, userId int) (_resp *resty.Response, data *os.File, err error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_resp, data, err = receiver.client.GetPublicDownloadAvatar(
			ctx,
			_headers,
			userId,
		)
		if err != nil {
			return errors.Wrap(err, "call GetPublicDownloadAvatar fail")
		}
		return nil
	}); _err != nil {
		// you can implement your fallback logic here
		if errors.Is(_err, rerrors.ErrCircuitOpen) {
			receiver.logger.Error(_err)
		}
		err = errors.Wrap(_err, "call GetPublicDownloadAvatar fail")
	}
	return
}
func (receiver *UsersvcClientProxy) GetLogout(ctx context.Context, _headers map[string]string) (_resp *resty.Response, data string, err error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_resp, data, err = receiver.client.GetLogout(
			ctx,
			_headers,
		)
		if err != nil {
			return errors.Wrap(err, "call GetLogout fail")
		}
		return nil
	}); _err != nil {
		// you can implement your fallback logic here
		if errors.Is(_err, rerrors.ErrCircuitOpen) {
			receiver.logger.Error(_err)
		}
		err = errors.Wrap(_err, "call GetLogout fail")
	}
	return
}
func (receiver *UsersvcClientProxy) PublicTokenValidate(ctx context.Context, _headers map[string]string, token string) (_resp *resty.Response, user vo.UserVo, err error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_resp, user, err = receiver.client.PublicTokenValidate(
			ctx,
			_headers,
			token,
		)
		if err != nil {
			return errors.Wrap(err, "call PublicTokenValidate fail")
		}
		return nil
	}); _err != nil {
		// you can implement your fallback logic here
		if errors.Is(_err, rerrors.ErrCircuitOpen) {
			receiver.logger.Error(_err)
		}
		err = errors.Wrap(_err, "call PublicTokenValidate fail")
	}
	return
}

type ProxyOption func(*UsersvcClientProxy)

func WithRunner(runner goresilience.Runner) ProxyOption {
	return func(proxy *UsersvcClientProxy) {
		proxy.runner = runner
	}
}

func WithLogger(logger *logrus.Logger) ProxyOption {
	return func(proxy *UsersvcClientProxy) {
		proxy.logger = logger
	}
}

func NewUsersvcClientProxy(client *UsersvcClient, opts ...ProxyOption) *UsersvcClientProxy {
	cp := &UsersvcClientProxy{
		client: client,
		logger: logrus.StandardLogger(),
	}

	for _, opt := range opts {
		opt(cp)
	}

	if cp.runner == nil {
		var mid []goresilience.Middleware
		mid = append(mid, metrics.NewMiddleware("github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-user_client", metrics.NewPrometheusRecorder(prometheus.DefaultRegisterer)))
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
