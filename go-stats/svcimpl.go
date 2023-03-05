/**
* Generated by go-doudou v2.0.3.
* You can edit it as your need.
 */
package service

import (
	"context"
	"go-doudou-tutorials/go-stats/config"
	pb "go-doudou-tutorials/go-stats/transport/grpc"
	"strconv"

	"github.com/shopspring/decimal"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/copier"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/numberutils"

	"go-doudou-tutorials/go-stats/vo"

	"github.com/brianvoe/gofakeit/v6"
)

var _ GoStats = (*GoStatsImpl)(nil)
var _ pb.GoStatsServiceServer = (*GoStatsImpl)(nil)

type GoStatsImpl struct {
	pb.UnimplementedGoStatsServiceServer
	conf *config.Config
}

func (receiver *GoStatsImpl) LargestRemainderRpc(ctx context.Context, request *pb.PercentageReqVo) (*pb.LargestRemainderRpcResponse, error) {
	var payload vo.PercentageReqVo
	copier.DeepCopy(request, &payload)
	data, err := receiver.LargestRemainder(ctx, payload)
	if err != nil {
		return nil, err
	}
	var ret pb.LargestRemainderRpcResponse
	err = copier.DeepCopy(data, &ret.Data)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

func NewGoStats(conf *config.Config) *GoStatsImpl {
	return &GoStatsImpl{
		conf: conf,
	}
}

func (receiver *GoStatsImpl) LargestRemainder(ctx context.Context, payload vo.PercentageReqVo) (data []vo.PercentageRespVo, err error) {
	if len(payload.Data) == 0 {
		return
	}
	input := make([]numberutils.Percentage, 0)
	for _, item := range payload.Data {
		input = append(input, numberutils.Percentage{
			Value: item.Value,
			Data:  item.Key,
		})
	}
	numberutils.LargestRemainder(input, int32(payload.Places))
	for _, item := range input {
		data = append(data, vo.PercentageRespVo{
			Value:            item.Value,
			Key:              item.Data.(string),
			Percent:          item.Percent,
			PercentFormatted: item.PercentFormatted,
		})
	}
	return
}

func (receiver *GoStatsImpl) GetShelves_ShelfBooks_Book(ctx context.Context, shelf int, book string) (data string, err error) {
	return strconv.Itoa(shelf) + ":" + book, err
}

func (receiver *GoStatsImpl) GetShelvesShelfBooksBookRpc(ctx context.Context, request *pb.GetShelvesShelfBooksBookRpcRequest) (*pb.GetShelvesShelfBooksBookRpcResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (receiver *GoStatsImpl) GetBook(ctx context.Context, price decimal.Decimal) (data string, err error) {
	var _result struct {
		Data string
	}
	_ = gofakeit.Struct(&_result)
	return _result.Data, nil
}