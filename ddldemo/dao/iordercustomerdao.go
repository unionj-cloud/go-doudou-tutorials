/**
* Generated by go-doudou v2.0.4.
* You can edit it as your need.
*/
package dao

import (
	"context"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/sqlext/query"
	"ddldemo/entity"
)

type IOrderCustomerDao interface {
	// single table CRUD operations
	Insert(ctx context.Context, data *entity.OrderCustomer) (int64, error)
	InsertIgnore(ctx context.Context, data *entity.OrderCustomer) (int64, error)
	BulkInsert(ctx context.Context, data []*entity.OrderCustomer) (int64, error)
	BulkInsertIgnore(ctx context.Context, data []*entity.OrderCustomer) (int64, error)
	Upsert(ctx context.Context, data *entity.OrderCustomer) (int64, error)
	BulkUpsert(ctx context.Context, data []*entity.OrderCustomer) (int64, error)
	BulkUpsertSelect(ctx context.Context, data []*entity.OrderCustomer, columns []string) (int64, error)
	UpsertNoneZero(ctx context.Context, data *entity.OrderCustomer) (int64, error)
	DeleteMany(ctx context.Context, where query.Where) (int64, error)
	Update(ctx context.Context, data *entity.OrderCustomer) (int64, error)
	UpdateNoneZero(ctx context.Context, data *entity.OrderCustomer) (int64, error)
	UpdateMany(ctx context.Context, data []*entity.OrderCustomer, where query.Where) (int64, error)
	UpdateManyNoneZero(ctx context.Context, data []*entity.OrderCustomer, where query.Where) (int64, error)
	Get(ctx context.Context, dest *entity.OrderCustomer, id string) error
	SelectMany(ctx context.Context, dest *[]entity.OrderCustomer, where query.Where) error
	CountMany(ctx context.Context, where query.Where) (int, error)
	PageMany(ctx context.Context, dest *OrderCustomerPageRet, page query.Page, where query.Where) error
	DeleteManySoft(ctx context.Context, where query.Where) (int64, error)

	// hooks
	BeforeSaveHook(ctx context.Context, data *entity.OrderCustomer)
	BeforeBulkSaveHook(ctx context.Context, data []*entity.OrderCustomer)
	AfterSaveHook(ctx context.Context, data *entity.OrderCustomer, lastInsertID int64, affected int64)
	AfterBulkSaveHook(ctx context.Context, data []*entity.OrderCustomer, lastInsertID int64, affected int64)
	BeforeUpdateManyHook(ctx context.Context, data []*entity.OrderCustomer, where *query.Where)
	AfterUpdateManyHook(ctx context.Context, data []*entity.OrderCustomer, where *query.Where, affected int64)
	BeforeDeleteManyHook(ctx context.Context, data []*entity.OrderCustomer, where *query.Where)
	AfterDeleteManyHook(ctx context.Context, data []*entity.OrderCustomer, where *query.Where, affected int64)
	BeforeReadManyHook(ctx context.Context, page *query.Page, where *query.Where)
}