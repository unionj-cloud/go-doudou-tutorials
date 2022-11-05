package service

import (
	"context"
	"gateway/config"
	"gateway/vo"

	"github.com/brianvoe/gofakeit/v6"
)

type GatewayImpl struct {
	conf *config.Config
}

func (receiver *GatewayImpl) PageUsers(ctx context.Context, query vo.PageQuery) (code int, data vo.PageRet, err error) {
	var _result struct {
		Code int
		Data vo.PageRet
	}
	_ = gofakeit.Struct(&_result)
	return _result.Code, _result.Data, nil
}

func NewGateway(conf *config.Config) Gateway {
	return &GatewayImpl{
		conf,
	}
}
