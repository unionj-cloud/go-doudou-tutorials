package service

import (
	"context"

	"github.com/unionj-cloud/go-doudou/v2/toolkit/copier"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/sqlext/query"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/sqlext/sortenum"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/sqlext/wrapper"

	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-task/dao"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-task/domain"

	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-task/config"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-task/vo"

	pb "github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-task/transport/grpc"
)

var _ pb.WordcloudTaskServiceServer = (*WordcloudTaskImpl)(nil)

type WordcloudTaskImpl struct {
	pb.UnimplementedWordcloudTaskServiceServer

	conf *config.Config
	db   wrapper.GddDB
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
	_, err = taskDao.UpdateNoneZero(ctx, &domain.WordCloudTask{
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
	_, err = taskDao.UpdateNoneZero(ctx, &domain.WordCloudTask{
		Id:     payload.TaskId,
		Status: 3,
		Error:  payload.Error,
	})
	if err != nil {
		return "", err
	}
	return "OK", nil
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
	var where query.Where
	if pq.Filter.UserId > 0 {
		where = where.And(query.C().Col("user_id").Eq(pq.Filter.UserId))
	}
	var dest dao.WordCloudTaskPageRet
	err = taskDao.PageMany(ctx, &dest, page, where)
	if err != nil {
		return vo.TaskPageRet{}, err
	}
	copier.DeepCopy(dest, &data)
	return data, nil
}

func NewWordcloudTask(conf *config.Config, db wrapper.GddDB) *WordcloudTaskImpl {
	return &WordcloudTaskImpl{
		conf: conf,
		db:   db,
	}
}

func (receiver *WordcloudTaskImpl) TaskSaveRpc(ctx context.Context, request *pb.TaskPayload) (*pb.TaskSaveRpcResponse, error) {
	payload := vo.TaskPayload{
		UserId: int(request.UserId),
		SrcUrl: request.SrcUrl,
		Lang:   request.Lang,
		Top:    int(request.Top),
	}
	data, err := receiver.TaskSave(ctx, payload)
	if err != nil {
		return nil, err
	}
	return &pb.TaskSaveRpcResponse{
		Data: int32(data),
	}, nil
}
func (receiver *WordcloudTaskImpl) TaskSuccessRpc(ctx context.Context, request *pb.TaskSuccess) (*pb.TaskSuccessRpcResponse, error) {
	payload := vo.TaskSuccess{
		TaskId: int(request.TaskId),
		ImgUrl: request.ImgUrl,
	}
	data, err := receiver.TaskSuccess(ctx, payload)
	if err != nil {
		return nil, err
	}
	return &pb.TaskSuccessRpcResponse{
		Data: data,
	}, nil
}
func (receiver *WordcloudTaskImpl) TaskFailRpc(ctx context.Context, request *pb.TaskFail) (*pb.TaskFailRpcResponse, error) {
	payload := vo.TaskFail{
		TaskId: int(request.TaskId),
		Error:  request.Error,
	}
	data, err := receiver.TaskFail(ctx, payload)
	if err != nil {
		return nil, err
	}
	return &pb.TaskFailRpcResponse{
		Data: data,
	}, nil
}
func (receiver *WordcloudTaskImpl) TaskPageRpc(ctx context.Context, request *pb.PageQuery) (*pb.TaskPageRet, error) {
	var payload vo.PageQuery
	copier.DeepCopy(request, &payload)
	//if request.Filter != nil {
	//	payload.Filter = vo.Filter{
	//		UserId: int(request.Filter.UserId),
	//	}
	//}
	//if request.Page != nil {
	//	payload.Page = vo.Page{
	//		PageNo: int(request.Page.PageNo),
	//		Size:   int(request.Page.Size),
	//	}
	//	for _, item := range request.Page.Orders {
	//		if item != nil {
	//			payload.Page.Orders = append(payload.Page.Orders, vo.Order{
	//				Col:  item.Col,
	//				Sort: item.Sort,
	//			})
	//		}
	//	}
	//}
	data, err := receiver.TaskPage(ctx, payload)
	if err != nil {
		return nil, err
	}
	var ret pb.TaskPageRet
	copier.DeepCopy(data, &ret)
	return &ret, nil
}
