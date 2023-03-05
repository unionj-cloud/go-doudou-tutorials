/**
* Generated by go-doudou v2.0.4.
* You can edit it as your need.
 */
package entity

import (
	"time"
)

//dd:table
type Knowledge struct {
	Id                   string     `dd:"pk;type:varchar(50);extra:comment 'ID';index:idx_cad_knowledge,1,asc"`
	DelFlag              *int8      `dd:"type:tinyint(1);extra:comment '是否被删除，0--否，1--是'"`
	Title                *string    `dd:"type:varchar(255);extra:comment '标题'"`
	Content              *string    `dd:"type:text;extra:comment '内容'"`
	ViewCount            *int       `dd:"type:int(11);default:'0';extra:comment '浏览次数';index:idx_cad_knowledge,2,asc"`
	Catalog              *string    `dd:"type:text;extra:comment '所属目录（分类字典code），多个时英文逗号分隔'"`
	Tag                  *string    `dd:"type:text;extra:comment '标签（数据字典code），多个时英文逗号分隔'"`
	Status               *string    `dd:"type:varchar(50);extra:comment '状态（数据字典code）'"`
	LastReleaseVersionId *string    `dd:"type:varchar(50);extra:comment '最后发布的版本id'"`
	LastReleaseTime      *time.Time `dd:"type:datetime;extra:comment '最后发布时间'"`
	LastReleaseBy        *string    `dd:"type:varchar(255);extra:comment '最后发布者username'"`
	CreateTime           *time.Time `dd:"type:datetime;extra:comment '创建时间'"`
	CreateBy             *string    `dd:"type:varchar(50);extra:comment '创建人username'"`
	UpdateTime           *time.Time `dd:"type:datetime;extra:comment '更新时间'"`
	UpdateBy             *string    `dd:"type:varchar(50);extra:comment '更新人username'"`
	SysOrgCode           *string    `dd:"type:varchar(255);extra:comment '所属部门'"`
	ReferenceDocuments   *string    `dd:"type:text;extra:comment '参考文献，JSON数组，[{label:文献文本, link:文献链接},...]'"`
	SortOrder            *int       `dd:"type:int(11);default:'0';extra:comment '同目录下的排序'"`
	ActualTitles         *string    `dd:"type:varchar(255);extra:comment '实际使用的标题'"`
}