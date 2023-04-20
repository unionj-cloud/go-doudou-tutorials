/**
* Generated by go-doudou v2.0.5.
* You can edit it as your need.
 */
package service

import (
	"context"
	"gormdemo/config"
	"gormdemo/dto"

	"github.com/brianvoe/gofakeit/v6"
)

var _ Gormdemo = (*GormdemoImpl)(nil)

type GormdemoImpl struct {
	conf *config.Config
}

func (receiver *GormdemoImpl) PageUsers(ctx context.Context, query dto.PageQuery) (data dto.PageRet, err error) {
	var _result struct {
		Data dto.PageRet
	}
	_ = gofakeit.Struct(&_result)
	return _result.Data, nil
}

func NewGormdemo(conf *config.Config) *GormdemoImpl {
	return &GormdemoImpl{
		conf: conf,
	}
}
