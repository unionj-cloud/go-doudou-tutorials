package service

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-bff/config"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-bff/vo"
	v3 "github.com/unionj-cloud/go-doudou/toolkit/openapi/v3"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type WordcloudBffImpl struct {
	conf        *config.Config
	minioClient *minio.Client
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

func (receiver *WordcloudBffImpl) Upload(ctx context.Context, file v3.FileModel) (data vo.UploadResult, err error) {
	defer file.Close()
	userId, _ := UserIdFromContext(ctx)
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

	bucketName := receiver.conf.BizConf.OssBucket
	objectName := strings.TrimSuffix(fileName, filepath.Ext(fileName))

	_, err = receiver.minioClient.FPutObject(ctx, bucketName, objectName, out, minio.PutObjectOptions{})
	if err != nil {
		return vo.UploadResult{}, err
	}

	srcUrl := getPublicOssUrl(receiver.conf.BizConf.OssEndpoint, bucketName, objectName)

	ret := vo.UploadResult{
		TaskId: 0,
		SrcUrl: srcUrl,
		ImgUrl: "",
	}

	return ret, nil
}

func NewWordcloudBff(conf *config.Config, minioClient *minio.Client) WordcloudBff {
	return &WordcloudBffImpl{
		conf,
		minioClient,
	}
}
