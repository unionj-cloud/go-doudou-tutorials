package service

import (
	"context"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-maker/vo"
)

type WordcloudMaker interface {
	// Make 生成词云图接口
	// make word cloud image
	Make(ctx context.Context, payload vo.MakePayload) (
		// 词云图链接
		// word cloud image url
		data string, err error)
}
