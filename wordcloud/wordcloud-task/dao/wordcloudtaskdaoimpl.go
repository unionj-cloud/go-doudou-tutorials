package dao

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-task/domain"
	"github.com/unionj-cloud/go-doudou/toolkit/sqlext/query"
	"github.com/unionj-cloud/go-doudou/toolkit/sqlext/wrapper"
	"github.com/unionj-cloud/go-doudou/toolkit/reflectutils"
	"github.com/unionj-cloud/go-doudou/toolkit/stringutils"
	"github.com/unionj-cloud/go-doudou/toolkit/templateutils"
	"strings"
	"math"
)

type WordCloudTaskDaoImpl struct {
	db wrapper.Querier
}

func NewWordCloudTaskDao(querier wrapper.Querier) WordCloudTaskDao {
	return WordCloudTaskDaoImpl{
		db: querier,
	}
}

func (receiver WordCloudTaskDaoImpl) Insert(ctx context.Context, data interface{}) (int64, error) {
	var (
		statement    string
		err          error
		result       sql.Result
		lastInsertID int64
	)
	if statement, err = templateutils.BlockMysql("wordcloudtaskdao.sql", wordcloudtaskdaosql, "InsertWordCloudTask", nil); err != nil {
		return 0, errors.Wrap(err, "")
	}
	if result, err = receiver.db.NamedExecContext(ctx, statement, data); err != nil {
		return 0, errors.Wrap(err, "")
	}
	if lastInsertID, err = result.LastInsertId(); err != nil {
		return 0, errors.Wrap(err, "")
	}
	if lastInsertID > 0 {
		if wordcloudtask, ok := data.(*domain.WordCloudTask); ok {
			wordcloudtask.Id = int(lastInsertID)
		}
	}
	return result.RowsAffected()
}

// Upsert With ON DUPLICATE KEY UPDATE, the affected-rows value per row is 1 if the row is inserted as a new row,
// 2 if an existing row is updated, and 0 if an existing row is set to its current values.
// If you specify the CLIENT_FOUND_ROWS flag to the mysql_real_connect() C API function when connecting to mysqld,
// the affected-rows value is 1 (not 0) if an existing row is set to its current values.
// https://dev.mysql.com/doc/refman/5.7/en/insert-on-duplicate.html
func (receiver WordCloudTaskDaoImpl) Upsert(ctx context.Context, data interface{}) (int64, error) {
	var (
		statement    string
		err          error
		result       sql.Result
		lastInsertID int64
	)
	if statement, err = templateutils.BlockMysql("wordcloudtaskdao.sql", wordcloudtaskdaosql, "UpsertWordCloudTask", nil); err != nil {
		return 0, errors.Wrap(err, "")
	}
	if result, err = receiver.db.NamedExecContext(ctx, statement, data); err != nil {
		return 0, errors.Wrap(err, "")
	}
	if lastInsertID, err = result.LastInsertId(); err != nil {
		return 0, errors.Wrap(err, "")
	}
	if lastInsertID > 0 {
		if wordcloudtask, ok := data.(*domain.WordCloudTask); ok {
			wordcloudtask.Id = int(lastInsertID)
		}
	}
	return result.RowsAffected()
}

func (receiver WordCloudTaskDaoImpl) UpsertNoneZero(ctx context.Context, data interface{}) (int64, error) {
	var (
		statement    string
		err          error
		result       sql.Result
		lastInsertID int64
	)
	value := reflectutils.ValueOf(data).Interface()
	if _, ok := value.(domain.WordCloudTask); !ok {
		return 0, errors.New("underlying type of data should be domain.WordCloudTask")
	}
	if statement, err = templateutils.BlockMysql("wordcloudtaskdao.sql", wordcloudtaskdaosql, "UpsertWordCloudTaskNoneZero", data); err != nil {
		return 0, errors.Wrap(err, "")
	}
	if result, err = receiver.db.NamedExecContext(ctx, statement, data); err != nil {
		return 0, errors.Wrap(err, "")
	}
	if lastInsertID, err = result.LastInsertId(); err != nil {
		return 0, errors.Wrap(err, "")
	}
	if lastInsertID > 0 {
		if wordcloudtask, ok := data.(*domain.WordCloudTask); ok {
			wordcloudtask.Id = int(lastInsertID)
		}
	}
	return result.RowsAffected()
}

func (receiver WordCloudTaskDaoImpl) DeleteMany(ctx context.Context, where query.Q) (int64, error) {
	var (
		err    error
		result sql.Result
		w      string
		args   []interface{}
	)
	w, args = where.Sql()
	if result, err = receiver.db.ExecContext(ctx, receiver.db.Rebind(fmt.Sprintf("delete from t_word_cloud_task where %s;", w)), args...); err != nil {
		return 0, errors.Wrap(err, "")
	}
	return result.RowsAffected()
}

func (receiver WordCloudTaskDaoImpl) Update(ctx context.Context, data interface{}) (int64, error) {
	var (
		statement string
		err       error
		result    sql.Result
	)
	if statement, err = templateutils.BlockMysql("wordcloudtaskdao.sql", wordcloudtaskdaosql, "UpdateWordCloudTask", nil); err != nil {
		return 0, errors.Wrap(err, "")
	}
	if result, err = receiver.db.NamedExecContext(ctx, statement, data); err != nil {
		return 0, errors.Wrap(err, "")
	}
	return result.RowsAffected()
}

func (receiver WordCloudTaskDaoImpl) UpdateNoneZero(ctx context.Context, data interface{}) (int64, error) {
	var (
		statement string
		err       error
		result    sql.Result
	)
	value := reflectutils.ValueOf(data).Interface()
	if _, ok := value.(domain.WordCloudTask); !ok {
		return 0, errors.New("underlying type of data should be domain.WordCloudTask")
	}
	if statement, err = templateutils.BlockMysql("wordcloudtaskdao.sql", wordcloudtaskdaosql, "UpdateWordCloudTaskNoneZero", data); err != nil {
		return 0, errors.Wrap(err, "")
	}
	if result, err = receiver.db.NamedExecContext(ctx, statement, data); err != nil {
		return 0, errors.Wrap(err, "")
	}
	return result.RowsAffected()
}

func (receiver WordCloudTaskDaoImpl) UpdateMany(ctx context.Context, data interface{}, where query.Q) (int64, error) {
	var (
		statement string
		err       error
		result    sql.Result
		q         string
		args      []interface{}
		wargs     []interface{}
		w         string
	)
	if statement, err = templateutils.BlockMysql("wordcloudtaskdao.sql", wordcloudtaskdaosql, "UpdateWordCloudTasks", nil); err != nil {
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

func (receiver WordCloudTaskDaoImpl) UpdateManyNoneZero(ctx context.Context, data interface{}, where query.Q) (int64, error) {
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
	if _, ok := value.(domain.WordCloudTask); !ok {
		return 0, errors.New("underlying type of data should be domain.WordCloudTask")
	}
	if statement, err = templateutils.BlockMysql("wordcloudtaskdao.sql", wordcloudtaskdaosql, "UpdateWordCloudTasksNoneZero", data); err != nil {
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

func (receiver WordCloudTaskDaoImpl) Get(ctx context.Context, id interface{}) (interface{}, error) {
	var (
		statement string
		err       error
		wordcloudtask      domain.WordCloudTask
	)
	if statement, err = templateutils.BlockMysql("wordcloudtaskdao.sql", wordcloudtaskdaosql, "GetWordCloudTask", nil); err != nil {
		return domain.WordCloudTask{}, errors.Wrap(err, "")
	}
	if err = receiver.db.GetContext(ctx, &wordcloudtask, receiver.db.Rebind(statement), id); err != nil {
		return domain.WordCloudTask{}, errors.Wrap(err, "")
	}
	return wordcloudtask, nil
}

func (receiver WordCloudTaskDaoImpl) SelectMany(ctx context.Context, where ...query.Q) (interface{}, error) {
	var (
		statements []string
		err       error
		wordcloudtasks     []domain.WordCloudTask
		args       []interface{}
	)
    statements = append(statements, "select * from t_word_cloud_task")
    if len(where) > 0 {
        statements = append(statements, "where")
        for _, item := range where {
			q, wargs := item.Sql()
			statements = append(statements, q)
			args = append(args, wargs...)
		}
    }
	if err = receiver.db.SelectContext(ctx, &wordcloudtasks, strings.Join(statements, " "), args...); err != nil {
		return nil, errors.Wrap(err, "")
	}
	return wordcloudtasks, nil
}

func (receiver WordCloudTaskDaoImpl) CountMany(ctx context.Context, where ...query.Q) (int, error) {
	var (
		statements []string
		err       error
		total     int
		args       []interface{}
	)
	statements = append(statements, "select count(1) from t_word_cloud_task")
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

func (receiver WordCloudTaskDaoImpl) PageMany(ctx context.Context, page query.Page, where ...query.Q) (query.PageRet, error) {
	var (
		statements []string
		err       error
		wordcloudtasks     []domain.WordCloudTask
		total     int
		args       []interface{}
	)
	statements = append(statements, "select * from t_word_cloud_task")
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
	if err = receiver.db.SelectContext(ctx, &wordcloudtasks, q, args...); err != nil {
		return query.PageRet{}, errors.Wrap(err, "")
	}
	
	args = nil
    statements = nil
	statements = append(statements, "select count(1) from t_word_cloud_task")
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
	pageRet.Items = wordcloudtasks
	pageRet.Total = total
	
	if pageRet.PageSize > 0 && math.Ceil(float64(total)/float64(pageRet.PageSize)) > float64(pageRet.PageNo) {
		pageRet.HasNext = true
	}

	return pageRet, nil
}