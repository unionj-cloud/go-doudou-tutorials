/**
* Generated by go-doudou v2.0.4.
* You can edit it as your need.
 */
package entity

import (
	"time"
)

//dd:table
type SquareTaskReject struct {
	Id         string     `dd:"pk;type:varchar(50);extra:comment '主键'"`
	TaskId     *string    `dd:"type:varchar(50);extra:comment '任务ID'"`
	Reason     *string    `dd:"type:text;extra:comment '原因'"`
	CreateTime *time.Time `dd:"type:datetime;extra:comment '创建时间'"`
	CreateBy   *string    `dd:"type:varchar(50);extra:comment '创建人id'"`
}
