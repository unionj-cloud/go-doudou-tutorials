// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"gorm.io/gorm"
)

const TableNameTAuthor = "t_author"

// TAuthor mapped from table <t_author>
type TAuthor struct {
	ID       int32          `gorm:"column:id;type:int(11);primaryKey;autoIncrement:true" json:"id,omitempty"`
	Book     *string        `gorm:"column:book;type:varchar(45)" json:"book,omitempty"`
	DeleteAt gorm.DeletedAt `gorm:"column:delete_at;type:datetime" json:"delete_at,omitempty"`
}

// TableName TAuthor's table name
func (*TAuthor) TableName() string {
	return TableNameTAuthor
}
