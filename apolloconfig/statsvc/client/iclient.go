package client

import (
	"context"

	"github.com/go-resty/resty/v2"
)

type IStatsvcClient interface {
	Add(ctx context.Context, _headers map[string]string, x int, y int) (_resp *resty.Response, data int, err error)
}
