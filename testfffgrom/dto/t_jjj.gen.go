// Code generated by gorm.io/gen. YOU CAN EDIT.
// Code generated by gorm.io/gen. YOU CAN EDIT.
// Code generated by gorm.io/gen. YOU CAN EDIT.

package dto

import (
	"gorm.io/gorm"
)

//go:generate go-doudou name --file $GOFILE

// TJjj mapped from table <t_jjj>
type TJjj struct {
	ID        int32          `json:"id"`
	JobName   *string        `json:"jobName"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}