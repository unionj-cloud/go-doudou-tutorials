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

func newTUser(db *gorm.DB, opts ...gen.DOOption) tUser {
	_tUser := tUser{}

	_tUser.tUserDo.UseDB(db, opts...)
	_tUser.tUserDo.UseModel(&model.TUser{})

	tableName := _tUser.tUserDo.TableName()
	_tUser.ALL = field.NewAsterisk(tableName)
	_tUser.ID = field.NewInt32(tableName, "id")
	_tUser.Username = field.NewString(tableName, "username")
	_tUser.Password = field.NewString(tableName, "password")
	_tUser.Name = field.NewString(tableName, "name")
	_tUser.Phone = field.NewString(tableName, "phone")
	_tUser.Dept = field.NewString(tableName, "dept")
	_tUser.Avatar = field.NewString(tableName, "avatar")
	_tUser.CreateAt = field.NewTime(tableName, "create_at")
	_tUser.UpdateAt = field.NewTime(tableName, "update_at")
	_tUser.DeleteAt = field.NewTime(tableName, "delete_at")

	_tUser.fillFieldMap()

	return _tUser
}

type tUser struct {
	tUserDo tUserDo

	ALL      field.Asterisk
	ID       field.Int32
	Username field.String // username
	Password field.String // password
	Name     field.String // real name
	Phone    field.String // phone number
	Dept     field.String // department
	Avatar   field.String // user avatar
	CreateAt field.Time
	UpdateAt field.Time
	DeleteAt field.Time

	fieldMap map[string]field.Expr
}

func (t tUser) Table(newTableName string) *tUser {
	t.tUserDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tUser) As(alias string) *tUser {
	t.tUserDo.DO = *(t.tUserDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tUser) updateTableName(table string) *tUser {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewInt32(table, "id")
	t.Username = field.NewString(table, "username")
	t.Password = field.NewString(table, "password")
	t.Name = field.NewString(table, "name")
	t.Phone = field.NewString(table, "phone")
	t.Dept = field.NewString(table, "dept")
	t.Avatar = field.NewString(table, "avatar")
	t.CreateAt = field.NewTime(table, "create_at")
	t.UpdateAt = field.NewTime(table, "update_at")
	t.DeleteAt = field.NewTime(table, "delete_at")

	t.fillFieldMap()

	return t
}

func (t *tUser) WithContext(ctx context.Context) ITUserDo { return t.tUserDo.WithContext(ctx) }

func (t tUser) TableName() string { return t.tUserDo.TableName() }

func (t tUser) Alias() string { return t.tUserDo.Alias() }

func (t tUser) Columns(cols ...field.Expr) gen.Columns { return t.tUserDo.Columns(cols...) }

func (t *tUser) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tUser) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 10)
	t.fieldMap["id"] = t.ID
	t.fieldMap["username"] = t.Username
	t.fieldMap["password"] = t.Password
	t.fieldMap["name"] = t.Name
	t.fieldMap["phone"] = t.Phone
	t.fieldMap["dept"] = t.Dept
	t.fieldMap["avatar"] = t.Avatar
	t.fieldMap["create_at"] = t.CreateAt
	t.fieldMap["update_at"] = t.UpdateAt
	t.fieldMap["delete_at"] = t.DeleteAt
}

func (t tUser) clone(db *gorm.DB) tUser {
	t.tUserDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tUser) replaceDB(db *gorm.DB) tUser {
	t.tUserDo.ReplaceDB(db)
	return t
}

type tUserDo struct{ gen.DO }

type ITUserDo interface {
	gen.SubQuery
	Debug() ITUserDo
	WithContext(ctx context.Context) ITUserDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITUserDo
	WriteDB() ITUserDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITUserDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITUserDo
	Not(conds ...gen.Condition) ITUserDo
	Or(conds ...gen.Condition) ITUserDo
	Select(conds ...field.Expr) ITUserDo
	Where(conds ...gen.Condition) ITUserDo
	Order(conds ...field.Expr) ITUserDo
	Distinct(cols ...field.Expr) ITUserDo
	Omit(cols ...field.Expr) ITUserDo
	Join(table schema.Tabler, on ...field.Expr) ITUserDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITUserDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITUserDo
	Group(cols ...field.Expr) ITUserDo
	Having(conds ...gen.Condition) ITUserDo
	Limit(limit int) ITUserDo
	Offset(offset int) ITUserDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITUserDo
	Unscoped() ITUserDo
	Create(values ...*model.TUser) error
	CreateInBatches(values []*model.TUser, batchSize int) error
	Save(values ...*model.TUser) error
	First() (*model.TUser, error)
	Take() (*model.TUser, error)
	Last() (*model.TUser, error)
	Find() ([]*model.TUser, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TUser, err error)
	FindInBatches(result *[]*model.TUser, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.TUser) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITUserDo
	Assign(attrs ...field.AssignExpr) ITUserDo
	Joins(fields ...field.RelationField) ITUserDo
	Preload(fields ...field.RelationField) ITUserDo
	FirstOrInit() (*model.TUser, error)
	FirstOrCreate() (*model.TUser, error)
	FindByPage(offset int, limit int) (result []*model.TUser, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITUserDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t tUserDo) Debug() ITUserDo {
	return t.withDO(t.DO.Debug())
}

func (t tUserDo) WithContext(ctx context.Context) ITUserDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tUserDo) ReadDB() ITUserDo {
	return t.Clauses(dbresolver.Read)
}

func (t tUserDo) WriteDB() ITUserDo {
	return t.Clauses(dbresolver.Write)
}

func (t tUserDo) Session(config *gorm.Session) ITUserDo {
	return t.withDO(t.DO.Session(config))
}

func (t tUserDo) Clauses(conds ...clause.Expression) ITUserDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tUserDo) Returning(value interface{}, columns ...string) ITUserDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tUserDo) Not(conds ...gen.Condition) ITUserDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tUserDo) Or(conds ...gen.Condition) ITUserDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tUserDo) Select(conds ...field.Expr) ITUserDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tUserDo) Where(conds ...gen.Condition) ITUserDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tUserDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ITUserDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t tUserDo) Order(conds ...field.Expr) ITUserDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tUserDo) Distinct(cols ...field.Expr) ITUserDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tUserDo) Omit(cols ...field.Expr) ITUserDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tUserDo) Join(table schema.Tabler, on ...field.Expr) ITUserDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tUserDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITUserDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tUserDo) RightJoin(table schema.Tabler, on ...field.Expr) ITUserDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tUserDo) Group(cols ...field.Expr) ITUserDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tUserDo) Having(conds ...gen.Condition) ITUserDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tUserDo) Limit(limit int) ITUserDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tUserDo) Offset(offset int) ITUserDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tUserDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITUserDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tUserDo) Unscoped() ITUserDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tUserDo) Create(values ...*model.TUser) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tUserDo) CreateInBatches(values []*model.TUser, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tUserDo) Save(values ...*model.TUser) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tUserDo) First() (*model.TUser, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TUser), nil
	}
}

func (t tUserDo) Take() (*model.TUser, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TUser), nil
	}
}

func (t tUserDo) Last() (*model.TUser, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TUser), nil
	}
}

func (t tUserDo) Find() ([]*model.TUser, error) {
	result, err := t.DO.Find()
	return result.([]*model.TUser), err
}

func (t tUserDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TUser, err error) {
	buf := make([]*model.TUser, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tUserDo) FindInBatches(result *[]*model.TUser, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tUserDo) Attrs(attrs ...field.AssignExpr) ITUserDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tUserDo) Assign(attrs ...field.AssignExpr) ITUserDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tUserDo) Joins(fields ...field.RelationField) ITUserDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tUserDo) Preload(fields ...field.RelationField) ITUserDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tUserDo) FirstOrInit() (*model.TUser, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TUser), nil
	}
}

func (t tUserDo) FirstOrCreate() (*model.TUser, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TUser), nil
	}
}

func (t tUserDo) FindByPage(offset int, limit int) (result []*model.TUser, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t tUserDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tUserDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tUserDo) Delete(models ...*model.TUser) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tUserDo) withDO(do gen.Dao) *tUserDo {
	t.DO = *do.(*gen.DO)
	return t
}