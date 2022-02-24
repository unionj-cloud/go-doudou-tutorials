package service

import (
	"context"
	"github.com/pkg/errors"
	ddhttp "github.com/unionj-cloud/go-doudou/framework/http"

	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-seg/internal/lib/en"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-seg/internal/lib/nlp"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-seg/internal/lib/zh"

	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-seg/config"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-seg/vo"
)

type WordcloudSegImpl struct {
	conf  *config.Config
	golac *zh.GoThulac
}

func (receiver *WordcloudSegImpl) Seg(ctx context.Context, payload vo.SegPayload) (rs vo.SegResult, err error) {
	size := len([]byte(payload.Text))
	limit := 1 << (10 * 2)
	if size > limit {
		return vo.SegResult{}, ddhttp.NewBizError(errors.New("text size is larger than 1M"), ddhttp.WithStatusCode(400))
	}

	var seg nlp.WordFreqResult
	var tokenizer nlp.Tokenizer
	switch vo.Lang(payload.Lang) {
	case "":
		fallthrough
	case vo.Zh:
		tokenizer = receiver.golac
	case vo.En:
		tokenizer = &en.EnglishTokenizer{}
	default:
		return vo.SegResult{}, errors.New("not support")
	}
	seg, err = tokenizer.DoSeg(ctx, payload.Text, nil, -1, -1)
	if err != nil {
		return vo.SegResult{}, err
	}
	return vo.SegResult{
		WordFreq: seg,
	}, nil
}

func NewWordcloudSeg(conf *config.Config, golac *zh.GoThulac) WordcloudSeg {
	return &WordcloudSegImpl{
		conf,
		golac,
	}
}
