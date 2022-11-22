package service

import (
	"context"

	"github.com/unionj-cloud/go-doudou/v2/framework/rest"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/pkg/errors"

	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-seg/internal/lib/en"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-seg/internal/lib/nlp"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-seg/internal/lib/zh"

	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-seg/config"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-seg/vo"

	pb "github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-seg/transport/grpc"
)

var _ pb.WordcloudSegServiceServer = (*WordcloudSegImpl)(nil)

type WordcloudSegImpl struct {
	pb.UnimplementedWordcloudSegServiceServer

	conf  *config.Config
	golac *zh.GoThulac
}

func (receiver *WordcloudSegImpl) Seg(ctx context.Context, payload vo.SegPayload) (rs vo.SegResult, err error) {
	size := len([]byte(payload.Text))
	limit := 1 << (10 * 2)
	if size > limit {
		return vo.SegResult{}, rest.NewBizError(errors.New("text size is larger than 1M"), rest.WithStatusCode(400))
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

func NewWordcloudSeg(conf *config.Config, golac *zh.GoThulac) *WordcloudSegImpl {
	return &WordcloudSegImpl{
		conf:  conf,
		golac: golac,
	}
}

func (receiver *WordcloudSegImpl) SegRpc(ctx context.Context, request *pb.SegPayload) (*pb.SegResult, error) {
	payload := vo.SegPayload{
		Text: request.Text,
		Lang: request.Lang,
	}
	rs, err := receiver.Seg(ctx, payload)
	if err != nil {
		return nil, err
	}
	var wordFreq []*pb.NestedAny
	for _, item := range rs.WordFreq {
		var nestedAny []*anypb.Any
		// item[0] is word
		wordWrapper := wrapperspb.String(item[0].(string))
		wordAny, _ := anypb.New(wordWrapper)
		nestedAny = append(nestedAny, wordAny)

		// item[1] is part of speech
		posWrapper := wrapperspb.String(item[1].(string))
		posAny, _ := anypb.New(posWrapper)
		nestedAny = append(nestedAny, posAny)

		// item[2] is frequency
		freqWrapper := wrapperspb.Double(item[2].(float64))
		freqAny, _ := anypb.New(freqWrapper)
		nestedAny = append(nestedAny, freqAny)

		wordFreq = append(wordFreq, &pb.NestedAny{
			NestedAny: nestedAny,
		})
	}
	return &pb.SegResult{
		WordFreq: wordFreq,
	}, nil
}
