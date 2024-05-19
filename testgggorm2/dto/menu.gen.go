// Code generated by gorm.io/gen. YOU CAN EDIT.
// Code generated by gorm.io/gen. YOU CAN EDIT.
// Code generated by gorm.io/gen. YOU CAN EDIT.

package dto

// Menu mapped from table <menu>
type Menu struct {
	ID        int32   `gorm:"column:id;type:int(11);primaryKey;autoIncrement:true;comment:编号" json:"id,omitempty"` // 编号
	Title     *string `gorm:"column:title;type:varchar(100);comment:标题" json:"title,omitempty"`                    // 标题
	Link      *string `gorm:"column:link;type:varchar(250);comment:连接" json:"link,omitempty"`                      // 连接
	Position  *int32  `gorm:"column:position;type:int(10);comment:位置" json:"position,omitempty"`                   // 位置
	IsOpennew *int32  `gorm:"column:is_opennew;type:int(10);comment:是否新打开" json:"isOpennew,omitempty"`             // 是否新打开
	Relation  *string `gorm:"column:relation;type:varchar(100);comment:关系" json:"relation,omitempty"`              // 关系
	Sort      *int32  `gorm:"column:sort;type:int(10);comment:排序" json:"sort,omitempty"`                           // 排序
	Status    *int32  `gorm:"column:status;type:int(10);comment:状态" json:"status,omitempty"`                       // 状态
	AddTime   *int32  `gorm:"column:add_time;type:int(10);comment:添加时间" json:"addTime,omitempty"`                  // 添加时间
}
