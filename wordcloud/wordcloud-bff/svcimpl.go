package service

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/iancoleman/strcase"

	"github.com/minio/minio-go/v7"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-bff/config"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-bff/vo"
	makerclient "github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-maker/client"
	makervo "github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-maker/vo"
	taskclient "github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-task/client"
	taskvo "github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-task/vo"
	userclient "github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-user/client"
	"github.com/unionj-cloud/go-doudou/toolkit/copier"
	v3 "github.com/unionj-cloud/go-doudou/toolkit/openapi/v3"
)

type WordcloudBffImpl struct {
	conf        *config.Config
	minioClient *minio.Client
	makerClient makerclient.IWordcloudMakerClient
	taskClient  taskclient.IWordcloudTaskClient
	userClient  userclient.IUsersvcClient
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

const tokenKey ctxKey = ctxKey(1)

func NewTokenContext(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, tokenKey, token)
}

func TokenFromContext(ctx context.Context) (string, bool) {
	token, ok := ctx.Value(tokenKey).(string)
	return token, ok
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

	_, taskId, err := receiver.taskClient.TaskSave(ctx, nil, tp)
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

	_, imgUrl, err := receiver.makerClient.Make(ctx, nil, mp)
	if err != nil {
		// if fail call TaskFail api
		old := err
		_, _, err = receiver.taskClient.TaskFail(ctx, nil, taskvo.TaskFail{
			TaskId: taskId,
			Error:  err.Error(),
		})
		if err != nil {
			return vo.UploadResult{}, err
		}
		return vo.UploadResult{}, old
	}

	// if success call TaskSuccess api
	_, _, err = receiver.taskClient.TaskSuccess(ctx, nil, taskvo.TaskSuccess{
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
	makerClient makerclient.IWordcloudMakerClient, taskClient taskclient.IWordcloudTaskClient, userClient userclient.IUsersvcClient) WordcloudBff {
	return &WordcloudBffImpl{
		conf,
		minioClient,
		makerClient,
		taskClient,
		userClient,
	}
}

func (receiver *WordcloudBffImpl) GetTaskPage(ctx context.Context, page int, pageSize int) (result vo.TaskPageRet, err error) {
	userId, _ := UserIdFromContext(ctx)
	pq := taskvo.PageQuery{
		Page: taskvo.Page{
			PageNo: page,
			Size:   pageSize,
		},
	}
	for i, item := range pq.Page.Orders {
		pq.Page.Orders[i].Col = strcase.ToSnake(item.Col)
	}
	pq.Filter.UserId = userId
	_, ret, err := receiver.taskClient.TaskPage(ctx, nil, pq)
	if err != nil {
		return vo.TaskPageRet{}, err
	}
	copier.DeepCopy(ret.PageRet, &result.PageRet)
	usermap := make(map[int]string)
	for _, item := range ret.Items {
		usermap[item.UserId] = ""
	}
	token, _ := TokenFromContext(ctx)
	headers := make(map[string]string)
	headers["Authorization"] = fmt.Sprintf("Bearer %s", token)
	for k, _ := range usermap {
		_, user, err := receiver.userClient.GetUser(ctx, headers, k)
		if err != nil {
			return vo.TaskPageRet{}, err
		}
		usermap[k] = user.Username
	}
	for _, item := range ret.Items {
		var taskVo vo.TaskVo
		copier.DeepCopy(item, &taskVo)
		taskVo.Username = usermap[item.UserId]
		result.Items = append(result.Items, taskVo)
	}
	return
}
