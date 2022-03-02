package service

import (
	"context"

	"github.com/unionj-cloud/go-doudou/toolkit/copier"
	"github.com/unionj-cloud/go-doudou/toolkit/sqlext/query"
	"github.com/unionj-cloud/go-doudou/toolkit/sqlext/sortenum"

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

func (receiver *WordcloudTaskImpl) TaskSave(ctx context.Context, payload vo.TaskPayload) (data int, err error) {
	taskDao := dao.NewWordCloudTaskDao(receiver.db)
	task := domain.WordCloudTask{
		SrcUrl: payload.SrcUrl,
		UserId: payload.UserId,
		Lang:   payload.Lang,
		Top:    payload.Top,
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

func (receiver *WordcloudTaskImpl) TaskPage(ctx context.Context, pq vo.PageQuery) (data vo.TaskPageRet, err error) {
	taskDao := dao.NewWordCloudTaskDao(receiver.db)
	var page query.Page
	for _, item := range pq.Page.Orders {
		page.Orders = append(page.Orders, query.Order{
			Col:  item.Col,
			Sort: sortenum.Sort(item.Sort),
		})
	}
	page.Size = pq.Page.Size
	pageNo := pq.Page.PageNo
	if pageNo < 1 {
		pageNo = 1
	}
	page.Offset = (pageNo - 1) * pq.Page.Size
	var where query.Q
	where = query.C().Col("delete_at").IsNull()
	if pq.Filter.UserId > 0 {
		where = query.C().Col("user_id").Eq(pq.Filter.UserId)
	}
	pr, err := taskDao.PageMany(ctx, page, where)
	if err != nil {
		return vo.TaskPageRet{}, err
	}
	copier.DeepCopy(pr, &data)
	return data, nil
}
