package client

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-bff/vo"
	v3 "github.com/unionj-cloud/go-doudou/toolkit/openapi/v3"
)

type IWordcloudBffClient interface {
	Upload(ctx context.Context, _headers map[string]string, file v3.FileModel, lang string, top *int) (_resp *resty.Response, data vo.UploadResult, err error)
	GetTaskPage(ctx context.Context, _headers map[string]string, page int, pageSize int) (_resp *resty.Response, result vo.TaskPageRet, err error)
}
