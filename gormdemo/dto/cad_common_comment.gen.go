// Code generated by gorm.io/gen. YOU CAN EDIT.
// Code generated by gorm.io/gen. YOU CAN EDIT.
// Code generated by gorm.io/gen. YOU CAN EDIT.

package dto

import (
	"time"
)

//go:generate go-doudou name --file $GOFILE

// CadCommonComment mapped from table <cad_common_comment>
type CadCommonComment struct {
	ID                  string    `json:"id"`
	DelFlag             int32     `json:"delFlag"`
	RelateBusinessTable string    `json:"relateBusinessTable"`
	RelateBusinessID    string    `json:"relateBusinessID"`
	Content             string    `json:"content"`
	ReplyTo             string    `json:"replyTo"`
	CreateTime          time.Time `json:"createTime"`
	CreateBy            string    `json:"createBy"`
	SysOrgCode          string    `json:"sysOrgCode"`
}
