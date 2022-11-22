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
	makerpb "github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-maker/transport/grpc"
	taskpb "github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-task/transport/grpc"
	userclient "github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-user/client"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/copier"
	v3 "github.com/unionj-cloud/go-doudou/v2/toolkit/openapi/v3"
)

var _ WordcloudBff = (*WordcloudBffImpl)(nil)

type WordcloudBffImpl struct {
	conf        *config.Config
	minioClient *minio.Client
	makerClient makerpb.WordcloudMakerServiceClient
	taskClient  taskpb.WordcloudTaskServiceClient
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
	tp := &taskpb.TaskPayload{
		UserId: int32(userId),
		SrcUrl: srcUrl,
		Lang:   lang,
		Top:    0,
	}
	if top != nil {
		tp.Top = int32(*top)
	}
	taskSaveResp, err := receiver.taskClient.TaskSaveRpc(ctx, tp)
	if err != nil {
		return vo.UploadResult{}, err
	}
	taskId := taskSaveResp.Data

	// generate word cloud image
	mp := &makerpb.MakePayload{
		SrcUrl: srcUrl,
		Lang:   lang,
	}
	if top != nil {
		mp.Top = int32(*top)
	}
	makeResp, err := receiver.makerClient.MakeRpc(ctx, mp)
	if err != nil {
		// if fail call TaskFail api
		old := err
		_, err = receiver.taskClient.TaskFailRpc(ctx, &taskpb.TaskFail{
			TaskId: taskId,
			Error:  err.Error(),
		})
		if err != nil {
			return vo.UploadResult{}, err
		}
		return vo.UploadResult{}, old
	}
	imgUrl := makeResp.Data

	// if success call TaskSuccess api
	_, err = receiver.taskClient.TaskSuccessRpc(ctx, &taskpb.TaskSuccess{
		TaskId: taskId,
		ImgUrl: imgUrl,
	})
	if err != nil {
		return vo.UploadResult{}, err
	}

	return vo.UploadResult{
		TaskId: int(taskId),
		SrcUrl: srcUrl,
		ImgUrl: imgUrl,
	}, nil
}

func NewWordcloudBff(conf *config.Config, minioClient *minio.Client,
	makerClient makerpb.WordcloudMakerServiceClient, taskClient taskpb.WordcloudTaskServiceClient, userClient userclient.IUsersvcClient) *WordcloudBffImpl {
	return &WordcloudBffImpl{
		conf:        conf,
		minioClient: minioClient,
		makerClient: makerClient,
		taskClient:  taskClient,
		userClient:  userClient,
	}
}

func (receiver *WordcloudBffImpl) GetTaskPage(ctx context.Context, page int, pageSize int) (result vo.TaskPageRet, err error) {
	userId, _ := UserIdFromContext(ctx)
	pq := taskpb.PageQuery{
		Page: &taskpb.Page{
			PageNo: int32(page),
			Size:   int32(pageSize),
		},
	}
	for i, item := range pq.Page.Orders {
		pq.Page.Orders[i].Col = strcase.ToSnake(item.Col)
	}
	pq.Filter.UserId = int32(userId)
	ret, err := receiver.taskClient.TaskPageRpc(ctx, &pq)
	if err != nil {
		return vo.TaskPageRet{}, err
	}
	copier.DeepCopy(ret, &result)
	token, _ := TokenFromContext(ctx)
	headers := make(map[string]string)
	headers["Authorization"] = fmt.Sprintf("Bearer %s", token)
	for i, item := range result.Items {
		_, user, err := receiver.userClient.GetUser(ctx, headers, item.UserId)
		if err != nil {
			return vo.TaskPageRet{}, err
		}
		result.Items[i].Username = user.Username
	}
	return
}
