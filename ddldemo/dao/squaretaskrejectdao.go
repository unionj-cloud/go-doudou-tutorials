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

var _ ISquareTaskRejectDao = (*SquareTaskRejectDao)(nil)

type SquareTaskRejectDao struct {
	db wrapper.GddDB
}

func NewSquareTaskRejectDao(querier wrapper.GddDB) *SquareTaskRejectDao {
	return &SquareTaskRejectDao{
		db: querier,
	}
}

func (receiver *SquareTaskRejectDao) BeforeSaveHook(ctx context.Context, data *entity.SquareTaskReject) {
	// implement your business logic
}

func (receiver *SquareTaskRejectDao) BeforeBulkSaveHook(ctx context.Context, data []*entity.SquareTaskReject) {
	// implement your business logic
}

func (receiver *SquareTaskRejectDao) AfterSaveHook(ctx context.Context, data *entity.SquareTaskReject, lastInsertID int64, affected int64) {
	// implement your business logic
}

func (receiver *SquareTaskRejectDao) AfterBulkSaveHook(ctx context.Context, data []*entity.SquareTaskReject, lastInsertID int64, affected int64) {
	// implement your business logic
}

func (receiver *SquareTaskRejectDao) BeforeUpdateManyHook(ctx context.Context, data []*entity.SquareTaskReject, where *query.Where) {
	// implement your business logic
}

func (receiver *SquareTaskRejectDao) AfterUpdateManyHook(ctx context.Context, data []*entity.SquareTaskReject, where *query.Where, affected int64) {
	// implement your business logic
}

func (receiver *SquareTaskRejectDao) BeforeDeleteManyHook(ctx context.Context, data []*entity.SquareTaskReject, where *query.Where) {
	// implement your business logic
}

func (receiver *SquareTaskRejectDao) AfterDeleteManyHook(ctx context.Context, data []*entity.SquareTaskReject, where *query.Where, affected int64) {
	// implement your business logic
}

func (receiver *SquareTaskRejectDao) BeforeReadManyHook(ctx context.Context, page *query.Page, where *query.Where) {
	// implement your business logic
}

func (receiver *SquareTaskRejectDao) Insert(ctx context.Context, data *entity.SquareTaskReject) (int64, error) {
	var (
		statement    string
		err          error
		result       sql.Result
		affected     int64
	)
	receiver.BeforeSaveHook(ctx, data)
	if statement, err = templateutils.BlockMysql("squaretaskrejectdao.sql", squaretaskrejectdaosql, "InsertSquareTaskReject", nil); err != nil {
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

func (receiver *SquareTaskRejectDao) InsertIgnore(ctx context.Context, data *entity.SquareTaskReject) (int64, error) {
	var (
		statement    string
		err          error
		result       sql.Result
		affected     int64
	)
	receiver.BeforeSaveHook(ctx, data)
	if statement, err = templateutils.BlockMysql("squaretaskrejectdao.sql", squaretaskrejectdaosql, "InsertIgnoreSquareTaskReject", nil); err != nil {
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

func (receiver *SquareTaskRejectDao) BulkInsert(ctx context.Context, data []*entity.SquareTaskReject) (int64, error) {
	var (
		statement    string
		err          error
		result       sql.Result
		affected     int64
	)
	receiver.BeforeBulkSaveHook(ctx, data)
	if statement, err = templateutils.BlockMysql("squaretaskrejectdao.sql", squaretaskrejectdaosql, "InsertSquareTaskReject", nil); err != nil {
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

func (receiver *SquareTaskRejectDao) BulkInsertIgnore(ctx context.Context, data []*entity.SquareTaskReject) (int64, error) {
	var (
		statement    string
		err          error
		result       sql.Result
		affected     int64
	)
	receiver.BeforeBulkSaveHook(ctx, data)
	if statement, err = templateutils.BlockMysql("squaretaskrejectdao.sql", squaretaskrejectdaosql, "InsertIgnoreSquareTaskReject", nil); err != nil {
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
func (receiver *SquareTaskRejectDao) Upsert(ctx context.Context, data *entity.SquareTaskReject) (int64, error) {
	var (
		statement    string
		err          error
		result       sql.Result
		affected     int64
	)
	receiver.BeforeSaveHook(ctx, data)
	if statement, err = templateutils.BlockMysql("squaretaskrejectdao.sql", squaretaskrejectdaosql, "UpsertSquareTaskReject", nil); err != nil {
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

func (receiver *SquareTaskRejectDao) BulkUpsert(ctx context.Context, data []*entity.SquareTaskReject) (int64, error) {
	var (
		statement    string
		updateClause string
		err          error
		result       sql.Result
		affected     int64
		args      []interface{}
	)
	receiver.BeforeBulkSaveHook(ctx, data)
	if statement, err = templateutils.BlockMysql("squaretaskrejectdao.sql", squaretaskrejectdaosql, "InsertSquareTaskReject", nil); err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	statement, args, err = receiver.db.BindNamed(statement, data)
	if err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	if updateClause, err = templateutils.BlockMysql("squaretaskrejectdao.sql", squaretaskrejectdaosql, "UpdateClauseSquareTaskReject", nil); err != nil {
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

func (receiver *SquareTaskRejectDao) BulkUpsertSelect(ctx context.Context, data []*entity.SquareTaskReject, columns []string) (int64, error) {
	var (
		statement    string
		updateClause string
		err          error
		result       sql.Result
		affected     int64
		args      []interface{}
	)
	receiver.BeforeBulkSaveHook(ctx, data)
	if statement, err = templateutils.BlockMysql("squaretaskrejectdao.sql", squaretaskrejectdaosql, "InsertSquareTaskReject", nil); err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	statement, args, err = receiver.db.BindNamed(statement, data)
	if err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	if updateClause, err = templateutils.BlockMysql("squaretaskrejectdao.sql", squaretaskrejectdaosql, "UpdateClauseSelectSquareTaskReject", struct {
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

func (receiver *SquareTaskRejectDao) UpsertNoneZero(ctx context.Context, data *entity.SquareTaskReject) (int64, error) {
	var (
		statement    string
		err          error
		result       sql.Result
		affected     int64
	)
	receiver.BeforeSaveHook(ctx, data)
	value := reflectutils.ValueOf(data).Interface()
	if _, ok := value.(entity.SquareTaskReject); !ok {
		return 0, errors.New("underlying type of data should be entity.SquareTaskReject")
	}
	if statement, err = templateutils.BlockMysql("squaretaskrejectdao.sql", squaretaskrejectdaosql, "UpsertSquareTaskRejectNoneZero", data); err != nil {
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

func (receiver *SquareTaskRejectDao) DeleteMany(ctx context.Context, where query.Where) (int64, error) {
	var (
		err    error
		result sql.Result
		w      string
		args   []interface{}
		affected int64
	)
	receiver.BeforeDeleteManyHook(ctx, nil, &where)
	w, args = where.Sql()
	if result, err = receiver.db.ExecContext(ctx, receiver.db.Rebind(fmt.Sprintf("delete from cad_square_task_reject where %s;", w)), args...); err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	if affected, err = result.RowsAffected(); err == nil {
		receiver.AfterDeleteManyHook(ctx, nil, &where, affected)
	}
	return affected, err
}

func (receiver *SquareTaskRejectDao) Update(ctx context.Context, data *entity.SquareTaskReject) (int64, error) {
	var (
		statement string
		err       error
		result    sql.Result
		affected  int64
	)
	receiver.BeforeSaveHook(ctx, data)
	if statement, err = templateutils.BlockMysql("squaretaskrejectdao.sql", squaretaskrejectdaosql, "UpdateSquareTaskReject", nil); err != nil {
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

func (receiver *SquareTaskRejectDao) UpdateNoneZero(ctx context.Context, data *entity.SquareTaskReject) (int64, error) {
	var (
		statement string
		err       error
		result    sql.Result
		affected  int64
	)
	receiver.BeforeSaveHook(ctx, data)
	value := reflectutils.ValueOf(data).Interface()
	if _, ok := value.(entity.SquareTaskReject); !ok {
		return 0, errors.New("underlying type of data should be entity.SquareTaskReject")
	}
	if statement, err = templateutils.BlockMysql("squaretaskrejectdao.sql", squaretaskrejectdaosql, "UpdateSquareTaskRejectNoneZero", data); err != nil {
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

func (receiver *SquareTaskRejectDao) UpdateMany(ctx context.Context, data []*entity.SquareTaskReject, where query.Where) (int64, error) {
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
	if statement, err = templateutils.BlockMysql("squaretaskrejectdao.sql", squaretaskrejectdaosql, "UpdateSquareTaskRejects", nil); err != nil {
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

func (receiver *SquareTaskRejectDao) UpdateManyNoneZero(ctx context.Context, data []*entity.SquareTaskReject, where query.Where) (int64, error) {
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
	if _, ok := value.(entity.SquareTaskReject); !ok {
		return 0, errors.New("underlying type of data should be entity.SquareTaskReject")
	}
	if statement, err = templateutils.BlockMysql("squaretaskrejectdao.sql", squaretaskrejectdaosql, "UpdateSquareTaskRejectsNoneZero", data); err != nil {
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

func (receiver *SquareTaskRejectDao) Get(ctx context.Context, dest *entity.SquareTaskReject, id string) error {
	var (
		statement string
		err       error
		squaretaskreject      entity.SquareTaskReject
	)
	if statement, err = templateutils.BlockMysql("squaretaskrejectdao.sql", squaretaskrejectdaosql, "GetSquareTaskReject", nil); err != nil {
		return errors.Wrap(err, caller.NewCaller().String())
	}
	if err = receiver.db.GetContext(ctx, &squaretaskreject, receiver.db.Rebind(statement), id); err != nil {
		return errors.Wrap(err, caller.NewCaller().String())
	}
	return nil
}

func (receiver *SquareTaskRejectDao) SelectMany(ctx context.Context, dest *[]entity.SquareTaskReject, where query.Where) error {
	var (
		statements []string
		err       error
		args       []interface{}
	)
	receiver.BeforeReadManyHook(ctx, nil, &where)
    statements = append(statements, "select * from cad_square_task_reject")
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

func (receiver *SquareTaskRejectDao) CountMany(ctx context.Context, where query.Where) (int, error) {
	var (
		statements []string
		err       error
		total     int
		args       []interface{}
	)
	receiver.BeforeReadManyHook(ctx, nil, &where)
	statements = append(statements, "select count(1) from cad_square_task_reject")
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

type SquareTaskRejectPageRet struct {
	Items    []entity.SquareTaskReject
	PageNo   int
	PageSize int
	Total    int
	HasNext  bool
}

func (receiver *SquareTaskRejectDao) PageMany(ctx context.Context, dest *SquareTaskRejectPageRet, page query.Page, where query.Where) error {
	var (
		statements []string
		err       error
		args       []interface{}
	)
	receiver.BeforeReadManyHook(ctx, &page, &where)
	statements = append(statements, "select * from cad_square_task_reject")
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
	statements = append(statements, "select count(1) from cad_square_task_reject")
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

func (receiver *SquareTaskRejectDao) DeleteManySoft(ctx context.Context, where query.Where) (int64, error) {
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
	if result, err = receiver.db.ExecContext(ctx, receiver.db.Rebind(fmt.Sprintf("update cad_square_task_reject set delete_at=? where %s;", w)), args...); err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	if affected, err = result.RowsAffected(); err == nil {
		receiver.AfterDeleteManyHook(ctx, nil, &where, affected)
	}
	return affected, err
}