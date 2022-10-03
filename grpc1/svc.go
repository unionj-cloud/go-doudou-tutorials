/**
* Generated by go-doudou v1.2.4.
* You can edit it as your need.
*/
package service

import (
	"context"
	"grpc1/vo"
)

//go:generate go-doudou svc http --handler -c --doc
//go:generate go-doudou svc grpc

type Grpc1 interface {
	// You can define your service methods as your need. Below is an example.
	// You can also add annotations here like @role(admin) to add meta data to routes for 
	// implementing your own middlewares
	PageUsers(ctx context.Context, query vo.PageQuery) (data vo.PageRet, err error)
}
