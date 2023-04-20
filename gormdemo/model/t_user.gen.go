// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTUser = "t_user"

// TUser mapped from table <t_user>
type TUser struct {
	ID       int32      `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Username string     `gorm:"column:username;not null" json:"username"` // username
	Password string     `gorm:"column:password;not null" json:"password"` // password
	Name     string     `gorm:"column:name;not null" json:"name"`         // real name
	Phone    string     `gorm:"column:phone;not null" json:"phone"`       // phone number
	Dept     string     `gorm:"column:dept;not null" json:"dept"`         // department
	Avatar   string     `gorm:"column:avatar;not null" json:"avatar"`     // user avatar
	CreateAt *time.Time `gorm:"column:create_at;default:CURRENT_TIMESTAMP" json:"create_at"`
	UpdateAt *time.Time `gorm:"column:update_at;default:CURRENT_TIMESTAMP" json:"update_at"`
	DeleteAt time.Time  `gorm:"column:delete_at" json:"delete_at"`
}

// TableName TUser's table name
func (*TUser) TableName() string {
	return TableNameTUser
}
