// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"fmt"
	"testgggorm2/config"

	"github.com/unionj-cloud/go-doudou/v2/toolkit/stringutils"
)

var TableNameProductColor string

func init() {
	if stringutils.IsNotEmpty(config.G_Config.Db.Name) {
		TableNameProductColor = fmt.Sprintf("%s.product_color", config.G_Config.Db.Name)
	} else {
		TableNameProductColor = "product_color"
	}
}

// ProductColor mapped from table <product_color>
type ProductColor struct {
	ID         int32   `gorm:"column:id;type:int(11);primaryKey;autoIncrement:true" json:"id,omitempty"`
	ColorName  *string `gorm:"column:color_name;type:varchar(100);comment:颜色名字" json:"colorName,omitempty"`  // 颜色名字
	ColorValue *string `gorm:"column:color_value;type:varchar(100);comment:颜色值" json:"colorValue,omitempty"` // 颜色值
	Status     *int32  `gorm:"column:status;type:tinyint(4);comment:状态" json:"status,omitempty"`             // 状态
	Checked    *int32  `gorm:"column:checked;type:tinyint(4);comment:是否检验" json:"checked,omitempty"`         // 是否检验
}

// TableName ProductColor's table name
func (*ProductColor) TableName() string {
	return TableNameProductColor
}