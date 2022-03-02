package service

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/minio/minio-go/v7"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-bff/config"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-bff/vo"
	makersvc "github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-maker"
	makervo "github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-maker/vo"
	tasksvc "github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-task"
	taskvo "github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-task/vo"
	v3 "github.com/unionj-cloud/go-doudou/toolkit/openapi/v3"
)

type WordcloudBffImpl struct {
	conf        *config.Config
	minioClient *minio.Client
	makerClient makersvc.WordcloudMaker
	taskClient  tasksvc.WordcloudTask
}

type ctxKey int

const userIdKey ctxKey = ctxKey(0)

func NewUserIdContext(ctx context.Context, id int) context.Context {
	return context.WithValue(ctx, userIdKey, id)
}

func UserIdFromContext(ctx context.Context) (int, bool) {
	userId, ok := ctx.Value(userIdKey).(int)
	return userId, ok
}

func getPublicOssUrl(endpoint, bucketName, objectName string) string {
	return fmt.Sprintf("http://%s/%s/%s", endpoint, bucketName, objectName)
}

func (receiver *WordcloudBffImpl) Upload(ctx context.Context, file v3.FileModel, lang string, top *int) (data vo.UploadResult, err error) {
	defer file.Close()
	// get user ID from context
	userId, _ := UserIdFromContext(ctx)

	// save uploaded file to disk
	_ = os.MkdirAll(receiver.conf.BizConf.Output, os.ModePerm)
	fileName := fmt.Sprintf("user%d_%s", userId, file.Filename)
	out := filepath.Join(receiver.conf.BizConf.Output, fileName)
	var f *os.File
	f, err = os.OpenFile(out, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return vo.UploadResult{}, err
	}
	defer f.Close()
	_, err = io.Copy(f, file.Reader)
	if err != nil {
		return vo.UploadResult{}, err
	}

	// save file from disk to minio storage
	bucketName := receiver.conf.BizConf.OssBucket
	objectName := strings.TrimSuffix(fileName, filepath.Ext(fileName))
	_, err = receiver.minioClient.FPutObject(ctx, bucketName, objectName, out, minio.PutObjectOptions{})
	if err != nil {
		return vo.UploadResult{}, err
	}
	srcUrl := getPublicOssUrl(receiver.conf.BizConf.OssEndpoint, bucketName, objectName)

	// save task
	tp := taskvo.TaskPayload{
		UserId: userId,
		SrcUrl: srcUrl,
		Lang:   lang,
	}

	if top != nil {
		tp.Top = *top
	}

	taskId, err := receiver.taskClient.TaskSave(ctx, tp)
	if err != nil {
		return vo.UploadResult{}, err
	}

	// generate word cloud image
	mp := makervo.MakePayload{
		SrcUrl: srcUrl,
		Lang:   lang,
	}

	if top != nil {
		mp.Top = *top
	}

	imgUrl, err := receiver.makerClient.Make(ctx, mp)
	if err != nil {
		// if fail call TaskFail api
		old := err
		_, err = receiver.taskClient.TaskFail(ctx, taskvo.TaskFail{
			TaskId: taskId,
			Error:  err.Error(),
		})
		if err != nil {
			return vo.UploadResult{}, err
		}
		return vo.UploadResult{}, old
	}

	// if success call TaskSuccess api
	_, err = receiver.taskClient.TaskSuccess(ctx, taskvo.TaskSuccess{
		TaskId: taskId,
		ImgUrl: imgUrl,
	})
	if err != nil {
		return vo.UploadResult{}, err
	}

	return vo.UploadResult{
		TaskId: taskId,
		SrcUrl: srcUrl,
		ImgUrl: imgUrl,
	}, nil
}

func NewWordcloudBff(conf *config.Config, minioClient *minio.Client,
	makerClient makersvc.WordcloudMaker, taskClient tasksvc.WordcloudTask) WordcloudBff {
	return &WordcloudBffImpl{
		conf,
		minioClient,
		makerClient,
		taskClient,
	}
}

func (receiver *WordcloudBffImpl) TaskPage(ctx context.Context, query vo.PageQuery) (data vo.TaskPageRet, err error) {
	userId, _ := UserIdFromContext(ctx)
	var pq taskvo.PageQuery

	receiver.taskClient.TaskPage(ctx, query)
}
