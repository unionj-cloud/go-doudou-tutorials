package dao

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-user/domain"
	"github.com/unionj-cloud/go-doudou/toolkit/sqlext/query"
	"github.com/unionj-cloud/go-doudou/toolkit/sqlext/wrapper"
	"github.com/unionj-cloud/go-doudou/toolkit/reflectutils"
	"github.com/unionj-cloud/go-doudou/toolkit/stringutils"
	"github.com/unionj-cloud/go-doudou/toolkit/templateutils"
	"strings"
	"math"
)

type InvalidTokenDaoImpl struct {
	db wrapper.Querier
}

func NewInvalidTokenDao(querier wrapper.Querier) InvalidTokenDao {
	return InvalidTokenDaoImpl{
		db: querier,
	}
}

func (receiver InvalidTokenDaoImpl) Insert(ctx context.Context, data interface{}) (int64, error) {
	var (
		statement    string
		err          error
		result       sql.Result
		lastInsertID int64
	)
	if statement, err = templateutils.BlockMysql("invalidtokendao.sql", invalidtokendaosql, "InsertInvalidToken", nil); err != nil {
		return 0, errors.Wrap(err, "")
	}
	if result, err = receiver.db.NamedExecContext(ctx, statement, data); err != nil {
		return 0, errors.Wrap(err, "")
	}
	if lastInsertID, err = result.LastInsertId(); err != nil {
		return 0, errors.Wrap(err, "")
	}
	if lastInsertID > 0 {
		if invalidtoken, ok := data.(*domain.InvalidToken); ok {
			invalidtoken.Id = int(lastInsertID)
		}
	}
	return result.RowsAffected()
}

// Upsert With ON DUPLICATE KEY UPDATE, the affected-rows value per row is 1 if the row is inserted as a new row,
// 2 if an existing row is updated, and 0 if an existing row is set to its current values.
// If you specify the CLIENT_FOUND_ROWS flag to the mysql_real_connect() C API function when connecting to mysqld,
// the affected-rows value is 1 (not 0) if an existing row is set to its current values.
// https://dev.mysql.com/doc/refman/5.7/en/insert-on-duplicate.html
func (receiver InvalidTokenDaoImpl) Upsert(ctx context.Context, data interface{}) (int64, error) {
	var (
		statement    string
		err          error
		result       sql.Result
		lastInsertID int64
	)
	if statement, err = templateutils.BlockMysql("invalidtokendao.sql", invalidtokendaosql, "UpsertInvalidToken", nil); err != nil {
		return 0, errors.Wrap(err, "")
	}
	if result, err = receiver.db.NamedExecContext(ctx, statement, data); err != nil {
		return 0, errors.Wrap(err, "")
	}
	if lastInsertID, err = result.LastInsertId(); err != nil {
		return 0, errors.Wrap(err, "")
	}
	if lastInsertID > 0 {
		if invalidtoken, ok := data.(*domain.InvalidToken); ok {
			invalidtoken.Id = int(lastInsertID)
		}
	}
	return result.RowsAffected()
}

func (receiver InvalidTokenDaoImpl) UpsertNoneZero(ctx context.Context, data interface{}) (int64, error) {
	var (
		statement    string
		err          error
		result       sql.Result
		lastInsertID int64
	)
	value := reflectutils.ValueOf(data).Interface()
	if _, ok := value.(domain.InvalidToken); !ok {
		return 0, errors.New("underlying type of data should be domain.InvalidToken")
	}
	if statement, err = templateutils.BlockMysql("invalidtokendao.sql", invalidtokendaosql, "UpsertInvalidTokenNoneZero", data); err != nil {
		return 0, errors.Wrap(err, "")
	}
	if result, err = receiver.db.NamedExecContext(ctx, statement, data); err != nil {
		return 0, errors.Wrap(err, "")
	}
	if lastInsertID, err = result.LastInsertId(); err != nil {
		return 0, errors.Wrap(err, "")
	}
	if lastInsertID > 0 {
		if invalidtoken, ok := data.(*domain.InvalidToken); ok {
			invalidtoken.Id = int(lastInsertID)
		}
	}
	return result.RowsAffected()
}

func (receiver InvalidTokenDaoImpl) DeleteMany(ctx context.Context, where query.Q) (int64, error) {
	var (
		err    error
		result sql.Result
		w      string
		args   []interface{}
	)
	w, args = where.Sql()
	if result, err = receiver.db.ExecContext(ctx, receiver.db.Rebind(fmt.Sprintf("delete from t_invalid_token where %s;", w)), args...); err != nil {
		return 0, errors.Wrap(err, "")
	}
	return result.RowsAffected()
}

func (receiver InvalidTokenDaoImpl) Update(ctx context.Context, data interface{}) (int64, error) {
	var (
		statement string
		err       error
		result    sql.Result
	)
	if statement, err = templateutils.BlockMysql("invalidtokendao.sql", invalidtokendaosql, "UpdateInvalidToken", nil); err != nil {
		return 0, errors.Wrap(err, "")
	}
	if result, err = receiver.db.NamedExecContext(ctx, statement, data); err != nil {
		return 0, errors.Wrap(err, "")
	}
	return result.RowsAffected()
}

func (receiver InvalidTokenDaoImpl) UpdateNoneZero(ctx context.Context, data interface{}) (int64, error) {
	var (
		statement string
		err       error
		result    sql.Result
	)
	value := reflectutils.ValueOf(data).Interface()
	if _, ok := value.(domain.InvalidToken); !ok {
		return 0, errors.New("underlying type of data should be domain.InvalidToken")
	}
	if statement, err = templateutils.BlockMysql("invalidtokendao.sql", invalidtokendaosql, "UpdateInvalidTokenNoneZero", data); err != nil {
		return 0, errors.Wrap(err, "")
	}
	if result, err = receiver.db.NamedExecContext(ctx, statement, data); err != nil {
		return 0, errors.Wrap(err, "")
	}
	return result.RowsAffected()
}

func (receiver InvalidTokenDaoImpl) UpdateMany(ctx context.Context, data interface{}, where query.Q) (int64, error) {
	var (
		statement string
		err       error
		result    sql.Result
		q         string
		args      []interface{}
		wargs     []interface{}
		w         string
	)
	if statement, err = templateutils.BlockMysql("invalidtokendao.sql", invalidtokendaosql, "UpdateInvalidTokens", nil); err != nil {
		return 0, errors.Wrap(err, "")
	}
	if q, args, err = receiver.db.BindNamed(statement, data); err != nil {
		return 0, errors.Wrap(err, "")
	}
	w, wargs = where.Sql()
	if stringutils.IsNotEmpty(w) {
		q += " where " + w
	}
	args = append(args, wargs...)
	if result, err = receiver.db.ExecContext(ctx, receiver.db.Rebind(q), args...); err != nil {
		return 0, errors.Wrap(err, "")
	}
	return result.RowsAffected()
}

func (receiver InvalidTokenDaoImpl) UpdateManyNoneZero(ctx context.Context, data interface{}, where query.Q) (int64, error) {
	var (
		statement string
		err       error
		result    sql.Result
		q         string
		args      []interface{}
		wargs     []interface{}
		w         string
	)
	value := reflectutils.ValueOf(data).Interface()
	if _, ok := value.(domain.InvalidToken); !ok {
		return 0, errors.New("underlying type of data should be domain.InvalidToken")
	}
	if statement, err = templateutils.BlockMysql("invalidtokendao.sql", invalidtokendaosql, "UpdateInvalidTokensNoneZero", data); err != nil {
		return 0, errors.Wrap(err, "")
	}
	if q, args, err = receiver.db.BindNamed(statement, data); err != nil {
		return 0, errors.Wrap(err, "")
	}
	w, wargs = where.Sql()
	if stringutils.IsNotEmpty(w) {
		q += " where " + w
	}
	args = append(args, wargs...)
	if result, err = receiver.db.ExecContext(ctx, receiver.db.Rebind(q), args...); err != nil {
		return 0, errors.Wrap(err, "")
	}
	return result.RowsAffected()
}

func (receiver InvalidTokenDaoImpl) Get(ctx context.Context, id interface{}) (interface{}, error) {
	var (
		statement string
		err       error
		invalidtoken      domain.InvalidToken
	)
	if statement, err = templateutils.BlockMysql("invalidtokendao.sql", invalidtokendaosql, "GetInvalidToken", nil); err != nil {
		return domain.InvalidToken{}, errors.Wrap(err, "")
	}
	if err = receiver.db.GetContext(ctx, &invalidtoken, receiver.db.Rebind(statement), id); err != nil {
		return domain.InvalidToken{}, errors.Wrap(err, "")
	}
	return invalidtoken, nil
}

func (receiver InvalidTokenDaoImpl) SelectMany(ctx context.Context, where ...query.Q) (interface{}, error) {
	var (
		statements []string
		err       error
		invalidtokens     []domain.InvalidToken
		args       []interface{}
	)
    statements = append(statements, "select * from t_invalid_token")
    if len(where) > 0 {
        statements = append(statements, "where")
        for _, item := range where {
			q, wargs := item.Sql()
			statements = append(statements, q)
			args = append(args, wargs...)
		}
    }
	if err = receiver.db.SelectContext(ctx, &invalidtokens, strings.Join(statements, " "), args...); err != nil {
		return nil, errors.Wrap(err, "")
	}
	return invalidtokens, nil
}

func (receiver InvalidTokenDaoImpl) CountMany(ctx context.Context, where ...query.Q) (int, error) {
	var (
		statements []string
		err       error
		total     int
		args       []interface{}
	)
	statements = append(statements, "select count(1) from t_invalid_token")
    if len(where) > 0 {
        statements = append(statements, "where")
        for _, item := range where {
			q, wargs := item.Sql()
			statements = append(statements, q)
			args = append(args, wargs...)
		}
    }
	if err = receiver.db.GetContext(ctx, &total, strings.Join(statements, " "), args...); err != nil {
		return 0, errors.Wrap(err, "")
	}
	return total, nil
}

func (receiver InvalidTokenDaoImpl) PageMany(ctx context.Context, page query.Page, where ...query.Q) (query.PageRet, error) {
	var (
		statements []string
		err       error
		invalidtokens     []domain.InvalidToken
		total     int
		args       []interface{}
	)
	statements = append(statements, "select * from t_invalid_token")
    if len(where) > 0 {
        statements = append(statements, "where")
        for _, item := range where {
			q, wargs := item.Sql()
			statements = append(statements, q)
			args = append(args, wargs...)
		}
    }
	p, pargs := page.Sql()
	statements = append(statements, p)
	args = append(args, pargs...)
	q := strings.Join(statements, " ")
	if err = receiver.db.SelectContext(ctx, &invalidtokens, q, args...); err != nil {
		return query.PageRet{}, errors.Wrap(err, "")
	}
	
	args = nil
    statements = nil
	statements = append(statements, "select count(1) from t_invalid_token")
    if len(where) > 0 {
        statements = append(statements, "where")
        for _, item := range where {
			q, wargs := item.Sql()
			statements = append(statements, q)
			args = append(args, wargs...)
		}
    }
	if err = receiver.db.GetContext(ctx, &total, strings.Join(statements, " "), args...); err != nil {
		return query.PageRet{}, errors.Wrap(err, "")
	}

	pageRet := query.NewPageRet(page)
	pageRet.Items = invalidtokens
	pageRet.Total = total
	
	if pageRet.PageSize > 0 && math.Ceil(float64(total)/float64(pageRet.PageSize)) > float64(pageRet.PageNo) {
		pageRet.HasNext = true
	}

	return pageRet, nil
}