package service

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-task/dao"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-task/domain"

	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-task/config"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-task/vo"
)

type WordcloudTaskImpl struct {
	conf *config.Config
	db   *sqlx.DB
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

func (receiver *WordcloudTaskImpl) TaskSave(ctx context.Context, srcUrl string) (data int, err error) {
	userId, _ := UserIdFromContext(ctx)
	taskDao := dao.NewWordCloudTaskDao(receiver.db)
	task := domain.WordCloudTask{
		SrcUrl: srcUrl,
		UserId: userId,
	}
	_, err = taskDao.Insert(ctx, &task)
	if err != nil {
		return 0, err
	}
	return task.Id, nil
}
func (receiver *WordcloudTaskImpl) TaskSuccess(ctx context.Context, payload vo.TaskSuccess) (data string, err error) {
	taskDao := dao.NewWordCloudTaskDao(receiver.db)
	_, err = taskDao.UpdateNoneZero(ctx, domain.WordCloudTask{
		Id:     payload.TaskId,
		ImgUrl: payload.ImgUrl,
		Status: 2,
	})
	if err != nil {
		return "", err
	}
	return "OK", nil
}
func (receiver *WordcloudTaskImpl) TaskFail(ctx context.Context, payload vo.TaskFail) (data string, err error) {
	taskDao := dao.NewWordCloudTaskDao(receiver.db)
	_, err = taskDao.UpdateNoneZero(ctx, domain.WordCloudTask{
		Id:     payload.TaskId,
		Status: 3,
		Error:  payload.Error,
	})
	if err != nil {
		return "", err
	}
	return "OK", nil
}

func NewWordcloudTask(conf *config.Config, db *sqlx.DB) WordcloudTask {
	return &WordcloudTaskImpl{
		conf,
		db,
	}
}
