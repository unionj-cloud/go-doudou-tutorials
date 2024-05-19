// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"fmt"
	"testgggorm2/config"

	"github.com/unionj-cloud/go-doudou/v2/toolkit/stringutils"
)

var TableNameCart string

func init() {
	if stringutils.IsNotEmpty(config.G_Config.Db.Name) {
		TableNameCart = fmt.Sprintf("%s.cart", config.G_Config.Db.Name)
	} else {
		TableNameCart = "cart"
	}
}

// Cart mapped from table <cart>
type Cart struct {
	ID             int32    `gorm:"column:id;type:int(11);primaryKey;autoIncrement:true;comment:主键" json:"id,omitempty"` // 主键
	Title          *string  `gorm:"column:title;type:varchar(250);comment:标题" json:"title,omitempty"`                    // 标题
	Price          *float64 `gorm:"column:price;type:decimal(10,2);default:0.00" json:"price,omitempty"`
	GoodsVersion   *string  `gorm:"column:goods_version;type:varchar(50);comment:版本" json:"goodsVersion,omitempty"`        // 版本
	Num            *int32   `gorm:"column:num;type:int(11);comment:数量" json:"num,omitempty"`                               // 数量
	ProductGift    *string  `gorm:"column:product_gift;type:varchar(100);comment:商品礼物" json:"productGift,omitempty"`       // 商品礼物
	ProductFitting *string  `gorm:"column:product_fitting;type:varchar(100);comment:商品搭配" json:"productFitting,omitempty"` // 商品搭配
	ProductColor   *string  `gorm:"column:product_color;type:varchar(50);comment:商品颜色" json:"productColor,omitempty"`      // 商品颜色
	ProductImg     *string  `gorm:"column:product_img;type:varchar(150);comment:商品图片" json:"productImg,omitempty"`         // 商品图片
	ProductAttr    *string  `gorm:"column:product_attr;type:varchar(100);comment:商品属性" json:"productAttr,omitempty"`       // 商品属性
}

// TableName Cart's table name
func (*Cart) TableName() string {
	return TableNameCart
}