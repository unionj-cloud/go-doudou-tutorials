package service

import "context"

//go:generate go-doudou svc http --handler --doc

type Annotation interface {
	GetGuest(ctx context.Context) (data string, err error)
	// @role(USER,ADMIN)
	GetUser(ctx context.Context) (data string, err error)
	// @role(ADMIN)
	GetAdmin(ctx context.Context) (data string, err error)
}
