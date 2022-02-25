package service

import (
	"context"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-bff/vo"
	v3 "github.com/unionj-cloud/go-doudou/toolkit/openapi/v3"
)

type WordcloudBff interface {
	Upload(ctx context.Context,
		// 文本文件
		// text file
		file v3.FileModel,
		// 语种
		// language
		lang string,
		// 取词频前top的词
		// only take top frequent words into word cloud
		top *int) (data vo.UploadResult, err error)
}
