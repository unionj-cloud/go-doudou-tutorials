package service

import (
	"context"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-bff/vo"
	v3 "github.com/unionj-cloud/go-doudou/toolkit/openapi/v3"
)

type WordcloudBff interface {
	Upload(ctx context.Context, file v3.FileModel) (data vo.UploadResult, err error)
}
