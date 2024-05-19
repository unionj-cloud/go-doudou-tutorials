// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"fmt"
	"testgggorm2/config"

	"github.com/unionj-cloud/go-doudou/v2/toolkit/stringutils"
)

var TableNameOrder string

func init() {
	if stringutils.IsNotEmpty(config.G_Config.Db.Name) {
		TableNameOrder = fmt.Sprintf("%s.order", config.G_Config.Db.Name)
	} else {
		TableNameOrder = "order"
	}
}

// Order mapped from table <order>
type Order struct {
	ID          int32    `gorm:"column:id;type:int(11);primaryKey;autoIncrement:true;comment:编号" json:"id,omitempty"`   // 编号
	OrderID     *string  `gorm:"column:order_id;type:varchar(100);comment:订单编号" json:"orderId,omitempty"`               // 订单编号
	UID         *int32   `gorm:"column:uid;type:int(11);comment:用户编号" json:"uid,omitempty"`                             // 用户编号
	AllPrice    *float64 `gorm:"column:all_price;type:decimal(10,2);default:0.00;comment:价格" json:"allPrice,omitempty"` // 价格
	Phone       *string  `gorm:"column:phone;type:varchar(30);comment:电话" json:"phone,omitempty"`                       // 电话
	Name        *string  `gorm:"column:name;type:varchar(100);comment:名字" json:"name,omitempty"`                        // 名字
	Address     *string  `gorm:"column:address;type:varchar(250);comment:地址" json:"address,omitempty"`                  // 地址
	Zipcode     *string  `gorm:"column:zipcode;type:varchar(30);comment:邮编" json:"zipcode,omitempty"`                   // 邮编
	PayStatus   *int32   `gorm:"column:pay_status;type:tinyint(4);comment:支付状态" json:"payStatus,omitempty"`             // 支付状态
	PayType     *int32   `gorm:"column:pay_type;type:tinyint(4);comment:支付类型" json:"payType,omitempty"`                 // 支付类型
	OrderStatus *int32   `gorm:"column:order_status;type:tinyint(4);comment:订单状态" json:"orderStatus,omitempty"`         // 订单状态
	AddTime     *int32   `gorm:"column:add_time;type:int(10);comment:添加时间" json:"addTime,omitempty"`                    // 添加时间
}

// TableName Order's table name
func (*Order) TableName() string {
	return TableNameOrder
}
