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

func newCadCommonComment(db *gorm.DB, opts ...gen.DOOption) cadCommonComment {
	_cadCommonComment := cadCommonComment{}

	_cadCommonComment.cadCommonCommentDo.UseDB(db, opts...)
	_cadCommonComment.cadCommonCommentDo.UseModel(&model.CadCommonComment{})

	tableName := _cadCommonComment.cadCommonCommentDo.TableName()
	_cadCommonComment.ALL = field.NewAsterisk(tableName)
	_cadCommonComment.ID = field.NewString(tableName, "id")
	_cadCommonComment.DelFlag = field.NewInt32(tableName, "del_flag")
	_cadCommonComment.RelateBusinessTable = field.NewString(tableName, "relate_business_table")
	_cadCommonComment.RelateBusinessID = field.NewString(tableName, "relate_business_id")
	_cadCommonComment.Content = field.NewString(tableName, "content")
	_cadCommonComment.ReplyTo = field.NewString(tableName, "reply_to")
	_cadCommonComment.CreateTime = field.NewTime(tableName, "create_time")
	_cadCommonComment.CreateBy = field.NewString(tableName, "create_by")
	_cadCommonComment.SysOrgCode = field.NewString(tableName, "sys_org_code")

	_cadCommonComment.fillFieldMap()

	return _cadCommonComment
}

type cadCommonComment struct {
	cadCommonCommentDo cadCommonCommentDo

	ALL                 field.Asterisk
	ID                  field.String
	DelFlag             field.Int32
	RelateBusinessTable field.String
	RelateBusinessID    field.String
	Content             field.String
	ReplyTo             field.String
	CreateTime          field.Time
	CreateBy            field.String
	SysOrgCode          field.String

	fieldMap map[string]field.Expr
}

func (c cadCommonComment) Table(newTableName string) *cadCommonComment {
	c.cadCommonCommentDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c cadCommonComment) As(alias string) *cadCommonComment {
	c.cadCommonCommentDo.DO = *(c.cadCommonCommentDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *cadCommonComment) updateTableName(table string) *cadCommonComment {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewString(table, "id")
	c.DelFlag = field.NewInt32(table, "del_flag")
	c.RelateBusinessTable = field.NewString(table, "relate_business_table")
	c.RelateBusinessID = field.NewString(table, "relate_business_id")
	c.Content = field.NewString(table, "content")
	c.ReplyTo = field.NewString(table, "reply_to")
	c.CreateTime = field.NewTime(table, "create_time")
	c.CreateBy = field.NewString(table, "create_by")
	c.SysOrgCode = field.NewString(table, "sys_org_code")

	c.fillFieldMap()

	return c
}

func (c *cadCommonComment) WithContext(ctx context.Context) ICadCommonCommentDo {
	return c.cadCommonCommentDo.WithContext(ctx)
}

func (c cadCommonComment) TableName() string { return c.cadCommonCommentDo.TableName() }

func (c cadCommonComment) Alias() string { return c.cadCommonCommentDo.Alias() }

func (c cadCommonComment) Columns(cols ...field.Expr) gen.Columns {
	return c.cadCommonCommentDo.Columns(cols...)
}

func (c *cadCommonComment) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *cadCommonComment) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 9)
	c.fieldMap["id"] = c.ID
	c.fieldMap["del_flag"] = c.DelFlag
	c.fieldMap["relate_business_table"] = c.RelateBusinessTable
	c.fieldMap["relate_business_id"] = c.RelateBusinessID
	c.fieldMap["content"] = c.Content
	c.fieldMap["reply_to"] = c.ReplyTo
	c.fieldMap["create_time"] = c.CreateTime
	c.fieldMap["create_by"] = c.CreateBy
	c.fieldMap["sys_org_code"] = c.SysOrgCode
}

func (c cadCommonComment) clone(db *gorm.DB) cadCommonComment {
	c.cadCommonCommentDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c cadCommonComment) replaceDB(db *gorm.DB) cadCommonComment {
	c.cadCommonCommentDo.ReplaceDB(db)
	return c
}

type cadCommonCommentDo struct{ gen.DO }

type ICadCommonCommentDo interface {
	gen.SubQuery
	Debug() ICadCommonCommentDo
	WithContext(ctx context.Context) ICadCommonCommentDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICadCommonCommentDo
	WriteDB() ICadCommonCommentDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICadCommonCommentDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICadCommonCommentDo
	Not(conds ...gen.Condition) ICadCommonCommentDo
	Or(conds ...gen.Condition) ICadCommonCommentDo
	Select(conds ...field.Expr) ICadCommonCommentDo
	Where(conds ...gen.Condition) ICadCommonCommentDo
	Order(conds ...field.Expr) ICadCommonCommentDo
	Distinct(cols ...field.Expr) ICadCommonCommentDo
	Omit(cols ...field.Expr) ICadCommonCommentDo
	Join(table schema.Tabler, on ...field.Expr) ICadCommonCommentDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICadCommonCommentDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICadCommonCommentDo
	Group(cols ...field.Expr) ICadCommonCommentDo
	Having(conds ...gen.Condition) ICadCommonCommentDo
	Limit(limit int) ICadCommonCommentDo
	Offset(offset int) ICadCommonCommentDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICadCommonCommentDo
	Unscoped() ICadCommonCommentDo
	Create(values ...*model.CadCommonComment) error
	CreateInBatches(values []*model.CadCommonComment, batchSize int) error
	Save(values ...*model.CadCommonComment) error
	First() (*model.CadCommonComment, error)
	Take() (*model.CadCommonComment, error)
	Last() (*model.CadCommonComment, error)
	Find() ([]*model.CadCommonComment, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CadCommonComment, err error)
	FindInBatches(result *[]*model.CadCommonComment, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.CadCommonComment) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICadCommonCommentDo
	Assign(attrs ...field.AssignExpr) ICadCommonCommentDo
	Joins(fields ...field.RelationField) ICadCommonCommentDo
	Preload(fields ...field.RelationField) ICadCommonCommentDo
	FirstOrInit() (*model.CadCommonComment, error)
	FirstOrCreate() (*model.CadCommonComment, error)
	FindByPage(offset int, limit int) (result []*model.CadCommonComment, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICadCommonCommentDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c cadCommonCommentDo) Debug() ICadCommonCommentDo {
	return c.withDO(c.DO.Debug())
}

func (c cadCommonCommentDo) WithContext(ctx context.Context) ICadCommonCommentDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c cadCommonCommentDo) ReadDB() ICadCommonCommentDo {
	return c.Clauses(dbresolver.Read)
}

func (c cadCommonCommentDo) WriteDB() ICadCommonCommentDo {
	return c.Clauses(dbresolver.Write)
}

func (c cadCommonCommentDo) Session(config *gorm.Session) ICadCommonCommentDo {
	return c.withDO(c.DO.Session(config))
}

func (c cadCommonCommentDo) Clauses(conds ...clause.Expression) ICadCommonCommentDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c cadCommonCommentDo) Returning(value interface{}, columns ...string) ICadCommonCommentDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c cadCommonCommentDo) Not(conds ...gen.Condition) ICadCommonCommentDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c cadCommonCommentDo) Or(conds ...gen.Condition) ICadCommonCommentDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c cadCommonCommentDo) Select(conds ...field.Expr) ICadCommonCommentDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c cadCommonCommentDo) Where(conds ...gen.Condition) ICadCommonCommentDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c cadCommonCommentDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ICadCommonCommentDo {
	return c.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (c cadCommonCommentDo) Order(conds ...field.Expr) ICadCommonCommentDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c cadCommonCommentDo) Distinct(cols ...field.Expr) ICadCommonCommentDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c cadCommonCommentDo) Omit(cols ...field.Expr) ICadCommonCommentDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c cadCommonCommentDo) Join(table schema.Tabler, on ...field.Expr) ICadCommonCommentDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c cadCommonCommentDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICadCommonCommentDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c cadCommonCommentDo) RightJoin(table schema.Tabler, on ...field.Expr) ICadCommonCommentDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c cadCommonCommentDo) Group(cols ...field.Expr) ICadCommonCommentDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c cadCommonCommentDo) Having(conds ...gen.Condition) ICadCommonCommentDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c cadCommonCommentDo) Limit(limit int) ICadCommonCommentDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c cadCommonCommentDo) Offset(offset int) ICadCommonCommentDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c cadCommonCommentDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICadCommonCommentDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c cadCommonCommentDo) Unscoped() ICadCommonCommentDo {
	return c.withDO(c.DO.Unscoped())
}

func (c cadCommonCommentDo) Create(values ...*model.CadCommonComment) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c cadCommonCommentDo) CreateInBatches(values []*model.CadCommonComment, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c cadCommonCommentDo) Save(values ...*model.CadCommonComment) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c cadCommonCommentDo) First() (*model.CadCommonComment, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.CadCommonComment), nil
	}
}

func (c cadCommonCommentDo) Take() (*model.CadCommonComment, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.CadCommonComment), nil
	}
}

func (c cadCommonCommentDo) Last() (*model.CadCommonComment, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.CadCommonComment), nil
	}
}

func (c cadCommonCommentDo) Find() ([]*model.CadCommonComment, error) {
	result, err := c.DO.Find()
	return result.([]*model.CadCommonComment), err
}

func (c cadCommonCommentDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CadCommonComment, err error) {
	buf := make([]*model.CadCommonComment, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c cadCommonCommentDo) FindInBatches(result *[]*model.CadCommonComment, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c cadCommonCommentDo) Attrs(attrs ...field.AssignExpr) ICadCommonCommentDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c cadCommonCommentDo) Assign(attrs ...field.AssignExpr) ICadCommonCommentDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c cadCommonCommentDo) Joins(fields ...field.RelationField) ICadCommonCommentDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c cadCommonCommentDo) Preload(fields ...field.RelationField) ICadCommonCommentDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c cadCommonCommentDo) FirstOrInit() (*model.CadCommonComment, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.CadCommonComment), nil
	}
}

func (c cadCommonCommentDo) FirstOrCreate() (*model.CadCommonComment, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.CadCommonComment), nil
	}
}

func (c cadCommonCommentDo) FindByPage(offset int, limit int) (result []*model.CadCommonComment, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c cadCommonCommentDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c cadCommonCommentDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c cadCommonCommentDo) Delete(models ...*model.CadCommonComment) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *cadCommonCommentDo) withDO(do gen.Dao) *cadCommonCommentDo {
	c.DO = *do.(*gen.DO)
	return c
}
