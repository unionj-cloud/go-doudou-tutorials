// Code generated by gorm.io/gen. YOU CAN EDIT.
// Code generated by gorm.io/gen. YOU CAN EDIT.
// Code generated by gorm.io/gen. YOU CAN EDIT.

package dto

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