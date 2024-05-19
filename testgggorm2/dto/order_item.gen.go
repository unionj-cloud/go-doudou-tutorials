// Code generated by gorm.io/gen. YOU CAN EDIT.
// Code generated by gorm.io/gen. YOU CAN EDIT.
// Code generated by gorm.io/gen. YOU CAN EDIT.

package dto

// OrderItem mapped from table <order_item>
type OrderItem struct {
	ID           int32    `gorm:"column:id;type:int(11);primaryKey;autoIncrement:true;comment:订单编号" json:"id,omitempty"`           // 订单编号
	OrderID      *int32   `gorm:"column:order_id;type:int(11);comment:订单编号" json:"orderId,omitempty"`                              // 订单编号
	UID          *int32   `gorm:"column:uid;type:int(11);comment:用户编号" json:"uid,omitempty"`                                       // 用户编号
	ProductTitle *string  `gorm:"column:product_title;type:varchar(100);comment:商品标题" json:"productTitle,omitempty"`               // 商品标题
	ProductID    *int32   `gorm:"column:product_id;type:int(11);comment:商品编号" json:"productId,omitempty"`                          // 商品编号
	ProductImg   *string  `gorm:"column:product_img;type:varchar(200);comment:商品图片" json:"productImg,omitempty"`                   // 商品图片
	ProductPrice *float64 `gorm:"column:product_price;type:decimal(10,2);default:0.00;comment:商品价格" json:"productPrice,omitempty"` // 商品价格
	ProductNum   *int32   `gorm:"column:product_num;type:int(10);comment:商品数量" json:"productNum,omitempty"`                        // 商品数量
	GoodsVersion *string  `gorm:"column:goods_version;type:varchar(100);comment:商品版本" json:"goodsVersion,omitempty"`               // 商品版本
	GoodsColor   *string  `gorm:"column:goods_color;type:varchar(100);comment:商品颜色" json:"goodsColor,omitempty"`                   // 商品颜色
	AddTime      *int32   `gorm:"column:add_time;type:int(10);comment:添加时间" json:"addTime,omitempty"`                              // 添加时间
}
