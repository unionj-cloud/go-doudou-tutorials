/**
* Generated by go-doudou v2.0.4.
* You can edit it as your need.
 */
package dao

var reportdbdaosql = `{{define "NoneZeroSet"}}
	{{- if .JimuReportId}}
	` + "`" + `jimu_report_id` + "`" + `=:jimu_report_id,
	{{- end}}
	{{- if .CreateBy}}
	` + "`" + `create_by` + "`" + `=:create_by,
	{{- end}}
	{{- if .UpdateBy}}
	` + "`" + `update_by` + "`" + `=:update_by,
	{{- end}}
	{{- if .CreateTime}}
	` + "`" + `create_time` + "`" + `=:create_time,
	{{- end}}
	{{- if .UpdateTime}}
	` + "`" + `update_time` + "`" + `=:update_time,
	{{- end}}
	{{- if .DbCode}}
	` + "`" + `db_code` + "`" + `=:db_code,
	{{- end}}
	{{- if .DbChName}}
	` + "`" + `db_ch_name` + "`" + `=:db_ch_name,
	{{- end}}
	{{- if .DbType}}
	` + "`" + `db_type` + "`" + `=:db_type,
	{{- end}}
	{{- if .DbTableName}}
	` + "`" + `db_table_name` + "`" + `=:db_table_name,
	{{- end}}
	{{- if .DbDynSql}}
	` + "`" + `db_dyn_sql` + "`" + `=:db_dyn_sql,
	{{- end}}
	{{- if .DbKey}}
	` + "`" + `db_key` + "`" + `=:db_key,
	{{- end}}
	{{- if .TbDbKey}}
	` + "`" + `tb_db_key` + "`" + `=:tb_db_key,
	{{- end}}
	{{- if .TbDbTableName}}
	` + "`" + `tb_db_table_name` + "`" + `=:tb_db_table_name,
	{{- end}}
	{{- if .JavaType}}
	` + "`" + `java_type` + "`" + `=:java_type,
	{{- end}}
	{{- if .JavaValue}}
	` + "`" + `java_value` + "`" + `=:java_value,
	{{- end}}
	{{- if .ApiUrl}}
	` + "`" + `api_url` + "`" + `=:api_url,
	{{- end}}
	{{- if .ApiMethod}}
	` + "`" + `api_method` + "`" + `=:api_method,
	{{- end}}
	{{- if .IsList}}
	` + "`" + `is_list` + "`" + `=:is_list,
	{{- end}}
	{{- if .IsPage}}
	` + "`" + `is_page` + "`" + `=:is_page,
	{{- end}}
	{{- if .DbSource}}
	` + "`" + `db_source` + "`" + `=:db_source,
	{{- end}}
	{{- if .DbSourceType}}
	` + "`" + `db_source_type` + "`" + `=:db_source_type,
	{{- end}}
	{{- if .JsonData}}
	` + "`" + `json_data` + "`" + `=:json_data,
	{{- end}}
	{{- if .ApiConvert}}
	` + "`" + `api_convert` + "`" + `=:api_convert,
	{{- end}}
{{end}}

{{define "InsertReportDb"}}
INSERT INTO ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `jimu_report_db` + "`" + `
(
` + "`" + `id` + "`" + `,
` + "`" + `jimu_report_id` + "`" + `,
` + "`" + `create_by` + "`" + `,
` + "`" + `update_by` + "`" + `,
` + "`" + `create_time` + "`" + `,
` + "`" + `update_time` + "`" + `,
` + "`" + `db_code` + "`" + `,
` + "`" + `db_ch_name` + "`" + `,
` + "`" + `db_type` + "`" + `,
` + "`" + `db_table_name` + "`" + `,
` + "`" + `db_dyn_sql` + "`" + `,
` + "`" + `db_key` + "`" + `,
` + "`" + `tb_db_key` + "`" + `,
` + "`" + `tb_db_table_name` + "`" + `,
` + "`" + `java_type` + "`" + `,
` + "`" + `java_value` + "`" + `,
` + "`" + `api_url` + "`" + `,
` + "`" + `api_method` + "`" + `,
` + "`" + `is_list` + "`" + `,
` + "`" + `is_page` + "`" + `,
` + "`" + `db_source` + "`" + `,
` + "`" + `db_source_type` + "`" + `,
` + "`" + `json_data` + "`" + `,
` + "`" + `api_convert` + "`" + `)
VALUES (
	   :id,
	   :jimu_report_id,
	   :create_by,
	   :update_by,
	   :create_time,
	   :update_time,
	   :db_code,
	   :db_ch_name,
	   :db_type,
	   :db_table_name,
	   :db_dyn_sql,
	   :db_key,
	   :tb_db_key,
	   :tb_db_table_name,
	   :java_type,
	   :java_value,
	   :api_url,
	   :api_method,
	   :is_list,
	   :is_page,
	   :db_source,
	   :db_source_type,
	   :json_data,
	   :api_convert)
{{end}}

{{define "UpdateReportDb"}}
UPDATE ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `jimu_report_db` + "`" + `
SET
	` + "`" + `jimu_report_id` + "`" + `=:jimu_report_id,
	` + "`" + `create_by` + "`" + `=:create_by,
	` + "`" + `update_by` + "`" + `=:update_by,
	` + "`" + `create_time` + "`" + `=:create_time,
	` + "`" + `update_time` + "`" + `=:update_time,
	` + "`" + `db_code` + "`" + `=:db_code,
	` + "`" + `db_ch_name` + "`" + `=:db_ch_name,
	` + "`" + `db_type` + "`" + `=:db_type,
	` + "`" + `db_table_name` + "`" + `=:db_table_name,
	` + "`" + `db_dyn_sql` + "`" + `=:db_dyn_sql,
	` + "`" + `db_key` + "`" + `=:db_key,
	` + "`" + `tb_db_key` + "`" + `=:tb_db_key,
	` + "`" + `tb_db_table_name` + "`" + `=:tb_db_table_name,
	` + "`" + `java_type` + "`" + `=:java_type,
	` + "`" + `java_value` + "`" + `=:java_value,
	` + "`" + `api_url` + "`" + `=:api_url,
	` + "`" + `api_method` + "`" + `=:api_method,
	` + "`" + `is_list` + "`" + `=:is_list,
	` + "`" + `is_page` + "`" + `=:is_page,
	` + "`" + `db_source` + "`" + `=:db_source,
	` + "`" + `db_source_type` + "`" + `=:db_source_type,
	` + "`" + `json_data` + "`" + `=:json_data,
	` + "`" + `api_convert` + "`" + `=:api_convert
WHERE
    ` + "`" + `id` + "`" + ` =:id
{{end}}

{{define "UpdateReportDbNoneZero"}}
UPDATE ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `jimu_report_db` + "`" + `
SET
    {{Eval "NoneZeroSet" . | TrimSuffix ","}}
WHERE
    ` + "`" + `id` + "`" + `=:id
{{end}}

{{define "UpsertReportDb"}}
INSERT INTO ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `jimu_report_db` + "`" + `
(
` + "`" + `id` + "`" + `,
` + "`" + `jimu_report_id` + "`" + `,
` + "`" + `create_by` + "`" + `,
` + "`" + `update_by` + "`" + `,
` + "`" + `create_time` + "`" + `,
` + "`" + `update_time` + "`" + `,
` + "`" + `db_code` + "`" + `,
` + "`" + `db_ch_name` + "`" + `,
` + "`" + `db_type` + "`" + `,
` + "`" + `db_table_name` + "`" + `,
` + "`" + `db_dyn_sql` + "`" + `,
` + "`" + `db_key` + "`" + `,
` + "`" + `tb_db_key` + "`" + `,
` + "`" + `tb_db_table_name` + "`" + `,
` + "`" + `java_type` + "`" + `,
` + "`" + `java_value` + "`" + `,
` + "`" + `api_url` + "`" + `,
` + "`" + `api_method` + "`" + `,
` + "`" + `is_list` + "`" + `,
` + "`" + `is_page` + "`" + `,
` + "`" + `db_source` + "`" + `,
` + "`" + `db_source_type` + "`" + `,
` + "`" + `json_data` + "`" + `,
` + "`" + `api_convert` + "`" + `)
VALUES (
        :id,
        :jimu_report_id,
        :create_by,
        :update_by,
        :create_time,
        :update_time,
        :db_code,
        :db_ch_name,
        :db_type,
        :db_table_name,
        :db_dyn_sql,
        :db_key,
        :tb_db_key,
        :tb_db_table_name,
        :java_type,
        :java_value,
        :api_url,
        :api_method,
        :is_list,
        :is_page,
        :db_source,
        :db_source_type,
        :json_data,
        :api_convert) ON DUPLICATE KEY
UPDATE
		` + "`" + `jimu_report_id` + "`" + `=:jimu_report_id,
		` + "`" + `create_by` + "`" + `=:create_by,
		` + "`" + `update_by` + "`" + `=:update_by,
		` + "`" + `create_time` + "`" + `=:create_time,
		` + "`" + `update_time` + "`" + `=:update_time,
		` + "`" + `db_code` + "`" + `=:db_code,
		` + "`" + `db_ch_name` + "`" + `=:db_ch_name,
		` + "`" + `db_type` + "`" + `=:db_type,
		` + "`" + `db_table_name` + "`" + `=:db_table_name,
		` + "`" + `db_dyn_sql` + "`" + `=:db_dyn_sql,
		` + "`" + `db_key` + "`" + `=:db_key,
		` + "`" + `tb_db_key` + "`" + `=:tb_db_key,
		` + "`" + `tb_db_table_name` + "`" + `=:tb_db_table_name,
		` + "`" + `java_type` + "`" + `=:java_type,
		` + "`" + `java_value` + "`" + `=:java_value,
		` + "`" + `api_url` + "`" + `=:api_url,
		` + "`" + `api_method` + "`" + `=:api_method,
		` + "`" + `is_list` + "`" + `=:is_list,
		` + "`" + `is_page` + "`" + `=:is_page,
		` + "`" + `db_source` + "`" + `=:db_source,
		` + "`" + `db_source_type` + "`" + `=:db_source_type,
		` + "`" + `json_data` + "`" + `=:json_data,
		` + "`" + `api_convert` + "`" + `=:api_convert
{{end}}

{{define "UpsertReportDbNoneZero"}}
INSERT INTO ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `jimu_report_db` + "`" + `
(
` + "`" + `id` + "`" + `,
` + "`" + `jimu_report_id` + "`" + `,
` + "`" + `create_by` + "`" + `,
` + "`" + `update_by` + "`" + `,
` + "`" + `create_time` + "`" + `,
` + "`" + `update_time` + "`" + `,
` + "`" + `db_code` + "`" + `,
` + "`" + `db_ch_name` + "`" + `,
` + "`" + `db_type` + "`" + `,
` + "`" + `db_table_name` + "`" + `,
` + "`" + `db_dyn_sql` + "`" + `,
` + "`" + `db_key` + "`" + `,
` + "`" + `tb_db_key` + "`" + `,
` + "`" + `tb_db_table_name` + "`" + `,
` + "`" + `java_type` + "`" + `,
` + "`" + `java_value` + "`" + `,
` + "`" + `api_url` + "`" + `,
` + "`" + `api_method` + "`" + `,
` + "`" + `is_list` + "`" + `,
` + "`" + `is_page` + "`" + `,
` + "`" + `db_source` + "`" + `,
` + "`" + `db_source_type` + "`" + `,
` + "`" + `json_data` + "`" + `,
` + "`" + `api_convert` + "`" + `)
VALUES (
        :id,
        :jimu_report_id,
        :create_by,
        :update_by,
        :create_time,
        :update_time,
        :db_code,
        :db_ch_name,
        :db_type,
        :db_table_name,
        :db_dyn_sql,
        :db_key,
        :tb_db_key,
        :tb_db_table_name,
        :java_type,
        :java_value,
        :api_url,
        :api_method,
        :is_list,
        :is_page,
        :db_source,
        :db_source_type,
        :json_data,
        :api_convert) ON DUPLICATE KEY
UPDATE
		{{Eval "NoneZeroSet" . | TrimSuffix ","}}
{{end}}

{{define "GetReportDb"}}
select *
from ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `jimu_report_db` + "`" + `
where ` + "`" + `id` + "`" + ` = ?
{{end}}

{{define "UpdateReportDbs"}}
UPDATE ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `jimu_report_db` + "`" + `
SET
	` + "`" + `jimu_report_id` + "`" + `=:jimu_report_id,
	` + "`" + `create_by` + "`" + `=:create_by,
	` + "`" + `update_by` + "`" + `=:update_by,
	` + "`" + `create_time` + "`" + `=:create_time,
	` + "`" + `update_time` + "`" + `=:update_time,
	` + "`" + `db_code` + "`" + `=:db_code,
	` + "`" + `db_ch_name` + "`" + `=:db_ch_name,
	` + "`" + `db_type` + "`" + `=:db_type,
	` + "`" + `db_table_name` + "`" + `=:db_table_name,
	` + "`" + `db_dyn_sql` + "`" + `=:db_dyn_sql,
	` + "`" + `db_key` + "`" + `=:db_key,
	` + "`" + `tb_db_key` + "`" + `=:tb_db_key,
	` + "`" + `tb_db_table_name` + "`" + `=:tb_db_table_name,
	` + "`" + `java_type` + "`" + `=:java_type,
	` + "`" + `java_value` + "`" + `=:java_value,
	` + "`" + `api_url` + "`" + `=:api_url,
	` + "`" + `api_method` + "`" + `=:api_method,
	` + "`" + `is_list` + "`" + `=:is_list,
	` + "`" + `is_page` + "`" + `=:is_page,
	` + "`" + `db_source` + "`" + `=:db_source,
	` + "`" + `db_source_type` + "`" + `=:db_source_type,
	` + "`" + `json_data` + "`" + `=:json_data,
	` + "`" + `api_convert` + "`" + `=:api_convert
{{end}}

{{define "UpdateReportDbsNoneZero"}}
UPDATE ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `jimu_report_db` + "`" + `
SET
    {{Eval "NoneZeroSet" . | TrimSuffix ","}}
{{end}}

{{define "InsertIgnoreReportDb"}}
INSERT IGNORE INTO ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `jimu_report_db` + "`" + `
(
` + "`" + `id` + "`" + `,
` + "`" + `jimu_report_id` + "`" + `,
` + "`" + `create_by` + "`" + `,
` + "`" + `update_by` + "`" + `,
` + "`" + `create_time` + "`" + `,
` + "`" + `update_time` + "`" + `,
` + "`" + `db_code` + "`" + `,
` + "`" + `db_ch_name` + "`" + `,
` + "`" + `db_type` + "`" + `,
` + "`" + `db_table_name` + "`" + `,
` + "`" + `db_dyn_sql` + "`" + `,
` + "`" + `db_key` + "`" + `,
` + "`" + `tb_db_key` + "`" + `,
` + "`" + `tb_db_table_name` + "`" + `,
` + "`" + `java_type` + "`" + `,
` + "`" + `java_value` + "`" + `,
` + "`" + `api_url` + "`" + `,
` + "`" + `api_method` + "`" + `,
` + "`" + `is_list` + "`" + `,
` + "`" + `is_page` + "`" + `,
` + "`" + `db_source` + "`" + `,
` + "`" + `db_source_type` + "`" + `,
` + "`" + `json_data` + "`" + `,
` + "`" + `api_convert` + "`" + `)
VALUES (
	   :id,
	   :jimu_report_id,
	   :create_by,
	   :update_by,
	   :create_time,
	   :update_time,
	   :db_code,
	   :db_ch_name,
	   :db_type,
	   :db_table_name,
	   :db_dyn_sql,
	   :db_key,
	   :tb_db_key,
	   :tb_db_table_name,
	   :java_type,
	   :java_value,
	   :api_url,
	   :api_method,
	   :is_list,
	   :is_page,
	   :db_source,
	   :db_source_type,
	   :json_data,
	   :api_convert)
{{end}}

{{define "UpdateClauseReportDb"}}
ON DUPLICATE KEY
UPDATE
		` + "`" + `jimu_report_id` + "`" + `=VALUES(jimu_report_id),
		` + "`" + `create_by` + "`" + `=VALUES(create_by),
		` + "`" + `update_by` + "`" + `=VALUES(update_by),
		` + "`" + `create_time` + "`" + `=VALUES(create_time),
		` + "`" + `update_time` + "`" + `=VALUES(update_time),
		` + "`" + `db_code` + "`" + `=VALUES(db_code),
		` + "`" + `db_ch_name` + "`" + `=VALUES(db_ch_name),
		` + "`" + `db_type` + "`" + `=VALUES(db_type),
		` + "`" + `db_table_name` + "`" + `=VALUES(db_table_name),
		` + "`" + `db_dyn_sql` + "`" + `=VALUES(db_dyn_sql),
		` + "`" + `db_key` + "`" + `=VALUES(db_key),
		` + "`" + `tb_db_key` + "`" + `=VALUES(tb_db_key),
		` + "`" + `tb_db_table_name` + "`" + `=VALUES(tb_db_table_name),
		` + "`" + `java_type` + "`" + `=VALUES(java_type),
		` + "`" + `java_value` + "`" + `=VALUES(java_value),
		` + "`" + `api_url` + "`" + `=VALUES(api_url),
		` + "`" + `api_method` + "`" + `=VALUES(api_method),
		` + "`" + `is_list` + "`" + `=VALUES(is_list),
		` + "`" + `is_page` + "`" + `=VALUES(is_page),
		` + "`" + `db_source` + "`" + `=VALUES(db_source),
		` + "`" + `db_source_type` + "`" + `=VALUES(db_source_type),
		` + "`" + `json_data` + "`" + `=VALUES(json_data),
		` + "`" + `api_convert` + "`" + `=VALUES(api_convert)
{{end}}

{{define "UpdateClauseSelectReportDb"}}
ON DUPLICATE KEY
UPDATE
		{{- range $i, $co := .Columns}}
		{{- if $i}},{{end}}
		` + "`" + `{{$co}}` + "`" + `=VALUES({{$co}})
		{{- end }}
{{end}}`