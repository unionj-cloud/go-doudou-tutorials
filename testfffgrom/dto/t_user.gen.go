// Code generated by gorm.io/gen. YOU CAN EDIT.
// Code generated by gorm.io/gen. YOU CAN EDIT.
// Code generated by gorm.io/gen. YOU CAN EDIT.

package dto

import (
	"github.com/unionj-cloud/go-doudou/v2/toolkit/customtypes"
	"gorm.io/gorm"
)

//go:generate go-doudou name --file $GOFILE

// TUser mapped from table <t_user>
type TUser struct {
	ID       int32             `json:"id"`
	Username string            `json:"username"` // username
	Password string            `json:"password"` // password
	Name     string            `json:"name"`     // real name
	Phone    string            `json:"phone"`    // phone number
	Dept     string            `json:"dept"`     // department
	Avatar   string            `json:"avatar"`   // user avatar
	CreateAt *customtypes.Time `json:"createAt"`
	UpdateAt *customtypes.Time `json:"updateAt"`
	DeleteAt gorm.DeletedAt    `json:"deleteAt"`
}
