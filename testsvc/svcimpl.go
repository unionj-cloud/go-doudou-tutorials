package service

import (
	"context"
	"fmt"
	"testsvc/config"
	"testsvc/vo"
)

type TestsvcImpl struct {
	conf *config.Config
}

func (receiver *TestsvcImpl) PageUsers(ctx context.Context, query vo.PageQuery) (data vo.PageRet, err error) {
	data.PageSize = query.Page.Size
	return
}

func NewTestsvc(conf *config.Config) Testsvc {
	return &TestsvcImpl{
		conf,
	}
}

func (receiver *TestsvcImpl) LogIn(ctx context.Context, payload *vo.LoginPayload) (data *vo.LoginResp, err error) {
	if payload != nil {
		data = &vo.LoginResp{}
		data.Token = fmt.Sprintf("%s:%s", payload.Username, payload.Password)
	}
	return
}

func (receiver *TestsvcImpl) LogIn2(ctx context.Context, payload vo.LoginPayload) (data *vo.LoginResp, err error) {
	data = &vo.LoginResp{}
	data.Token = fmt.Sprintf("%s:%s", payload.Username, payload.Password)
	return
}
