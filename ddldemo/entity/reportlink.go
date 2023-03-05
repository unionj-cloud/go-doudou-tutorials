/**
* Generated by go-doudou v2.0.4.
* You can edit it as your need.
 */
package entity

//dd:table
type ReportLink struct {
	Id          string  `dd:"pk;type:varchar(32);extra:comment '主键id'"`
	ReportId    *string `dd:"type:varchar(32);extra:comment '积木设计器id'"`
	Parameter   *string `dd:"type:text;extra:comment '参数'"`
	EjectType   *string `dd:"type:varchar(1);extra:comment '弹出方式（0 当前页面 1 新窗口）'"`
	LinkName    *string `dd:"type:varchar(255);extra:comment '链接名称'"`
	ApiMethod   *string `dd:"type:varchar(1);extra:comment '请求方法0-get,1-post'"`
	LinkType    *string `dd:"type:varchar(1);extra:comment '链接方式(0 网络报表 1 网络连接 2 图表联动)'"`
	ApiUrl      *string `dd:"type:varchar(1000);extra:comment '外网api'"`
	LinkChartId *string `dd:"type:varchar(50);extra:comment '联动图表的ID'"`
}