package service

import (
	"context"
	"testsvc/vo"
)

//go:generate go-doudou svc http --handler --doc

type Testsvc interface {
	PageUsers(ctx context.Context, query vo.PageQuery) (data vo.PageRet, err error)
	LogIn(ctx context.Context, payload *vo.LoginPayload) (data *vo.LoginResp, err error)
	LogIn2(ctx context.Context, payload vo.LoginPayload) (data *vo.LoginResp, err error)
}
