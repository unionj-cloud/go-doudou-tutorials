// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"gorm.io/gorm"
)

const TableNameTPost = "t_post"

// TPost mapped from table <t_post>
type TPost struct {
	ID       int32          `gorm:"column:id;type:int(11);primaryKey;autoIncrement:true" json:"id,omitempty"`
	Title    *string        `gorm:"column:title;type:varchar(45)" json:"title,omitempty"`
	DeleteAt gorm.DeletedAt `gorm:"column:delete_at;type:datetime" json:"delete_at,omitempty"`
}

// TableName TPost's table name
func (*TPost) TableName() string {
	return TableNameTPost
}
