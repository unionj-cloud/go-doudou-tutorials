/**
* Generated by go-doudou v2.0.5.
* You can edit it as your need.
 */
package entity

import (
	"time"
)

//dd:table
type CommonFile struct {
	Id                  *string    `dd:"type:varchar(50)"`
	DelFlag             *int8      `dd:"type:tinyint(4)"`
	RelateBusinessTable *string    `dd:"type:varchar(255)"`
	RelateBusinessId    *string    `dd:"type:varchar(50)"`
	FileName            *string    `dd:"type:varchar(255)"`
	FileType            *string    `dd:"type:varchar(255)"`
	FileSize            *int64     `dd:"type:bigint(20)"`
	FilePath            *string    `dd:"type:varchar(255)"`
	CreateTime          *time.Time `dd:"type:datetime"`
	CreateBy            *string    `dd:"type:varchar(50)"`
	SysOrgCode          *string    `dd:"type:varchar(255)"`
}