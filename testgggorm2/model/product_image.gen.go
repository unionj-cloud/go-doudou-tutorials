// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"fmt"
	"testgggorm2/config"

	"github.com/unionj-cloud/go-doudou/v2/toolkit/stringutils"
)

var TableNameProductImage string

func init() {
	if stringutils.IsNotEmpty(config.G_Config.Db.Name) {
		TableNameProductImage = fmt.Sprintf("%s.product_image", config.G_Config.Db.Name)
	} else {
		TableNameProductImage = "product_image"
	}
}

// ProductImage mapped from table <product_image>
type ProductImage struct {
	ID        int32   `gorm:"column:id;type:int(11);primaryKey;autoIncrement:true" json:"id,omitempty"`
	ProductID *int32  `gorm:"column:product_id;type:int(11);comment:商品编号" json:"productId,omitempty"` // 商品编号
	ImgURL    *string `gorm:"column:img_url;type:varchar(250);comment:图片地址" json:"imgUrl,omitempty"`  // 图片地址
	ColorID   *int32  `gorm:"column:color_id;type:int(11);comment:颜色编号" json:"colorId,omitempty"`     // 颜色编号
	Sort      *int32  `gorm:"column:sort;type:int(10);comment:排序" json:"sort,omitempty"`              // 排序
	AddTime   *int32  `gorm:"column:add_time;type:int(10);comment:添加时间" json:"addTime,omitempty"`     // 添加时间
	Status    *int32  `gorm:"column:status;type:tinyint(4);comment:状态" json:"status,omitempty"`       // 状态
}

// TableName ProductImage's table name
func (*ProductImage) TableName() string {
	return TableNameProductImage
}