// Code generated by gorm.io/gen. YOU CAN EDIT.
// Code generated by gorm.io/gen. YOU CAN EDIT.
// Code generated by gorm.io/gen. YOU CAN EDIT.

package dto

import (
	"time"
)

//go:generate go-doudou name --file $GOFILE

// CadCommonFile mapped from table <cad_common_file>
type CadCommonFile struct {
	ID                  *string    `json:"id"`
	DelFlag             *int32     `json:"delFlag"`
	RelateBusinessTable *string    `json:"relateBusinessTable"`
	RelateBusinessID    *string    `json:"relateBusinessID"`
	FileName            *string    `json:"fileName"`
	FileType            *string    `json:"fileType"`
	FileSize            *int64     `json:"fileSize"`
	FilePath            *string    `json:"filePath"`
	CreateTime          *time.Time `json:"createTime"`
	CreateBy            *string    `json:"createBy"`
	SysOrgCode          *string    `json:"sysOrgCode"`
}
