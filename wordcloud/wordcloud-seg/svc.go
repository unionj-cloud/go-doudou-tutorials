package service

import (
	"context"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-seg/vo"
)

type WordcloudSeg interface {
	Seg(ctx context.Context, payload vo.SegPayload) (rs vo.SegResult, err error)
}
