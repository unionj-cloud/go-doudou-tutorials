/**
* Generated by go-doudou v2.0.1.
* You can edit it as your need.
 */
package dao

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-user/domain"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/caller"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/reflectutils"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/sqlext/query"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/sqlext/wrapper"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/stringutils"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/templateutils"
	"math"
	"strings"
	"time"
)

type UserDao struct {
	db wrapper.GddDB
}

func NewUserDao(querier wrapper.GddDB) UserDao {
	return UserDao{
		db: querier,
	}
}

func (receiver UserDao) BeforeSaveHook(ctx context.Context, data *domain.User) {
	// implement your business logic
}

func (receiver UserDao) BeforeBulkSaveHook(ctx context.Context, data []*domain.User) {
	// implement your business logic
}

func (receiver UserDao) AfterSaveHook(ctx context.Context, data *domain.User, lastInsertID int64, affected int64) {
	// implement your business logic
}

func (receiver UserDao) AfterBulkSaveHook(ctx context.Context, data []*domain.User, lastInsertID int64, affected int64) {
	// implement your business logic
}

func (receiver UserDao) BeforeUpdateManyHook(ctx context.Context, data []*domain.User, where *query.Where) {
	// implement your business logic
}

func (receiver UserDao) AfterUpdateManyHook(ctx context.Context, data []*domain.User, where *query.Where, affected int64) {
	// implement your business logic
}

func (receiver UserDao) BeforeDeleteManyHook(ctx context.Context, data []*domain.User, where *query.Where) {
	// implement your business logic
}

func (receiver UserDao) AfterDeleteManyHook(ctx context.Context, data []*domain.User, where *query.Where, affected int64) {
	// implement your business logic
}

func (receiver UserDao) BeforeReadManyHook(ctx context.Context, page *query.Page, where *query.Where) {
	// implement your business logic
	*where = where.And(query.C().Col("delete_at").IsNull())
}

func (receiver UserDao) Insert(ctx context.Context, data *domain.User) (int64, error) {
	var (
		statement    string
		err          error
		result       sql.Result
		lastInsertID int64
		affected     int64
	)
	receiver.BeforeSaveHook(ctx, data)
	if statement, err = templateutils.BlockMysql("userdao.sql", userdaosql, "InsertUser", nil); err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	if result, err = receiver.db.NamedExecContext(ctx, statement, data); err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	if lastInsertID, err = result.LastInsertId(); err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	if lastInsertID > 0 {
		data.Id = int(lastInsertID)
	}
	if affected, err = result.RowsAffected(); err == nil {
		receiver.AfterSaveHook(ctx, data, lastInsertID, affected)
	}
	return affected, err
}

func (receiver UserDao) InsertIgnore(ctx context.Context, data *domain.User) (int64, error) {
	var (
		statement string
		err       error
		result    sql.Result
		affected  int64
	)
	receiver.BeforeSaveHook(ctx, data)
	if statement, err = templateutils.BlockMysql("userdao.sql", userdaosql, "InsertIgnoreUser", nil); err != nil {
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

func (receiver UserDao) BulkInsert(ctx context.Context, data []*domain.User) (int64, error) {
	var (
		statement    string
		err          error
		result       sql.Result
		lastInsertID int64
		affected     int64
	)
	receiver.BeforeBulkSaveHook(ctx, data)
	if statement, err = templateutils.BlockMysql("userdao.sql", userdaosql, "InsertUser", nil); err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	if result, err = receiver.db.NamedExecContext(ctx, statement, data); err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	if lastInsertID, err = result.LastInsertId(); err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	if lastInsertID > 0 {
		for i, item := range data {
			item.Id = int(lastInsertID) + int(i)
		}
	}
	if affected, err = result.RowsAffected(); err == nil {
		receiver.AfterBulkSaveHook(ctx, data, lastInsertID, affected)
	}
	return affected, err
}

func (receiver UserDao) BulkInsertIgnore(ctx context.Context, data []*domain.User) (int64, error) {
	var (
		statement string
		err       error
		result    sql.Result
		affected  int64
	)
	receiver.BeforeBulkSaveHook(ctx, data)
	if statement, err = templateutils.BlockMysql("userdao.sql", userdaosql, "InsertIgnoreUser", nil); err != nil {
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
func (receiver UserDao) Upsert(ctx context.Context, data *domain.User) (int64, error) {
	var (
		statement    string
		err          error
		result       sql.Result
		lastInsertID int64
		affected     int64
	)
	receiver.BeforeSaveHook(ctx, data)
	if statement, err = templateutils.BlockMysql("userdao.sql", userdaosql, "UpsertUser", nil); err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	if result, err = receiver.db.NamedExecContext(ctx, statement, data); err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	if lastInsertID, err = result.LastInsertId(); err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	if lastInsertID > 0 {
		data.Id = int(lastInsertID)
	}
	if affected, err = result.RowsAffected(); err == nil {
		receiver.AfterSaveHook(ctx, data, lastInsertID, affected)
	}
	return affected, err
}

func (receiver UserDao) BulkUpsert(ctx context.Context, data []*domain.User) (int64, error) {
	var (
		statement    string
		updateClause string
		err          error
		result       sql.Result
		affected     int64
		args         []interface{}
	)
	receiver.BeforeBulkSaveHook(ctx, data)
	if statement, err = templateutils.BlockMysql("userdao.sql", userdaosql, "InsertUser", nil); err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	statement, args, err = receiver.db.BindNamed(statement, data)
	if err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	if updateClause, err = templateutils.BlockMysql("userdao.sql", userdaosql, "UpdateClauseUser", nil); err != nil {
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

func (receiver UserDao) BulkUpsertSelect(ctx context.Context, data []*domain.User, columns []string) (int64, error) {
	var (
		statement    string
		updateClause string
		err          error
		result       sql.Result
		affected     int64
		args         []interface{}
	)
	receiver.BeforeBulkSaveHook(ctx, data)
	if statement, err = templateutils.BlockMysql("userdao.sql", userdaosql, "InsertUser", nil); err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	statement, args, err = receiver.db.BindNamed(statement, data)
	if err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	if updateClause, err = templateutils.BlockMysql("userdao.sql", userdaosql, "UpdateClauseSelectUser", struct {
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

func (receiver UserDao) UpsertNoneZero(ctx context.Context, data *domain.User) (int64, error) {
	var (
		statement    string
		err          error
		result       sql.Result
		lastInsertID int64
		affected     int64
	)
	receiver.BeforeSaveHook(ctx, data)
	value := reflectutils.ValueOf(data).Interface()
	if _, ok := value.(domain.User); !ok {
		return 0, errors.New("underlying type of data should be domain.User")
	}
	if statement, err = templateutils.BlockMysql("userdao.sql", userdaosql, "UpsertUserNoneZero", data); err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	if result, err = receiver.db.NamedExecContext(ctx, statement, data); err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	if lastInsertID, err = result.LastInsertId(); err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	if lastInsertID > 0 {
		data.Id = int(lastInsertID)
	}
	if affected, err = result.RowsAffected(); err == nil {
		receiver.AfterSaveHook(ctx, data, lastInsertID, affected)
	}
	return affected, err
}

func (receiver UserDao) DeleteMany(ctx context.Context, where query.Where) (int64, error) {
	var (
		err      error
		result   sql.Result
		w        string
		args     []interface{}
		affected int64
	)
	receiver.BeforeDeleteManyHook(ctx, nil, &where)
	w, args = where.Sql()
	if result, err = receiver.db.ExecContext(ctx, receiver.db.Rebind(fmt.Sprintf("delete from t_user where %s;", w)), args...); err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	if affected, err = result.RowsAffected(); err == nil {
		receiver.AfterDeleteManyHook(ctx, nil, &where, affected)
	}
	return affected, err
}

func (receiver UserDao) Update(ctx context.Context, data *domain.User) (int64, error) {
	var (
		statement string
		err       error
		result    sql.Result
		affected  int64
	)
	receiver.BeforeSaveHook(ctx, data)
	if statement, err = templateutils.BlockMysql("userdao.sql", userdaosql, "UpdateUser", nil); err != nil {
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

func (receiver UserDao) UpdateNoneZero(ctx context.Context, data *domain.User) (int64, error) {
	var (
		statement string
		err       error
		result    sql.Result
		affected  int64
	)
	receiver.BeforeSaveHook(ctx, data)
	value := reflectutils.ValueOf(data).Interface()
	if _, ok := value.(domain.User); !ok {
		return 0, errors.New("underlying type of data should be domain.User")
	}
	if statement, err = templateutils.BlockMysql("userdao.sql", userdaosql, "UpdateUserNoneZero", data); err != nil {
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

func (receiver UserDao) UpdateMany(ctx context.Context, data []*domain.User, where query.Where) (int64, error) {
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
	if statement, err = templateutils.BlockMysql("userdao.sql", userdaosql, "UpdateUsers", nil); err != nil {
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

func (receiver UserDao) UpdateManyNoneZero(ctx context.Context, data []*domain.User, where query.Where) (int64, error) {
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
	if _, ok := value.(domain.User); !ok {
		return 0, errors.New("underlying type of data should be domain.User")
	}
	if statement, err = templateutils.BlockMysql("userdao.sql", userdaosql, "UpdateUsersNoneZero", data); err != nil {
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

func (receiver UserDao) Get(ctx context.Context, dest *domain.User, id int) error {
	var (
		statement string
		err       error
		user      domain.User
	)
	if statement, err = templateutils.BlockMysql("userdao.sql", userdaosql, "GetUser", nil); err != nil {
		return errors.Wrap(err, caller.NewCaller().String())
	}
	if err = receiver.db.GetContext(ctx, &user, receiver.db.Rebind(statement), id); err != nil {
		return errors.Wrap(err, caller.NewCaller().String())
	}
	return nil
}

func (receiver UserDao) SelectMany(ctx context.Context, dest *[]domain.User, where query.Where) error {
	var (
		statements []string
		err        error
		args       []interface{}
	)
	receiver.BeforeReadManyHook(ctx, nil, &where)
	statements = append(statements, "select * from t_user")
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

func (receiver UserDao) CountMany(ctx context.Context, where query.Where) (int, error) {
	var (
		statements []string
		err        error
		total      int
		args       []interface{}
	)
	receiver.BeforeReadManyHook(ctx, nil, &where)
	statements = append(statements, "select count(1) from t_user")
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

type UserPageRet struct {
	Items    []domain.User
	PageNo   int
	PageSize int
	Total    int
	HasNext  bool
}

func (receiver UserDao) PageMany(ctx context.Context, dest *UserPageRet, page query.Page, where query.Where) error {
	var (
		statements []string
		err        error
		args       []interface{}
	)
	receiver.BeforeReadManyHook(ctx, &page, &where)
	statements = append(statements, "select * from t_user")
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
	statements = append(statements, "select count(1) from t_user")
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

func (receiver UserDao) DeleteManySoft(ctx context.Context, where query.Where) (int64, error) {
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
	if result, err = receiver.db.ExecContext(ctx, receiver.db.Rebind(fmt.Sprintf("update t_user set delete_at=? where %s;", w)), args...); err != nil {
		return 0, errors.Wrap(err, caller.NewCaller().String())
	}
	if affected, err = result.RowsAffected(); err == nil {
		receiver.AfterDeleteManyHook(ctx, nil, &where, affected)
	}
	return affected, err
}
