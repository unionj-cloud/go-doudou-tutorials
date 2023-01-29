/**
* Generated by go-doudou v2.0.4.
* You can edit it as your need.
 */
package entity

import (
	"time"
)

//dd:table
type CommonLog struct {
	Id                  string     `dd:"pk;type:varchar(50);extra:comment '评论ID'"`
	RelateBusinessTable *string    `dd:"type:varchar(255);extra:comment '关联业务表'"`
	RelateBusinessId    *string    `dd:"type:varchar(50);extra:comment '关联业务ID'"`
	Type                *string    `dd:"type:varchar(255);extra:comment '操作类型'"`
	SubType             *string    `dd:"type:varchar(255);extra:comment '操作子类型'"`
	Desc                *string    `dd:"type:varchar(255);extra:comment '操作描述'"`
	CreateTime          *time.Time `dd:"type:datetime;extra:comment '创建时间'"`
	CreateBy            *string    `dd:"type:varchar(50);extra:comment '创建人id'"`
	SysOrgCode          *string    `dd:"type:varchar(255);extra:comment '所属部门'"`
}
