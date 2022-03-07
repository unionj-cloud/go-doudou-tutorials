package client

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-seg/vo"
)

type IWordcloudSegClient interface {
	Seg(ctx context.Context, _headers map[string]string, payload vo.SegPayload) (_resp *resty.Response, rs vo.SegResult, err error)
}
