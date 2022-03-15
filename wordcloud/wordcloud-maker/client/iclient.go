package client

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-maker/vo"
)

type IWordcloudMakerClient interface {
	Make(ctx context.Context, _headers map[string]string, payload vo.MakePayload) (_resp *resty.Response, data string, err error)
}
