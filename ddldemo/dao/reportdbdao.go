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

var _ IReportDbDao = (*ReportDbDao)(nil)

type ReportDbDao struct {
	db wrapper.GddDB
}

func NewReportDbDao(querier wrapper.GddDB) *ReportDbDao {
	return &ReportDbDao{
		db: querier,
	}
}

func (receiver *ReportDbDao) BeforeSaveHook(ctx context.Context, data *entity.ReportDb) {
	// implement your business logic
}

func (receiver *ReportDbDao) BeforeBulkSaveHook(ctx context.Context, data []*entity.ReportDb) {
	// implement your business logic
}

func (receiver *ReportDbDao) AfterSaveHook(ctx context.Context, data *entity.ReportDb, lastInsertID int64, affected int64) {
	// implement your business logic
}

func (receiver *ReportDbDao) AfterBulkSaveHook(ctx context.Context, data []*entity.ReportDb, lastInsertID int64, affected int64) {
	// implement your business logic
}

func (receiver *ReportDbDao) BeforeUpdateManyHook(ctx context.Context, data []*entity.ReportDb, where *query.Where) {
	// implement your business logic
}

func (receiver *ReportDbDao) AfterUpdateManyHook(ctx context.Context, data []*entity.ReportDb, where *query.Where, affected int64) {
	// implement your business logic
}

func (receiver *ReportDbDao) BeforeDeleteManyHook(ctx context.Context, data []*entity.ReportDb, where *query.Where) {
	// implement your business logic
}

func (receiver *ReportDbDao) AfterDeleteManyHook(ctx context.Context, data []*entity.ReportDb, where *query.Where, affected int64) {
	// implement your business logic
}

func (receiver *ReportDbDao) BeforeReadManyHook(ctx context.Context, page *query.Page, where *query.Where) {
	// implement your business logic
}

func (receiver *ReportDbDao) Insert(ctx context.Context, data *entity.ReportDb) (int64, error) {
	var (
		statement    string
		err          error
		result       sql.Result
		affected     int64
	)
	receiver.BeforeSaveHook(ctx, data)
	if statement, err = templateutils.BlockMysql("reportdbdao.sql", reportdbdaosql, "InsertReportDb", nil); err != nil {
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

func (receiver *ReportDbDao) InsertIgnore(ctx context.Context, data *entity.ReportDb) (int64, error) {
	var (
		statement    string
		err          error
		result       sql.Result
		affected     int64
	)
	receiver.BeforeSaveHook(ctx, data)
	if statement, err = templateutils.BlockMysql("reportdbdao.sql", reportdbdaosql, "InsertIgnoreReportDb", nil); err != nil {
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

func (receiver *ReportDbDao) BulkInsert(ctx context.Context, data []*entity.ReportDb) (int64, error) {
	var (
		statement    string
		err          error
		result       sql.Result
		affected     int64
	)
	receiver.BeforeBulkSaveHook(ctx, data)
	if statement, err = templateutils.BlockMysql("reportdbdao.sql", reportdbdaosql, "InsertReportDb", nil); err != nil {
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

func (receiver *ReportDbDao) BulkInsertIgnore(ctx context.Context, data []*entity.ReportDb) (int64, error) {
	var (
		statement    string
		err          error
		result       sql.Result
		affected     int64
	)
	receiver.BeforeBulkSaveHook(ctx, data)
	if statement, err = templateutils.BlockMysql("reportdbdao.sql", reportdbdaosql, "InsertIgnoreReportDb", nil); err != nil {
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
func (receiver *ReportDbDao) Upsert(ctx context.Context, data *entity.ReportDb) (int64, error) {
	var (
		statement    string
		err          error
		result       sql.Result
		affected     int64
	)
	receiver.BeforeSaveHook(ctx, data)
	if statement, err = templateutils.BlockMysql("reportdbdao.sql", reportdbdaosql, "UpsertReportDb", nil); err != nil {
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

func (receiver *ReportDbDao) BulkUpsert(ctx context.Context, data []*entity.ReportDb) (int64, error) {
	var (
		statement    string
		updateClause string
		err          error
		result       sql.Result
		affected     int64
		args      []interface{}
	)
	receiver.BeforeBulkSaveHook(ctx, data)
	if statement, err = templateutils.BlockMysql("reportdbdao.sql", reportdbdaosql, "InsertReportDb", nil); err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	statement, args, err = receiver.db.BindNamed(statement, data)
	if err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	if updateClause, err = templateutils.BlockMysql("reportdbdao.sql", reportdbdaosql, "UpdateClauseReportDb", nil); err != nil {
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

func (receiver *ReportDbDao) BulkUpsertSelect(ctx context.Context, data []*entity.ReportDb, columns []string) (int64, error) {
	var (
		statement    string
		updateClause string
		err          error
		result       sql.Result
		affected     int64
		args      []interface{}
	)
	receiver.BeforeBulkSaveHook(ctx, data)
	if statement, err = templateutils.BlockMysql("reportdbdao.sql", reportdbdaosql, "InsertReportDb", nil); err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	statement, args, err = receiver.db.BindNamed(statement, data)
	if err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	if updateClause, err = templateutils.BlockMysql("reportdbdao.sql", reportdbdaosql, "UpdateClauseSelectReportDb", struct {
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

func (receiver *ReportDbDao) UpsertNoneZero(ctx context.Context, data *entity.ReportDb) (int64, error) {
	var (
		statement    string
		err          error
		result       sql.Result
		affected     int64
	)
	receiver.BeforeSaveHook(ctx, data)
	value := reflectutils.ValueOf(data).Interface()
	if _, ok := value.(entity.ReportDb); !ok {
		return 0, errors.New("underlying type of data should be entity.ReportDb")
	}
	if statement, err = templateutils.BlockMysql("reportdbdao.sql", reportdbdaosql, "UpsertReportDbNoneZero", data); err != nil {
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

func (receiver *ReportDbDao) DeleteMany(ctx context.Context, where query.Where) (int64, error) {
	var (
		err    error
		result sql.Result
		w      string
		args   []interface{}
		affected int64
	)
	receiver.BeforeDeleteManyHook(ctx, nil, &where)
	w, args = where.Sql()
	if result, err = receiver.db.ExecContext(ctx, receiver.db.Rebind(fmt.Sprintf("delete from jimu_report_db where %s;", w)), args...); err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	if affected, err = result.RowsAffected(); err == nil {
		receiver.AfterDeleteManyHook(ctx, nil, &where, affected)
	}
	return affected, err
}

func (receiver *ReportDbDao) Update(ctx context.Context, data *entity.ReportDb) (int64, error) {
	var (
		statement string
		err       error
		result    sql.Result
		affected  int64
	)
	receiver.BeforeSaveHook(ctx, data)
	if statement, err = templateutils.BlockMysql("reportdbdao.sql", reportdbdaosql, "UpdateReportDb", nil); err != nil {
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

func (receiver *ReportDbDao) UpdateNoneZero(ctx context.Context, data *entity.ReportDb) (int64, error) {
	var (
		statement string
		err       error
		result    sql.Result
		affected  int64
	)
	receiver.BeforeSaveHook(ctx, data)
	value := reflectutils.ValueOf(data).Interface()
	if _, ok := value.(entity.ReportDb); !ok {
		return 0, errors.New("underlying type of data should be entity.ReportDb")
	}
	if statement, err = templateutils.BlockMysql("reportdbdao.sql", reportdbdaosql, "UpdateReportDbNoneZero", data); err != nil {
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

func (receiver *ReportDbDao) UpdateMany(ctx context.Context, data []*entity.ReportDb, where query.Where) (int64, error) {
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
	if statement, err = templateutils.BlockMysql("reportdbdao.sql", reportdbdaosql, "UpdateReportDbs", nil); err != nil {
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

func (receiver *ReportDbDao) UpdateManyNoneZero(ctx context.Context, data []*entity.ReportDb, where query.Where) (int64, error) {
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
	if _, ok := value.(entity.ReportDb); !ok {
		return 0, errors.New("underlying type of data should be entity.ReportDb")
	}
	if statement, err = templateutils.BlockMysql("reportdbdao.sql", reportdbdaosql, "UpdateReportDbsNoneZero", data); err != nil {
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

func (receiver *ReportDbDao) Get(ctx context.Context, dest *entity.ReportDb, id string) error {
	var (
		statement string
		err       error
		reportdb      entity.ReportDb
	)
	if statement, err = templateutils.BlockMysql("reportdbdao.sql", reportdbdaosql, "GetReportDb", nil); err != nil {
		return errors.Wrap(err, caller.NewCaller().String())
	}
	if err = receiver.db.GetContext(ctx, &reportdb, receiver.db.Rebind(statement), id); err != nil {
		return errors.Wrap(err, caller.NewCaller().String())
	}
	return nil
}

func (receiver *ReportDbDao) SelectMany(ctx context.Context, dest *[]entity.ReportDb, where query.Where) error {
	var (
		statements []string
		err       error
		args       []interface{}
	)
	receiver.BeforeReadManyHook(ctx, nil, &where)
    statements = append(statements, "select * from jimu_report_db")
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

func (receiver *ReportDbDao) CountMany(ctx context.Context, where query.Where) (int, error) {
	var (
		statements []string
		err       error
		total     int
		args       []interface{}
	)
	receiver.BeforeReadManyHook(ctx, nil, &where)
	statements = append(statements, "select count(1) from jimu_report_db")
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

type ReportDbPageRet struct {
	Items    []entity.ReportDb
	PageNo   int
	PageSize int
	Total    int
	HasNext  bool
}

func (receiver *ReportDbDao) PageMany(ctx context.Context, dest *ReportDbPageRet, page query.Page, where query.Where) error {
	var (
		statements []string
		err       error
		args       []interface{}
	)
	receiver.BeforeReadManyHook(ctx, &page, &where)
	statements = append(statements, "select * from jimu_report_db")
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
	statements = append(statements, "select count(1) from jimu_report_db")
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

func (receiver *ReportDbDao) DeleteManySoft(ctx context.Context, where query.Where) (int64, error) {
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
	if result, err = receiver.db.ExecContext(ctx, receiver.db.Rebind(fmt.Sprintf("update jimu_report_db set delete_at=? where %s;", w)), args...); err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	if affected, err = result.RowsAffected(); err == nil {
		receiver.AfterDeleteManyHook(ctx, nil, &where, affected)
	}
	return affected, err
}