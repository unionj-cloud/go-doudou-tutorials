package service

import (
	"context"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-task/vo"
)

type WordcloudTask interface {
	// TaskSave save word cloud task 保存词云图任务
	TaskSave(ctx context.Context, payload vo.TaskPayload) (
		// task ID 任务ID
		data int, err error)

	// TaskSuccess task success 任务执行成功接口
	TaskSuccess(ctx context.Context, payload vo.TaskSuccess) (
		// return OK if success
		// 成功返回OK
		data string, err error)

	// TaskFail task fail 任务执行失败接口
	TaskFail(ctx context.Context, payload vo.TaskFail) (
		// return OK if success
		// 成功返回OK
		data string, err error)
}
