package dao

import (
	"context"
	"github.com/unionj-cloud/go-doudou/toolkit/sqlext/query"
)

type Base interface {
	Insert(ctx context.Context, data interface{}) (int64, error)
	Upsert(ctx context.Context, data interface{}) (int64, error)
	UpsertNoneZero(ctx context.Context, data interface{}) (int64, error)
	DeleteMany(ctx context.Context, where query.Q) (int64, error)
	Update(ctx context.Context, data interface{}) (int64, error)
	UpdateNoneZero(ctx context.Context, data interface{}) (int64, error)
	UpdateMany(ctx context.Context, data interface{}, where query.Q) (int64, error)
	UpdateManyNoneZero(ctx context.Context, data interface{}, where query.Q) (int64, error)
	Get(ctx context.Context, id interface{}) (interface{}, error)
	SelectMany(ctx context.Context, where ...query.Q) (interface{}, error)
	CountMany(ctx context.Context, where ...query.Q) (int, error)
	PageMany(ctx context.Context, page query.Page, where ...query.Q) (query.PageRet, error)
}
