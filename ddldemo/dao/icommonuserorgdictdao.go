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

type ICommonUserOrgDictDao interface {
	// single table CRUD operations
	Insert(ctx context.Context, data *entity.CommonUserOrgDict) (int64, error)
	InsertIgnore(ctx context.Context, data *entity.CommonUserOrgDict) (int64, error)
	BulkInsert(ctx context.Context, data []*entity.CommonUserOrgDict) (int64, error)
	BulkInsertIgnore(ctx context.Context, data []*entity.CommonUserOrgDict) (int64, error)
	Upsert(ctx context.Context, data *entity.CommonUserOrgDict) (int64, error)
	BulkUpsert(ctx context.Context, data []*entity.CommonUserOrgDict) (int64, error)
	BulkUpsertSelect(ctx context.Context, data []*entity.CommonUserOrgDict, columns []string) (int64, error)
	UpsertNoneZero(ctx context.Context, data *entity.CommonUserOrgDict) (int64, error)
	DeleteMany(ctx context.Context, where query.Where) (int64, error)
	Update(ctx context.Context, data *entity.CommonUserOrgDict) (int64, error)
	UpdateNoneZero(ctx context.Context, data *entity.CommonUserOrgDict) (int64, error)
	UpdateMany(ctx context.Context, data []*entity.CommonUserOrgDict, where query.Where) (int64, error)
	UpdateManyNoneZero(ctx context.Context, data []*entity.CommonUserOrgDict, where query.Where) (int64, error)
	Get(ctx context.Context, dest *entity.CommonUserOrgDict, id string) error
	SelectMany(ctx context.Context, dest *[]entity.CommonUserOrgDict, where query.Where) error
	CountMany(ctx context.Context, where query.Where) (int, error)
	PageMany(ctx context.Context, dest *CommonUserOrgDictPageRet, page query.Page, where query.Where) error
	DeleteManySoft(ctx context.Context, where query.Where) (int64, error)

	// hooks
	BeforeSaveHook(ctx context.Context, data *entity.CommonUserOrgDict)
	BeforeBulkSaveHook(ctx context.Context, data []*entity.CommonUserOrgDict)
	AfterSaveHook(ctx context.Context, data *entity.CommonUserOrgDict, lastInsertID int64, affected int64)
	AfterBulkSaveHook(ctx context.Context, data []*entity.CommonUserOrgDict, lastInsertID int64, affected int64)
	BeforeUpdateManyHook(ctx context.Context, data []*entity.CommonUserOrgDict, where *query.Where)
	AfterUpdateManyHook(ctx context.Context, data []*entity.CommonUserOrgDict, where *query.Where, affected int64)
	BeforeDeleteManyHook(ctx context.Context, data []*entity.CommonUserOrgDict, where *query.Where)
	AfterDeleteManyHook(ctx context.Context, data []*entity.CommonUserOrgDict, where *query.Where, affected int64)
	BeforeReadManyHook(ctx context.Context, page *query.Page, where *query.Where)
}