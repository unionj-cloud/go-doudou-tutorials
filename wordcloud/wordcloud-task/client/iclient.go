package client

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-task/vo"
)

type IWordcloudTaskClient interface {
	TaskSave(ctx context.Context, _headers map[string]string, payload vo.TaskPayload) (_resp *resty.Response, data int, err error)
	TaskSuccess(ctx context.Context, _headers map[string]string, payload vo.TaskSuccess) (_resp *resty.Response, data string, err error)
	TaskFail(ctx context.Context, _headers map[string]string, payload vo.TaskFail) (_resp *resty.Response, data string, err error)
	TaskPage(ctx context.Context, _headers map[string]string, query vo.PageQuery) (_resp *resty.Response, data vo.TaskPageRet, err error)
}
