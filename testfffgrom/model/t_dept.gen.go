// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"gorm.io/gorm"
)

const TableNameTDept = "t_dept"

// TDept mapped from table <t_dept>
type TDept struct {
	ID        int32          `gorm:"column:id;type:int(11);primaryKey;autoIncrement:true" json:"id,omitempty"`
	Name      string         `gorm:"column:name;type:varchar(45);not null" json:"name,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime" json:"deleted_at,omitempty"`
}

// TableName TDept's table name
func (*TDept) TableName() string {
	return TableNameTDept
}