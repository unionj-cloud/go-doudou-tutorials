package service

import (
	"context"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-bff/vo"
	v3 "github.com/unionj-cloud/go-doudou/toolkit/openapi/v3"
)

type WordcloudBff interface {
	// 上传文本文件生成词云图接口
	// upload text file to generate word cloud image
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

	// 词云图任务分页接口
	// word cloud task pagination
	TaskPage(ctx context.Context,
		// 分页请求参数
		// pagination parameter
		query vo.PageQuery) (
		// 分页结果
		// pagination result
		data vo.PageRet,
		// 错误信息
		// error
		err error)
}
