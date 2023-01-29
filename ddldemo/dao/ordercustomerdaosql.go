/**
* Generated by go-doudou v2.0.4.
* You can edit it as your need.
 */
package dao

var ordercustomerdaosql = `{{define "NoneZeroSet"}}
	{{- if .CreateBy}}
	` + "`" + `create_by` + "`" + `=:create_by,
	{{- end}}
	{{- if .CreateTime}}
	` + "`" + `create_time` + "`" + `=:create_time,
	{{- end}}
	{{- if .UpdateBy}}
	` + "`" + `update_by` + "`" + `=:update_by,
	{{- end}}
	{{- if .UpdateTime}}
	` + "`" + `update_time` + "`" + `=:update_time,
	{{- end}}
	{{- if .SysOrgCode}}
	` + "`" + `sys_org_code` + "`" + `=:sys_org_code,
	{{- end}}
	{{- if .Name}}
	` + "`" + `name` + "`" + `=:name,
	{{- end}}
	{{- if .Sex}}
	` + "`" + `sex` + "`" + `=:sex,
	{{- end}}
	{{- if .Birthday}}
	` + "`" + `birthday` + "`" + `=:birthday,
	{{- end}}
	{{- if .Age}}
	` + "`" + `age` + "`" + `=:age,
	{{- end}}
	{{- if .Address}}
	` + "`" + `address` + "`" + `=:address,
	{{- end}}
	{{- if .OrderMainId}}
	` + "`" + `order_main_id` + "`" + `=:order_main_id,
	{{- end}}
{{end}}

{{define "InsertOrderCustomer"}}
INSERT INTO ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `ces_order_customer` + "`" + `
(
` + "`" + `id` + "`" + `,
` + "`" + `create_by` + "`" + `,
` + "`" + `create_time` + "`" + `,
` + "`" + `update_by` + "`" + `,
` + "`" + `update_time` + "`" + `,
` + "`" + `sys_org_code` + "`" + `,
` + "`" + `name` + "`" + `,
` + "`" + `sex` + "`" + `,
` + "`" + `birthday` + "`" + `,
` + "`" + `age` + "`" + `,
` + "`" + `address` + "`" + `,
` + "`" + `order_main_id` + "`" + `)
VALUES (
	   :id,
	   :create_by,
	   :create_time,
	   :update_by,
	   :update_time,
	   :sys_org_code,
	   :name,
	   :sex,
	   :birthday,
	   :age,
	   :address,
	   :order_main_id)
{{end}}

{{define "UpdateOrderCustomer"}}
UPDATE ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `ces_order_customer` + "`" + `
SET
	` + "`" + `create_by` + "`" + `=:create_by,
	` + "`" + `create_time` + "`" + `=:create_time,
	` + "`" + `update_by` + "`" + `=:update_by,
	` + "`" + `update_time` + "`" + `=:update_time,
	` + "`" + `sys_org_code` + "`" + `=:sys_org_code,
	` + "`" + `name` + "`" + `=:name,
	` + "`" + `sex` + "`" + `=:sex,
	` + "`" + `birthday` + "`" + `=:birthday,
	` + "`" + `age` + "`" + `=:age,
	` + "`" + `address` + "`" + `=:address,
	` + "`" + `order_main_id` + "`" + `=:order_main_id
WHERE
    ` + "`" + `id` + "`" + ` =:id
{{end}}

{{define "UpdateOrderCustomerNoneZero"}}
UPDATE ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `ces_order_customer` + "`" + `
SET
    {{Eval "NoneZeroSet" . | TrimSuffix ","}}
WHERE
    ` + "`" + `id` + "`" + `=:id
{{end}}

{{define "UpsertOrderCustomer"}}
INSERT INTO ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `ces_order_customer` + "`" + `
(
` + "`" + `id` + "`" + `,
` + "`" + `create_by` + "`" + `,
` + "`" + `create_time` + "`" + `,
` + "`" + `update_by` + "`" + `,
` + "`" + `update_time` + "`" + `,
` + "`" + `sys_org_code` + "`" + `,
` + "`" + `name` + "`" + `,
` + "`" + `sex` + "`" + `,
` + "`" + `birthday` + "`" + `,
` + "`" + `age` + "`" + `,
` + "`" + `address` + "`" + `,
` + "`" + `order_main_id` + "`" + `)
VALUES (
        :id,
        :create_by,
        :create_time,
        :update_by,
        :update_time,
        :sys_org_code,
        :name,
        :sex,
        :birthday,
        :age,
        :address,
        :order_main_id) ON DUPLICATE KEY
UPDATE
		` + "`" + `create_by` + "`" + `=:create_by,
		` + "`" + `create_time` + "`" + `=:create_time,
		` + "`" + `update_by` + "`" + `=:update_by,
		` + "`" + `update_time` + "`" + `=:update_time,
		` + "`" + `sys_org_code` + "`" + `=:sys_org_code,
		` + "`" + `name` + "`" + `=:name,
		` + "`" + `sex` + "`" + `=:sex,
		` + "`" + `birthday` + "`" + `=:birthday,
		` + "`" + `age` + "`" + `=:age,
		` + "`" + `address` + "`" + `=:address,
		` + "`" + `order_main_id` + "`" + `=:order_main_id
{{end}}

{{define "UpsertOrderCustomerNoneZero"}}
INSERT INTO ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `ces_order_customer` + "`" + `
(
` + "`" + `id` + "`" + `,
` + "`" + `create_by` + "`" + `,
` + "`" + `create_time` + "`" + `,
` + "`" + `update_by` + "`" + `,
` + "`" + `update_time` + "`" + `,
` + "`" + `sys_org_code` + "`" + `,
` + "`" + `name` + "`" + `,
` + "`" + `sex` + "`" + `,
` + "`" + `birthday` + "`" + `,
` + "`" + `age` + "`" + `,
` + "`" + `address` + "`" + `,
` + "`" + `order_main_id` + "`" + `)
VALUES (
        :id,
        :create_by,
        :create_time,
        :update_by,
        :update_time,
        :sys_org_code,
        :name,
        :sex,
        :birthday,
        :age,
        :address,
        :order_main_id) ON DUPLICATE KEY
UPDATE
		{{Eval "NoneZeroSet" . | TrimSuffix ","}}
{{end}}

{{define "GetOrderCustomer"}}
select *
from ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `ces_order_customer` + "`" + `
where ` + "`" + `id` + "`" + ` = ?
{{end}}

{{define "UpdateOrderCustomers"}}
UPDATE ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `ces_order_customer` + "`" + `
SET
	` + "`" + `create_by` + "`" + `=:create_by,
	` + "`" + `create_time` + "`" + `=:create_time,
	` + "`" + `update_by` + "`" + `=:update_by,
	` + "`" + `update_time` + "`" + `=:update_time,
	` + "`" + `sys_org_code` + "`" + `=:sys_org_code,
	` + "`" + `name` + "`" + `=:name,
	` + "`" + `sex` + "`" + `=:sex,
	` + "`" + `birthday` + "`" + `=:birthday,
	` + "`" + `age` + "`" + `=:age,
	` + "`" + `address` + "`" + `=:address,
	` + "`" + `order_main_id` + "`" + `=:order_main_id
{{end}}

{{define "UpdateOrderCustomersNoneZero"}}
UPDATE ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `ces_order_customer` + "`" + `
SET
    {{Eval "NoneZeroSet" . | TrimSuffix ","}}
{{end}}

{{define "InsertIgnoreOrderCustomer"}}
INSERT IGNORE INTO ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `ces_order_customer` + "`" + `
(
` + "`" + `id` + "`" + `,
` + "`" + `create_by` + "`" + `,
` + "`" + `create_time` + "`" + `,
` + "`" + `update_by` + "`" + `,
` + "`" + `update_time` + "`" + `,
` + "`" + `sys_org_code` + "`" + `,
` + "`" + `name` + "`" + `,
` + "`" + `sex` + "`" + `,
` + "`" + `birthday` + "`" + `,
` + "`" + `age` + "`" + `,
` + "`" + `address` + "`" + `,
` + "`" + `order_main_id` + "`" + `)
VALUES (
	   :id,
	   :create_by,
	   :create_time,
	   :update_by,
	   :update_time,
	   :sys_org_code,
	   :name,
	   :sex,
	   :birthday,
	   :age,
	   :address,
	   :order_main_id)
{{end}}

{{define "UpdateClauseOrderCustomer"}}
ON DUPLICATE KEY
UPDATE
		` + "`" + `create_by` + "`" + `=VALUES(create_by),
		` + "`" + `create_time` + "`" + `=VALUES(create_time),
		` + "`" + `update_by` + "`" + `=VALUES(update_by),
		` + "`" + `update_time` + "`" + `=VALUES(update_time),
		` + "`" + `sys_org_code` + "`" + `=VALUES(sys_org_code),
		` + "`" + `name` + "`" + `=VALUES(name),
		` + "`" + `sex` + "`" + `=VALUES(sex),
		` + "`" + `birthday` + "`" + `=VALUES(birthday),
		` + "`" + `age` + "`" + `=VALUES(age),
		` + "`" + `address` + "`" + `=VALUES(address),
		` + "`" + `order_main_id` + "`" + `=VALUES(order_main_id)
{{end}}

{{define "UpdateClauseSelectOrderCustomer"}}
ON DUPLICATE KEY
UPDATE
		{{- range $i, $co := .Columns}}
		{{- if $i}},{{end}}
		` + "`" + `{{$co}}` + "`" + `=VALUES({{$co}})
		{{- end }}
{{end}}`
