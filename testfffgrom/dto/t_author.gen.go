// Code generated by gorm.io/gen. YOU CAN EDIT.
// Code generated by gorm.io/gen. YOU CAN EDIT.
// Code generated by gorm.io/gen. YOU CAN EDIT.

package dto

import (
	"gorm.io/gorm"
)

//go:generate go-doudou name --file $GOFILE

// TAuthor mapped from table <t_author>
type TAuthor struct {
	ID       int32          `json:"id"`
	Book     *string        `json:"book"`
	DeleteAt gorm.DeletedAt `json:"deleteAt"`
}
