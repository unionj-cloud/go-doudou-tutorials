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

func newProductAttr(db *gorm.DB, opts ...gormgen.DOOption) productAttr {
	_productAttr := productAttr{}

	_productAttr.productAttrDo.UseDB(db, opts...)
	_productAttr.productAttrDo.UseModel(&model.ProductAttr{})

	tableName := _productAttr.productAttrDo.TableName()
	_productAttr.ALL = field.NewAsterisk(tableName)
	_productAttr.ID = field.NewInt32(tableName, "id")
	_productAttr.ProductID = field.NewInt32(tableName, "product_id")
	_productAttr.AttributeCateID = field.NewInt32(tableName, "attribute_cate_id")
	_productAttr.AttributeID = field.NewInt32(tableName, "attribute_id")
	_productAttr.AttributeTitle = field.NewString(tableName, "attribute_title")
	_productAttr.AttributeType = field.NewInt32(tableName, "attribute_type")
	_productAttr.AttributeValue = field.NewString(tableName, "attribute_value")
	_productAttr.Sort = field.NewInt32(tableName, "sort")
	_productAttr.AddTime = field.NewInt32(tableName, "add_time")
	_productAttr.Status = field.NewInt32(tableName, "status")

	_productAttr.fillFieldMap()

	return _productAttr
}

type productAttr struct {
	productAttrDo productAttrDo

	ALL             field.Asterisk
	ID              field.Int32
	ProductID       field.Int32  // 商品编号
	AttributeCateID field.Int32  // 属性分类编号
	AttributeID     field.Int32  // 属性编号
	AttributeTitle  field.String // 属性标题
	AttributeType   field.Int32  // 属性类型
	AttributeValue  field.String // 属性值
	Sort            field.Int32  // 排序
	AddTime         field.Int32  // 添加时间
	Status          field.Int32  // 状态

	fieldMap map[string]field.Expr
}

func (p productAttr) Table(newTableName string) *productAttr {
	p.productAttrDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p productAttr) As(alias string) *productAttr {
	p.productAttrDo.DO = *(p.productAttrDo.As(alias).(*gormgen.DO))
	return p.updateTableName(alias)
}

func (p *productAttr) updateTableName(table string) *productAttr {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt32(table, "id")
	p.ProductID = field.NewInt32(table, "product_id")
	p.AttributeCateID = field.NewInt32(table, "attribute_cate_id")
	p.AttributeID = field.NewInt32(table, "attribute_id")
	p.AttributeTitle = field.NewString(table, "attribute_title")
	p.AttributeType = field.NewInt32(table, "attribute_type")
	p.AttributeValue = field.NewString(table, "attribute_value")
	p.Sort = field.NewInt32(table, "sort")
	p.AddTime = field.NewInt32(table, "add_time")
	p.Status = field.NewInt32(table, "status")

	p.fillFieldMap()

	return p
}

func (p *productAttr) WithContext(ctx context.Context) IProductAttrDo {
	return p.productAttrDo.WithContext(ctx)
}

func (p productAttr) TableName() string { return p.productAttrDo.TableName() }

func (p productAttr) Alias() string { return p.productAttrDo.Alias() }

func (p productAttr) Columns(cols ...field.Expr) gormgen.Columns {
	return p.productAttrDo.Columns(cols...)
}

func (p *productAttr) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *productAttr) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 10)
	p.fieldMap["id"] = p.ID
	p.fieldMap["product_id"] = p.ProductID
	p.fieldMap["attribute_cate_id"] = p.AttributeCateID
	p.fieldMap["attribute_id"] = p.AttributeID
	p.fieldMap["attribute_title"] = p.AttributeTitle
	p.fieldMap["attribute_type"] = p.AttributeType
	p.fieldMap["attribute_value"] = p.AttributeValue
	p.fieldMap["sort"] = p.Sort
	p.fieldMap["add_time"] = p.AddTime
	p.fieldMap["status"] = p.Status
}

func (p productAttr) clone(db *gorm.DB) productAttr {
	p.productAttrDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p productAttr) replaceDB(db *gorm.DB) productAttr {
	p.productAttrDo.ReplaceDB(db)
	return p
}

type productAttrDo struct{ gormgen.DO }

type IProductAttrDo interface {
	gormgen.SubQuery
	Debug() IProductAttrDo
	WithContext(ctx context.Context) IProductAttrDo
	WithResult(fc func(tx gormgen.Dao)) gormgen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IProductAttrDo
	WriteDB() IProductAttrDo
	As(alias string) gormgen.Dao
	Session(config *gorm.Session) IProductAttrDo
	Columns(cols ...field.Expr) gormgen.Columns
	Clauses(conds ...clause.Expression) IProductAttrDo
	Not(conds ...gormgen.Condition) IProductAttrDo
	Or(conds ...gormgen.Condition) IProductAttrDo
	Select(conds ...field.Expr) IProductAttrDo
	Where(conds ...gormgen.Condition) IProductAttrDo
	Order(conds ...field.Expr) IProductAttrDo
	Distinct(cols ...field.Expr) IProductAttrDo
	Omit(cols ...field.Expr) IProductAttrDo
	Join(table schema.Tabler, on ...field.Expr) IProductAttrDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IProductAttrDo
	RightJoin(table schema.Tabler, on ...field.Expr) IProductAttrDo
	Group(cols ...field.Expr) IProductAttrDo
	Having(conds ...gormgen.Condition) IProductAttrDo
	Limit(limit int) IProductAttrDo
	Offset(offset int) IProductAttrDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gormgen.Dao) gormgen.Dao) IProductAttrDo
	Unscoped() IProductAttrDo
	Create(values ...*model.ProductAttr) error
	CreateInBatches(values []*model.ProductAttr, batchSize int) error
	Save(values ...*model.ProductAttr) error
	First() (*model.ProductAttr, error)
	Take() (*model.ProductAttr, error)
	Last() (*model.ProductAttr, error)
	Find() ([]*model.ProductAttr, error)
	FindInBatch(batchSize int, fc func(tx gormgen.Dao, batch int) error) (results []*model.ProductAttr, err error)
	FindInBatches(result *[]*model.ProductAttr, batchSize int, fc func(tx gormgen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.ProductAttr) (info gormgen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gormgen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gormgen.ResultInfo, err error)
	Updates(value interface{}) (info gormgen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gormgen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gormgen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gormgen.ResultInfo, err error)
	UpdateFrom(q gormgen.SubQuery) gormgen.Dao
	Attrs(attrs ...field.AssignExpr) IProductAttrDo
	Assign(attrs ...field.AssignExpr) IProductAttrDo
	Joins(fields ...field.RelationField) IProductAttrDo
	Preload(fields ...field.RelationField) IProductAttrDo
	FirstOrInit() (*model.ProductAttr, error)
	FirstOrCreate() (*model.ProductAttr, error)
	FindByPage(offset int, limit int) (result []*model.ProductAttr, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IProductAttrDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p productAttrDo) Debug() IProductAttrDo {
	return p.withDO(p.DO.Debug())
}

func (p productAttrDo) WithContext(ctx context.Context) IProductAttrDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p productAttrDo) ReadDB() IProductAttrDo {
	return p.Clauses(dbresolver.Read)
}

func (p productAttrDo) WriteDB() IProductAttrDo {
	return p.Clauses(dbresolver.Write)
}

func (p productAttrDo) Session(config *gorm.Session) IProductAttrDo {
	return p.withDO(p.DO.Session(config))
}

func (p productAttrDo) Clauses(conds ...clause.Expression) IProductAttrDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p productAttrDo) Returning(value interface{}, columns ...string) IProductAttrDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p productAttrDo) Not(conds ...gormgen.Condition) IProductAttrDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p productAttrDo) Or(conds ...gormgen.Condition) IProductAttrDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p productAttrDo) Select(conds ...field.Expr) IProductAttrDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p productAttrDo) Where(conds ...gormgen.Condition) IProductAttrDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p productAttrDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IProductAttrDo {
	return p.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (p productAttrDo) Order(conds ...field.Expr) IProductAttrDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p productAttrDo) Distinct(cols ...field.Expr) IProductAttrDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p productAttrDo) Omit(cols ...field.Expr) IProductAttrDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p productAttrDo) Join(table schema.Tabler, on ...field.Expr) IProductAttrDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p productAttrDo) LeftJoin(table schema.Tabler, on ...field.Expr) IProductAttrDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p productAttrDo) RightJoin(table schema.Tabler, on ...field.Expr) IProductAttrDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p productAttrDo) Group(cols ...field.Expr) IProductAttrDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p productAttrDo) Having(conds ...gormgen.Condition) IProductAttrDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p productAttrDo) Limit(limit int) IProductAttrDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p productAttrDo) Offset(offset int) IProductAttrDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p productAttrDo) Scopes(funcs ...func(gormgen.Dao) gormgen.Dao) IProductAttrDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p productAttrDo) Unscoped() IProductAttrDo {
	return p.withDO(p.DO.Unscoped())
}

func (p productAttrDo) Create(values ...*model.ProductAttr) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p productAttrDo) CreateInBatches(values []*model.ProductAttr, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p productAttrDo) Save(values ...*model.ProductAttr) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p productAttrDo) First() (*model.ProductAttr, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProductAttr), nil
	}
}

func (p productAttrDo) Take() (*model.ProductAttr, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProductAttr), nil
	}
}

func (p productAttrDo) Last() (*model.ProductAttr, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProductAttr), nil
	}
}

func (p productAttrDo) Find() ([]*model.ProductAttr, error) {
	result, err := p.DO.Find()
	return result.([]*model.ProductAttr), err
}

func (p productAttrDo) FindInBatch(batchSize int, fc func(tx gormgen.Dao, batch int) error) (results []*model.ProductAttr, err error) {
	buf := make([]*model.ProductAttr, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gormgen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p productAttrDo) FindInBatches(result *[]*model.ProductAttr, batchSize int, fc func(tx gormgen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p productAttrDo) Attrs(attrs ...field.AssignExpr) IProductAttrDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p productAttrDo) Assign(attrs ...field.AssignExpr) IProductAttrDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p productAttrDo) Joins(fields ...field.RelationField) IProductAttrDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p productAttrDo) Preload(fields ...field.RelationField) IProductAttrDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p productAttrDo) FirstOrInit() (*model.ProductAttr, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProductAttr), nil
	}
}

func (p productAttrDo) FirstOrCreate() (*model.ProductAttr, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProductAttr), nil
	}
}

func (p productAttrDo) FindByPage(offset int, limit int) (result []*model.ProductAttr, count int64, err error) {
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

func (p productAttrDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p productAttrDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p productAttrDo) Delete(models ...*model.ProductAttr) (result gormgen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *productAttrDo) withDO(do gormgen.Dao) *productAttrDo {
	p.DO = *do.(*gormgen.DO)
	return p
}