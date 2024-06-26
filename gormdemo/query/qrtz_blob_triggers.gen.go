// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"gormdemo/model"
)

func newQrtzBlobTrigger(db *gorm.DB, opts ...gen.DOOption) qrtzBlobTrigger {
	_qrtzBlobTrigger := qrtzBlobTrigger{}

	_qrtzBlobTrigger.qrtzBlobTriggerDo.UseDB(db, opts...)
	_qrtzBlobTrigger.qrtzBlobTriggerDo.UseModel(&model.QrtzBlobTrigger{})

	tableName := _qrtzBlobTrigger.qrtzBlobTriggerDo.TableName()
	_qrtzBlobTrigger.ALL = field.NewAsterisk(tableName)
	_qrtzBlobTrigger.SCHEDNAME = field.NewString(tableName, "SCHED_NAME")
	_qrtzBlobTrigger.TRIGGERNAME = field.NewString(tableName, "TRIGGER_NAME")
	_qrtzBlobTrigger.TRIGGERGROUP = field.NewString(tableName, "TRIGGER_GROUP")
	_qrtzBlobTrigger.BLOBDATA = field.NewBytes(tableName, "BLOB_DATA")

	_qrtzBlobTrigger.fillFieldMap()

	return _qrtzBlobTrigger
}

type qrtzBlobTrigger struct {
	qrtzBlobTriggerDo qrtzBlobTriggerDo

	ALL          field.Asterisk
	SCHEDNAME    field.String
	TRIGGERNAME  field.String
	TRIGGERGROUP field.String
	BLOBDATA     field.Bytes

	fieldMap map[string]field.Expr
}

func (q qrtzBlobTrigger) Table(newTableName string) *qrtzBlobTrigger {
	q.qrtzBlobTriggerDo.UseTable(newTableName)
	return q.updateTableName(newTableName)
}

func (q qrtzBlobTrigger) As(alias string) *qrtzBlobTrigger {
	q.qrtzBlobTriggerDo.DO = *(q.qrtzBlobTriggerDo.As(alias).(*gen.DO))
	return q.updateTableName(alias)
}

func (q *qrtzBlobTrigger) updateTableName(table string) *qrtzBlobTrigger {
	q.ALL = field.NewAsterisk(table)
	q.SCHEDNAME = field.NewString(table, "SCHED_NAME")
	q.TRIGGERNAME = field.NewString(table, "TRIGGER_NAME")
	q.TRIGGERGROUP = field.NewString(table, "TRIGGER_GROUP")
	q.BLOBDATA = field.NewBytes(table, "BLOB_DATA")

	q.fillFieldMap()

	return q
}

func (q *qrtzBlobTrigger) WithContext(ctx context.Context) IQrtzBlobTriggerDo {
	return q.qrtzBlobTriggerDo.WithContext(ctx)
}

func (q qrtzBlobTrigger) TableName() string { return q.qrtzBlobTriggerDo.TableName() }

func (q qrtzBlobTrigger) Alias() string { return q.qrtzBlobTriggerDo.Alias() }

func (q qrtzBlobTrigger) Columns(cols ...field.Expr) gen.Columns {
	return q.qrtzBlobTriggerDo.Columns(cols...)
}

func (q *qrtzBlobTrigger) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := q.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (q *qrtzBlobTrigger) fillFieldMap() {
	q.fieldMap = make(map[string]field.Expr, 4)
	q.fieldMap["SCHED_NAME"] = q.SCHEDNAME
	q.fieldMap["TRIGGER_NAME"] = q.TRIGGERNAME
	q.fieldMap["TRIGGER_GROUP"] = q.TRIGGERGROUP
	q.fieldMap["BLOB_DATA"] = q.BLOBDATA
}

func (q qrtzBlobTrigger) clone(db *gorm.DB) qrtzBlobTrigger {
	q.qrtzBlobTriggerDo.ReplaceConnPool(db.Statement.ConnPool)
	return q
}

func (q qrtzBlobTrigger) replaceDB(db *gorm.DB) qrtzBlobTrigger {
	q.qrtzBlobTriggerDo.ReplaceDB(db)
	return q
}

type qrtzBlobTriggerDo struct{ gen.DO }

type IQrtzBlobTriggerDo interface {
	gen.SubQuery
	Debug() IQrtzBlobTriggerDo
	WithContext(ctx context.Context) IQrtzBlobTriggerDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IQrtzBlobTriggerDo
	WriteDB() IQrtzBlobTriggerDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IQrtzBlobTriggerDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IQrtzBlobTriggerDo
	Not(conds ...gen.Condition) IQrtzBlobTriggerDo
	Or(conds ...gen.Condition) IQrtzBlobTriggerDo
	Select(conds ...field.Expr) IQrtzBlobTriggerDo
	Where(conds ...gen.Condition) IQrtzBlobTriggerDo
	Order(conds ...field.Expr) IQrtzBlobTriggerDo
	Distinct(cols ...field.Expr) IQrtzBlobTriggerDo
	Omit(cols ...field.Expr) IQrtzBlobTriggerDo
	Join(table schema.Tabler, on ...field.Expr) IQrtzBlobTriggerDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IQrtzBlobTriggerDo
	RightJoin(table schema.Tabler, on ...field.Expr) IQrtzBlobTriggerDo
	Group(cols ...field.Expr) IQrtzBlobTriggerDo
	Having(conds ...gen.Condition) IQrtzBlobTriggerDo
	Limit(limit int) IQrtzBlobTriggerDo
	Offset(offset int) IQrtzBlobTriggerDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IQrtzBlobTriggerDo
	Unscoped() IQrtzBlobTriggerDo
	Create(values ...*model.QrtzBlobTrigger) error
	CreateInBatches(values []*model.QrtzBlobTrigger, batchSize int) error
	Save(values ...*model.QrtzBlobTrigger) error
	First() (*model.QrtzBlobTrigger, error)
	Take() (*model.QrtzBlobTrigger, error)
	Last() (*model.QrtzBlobTrigger, error)
	Find() ([]*model.QrtzBlobTrigger, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.QrtzBlobTrigger, err error)
	FindInBatches(result *[]*model.QrtzBlobTrigger, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.QrtzBlobTrigger) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IQrtzBlobTriggerDo
	Assign(attrs ...field.AssignExpr) IQrtzBlobTriggerDo
	Joins(fields ...field.RelationField) IQrtzBlobTriggerDo
	Preload(fields ...field.RelationField) IQrtzBlobTriggerDo
	FirstOrInit() (*model.QrtzBlobTrigger, error)
	FirstOrCreate() (*model.QrtzBlobTrigger, error)
	FindByPage(offset int, limit int) (result []*model.QrtzBlobTrigger, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IQrtzBlobTriggerDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (q qrtzBlobTriggerDo) Debug() IQrtzBlobTriggerDo {
	return q.withDO(q.DO.Debug())
}

func (q qrtzBlobTriggerDo) WithContext(ctx context.Context) IQrtzBlobTriggerDo {
	return q.withDO(q.DO.WithContext(ctx))
}

func (q qrtzBlobTriggerDo) ReadDB() IQrtzBlobTriggerDo {
	return q.Clauses(dbresolver.Read)
}

func (q qrtzBlobTriggerDo) WriteDB() IQrtzBlobTriggerDo {
	return q.Clauses(dbresolver.Write)
}

func (q qrtzBlobTriggerDo) Session(config *gorm.Session) IQrtzBlobTriggerDo {
	return q.withDO(q.DO.Session(config))
}

func (q qrtzBlobTriggerDo) Clauses(conds ...clause.Expression) IQrtzBlobTriggerDo {
	return q.withDO(q.DO.Clauses(conds...))
}

func (q qrtzBlobTriggerDo) Returning(value interface{}, columns ...string) IQrtzBlobTriggerDo {
	return q.withDO(q.DO.Returning(value, columns...))
}

func (q qrtzBlobTriggerDo) Not(conds ...gen.Condition) IQrtzBlobTriggerDo {
	return q.withDO(q.DO.Not(conds...))
}

func (q qrtzBlobTriggerDo) Or(conds ...gen.Condition) IQrtzBlobTriggerDo {
	return q.withDO(q.DO.Or(conds...))
}

func (q qrtzBlobTriggerDo) Select(conds ...field.Expr) IQrtzBlobTriggerDo {
	return q.withDO(q.DO.Select(conds...))
}

func (q qrtzBlobTriggerDo) Where(conds ...gen.Condition) IQrtzBlobTriggerDo {
	return q.withDO(q.DO.Where(conds...))
}

func (q qrtzBlobTriggerDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IQrtzBlobTriggerDo {
	return q.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (q qrtzBlobTriggerDo) Order(conds ...field.Expr) IQrtzBlobTriggerDo {
	return q.withDO(q.DO.Order(conds...))
}

func (q qrtzBlobTriggerDo) Distinct(cols ...field.Expr) IQrtzBlobTriggerDo {
	return q.withDO(q.DO.Distinct(cols...))
}

func (q qrtzBlobTriggerDo) Omit(cols ...field.Expr) IQrtzBlobTriggerDo {
	return q.withDO(q.DO.Omit(cols...))
}

func (q qrtzBlobTriggerDo) Join(table schema.Tabler, on ...field.Expr) IQrtzBlobTriggerDo {
	return q.withDO(q.DO.Join(table, on...))
}

func (q qrtzBlobTriggerDo) LeftJoin(table schema.Tabler, on ...field.Expr) IQrtzBlobTriggerDo {
	return q.withDO(q.DO.LeftJoin(table, on...))
}

func (q qrtzBlobTriggerDo) RightJoin(table schema.Tabler, on ...field.Expr) IQrtzBlobTriggerDo {
	return q.withDO(q.DO.RightJoin(table, on...))
}

func (q qrtzBlobTriggerDo) Group(cols ...field.Expr) IQrtzBlobTriggerDo {
	return q.withDO(q.DO.Group(cols...))
}

func (q qrtzBlobTriggerDo) Having(conds ...gen.Condition) IQrtzBlobTriggerDo {
	return q.withDO(q.DO.Having(conds...))
}

func (q qrtzBlobTriggerDo) Limit(limit int) IQrtzBlobTriggerDo {
	return q.withDO(q.DO.Limit(limit))
}

func (q qrtzBlobTriggerDo) Offset(offset int) IQrtzBlobTriggerDo {
	return q.withDO(q.DO.Offset(offset))
}

func (q qrtzBlobTriggerDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IQrtzBlobTriggerDo {
	return q.withDO(q.DO.Scopes(funcs...))
}

func (q qrtzBlobTriggerDo) Unscoped() IQrtzBlobTriggerDo {
	return q.withDO(q.DO.Unscoped())
}

func (q qrtzBlobTriggerDo) Create(values ...*model.QrtzBlobTrigger) error {
	if len(values) == 0 {
		return nil
	}
	return q.DO.Create(values)
}

func (q qrtzBlobTriggerDo) CreateInBatches(values []*model.QrtzBlobTrigger, batchSize int) error {
	return q.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (q qrtzBlobTriggerDo) Save(values ...*model.QrtzBlobTrigger) error {
	if len(values) == 0 {
		return nil
	}
	return q.DO.Save(values)
}

func (q qrtzBlobTriggerDo) First() (*model.QrtzBlobTrigger, error) {
	if result, err := q.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.QrtzBlobTrigger), nil
	}
}

func (q qrtzBlobTriggerDo) Take() (*model.QrtzBlobTrigger, error) {
	if result, err := q.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.QrtzBlobTrigger), nil
	}
}

func (q qrtzBlobTriggerDo) Last() (*model.QrtzBlobTrigger, error) {
	if result, err := q.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.QrtzBlobTrigger), nil
	}
}

func (q qrtzBlobTriggerDo) Find() ([]*model.QrtzBlobTrigger, error) {
	result, err := q.DO.Find()
	return result.([]*model.QrtzBlobTrigger), err
}

func (q qrtzBlobTriggerDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.QrtzBlobTrigger, err error) {
	buf := make([]*model.QrtzBlobTrigger, 0, batchSize)
	err = q.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (q qrtzBlobTriggerDo) FindInBatches(result *[]*model.QrtzBlobTrigger, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return q.DO.FindInBatches(result, batchSize, fc)
}

func (q qrtzBlobTriggerDo) Attrs(attrs ...field.AssignExpr) IQrtzBlobTriggerDo {
	return q.withDO(q.DO.Attrs(attrs...))
}

func (q qrtzBlobTriggerDo) Assign(attrs ...field.AssignExpr) IQrtzBlobTriggerDo {
	return q.withDO(q.DO.Assign(attrs...))
}

func (q qrtzBlobTriggerDo) Joins(fields ...field.RelationField) IQrtzBlobTriggerDo {
	for _, _f := range fields {
		q = *q.withDO(q.DO.Joins(_f))
	}
	return &q
}

func (q qrtzBlobTriggerDo) Preload(fields ...field.RelationField) IQrtzBlobTriggerDo {
	for _, _f := range fields {
		q = *q.withDO(q.DO.Preload(_f))
	}
	return &q
}

func (q qrtzBlobTriggerDo) FirstOrInit() (*model.QrtzBlobTrigger, error) {
	if result, err := q.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.QrtzBlobTrigger), nil
	}
}

func (q qrtzBlobTriggerDo) FirstOrCreate() (*model.QrtzBlobTrigger, error) {
	if result, err := q.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.QrtzBlobTrigger), nil
	}
}

func (q qrtzBlobTriggerDo) FindByPage(offset int, limit int) (result []*model.QrtzBlobTrigger, count int64, err error) {
	result, err = q.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = q.Offset(-1).Limit(-1).Count()
	return
}

func (q qrtzBlobTriggerDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = q.Count()
	if err != nil {
		return
	}

	err = q.Offset(offset).Limit(limit).Scan(result)
	return
}

func (q qrtzBlobTriggerDo) Scan(result interface{}) (err error) {
	return q.DO.Scan(result)
}

func (q qrtzBlobTriggerDo) Delete(models ...*model.QrtzBlobTrigger) (result gen.ResultInfo, err error) {
	return q.DO.Delete(models)
}

func (q *qrtzBlobTriggerDo) withDO(do gen.Dao) *qrtzBlobTriggerDo {
	q.DO = *do.(*gen.DO)
	return q
}
