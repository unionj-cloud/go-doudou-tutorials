// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"github.com/unionj-cloud/go-doudou/v2/toolkit/gormgen"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/gormgen/field"

	"gorm.io/plugin/dbresolver"

	"testgggorm2/model"
)

func newProductCate(db *gorm.DB, opts ...gormgen.DOOption) productCate {
	_productCate := productCate{}

	_productCate.productCateDo.UseDB(db, opts...)
	_productCate.productCateDo.UseModel(&model.ProductCate{})

	tableName := _productCate.productCateDo.TableName()
	_productCate.ALL = field.NewAsterisk(tableName)
	_productCate.ID = field.NewInt32(tableName, "id")
	_productCate.Title = field.NewString(tableName, "title")
	_productCate.CateImg = field.NewString(tableName, "cate_img")
	_productCate.Link = field.NewString(tableName, "link")
	_productCate.Template = field.NewString(tableName, "template")
	_productCate.Pid = field.NewInt32(tableName, "pid")
	_productCate.SubTitle = field.NewString(tableName, "sub_title")
	_productCate.Keywords = field.NewString(tableName, "keywords")
	_productCate.Description = field.NewString(tableName, "description")
	_productCate.Sort = field.NewInt32(tableName, "sort")
	_productCate.Status = field.NewInt32(tableName, "status")
	_productCate.AddTime = field.NewInt32(tableName, "add_time")

	_productCate.fillFieldMap()

	return _productCate
}

type productCate struct {
	productCateDo productCateDo

	ALL         field.Asterisk
	ID          field.Int32
	Title       field.String // 分类名称
	CateImg     field.String // 分类图片
	Link        field.String // 链接
	Template    field.String // 模版
	Pid         field.Int32  // 父编号
	SubTitle    field.String // 子标题
	Keywords    field.String // 关键字
	Description field.String // 描述
	Sort        field.Int32  // 排序
	Status      field.Int32  // 状态
	AddTime     field.Int32  // 添加时间

	fieldMap map[string]field.Expr
}

func (p productCate) Table(newTableName string) *productCate {
	p.productCateDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p productCate) As(alias string) *productCate {
	p.productCateDo.DO = *(p.productCateDo.As(alias).(*gormgen.DO))
	return p.updateTableName(alias)
}

func (p *productCate) updateTableName(table string) *productCate {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt32(table, "id")
	p.Title = field.NewString(table, "title")
	p.CateImg = field.NewString(table, "cate_img")
	p.Link = field.NewString(table, "link")
	p.Template = field.NewString(table, "template")
	p.Pid = field.NewInt32(table, "pid")
	p.SubTitle = field.NewString(table, "sub_title")
	p.Keywords = field.NewString(table, "keywords")
	p.Description = field.NewString(table, "description")
	p.Sort = field.NewInt32(table, "sort")
	p.Status = field.NewInt32(table, "status")
	p.AddTime = field.NewInt32(table, "add_time")

	p.fillFieldMap()

	return p
}

func (p *productCate) WithContext(ctx context.Context) IProductCateDo {
	return p.productCateDo.WithContext(ctx)
}

func (p productCate) TableName() string { return p.productCateDo.TableName() }

func (p productCate) Alias() string { return p.productCateDo.Alias() }

func (p productCate) Columns(cols ...field.Expr) gormgen.Columns {
	return p.productCateDo.Columns(cols...)
}

func (p *productCate) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *productCate) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 12)
	p.fieldMap["id"] = p.ID
	p.fieldMap["title"] = p.Title
	p.fieldMap["cate_img"] = p.CateImg
	p.fieldMap["link"] = p.Link
	p.fieldMap["template"] = p.Template
	p.fieldMap["pid"] = p.Pid
	p.fieldMap["sub_title"] = p.SubTitle
	p.fieldMap["keywords"] = p.Keywords
	p.fieldMap["description"] = p.Description
	p.fieldMap["sort"] = p.Sort
	p.fieldMap["status"] = p.Status
	p.fieldMap["add_time"] = p.AddTime
}

func (p productCate) clone(db *gorm.DB) productCate {
	p.productCateDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p productCate) replaceDB(db *gorm.DB) productCate {
	p.productCateDo.ReplaceDB(db)
	return p
}

type productCateDo struct{ gormgen.DO }

type IProductCateDo interface {
	gormgen.SubQuery
	Debug() IProductCateDo
	WithContext(ctx context.Context) IProductCateDo
	WithResult(fc func(tx gormgen.Dao)) gormgen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IProductCateDo
	WriteDB() IProductCateDo
	As(alias string) gormgen.Dao
	Session(config *gorm.Session) IProductCateDo
	Columns(cols ...field.Expr) gormgen.Columns
	Clauses(conds ...clause.Expression) IProductCateDo
	Not(conds ...gormgen.Condition) IProductCateDo
	Or(conds ...gormgen.Condition) IProductCateDo
	Select(conds ...field.Expr) IProductCateDo
	Where(conds ...gormgen.Condition) IProductCateDo
	Order(conds ...field.Expr) IProductCateDo
	Distinct(cols ...field.Expr) IProductCateDo
	Omit(cols ...field.Expr) IProductCateDo
	Join(table schema.Tabler, on ...field.Expr) IProductCateDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IProductCateDo
	RightJoin(table schema.Tabler, on ...field.Expr) IProductCateDo
	Group(cols ...field.Expr) IProductCateDo
	Having(conds ...gormgen.Condition) IProductCateDo
	Limit(limit int) IProductCateDo
	Offset(offset int) IProductCateDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gormgen.Dao) gormgen.Dao) IProductCateDo
	Unscoped() IProductCateDo
	Create(values ...*model.ProductCate) error
	CreateInBatches(values []*model.ProductCate, batchSize int) error
	Save(values ...*model.ProductCate) error
	First() (*model.ProductCate, error)
	Take() (*model.ProductCate, error)
	Last() (*model.ProductCate, error)
	Find() ([]*model.ProductCate, error)
	FindInBatch(batchSize int, fc func(tx gormgen.Dao, batch int) error) (results []*model.ProductCate, err error)
	FindInBatches(result *[]*model.ProductCate, batchSize int, fc func(tx gormgen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.ProductCate) (info gormgen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gormgen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gormgen.ResultInfo, err error)
	Updates(value interface{}) (info gormgen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gormgen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gormgen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gormgen.ResultInfo, err error)
	UpdateFrom(q gormgen.SubQuery) gormgen.Dao
	Attrs(attrs ...field.AssignExpr) IProductCateDo
	Assign(attrs ...field.AssignExpr) IProductCateDo
	Joins(fields ...field.RelationField) IProductCateDo
	Preload(fields ...field.RelationField) IProductCateDo
	FirstOrInit() (*model.ProductCate, error)
	FirstOrCreate() (*model.ProductCate, error)
	FindByPage(offset int, limit int) (result []*model.ProductCate, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IProductCateDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p productCateDo) Debug() IProductCateDo {
	return p.withDO(p.DO.Debug())
}

func (p productCateDo) WithContext(ctx context.Context) IProductCateDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p productCateDo) ReadDB() IProductCateDo {
	return p.Clauses(dbresolver.Read)
}

func (p productCateDo) WriteDB() IProductCateDo {
	return p.Clauses(dbresolver.Write)
}

func (p productCateDo) Session(config *gorm.Session) IProductCateDo {
	return p.withDO(p.DO.Session(config))
}

func (p productCateDo) Clauses(conds ...clause.Expression) IProductCateDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p productCateDo) Returning(value interface{}, columns ...string) IProductCateDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p productCateDo) Not(conds ...gormgen.Condition) IProductCateDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p productCateDo) Or(conds ...gormgen.Condition) IProductCateDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p productCateDo) Select(conds ...field.Expr) IProductCateDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p productCateDo) Where(conds ...gormgen.Condition) IProductCateDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p productCateDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IProductCateDo {
	return p.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (p productCateDo) Order(conds ...field.Expr) IProductCateDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p productCateDo) Distinct(cols ...field.Expr) IProductCateDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p productCateDo) Omit(cols ...field.Expr) IProductCateDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p productCateDo) Join(table schema.Tabler, on ...field.Expr) IProductCateDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p productCateDo) LeftJoin(table schema.Tabler, on ...field.Expr) IProductCateDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p productCateDo) RightJoin(table schema.Tabler, on ...field.Expr) IProductCateDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p productCateDo) Group(cols ...field.Expr) IProductCateDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p productCateDo) Having(conds ...gormgen.Condition) IProductCateDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p productCateDo) Limit(limit int) IProductCateDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p productCateDo) Offset(offset int) IProductCateDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p productCateDo) Scopes(funcs ...func(gormgen.Dao) gormgen.Dao) IProductCateDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p productCateDo) Unscoped() IProductCateDo {
	return p.withDO(p.DO.Unscoped())
}

func (p productCateDo) Create(values ...*model.ProductCate) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p productCateDo) CreateInBatches(values []*model.ProductCate, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p productCateDo) Save(values ...*model.ProductCate) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p productCateDo) First() (*model.ProductCate, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProductCate), nil
	}
}

func (p productCateDo) Take() (*model.ProductCate, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProductCate), nil
	}
}

func (p productCateDo) Last() (*model.ProductCate, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProductCate), nil
	}
}

func (p productCateDo) Find() ([]*model.ProductCate, error) {
	result, err := p.DO.Find()
	return result.([]*model.ProductCate), err
}

func (p productCateDo) FindInBatch(batchSize int, fc func(tx gormgen.Dao, batch int) error) (results []*model.ProductCate, err error) {
	buf := make([]*model.ProductCate, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gormgen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p productCateDo) FindInBatches(result *[]*model.ProductCate, batchSize int, fc func(tx gormgen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p productCateDo) Attrs(attrs ...field.AssignExpr) IProductCateDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p productCateDo) Assign(attrs ...field.AssignExpr) IProductCateDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p productCateDo) Joins(fields ...field.RelationField) IProductCateDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p productCateDo) Preload(fields ...field.RelationField) IProductCateDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p productCateDo) FirstOrInit() (*model.ProductCate, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProductCate), nil
	}
}

func (p productCateDo) FirstOrCreate() (*model.ProductCate, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProductCate), nil
	}
}

func (p productCateDo) FindByPage(offset int, limit int) (result []*model.ProductCate, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p productCateDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p productCateDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p productCateDo) Delete(models ...*model.ProductCate) (result gormgen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *productCateDo) withDO(do gormgen.Dao) *productCateDo {
	p.DO = *do.(*gormgen.DO)
	return p
}
