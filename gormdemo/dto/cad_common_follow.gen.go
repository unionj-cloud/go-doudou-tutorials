// Code generated by gorm.io/gen. YOU CAN EDIT.
// Code generated by gorm.io/gen. YOU CAN EDIT.
// Code generated by gorm.io/gen. YOU CAN EDIT.

package dto

import (
	"time"
)

//go:generate go-doudou name --file $GOFILE

// CadCommonFollow mapped from table <cad_common_follow>
type CadCommonFollow struct {
	ID                  string    `json:"id"`
	RelateBusinessTable string    `json:"relateBusinessTable"`
	RelateBusinessID    string    `json:"relateBusinessID"`
	CreateTime          time.Time `json:"createTime"`
	CreateBy            string    `json:"createBy"`
	SysOrgCode          string    `json:"sysOrgCode"`
}
