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

func newSetting(db *gorm.DB, opts ...gormgen.DOOption) setting {
	_setting := setting{}

	_setting.settingDo.UseDB(db, opts...)
	_setting.settingDo.UseModel(&model.Setting{})

	tableName := _setting.settingDo.TableName()
	_setting.ALL = field.NewAsterisk(tableName)
	_setting.ID = field.NewInt32(tableName, "id")
	_setting.SiteTitle = field.NewString(tableName, "site_title")
	_setting.SiteLogo = field.NewString(tableName, "site_logo")
	_setting.SiteKeywords = field.NewString(tableName, "site_keywords")
	_setting.SiteDescription = field.NewString(tableName, "site_description")
	_setting.NoPicture = field.NewString(tableName, "no_picture")
	_setting.SiteIcp = field.NewString(tableName, "site_icp")
	_setting.SiteTel = field.NewString(tableName, "site_tel")
	_setting.SearchKeywords = field.NewString(tableName, "search_keywords")
	_setting.TongjiCode = field.NewString(tableName, "tongji_code")
	_setting.Appid = field.NewString(tableName, "appid")
	_setting.AppSecret = field.NewString(tableName, "app_secret")
	_setting.EndPoint = field.NewString(tableName, "end_point")
	_setting.BucketName = field.NewString(tableName, "bucket_name")
	_setting.OssStatus = field.NewInt32(tableName, "oss_status")

	_setting.fillFieldMap()

	return _setting
}

type setting struct {
	settingDo settingDo

	ALL             field.Asterisk
	ID              field.Int32
	SiteTitle       field.String // 商城名称
	SiteLogo        field.String // 商城图标
	SiteKeywords    field.String // 商城关键字
	SiteDescription field.String // 商城描述
	NoPicture       field.String // 没有图片显示
	SiteIcp         field.String // 商城ICP
	SiteTel         field.String // 商城手机号
	SearchKeywords  field.String // 搜索关键字
	TongjiCode      field.String // 统计编码
	Appid           field.String // oss appid
	AppSecret       field.String // oss app_secret
	EndPoint        field.String // oss 终端点
	BucketName      field.String // oss 桶名称
	OssStatus       field.Int32  // oss 状态

	fieldMap map[string]field.Expr
}

func (s setting) Table(newTableName string) *setting {
	s.settingDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s setting) As(alias string) *setting {
	s.settingDo.DO = *(s.settingDo.As(alias).(*gormgen.DO))
	return s.updateTableName(alias)
}

func (s *setting) updateTableName(table string) *setting {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt32(table, "id")
	s.SiteTitle = field.NewString(table, "site_title")
	s.SiteLogo = field.NewString(table, "site_logo")
	s.SiteKeywords = field.NewString(table, "site_keywords")
	s.SiteDescription = field.NewString(table, "site_description")
	s.NoPicture = field.NewString(table, "no_picture")
	s.SiteIcp = field.NewString(table, "site_icp")
	s.SiteTel = field.NewString(table, "site_tel")
	s.SearchKeywords = field.NewString(table, "search_keywords")
	s.TongjiCode = field.NewString(table, "tongji_code")
	s.Appid = field.NewString(table, "appid")
	s.AppSecret = field.NewString(table, "app_secret")
	s.EndPoint = field.NewString(table, "end_point")
	s.BucketName = field.NewString(table, "bucket_name")
	s.OssStatus = field.NewInt32(table, "oss_status")

	s.fillFieldMap()

	return s
}

func (s *setting) WithContext(ctx context.Context) ISettingDo { return s.settingDo.WithContext(ctx) }

func (s setting) TableName() string { return s.settingDo.TableName() }

func (s setting) Alias() string { return s.settingDo.Alias() }

func (s setting) Columns(cols ...field.Expr) gormgen.Columns { return s.settingDo.Columns(cols...) }

func (s *setting) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *setting) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 15)
	s.fieldMap["id"] = s.ID
	s.fieldMap["site_title"] = s.SiteTitle
	s.fieldMap["site_logo"] = s.SiteLogo
	s.fieldMap["site_keywords"] = s.SiteKeywords
	s.fieldMap["site_description"] = s.SiteDescription
	s.fieldMap["no_picture"] = s.NoPicture
	s.fieldMap["site_icp"] = s.SiteIcp
	s.fieldMap["site_tel"] = s.SiteTel
	s.fieldMap["search_keywords"] = s.SearchKeywords
	s.fieldMap["tongji_code"] = s.TongjiCode
	s.fieldMap["appid"] = s.Appid
	s.fieldMap["app_secret"] = s.AppSecret
	s.fieldMap["end_point"] = s.EndPoint
	s.fieldMap["bucket_name"] = s.BucketName
	s.fieldMap["oss_status"] = s.OssStatus
}

func (s setting) clone(db *gorm.DB) setting {
	s.settingDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s setting) replaceDB(db *gorm.DB) setting {
	s.settingDo.ReplaceDB(db)
	return s
}

type settingDo struct{ gormgen.DO }

type ISettingDo interface {
	gormgen.SubQuery
	Debug() ISettingDo
	WithContext(ctx context.Context) ISettingDo
	WithResult(fc func(tx gormgen.Dao)) gormgen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISettingDo
	WriteDB() ISettingDo
	As(alias string) gormgen.Dao
	Session(config *gorm.Session) ISettingDo
	Columns(cols ...field.Expr) gormgen.Columns
	Clauses(conds ...clause.Expression) ISettingDo
	Not(conds ...gormgen.Condition) ISettingDo
	Or(conds ...gormgen.Condition) ISettingDo
	Select(conds ...field.Expr) ISettingDo
	Where(conds ...gormgen.Condition) ISettingDo
	Order(conds ...field.Expr) ISettingDo
	Distinct(cols ...field.Expr) ISettingDo
	Omit(cols ...field.Expr) ISettingDo
	Join(table schema.Tabler, on ...field.Expr) ISettingDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISettingDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISettingDo
	Group(cols ...field.Expr) ISettingDo
	Having(conds ...gormgen.Condition) ISettingDo
	Limit(limit int) ISettingDo
	Offset(offset int) ISettingDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gormgen.Dao) gormgen.Dao) ISettingDo
	Unscoped() ISettingDo
	Create(values ...*model.Setting) error
	CreateInBatches(values []*model.Setting, batchSize int) error
	Save(values ...*model.Setting) error
	First() (*model.Setting, error)
	Take() (*model.Setting, error)
	Last() (*model.Setting, error)
	Find() ([]*model.Setting, error)
	FindInBatch(batchSize int, fc func(tx gormgen.Dao, batch int) error) (results []*model.Setting, err error)
	FindInBatches(result *[]*model.Setting, batchSize int, fc func(tx gormgen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Setting) (info gormgen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gormgen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gormgen.ResultInfo, err error)
	Updates(value interface{}) (info gormgen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gormgen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gormgen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gormgen.ResultInfo, err error)
	UpdateFrom(q gormgen.SubQuery) gormgen.Dao
	Attrs(attrs ...field.AssignExpr) ISettingDo
	Assign(attrs ...field.AssignExpr) ISettingDo
	Joins(fields ...field.RelationField) ISettingDo
	Preload(fields ...field.RelationField) ISettingDo
	FirstOrInit() (*model.Setting, error)
	FirstOrCreate() (*model.Setting, error)
	FindByPage(offset int, limit int) (result []*model.Setting, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISettingDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s settingDo) Debug() ISettingDo {
	return s.withDO(s.DO.Debug())
}

func (s settingDo) WithContext(ctx context.Context) ISettingDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s settingDo) ReadDB() ISettingDo {
	return s.Clauses(dbresolver.Read)
}

func (s settingDo) WriteDB() ISettingDo {
	return s.Clauses(dbresolver.Write)
}

func (s settingDo) Session(config *gorm.Session) ISettingDo {
	return s.withDO(s.DO.Session(config))
}

func (s settingDo) Clauses(conds ...clause.Expression) ISettingDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s settingDo) Returning(value interface{}, columns ...string) ISettingDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s settingDo) Not(conds ...gormgen.Condition) ISettingDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s settingDo) Or(conds ...gormgen.Condition) ISettingDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s settingDo) Select(conds ...field.Expr) ISettingDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s settingDo) Where(conds ...gormgen.Condition) ISettingDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s settingDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ISettingDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s settingDo) Order(conds ...field.Expr) ISettingDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s settingDo) Distinct(cols ...field.Expr) ISettingDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s settingDo) Omit(cols ...field.Expr) ISettingDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s settingDo) Join(table schema.Tabler, on ...field.Expr) ISettingDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s settingDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISettingDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s settingDo) RightJoin(table schema.Tabler, on ...field.Expr) ISettingDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s settingDo) Group(cols ...field.Expr) ISettingDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s settingDo) Having(conds ...gormgen.Condition) ISettingDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s settingDo) Limit(limit int) ISettingDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s settingDo) Offset(offset int) ISettingDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s settingDo) Scopes(funcs ...func(gormgen.Dao) gormgen.Dao) ISettingDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s settingDo) Unscoped() ISettingDo {
	return s.withDO(s.DO.Unscoped())
}

func (s settingDo) Create(values ...*model.Setting) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s settingDo) CreateInBatches(values []*model.Setting, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s settingDo) Save(values ...*model.Setting) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s settingDo) First() (*model.Setting, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Setting), nil
	}
}

func (s settingDo) Take() (*model.Setting, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Setting), nil
	}
}

func (s settingDo) Last() (*model.Setting, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Setting), nil
	}
}

func (s settingDo) Find() ([]*model.Setting, error) {
	result, err := s.DO.Find()
	return result.([]*model.Setting), err
}

func (s settingDo) FindInBatch(batchSize int, fc func(tx gormgen.Dao, batch int) error) (results []*model.Setting, err error) {
	buf := make([]*model.Setting, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gormgen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s settingDo) FindInBatches(result *[]*model.Setting, batchSize int, fc func(tx gormgen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s settingDo) Attrs(attrs ...field.AssignExpr) ISettingDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s settingDo) Assign(attrs ...field.AssignExpr) ISettingDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s settingDo) Joins(fields ...field.RelationField) ISettingDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s settingDo) Preload(fields ...field.RelationField) ISettingDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s settingDo) FirstOrInit() (*model.Setting, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Setting), nil
	}
}

func (s settingDo) FirstOrCreate() (*model.Setting, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Setting), nil
	}
}

func (s settingDo) FindByPage(offset int, limit int) (result []*model.Setting, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s settingDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s settingDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s settingDo) Delete(models ...*model.Setting) (result gormgen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *settingDo) withDO(do gormgen.Dao) *settingDo {
	s.DO = *do.(*gormgen.DO)
	return s
}
