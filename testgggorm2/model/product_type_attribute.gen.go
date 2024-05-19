// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"fmt"
	"testgggorm2/config"

	"github.com/unionj-cloud/go-doudou/v2/toolkit/stringutils"
)

var TableNameProductTypeAttribute string

func init() {
	if stringutils.IsNotEmpty(config.G_Config.Db.Name) {
		TableNameProductTypeAttribute = fmt.Sprintf("%s.product_type_attribute", config.G_Config.Db.Name)
	} else {
		TableNameProductTypeAttribute = "product_type_attribute"
	}
}

// ProductTypeAttribute mapped from table <product_type_attribute>
type ProductTypeAttribute struct {
	ID        int32   `gorm:"column:id;type:int(11);primaryKey;autoIncrement:true" json:"id,omitempty"`
	CateID    *int32  `gorm:"column:cate_id;type:int(11);comment:分类编号" json:"cateId,omitempty"`           // 分类编号
	Title     *string `gorm:"column:title;type:varchar(100);comment:标题" json:"title,omitempty"`           // 标题
	AttrType  *int32  `gorm:"column:attr_type;type:tinyint(4);comment:属性类型" json:"attrType,omitempty"`    // 属性类型
	AttrValue *string `gorm:"column:attr_value;type:varchar(100);comment:属性值" json:"attrValue,omitempty"` // 属性值
	Status    *int32  `gorm:"column:status;type:tinyint(4);comment:状态" json:"status,omitempty"`           // 状态
	Sort      *int32  `gorm:"column:sort;type:int(10);comment:排序" json:"sort,omitempty"`                  // 排序
	AddTime   *int32  `gorm:"column:add_time;type:int(10);comment:添加时间" json:"addTime,omitempty"`         // 添加时间
}

// TableName ProductTypeAttribute's table name
func (*ProductTypeAttribute) TableName() string {
	return TableNameProductTypeAttribute
}