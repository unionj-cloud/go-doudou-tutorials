/**
* Generated by go-doudou v2.0.4.
* You can edit it as your need.
*/
package dao

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"ddldemo/entity"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/caller"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/sqlext/query"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/sqlext/wrapper"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/reflectutils"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/stringutils"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/templateutils"
	"strings"
	"math"
	"time"
)

var _ IOrderMainDao = (*OrderMainDao)(nil)

type OrderMainDao struct {
	db wrapper.GddDB
}

func NewOrderMainDao(querier wrapper.GddDB) *OrderMainDao {
	return &OrderMainDao{
		db: querier,
	}
}

func (receiver *OrderMainDao) BeforeSaveHook(ctx context.Context, data *entity.OrderMain) {
	// implement your business logic
}

func (receiver *OrderMainDao) BeforeBulkSaveHook(ctx context.Context, data []*entity.OrderMain) {
	// implement your business logic
}

func (receiver *OrderMainDao) AfterSaveHook(ctx context.Context, data *entity.OrderMain, lastInsertID int64, affected int64) {
	// implement your business logic
}

func (receiver *OrderMainDao) AfterBulkSaveHook(ctx context.Context, data []*entity.OrderMain, lastInsertID int64, affected int64) {
	// implement your business logic
}

func (receiver *OrderMainDao) BeforeUpdateManyHook(ctx context.Context, data []*entity.OrderMain, where *query.Where) {
	// implement your business logic
}

func (receiver *OrderMainDao) AfterUpdateManyHook(ctx context.Context, data []*entity.OrderMain, where *query.Where, affected int64) {
	// implement your business logic
}

func (receiver *OrderMainDao) BeforeDeleteManyHook(ctx context.Context, data []*entity.OrderMain, where *query.Where) {
	// implement your business logic
}

func (receiver *OrderMainDao) AfterDeleteManyHook(ctx context.Context, data []*entity.OrderMain, where *query.Where, affected int64) {
	// implement your business logic
}

func (receiver *OrderMainDao) BeforeReadManyHook(ctx context.Context, page *query.Page, where *query.Where) {
	// implement your business logic
}

func (receiver *OrderMainDao) Insert(ctx context.Context, data *entity.OrderMain) (int64, error) {
	var (
		statement    string
		err          error
		result       sql.Result
		affected     int64
	)
	receiver.BeforeSaveHook(ctx, data)
	if statement, err = templateutils.BlockMysql("ordermaindao.sql", ordermaindaosql, "InsertOrderMain", nil); err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	if result, err = receiver.db.NamedExecContext(ctx, statement, data); err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	if affected, err = result.RowsAffected(); err == nil {
		receiver.AfterSaveHook(ctx, data, 0, affected)
	}
	return affected, err
}

func (receiver *OrderMainDao) InsertIgnore(ctx context.Context, data *entity.OrderMain) (int64, error) {
	var (
		statement    string
		err          error
		result       sql.Result
		affected     int64
	)
	receiver.BeforeSaveHook(ctx, data)
	if statement, err = templateutils.BlockMysql("ordermaindao.sql", ordermaindaosql, "InsertIgnoreOrderMain", nil); err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	if result, err = receiver.db.NamedExecContext(ctx, statement, data); err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	if affected, err = result.RowsAffected(); err == nil {
		receiver.AfterSaveHook(ctx, data, 0, affected)
	}
	return affected, err
}

func (receiver *OrderMainDao) BulkInsert(ctx context.Context, data []*entity.OrderMain) (int64, error) {
	var (
		statement    string
		err          error
		result       sql.Result
		affected     int64
	)
	receiver.BeforeBulkSaveHook(ctx, data)
	if statement, err = templateutils.BlockMysql("ordermaindao.sql", ordermaindaosql, "InsertOrderMain", nil); err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	if result, err = receiver.db.NamedExecContext(ctx, statement, data); err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	if affected, err = result.RowsAffected(); err == nil {
		receiver.AfterBulkSaveHook(ctx, data, 0, affected)
	}
	return affected, err
}

func (receiver *OrderMainDao) BulkInsertIgnore(ctx context.Context, data []*entity.OrderMain) (int64, error) {
	var (
		statement    string
		err          error
		result       sql.Result
		affected     int64
	)
	receiver.BeforeBulkSaveHook(ctx, data)
	if statement, err = templateutils.BlockMysql("ordermaindao.sql", ordermaindaosql, "InsertIgnoreOrderMain", nil); err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	if result, err = receiver.db.NamedExecContext(ctx, statement, data); err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	if affected, err = result.RowsAffected(); err == nil {
		receiver.AfterBulkSaveHook(ctx, data, 0, affected)
	}
	return affected, err
}

// Upsert With ON DUPLICATE KEY UPDATE, the affected-rows value per row is 1 if the row is inserted as a new row,
// 2 if an existing row is updated, and 0 if an existing row is set to its current values.
// If you specify the CLIENT_FOUND_ROWS flag to the mysql_real_connect() C API function when connecting to mysqld,
// the affected-rows value is 1 (not 0) if an existing row is set to its current values.
// https://dev.mysql.com/doc/refman/5.7/en/insert-on-duplicate.html
func (receiver *OrderMainDao) Upsert(ctx context.Context, data *entity.OrderMain) (int64, error) {
	var (
		statement    string
		err          error
		result       sql.Result
		affected     int64
	)
	receiver.BeforeSaveHook(ctx, data)
	if statement, err = templateutils.BlockMysql("ordermaindao.sql", ordermaindaosql, "UpsertOrderMain", nil); err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	if result, err = receiver.db.NamedExecContext(ctx, statement, data); err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	if affected, err = result.RowsAffected(); err == nil {
		receiver.AfterSaveHook(ctx, data, 0, affected)
	}
	return affected, err
}

func (receiver *OrderMainDao) BulkUpsert(ctx context.Context, data []*entity.OrderMain) (int64, error) {
	var (
		statement    string
		updateClause string
		err          error
		result       sql.Result
		affected     int64
		args      []interface{}
	)
	receiver.BeforeBulkSaveHook(ctx, data)
	if statement, err = templateutils.BlockMysql("ordermaindao.sql", ordermaindaosql, "InsertOrderMain", nil); err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	statement, args, err = receiver.db.BindNamed(statement, data)
	if err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	if updateClause, err = templateutils.BlockMysql("ordermaindao.sql", ordermaindaosql, "UpdateClauseOrderMain", nil); err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	statement += "\n" + updateClause
	if result, err = receiver.db.ExecContext(ctx, statement, args...); err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	if affected, err = result.RowsAffected(); err == nil {
		receiver.AfterBulkSaveHook(ctx, data, 0, affected)
	}
	return affected, err
}

func (receiver *OrderMainDao) BulkUpsertSelect(ctx context.Context, data []*entity.OrderMain, columns []string) (int64, error) {
	var (
		statement    string
		updateClause string
		err          error
		result       sql.Result
		affected     int64
		args      []interface{}
	)
	receiver.BeforeBulkSaveHook(ctx, data)
	if statement, err = templateutils.BlockMysql("ordermaindao.sql", ordermaindaosql, "InsertOrderMain", nil); err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	statement, args, err = receiver.db.BindNamed(statement, data)
	if err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	if updateClause, err = templateutils.BlockMysql("ordermaindao.sql", ordermaindaosql, "UpdateClauseSelectOrderMain", struct {
		Columns []string
	}{
		Columns: columns,
	}); err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	statement += "\n" + updateClause
	if result, err = receiver.db.ExecContext(ctx, statement, args...); err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	if affected, err = result.RowsAffected(); err == nil {
		receiver.AfterBulkSaveHook(ctx, data, 0, affected)
	}
	return affected, err
}

func (receiver *OrderMainDao) UpsertNoneZero(ctx context.Context, data *entity.OrderMain) (int64, error) {
	var (
		statement    string
		err          error
		result       sql.Result
		affected     int64
	)
	receiver.BeforeSaveHook(ctx, data)
	value := reflectutils.ValueOf(data).Interface()
	if _, ok := value.(entity.OrderMain); !ok {
		return 0, errors.New("underlying type of data should be entity.OrderMain")
	}
	if statement, err = templateutils.BlockMysql("ordermaindao.sql", ordermaindaosql, "UpsertOrderMainNoneZero", data); err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	if result, err = receiver.db.NamedExecContext(ctx, statement, data); err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	if affected, err = result.RowsAffected(); err == nil {
		receiver.AfterSaveHook(ctx, data, 0, affected)
	}
	return affected, err
}

func (receiver *OrderMainDao) DeleteMany(ctx context.Context, where query.Where) (int64, error) {
	var (
		err    error
		result sql.Result
		w      string
		args   []interface{}
		affected int64
	)
	receiver.BeforeDeleteManyHook(ctx, nil, &where)
	w, args = where.Sql()
	if result, err = receiver.db.ExecContext(ctx, receiver.db.Rebind(fmt.Sprintf("delete from ces_order_main where %s;", w)), args...); err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	if affected, err = result.RowsAffected(); err == nil {
		receiver.AfterDeleteManyHook(ctx, nil, &where, affected)
	}
	return affected, err
}

func (receiver *OrderMainDao) Update(ctx context.Context, data *entity.OrderMain) (int64, error) {
	var (
		statement string
		err       error
		result    sql.Result
		affected  int64
	)
	receiver.BeforeSaveHook(ctx, data)
	if statement, err = templateutils.BlockMysql("ordermaindao.sql", ordermaindaosql, "UpdateOrderMain", nil); err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	if result, err = receiver.db.NamedExecContext(ctx, statement, data); err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	if affected, err = result.RowsAffected(); err == nil {
		receiver.AfterSaveHook(ctx, data, 0, affected)
	}
	return affected, err
}

func (receiver *OrderMainDao) UpdateNoneZero(ctx context.Context, data *entity.OrderMain) (int64, error) {
	var (
		statement string
		err       error
		result    sql.Result
		affected  int64
	)
	receiver.BeforeSaveHook(ctx, data)
	value := reflectutils.ValueOf(data).Interface()
	if _, ok := value.(entity.OrderMain); !ok {
		return 0, errors.New("underlying type of data should be entity.OrderMain")
	}
	if statement, err = templateutils.BlockMysql("ordermaindao.sql", ordermaindaosql, "UpdateOrderMainNoneZero", data); err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	if result, err = receiver.db.NamedExecContext(ctx, statement, data); err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	if affected, err = result.RowsAffected(); err == nil {
		receiver.AfterSaveHook(ctx, data, 0, affected)
	}
	return affected, err
}

func (receiver *OrderMainDao) UpdateMany(ctx context.Context, data []*entity.OrderMain, where query.Where) (int64, error) {
	var (
		statement string
		err       error
		result    sql.Result
		q         string
		args      []interface{}
		wargs     []interface{}
		w         string
		affected  int64
	)
	receiver.BeforeUpdateManyHook(ctx, data, &where)
	if statement, err = templateutils.BlockMysql("ordermaindao.sql", ordermaindaosql, "UpdateOrderMains", nil); err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	if q, args, err = receiver.db.BindNamed(statement, data); err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	w, wargs = where.Sql()
	if stringutils.IsNotEmpty(w) {
		q += " where " + w
	}
	args = append(args, wargs...)
	if result, err = receiver.db.ExecContext(ctx, receiver.db.Rebind(q), args...); err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	if affected, err = result.RowsAffected(); err == nil {
		receiver.AfterUpdateManyHook(ctx, data, &where, affected)
	}
	return affected, err
}

func (receiver *OrderMainDao) UpdateManyNoneZero(ctx context.Context, data []*entity.OrderMain, where query.Where) (int64, error) {
	var (
		statement string
		err       error
		result    sql.Result
		q         string
		args      []interface{}
		wargs     []interface{}
		w         string
		affected  int64
	)
	receiver.BeforeUpdateManyHook(ctx, data, &where)
	value := reflectutils.ValueOf(data).Interface()
	if _, ok := value.(entity.OrderMain); !ok {
		return 0, errors.New("underlying type of data should be entity.OrderMain")
	}
	if statement, err = templateutils.BlockMysql("ordermaindao.sql", ordermaindaosql, "UpdateOrderMainsNoneZero", data); err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	if q, args, err = receiver.db.BindNamed(statement, data); err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	w, wargs = where.Sql()
	if stringutils.IsNotEmpty(w) {
		q += " where " + w
	}
	args = append(args, wargs...)
	if result, err = receiver.db.ExecContext(ctx, receiver.db.Rebind(q), args...); err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	if affected, err = result.RowsAffected(); err == nil {
		receiver.AfterUpdateManyHook(ctx, data, &where, affected)
	}
	return affected, err
}

func (receiver *OrderMainDao) Get(ctx context.Context, dest *entity.OrderMain, id string) error {
	var (
		statement string
		err       error
		ordermain      entity.OrderMain
	)
	if statement, err = templateutils.BlockMysql("ordermaindao.sql", ordermaindaosql, "GetOrderMain", nil); err != nil {
		return errors.Wrap(err, caller.NewCaller().String())
	}
	if err = receiver.db.GetContext(ctx, &ordermain, receiver.db.Rebind(statement), id); err != nil {
		return errors.Wrap(err, caller.NewCaller().String())
	}
	return nil
}

func (receiver *OrderMainDao) SelectMany(ctx context.Context, dest *[]entity.OrderMain, where query.Where) error {
	var (
		statements []string
		err       error
		args       []interface{}
	)
	receiver.BeforeReadManyHook(ctx, nil, &where)
    statements = append(statements, "select * from ces_order_main")
	if !where.IsEmpty() {
		statements = append(statements, "where")
		q, wargs := where.Sql()
		statements = append(statements, q)
		args = append(args, wargs...)
	}
	sqlStr := strings.TrimSpace(strings.TrimSuffix(strings.TrimSpace(strings.Join(statements, " ")), "where"))
	if err = receiver.db.SelectContext(ctx, dest, receiver.db.Rebind(sqlStr), args...); err != nil {
		return errors.Wrap(err, caller.NewCaller().String())
	}
	return nil
}

func (receiver *OrderMainDao) CountMany(ctx context.Context, where query.Where) (int, error) {
	var (
		statements []string
		err       error
		total     int
		args       []interface{}
	)
	receiver.BeforeReadManyHook(ctx, nil, &where)
	statements = append(statements, "select count(1) from ces_order_main")
    if !where.IsEmpty() {
		statements = append(statements, "where")
		q, wargs := where.Sql()
		statements = append(statements, q)
		args = append(args, wargs...)
	}
	sqlStr := strings.TrimSpace(strings.TrimSuffix(strings.TrimSpace(strings.Join(statements, " ")), "where"))
	if err = receiver.db.GetContext(ctx, &total, receiver.db.Rebind(sqlStr), args...); err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	return total, nil
}

type OrderMainPageRet struct {
	Items    []entity.OrderMain
	PageNo   int
	PageSize int
	Total    int
	HasNext  bool
}

func (receiver *OrderMainDao) PageMany(ctx context.Context, dest *OrderMainPageRet, page query.Page, where query.Where) error {
	var (
		statements []string
		err       error
		args       []interface{}
	)
	receiver.BeforeReadManyHook(ctx, &page, &where)
	statements = append(statements, "select * from ces_order_main")
    if !where.IsEmpty() {
		statements = append(statements, "where")
		q, wargs := where.Sql()
		statements = append(statements, q)
		args = append(args, wargs...)
	}
	p, pargs := page.Sql()
	statements = append(statements, p)
	args = append(args, pargs...)
	sqlStr := strings.TrimSpace(strings.TrimSuffix(strings.TrimSpace(strings.Join(statements, " ")), "where"))
	if err = receiver.db.SelectContext(ctx, &dest.Items, receiver.db.Rebind(sqlStr), args...); err != nil {
		return errors.Wrap(err, caller.NewCaller().String())
	}
	
	args = nil
    statements = nil
	statements = append(statements, "select count(1) from ces_order_main")
    if !where.IsEmpty() {
		statements = append(statements, "where")
		q, wargs := where.Sql()
		statements = append(statements, q)
		args = append(args, wargs...)
	}
	sqlStr = strings.TrimSpace(strings.TrimSuffix(strings.TrimSpace(strings.Join(statements, " ")), "where"))
	if err = receiver.db.GetContext(ctx, &dest.Total, receiver.db.Rebind(sqlStr), args...); err != nil {
		return errors.Wrap(err, caller.NewCaller().String())
	}

	pageNo := 1
	if page.Size > 0 {
		pageNo = page.Offset/page.Size + 1
	}
	dest.PageNo = pageNo
	dest.PageSize = page.Size
	if dest.PageSize > 0 && math.Ceil(float64(dest.Total)/float64(dest.PageSize)) > float64(dest.PageNo) {
		dest.HasNext = true
	}
	return nil
}

func (receiver *OrderMainDao) DeleteManySoft(ctx context.Context, where query.Where) (int64, error) {
	var (
		err      error
		result   sql.Result
		w        string
		args     []interface{}
		affected int64
	)
	receiver.BeforeDeleteManyHook(ctx, nil, &where)
	w, args = where.Sql()
	args = append([]interface{}{time.Now()}, args...)
	if result, err = receiver.db.ExecContext(ctx, receiver.db.Rebind(fmt.Sprintf("update ces_order_main set delete_at=? where %s;", w)), args...); err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	if affected, err = result.RowsAffected(); err == nil {
		receiver.AfterDeleteManyHook(ctx, nil, &where, affected)
	}
	return affected, err
}