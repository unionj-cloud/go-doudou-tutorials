package service

import (
	"context"

	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-seg/config"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-seg/internal/lib"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-seg/vo"
)

type WordcloudSegImpl struct {
	conf  *config.Config
	golac *lib.GoThulac
}

func (receiver *WordcloudSegImpl) Seg(ctx context.Context, payload vo.SegPayload) (rs vo.SegResult, err error) {
	seg, err := receiver.golac.DoSeg(ctx, payload.Text, lib.Pos, -1, -1)
	if err != nil {
		return vo.SegResult{}, err
	}
	return vo.SegResult{
		WordFreq: seg,
	}, nil
}

func NewWordcloudSeg(conf *config.Config, golac *lib.GoThulac) WordcloudSeg {
	return &WordcloudSegImpl{
		conf,
		golac,
	}
}
