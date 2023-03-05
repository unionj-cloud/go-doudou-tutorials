/**
* Generated by go-doudou v2.0.4.
* You can edit it as your need.
 */
package dao

var fieldkongjdaosql = `{{define "NoneZeroSet"}}
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
	{{- if .Radio}}
	` + "`" + `radio` + "`" + `=:radio,
	{{- end}}
	{{- if .Checkbox}}
	` + "`" + `checkbox` + "`" + `=:checkbox,
	{{- end}}
	{{- if .SelMut}}
	` + "`" + `sel_mut` + "`" + `=:sel_mut,
	{{- end}}
	{{- if .SelSearch}}
	` + "`" + `sel_search` + "`" + `=:sel_search,
	{{- end}}
	{{- if .Birthday}}
	` + "`" + `birthday` + "`" + `=:birthday,
	{{- end}}
	{{- if .Pic}}
	` + "`" + `pic` + "`" + `=:pic,
	{{- end}}
	{{- if .Files}}
	` + "`" + `files` + "`" + `=:files,
	{{- end}}
	{{- if .Remakr}}
	` + "`" + `remakr` + "`" + `=:remakr,
	{{- end}}
	{{- if .Fuwenb}}
	` + "`" + `fuwenb` + "`" + `=:fuwenb,
	{{- end}}
	{{- if .UserSel}}
	` + "`" + `user_sel` + "`" + `=:user_sel,
	{{- end}}
	{{- if .DepSel}}
	` + "`" + `dep_sel` + "`" + `=:dep_sel,
	{{- end}}
	{{- if .Ddd}}
	` + "`" + `ddd` + "`" + `=:ddd,
	{{- end}}
{{end}}

{{define "InsertFieldKongj"}}
INSERT INTO ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `ces_field_kongj` + "`" + `
(
` + "`" + `id` + "`" + `,
` + "`" + `create_by` + "`" + `,
` + "`" + `create_time` + "`" + `,
` + "`" + `update_by` + "`" + `,
` + "`" + `update_time` + "`" + `,
` + "`" + `sys_org_code` + "`" + `,
` + "`" + `name` + "`" + `,
` + "`" + `sex` + "`" + `,
` + "`" + `radio` + "`" + `,
` + "`" + `checkbox` + "`" + `,
` + "`" + `sel_mut` + "`" + `,
` + "`" + `sel_search` + "`" + `,
` + "`" + `birthday` + "`" + `,
` + "`" + `pic` + "`" + `,
` + "`" + `files` + "`" + `,
` + "`" + `remakr` + "`" + `,
` + "`" + `fuwenb` + "`" + `,
` + "`" + `user_sel` + "`" + `,
` + "`" + `dep_sel` + "`" + `,
` + "`" + `ddd` + "`" + `)
VALUES (
	   :id,
	   :create_by,
	   :create_time,
	   :update_by,
	   :update_time,
	   :sys_org_code,
	   :name,
	   :sex,
	   :radio,
	   :checkbox,
	   :sel_mut,
	   :sel_search,
	   :birthday,
	   :pic,
	   :files,
	   :remakr,
	   :fuwenb,
	   :user_sel,
	   :dep_sel,
	   :ddd)
{{end}}

{{define "UpdateFieldKongj"}}
UPDATE ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `ces_field_kongj` + "`" + `
SET
	` + "`" + `create_by` + "`" + `=:create_by,
	` + "`" + `create_time` + "`" + `=:create_time,
	` + "`" + `update_by` + "`" + `=:update_by,
	` + "`" + `update_time` + "`" + `=:update_time,
	` + "`" + `sys_org_code` + "`" + `=:sys_org_code,
	` + "`" + `name` + "`" + `=:name,
	` + "`" + `sex` + "`" + `=:sex,
	` + "`" + `radio` + "`" + `=:radio,
	` + "`" + `checkbox` + "`" + `=:checkbox,
	` + "`" + `sel_mut` + "`" + `=:sel_mut,
	` + "`" + `sel_search` + "`" + `=:sel_search,
	` + "`" + `birthday` + "`" + `=:birthday,
	` + "`" + `pic` + "`" + `=:pic,
	` + "`" + `files` + "`" + `=:files,
	` + "`" + `remakr` + "`" + `=:remakr,
	` + "`" + `fuwenb` + "`" + `=:fuwenb,
	` + "`" + `user_sel` + "`" + `=:user_sel,
	` + "`" + `dep_sel` + "`" + `=:dep_sel,
	` + "`" + `ddd` + "`" + `=:ddd
WHERE
    ` + "`" + `id` + "`" + ` =:id
{{end}}

{{define "UpdateFieldKongjNoneZero"}}
UPDATE ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `ces_field_kongj` + "`" + `
SET
    {{Eval "NoneZeroSet" . | TrimSuffix ","}}
WHERE
    ` + "`" + `id` + "`" + `=:id
{{end}}

{{define "UpsertFieldKongj"}}
INSERT INTO ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `ces_field_kongj` + "`" + `
(
` + "`" + `id` + "`" + `,
` + "`" + `create_by` + "`" + `,
` + "`" + `create_time` + "`" + `,
` + "`" + `update_by` + "`" + `,
` + "`" + `update_time` + "`" + `,
` + "`" + `sys_org_code` + "`" + `,
` + "`" + `name` + "`" + `,
` + "`" + `sex` + "`" + `,
` + "`" + `radio` + "`" + `,
` + "`" + `checkbox` + "`" + `,
` + "`" + `sel_mut` + "`" + `,
` + "`" + `sel_search` + "`" + `,
` + "`" + `birthday` + "`" + `,
` + "`" + `pic` + "`" + `,
` + "`" + `files` + "`" + `,
` + "`" + `remakr` + "`" + `,
` + "`" + `fuwenb` + "`" + `,
` + "`" + `user_sel` + "`" + `,
` + "`" + `dep_sel` + "`" + `,
` + "`" + `ddd` + "`" + `)
VALUES (
        :id,
        :create_by,
        :create_time,
        :update_by,
        :update_time,
        :sys_org_code,
        :name,
        :sex,
        :radio,
        :checkbox,
        :sel_mut,
        :sel_search,
        :birthday,
        :pic,
        :files,
        :remakr,
        :fuwenb,
        :user_sel,
        :dep_sel,
        :ddd) ON DUPLICATE KEY
UPDATE
		` + "`" + `create_by` + "`" + `=:create_by,
		` + "`" + `create_time` + "`" + `=:create_time,
		` + "`" + `update_by` + "`" + `=:update_by,
		` + "`" + `update_time` + "`" + `=:update_time,
		` + "`" + `sys_org_code` + "`" + `=:sys_org_code,
		` + "`" + `name` + "`" + `=:name,
		` + "`" + `sex` + "`" + `=:sex,
		` + "`" + `radio` + "`" + `=:radio,
		` + "`" + `checkbox` + "`" + `=:checkbox,
		` + "`" + `sel_mut` + "`" + `=:sel_mut,
		` + "`" + `sel_search` + "`" + `=:sel_search,
		` + "`" + `birthday` + "`" + `=:birthday,
		` + "`" + `pic` + "`" + `=:pic,
		` + "`" + `files` + "`" + `=:files,
		` + "`" + `remakr` + "`" + `=:remakr,
		` + "`" + `fuwenb` + "`" + `=:fuwenb,
		` + "`" + `user_sel` + "`" + `=:user_sel,
		` + "`" + `dep_sel` + "`" + `=:dep_sel,
		` + "`" + `ddd` + "`" + `=:ddd
{{end}}

{{define "UpsertFieldKongjNoneZero"}}
INSERT INTO ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `ces_field_kongj` + "`" + `
(
` + "`" + `id` + "`" + `,
` + "`" + `create_by` + "`" + `,
` + "`" + `create_time` + "`" + `,
` + "`" + `update_by` + "`" + `,
` + "`" + `update_time` + "`" + `,
` + "`" + `sys_org_code` + "`" + `,
` + "`" + `name` + "`" + `,
` + "`" + `sex` + "`" + `,
` + "`" + `radio` + "`" + `,
` + "`" + `checkbox` + "`" + `,
` + "`" + `sel_mut` + "`" + `,
` + "`" + `sel_search` + "`" + `,
` + "`" + `birthday` + "`" + `,
` + "`" + `pic` + "`" + `,
` + "`" + `files` + "`" + `,
` + "`" + `remakr` + "`" + `,
` + "`" + `fuwenb` + "`" + `,
` + "`" + `user_sel` + "`" + `,
` + "`" + `dep_sel` + "`" + `,
` + "`" + `ddd` + "`" + `)
VALUES (
        :id,
        :create_by,
        :create_time,
        :update_by,
        :update_time,
        :sys_org_code,
        :name,
        :sex,
        :radio,
        :checkbox,
        :sel_mut,
        :sel_search,
        :birthday,
        :pic,
        :files,
        :remakr,
        :fuwenb,
        :user_sel,
        :dep_sel,
        :ddd) ON DUPLICATE KEY
UPDATE
		{{Eval "NoneZeroSet" . | TrimSuffix ","}}
{{end}}

{{define "GetFieldKongj"}}
select *
from ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `ces_field_kongj` + "`" + `
where ` + "`" + `id` + "`" + ` = ?
{{end}}

{{define "UpdateFieldKongjs"}}
UPDATE ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `ces_field_kongj` + "`" + `
SET
	` + "`" + `create_by` + "`" + `=:create_by,
	` + "`" + `create_time` + "`" + `=:create_time,
	` + "`" + `update_by` + "`" + `=:update_by,
	` + "`" + `update_time` + "`" + `=:update_time,
	` + "`" + `sys_org_code` + "`" + `=:sys_org_code,
	` + "`" + `name` + "`" + `=:name,
	` + "`" + `sex` + "`" + `=:sex,
	` + "`" + `radio` + "`" + `=:radio,
	` + "`" + `checkbox` + "`" + `=:checkbox,
	` + "`" + `sel_mut` + "`" + `=:sel_mut,
	` + "`" + `sel_search` + "`" + `=:sel_search,
	` + "`" + `birthday` + "`" + `=:birthday,
	` + "`" + `pic` + "`" + `=:pic,
	` + "`" + `files` + "`" + `=:files,
	` + "`" + `remakr` + "`" + `=:remakr,
	` + "`" + `fuwenb` + "`" + `=:fuwenb,
	` + "`" + `user_sel` + "`" + `=:user_sel,
	` + "`" + `dep_sel` + "`" + `=:dep_sel,
	` + "`" + `ddd` + "`" + `=:ddd
{{end}}

{{define "UpdateFieldKongjsNoneZero"}}
UPDATE ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `ces_field_kongj` + "`" + `
SET
    {{Eval "NoneZeroSet" . | TrimSuffix ","}}
{{end}}

{{define "InsertIgnoreFieldKongj"}}
INSERT IGNORE INTO ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `ces_field_kongj` + "`" + `
(
` + "`" + `id` + "`" + `,
` + "`" + `create_by` + "`" + `,
` + "`" + `create_time` + "`" + `,
` + "`" + `update_by` + "`" + `,
` + "`" + `update_time` + "`" + `,
` + "`" + `sys_org_code` + "`" + `,
` + "`" + `name` + "`" + `,
` + "`" + `sex` + "`" + `,
` + "`" + `radio` + "`" + `,
` + "`" + `checkbox` + "`" + `,
` + "`" + `sel_mut` + "`" + `,
` + "`" + `sel_search` + "`" + `,
` + "`" + `birthday` + "`" + `,
` + "`" + `pic` + "`" + `,
` + "`" + `files` + "`" + `,
` + "`" + `remakr` + "`" + `,
` + "`" + `fuwenb` + "`" + `,
` + "`" + `user_sel` + "`" + `,
` + "`" + `dep_sel` + "`" + `,
` + "`" + `ddd` + "`" + `)
VALUES (
	   :id,
	   :create_by,
	   :create_time,
	   :update_by,
	   :update_time,
	   :sys_org_code,
	   :name,
	   :sex,
	   :radio,
	   :checkbox,
	   :sel_mut,
	   :sel_search,
	   :birthday,
	   :pic,
	   :files,
	   :remakr,
	   :fuwenb,
	   :user_sel,
	   :dep_sel,
	   :ddd)
{{end}}

{{define "UpdateClauseFieldKongj"}}
ON DUPLICATE KEY
UPDATE
		` + "`" + `create_by` + "`" + `=VALUES(create_by),
		` + "`" + `create_time` + "`" + `=VALUES(create_time),
		` + "`" + `update_by` + "`" + `=VALUES(update_by),
		` + "`" + `update_time` + "`" + `=VALUES(update_time),
		` + "`" + `sys_org_code` + "`" + `=VALUES(sys_org_code),
		` + "`" + `name` + "`" + `=VALUES(name),
		` + "`" + `sex` + "`" + `=VALUES(sex),
		` + "`" + `radio` + "`" + `=VALUES(radio),
		` + "`" + `checkbox` + "`" + `=VALUES(checkbox),
		` + "`" + `sel_mut` + "`" + `=VALUES(sel_mut),
		` + "`" + `sel_search` + "`" + `=VALUES(sel_search),
		` + "`" + `birthday` + "`" + `=VALUES(birthday),
		` + "`" + `pic` + "`" + `=VALUES(pic),
		` + "`" + `files` + "`" + `=VALUES(files),
		` + "`" + `remakr` + "`" + `=VALUES(remakr),
		` + "`" + `fuwenb` + "`" + `=VALUES(fuwenb),
		` + "`" + `user_sel` + "`" + `=VALUES(user_sel),
		` + "`" + `dep_sel` + "`" + `=VALUES(dep_sel),
		` + "`" + `ddd` + "`" + `=VALUES(ddd)
{{end}}

{{define "UpdateClauseSelectFieldKongj"}}
ON DUPLICATE KEY
UPDATE
		{{- range $i, $co := .Columns}}
		{{- if $i}},{{end}}
		` + "`" + `{{$co}}` + "`" + `=VALUES({{$co}})
		{{- end }}
{{end}}`