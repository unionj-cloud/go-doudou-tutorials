package service

import "context"

//go:generate go-doudou svc http --handler --doc

type Annotation interface {
	// 此接口可公开访问，无需校验登录和权限
	GetGuest(ctx context.Context) (data string, err error)
	// 此接口只有登录用户有权访问
	// @role(USER,ADMIN)
	GetUser(ctx context.Context) (data string, err error)
	// 此接口只有管理员有权访问
	// @role(ADMIN)
	GetAdmin(ctx context.Context) (data string, err error)
}
