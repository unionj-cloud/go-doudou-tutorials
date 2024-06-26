// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"fmt"
	"testgggorm2/config"

	"github.com/unionj-cloud/go-doudou/v2/toolkit/stringutils"
)

var TableNameProductType string

func init() {
	if stringutils.IsNotEmpty(config.G_Config.Db.Name) {
		TableNameProductType = fmt.Sprintf("%s.product_type", config.G_Config.Db.Name)
	} else {
		TableNameProductType = "product_type"
	}
}

// ProductType mapped from table <product_type>
type ProductType struct {
	ID          int32   `gorm:"column:id;type:int(11);primaryKey;autoIncrement:true" json:"id,omitempty"`
	Title       *string `gorm:"column:title;type:varchar(100);comment:标题" json:"title,omitempty"`             // 标题
	Description *string `gorm:"column:description;type:varchar(500);comment:描述" json:"description,omitempty"` // 描述
	Status      *int32  `gorm:"column:status;type:tinyint(4);comment:状态" json:"status,omitempty"`             // 状态
	AddTime     *int32  `gorm:"column:add_time;type:int(10);comment:添加时间" json:"addTime,omitempty"`           // 添加时间
}

// TableName ProductType's table name
func (*ProductType) TableName() string {
	return TableNameProductType
}
