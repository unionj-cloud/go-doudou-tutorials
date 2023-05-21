// Code generated by gorm.io/gen. YOU CAN EDIT.
// Code generated by gorm.io/gen. YOU CAN EDIT.
// Code generated by gorm.io/gen. YOU CAN EDIT.

package dto

import (
	"time"
)

//go:generate go-doudou name --file $GOFILE

// TInvalidToken mapped from table <t_invalid_token>
type TInvalidToken struct {
	ID       int32      `json:"id"`
	Token    string     `json:"token"`    // revoked token
	ExpireAt time.Time  `json:"expireAt"` // expire time
	CreateAt *time.Time `json:"createAt"`
	UpdateAt *time.Time `json:"updateAt"`
	DeleteAt time.Time  `json:"deleteAt"`
}