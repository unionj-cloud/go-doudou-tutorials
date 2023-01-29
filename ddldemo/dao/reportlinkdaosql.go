/**
* Generated by go-doudou v2.0.4.
* You can edit it as your need.
 */
package dao

var reportlinkdaosql = `{{define "NoneZeroSet"}}
	{{- if .ReportId}}
	` + "`" + `report_id` + "`" + `=:report_id,
	{{- end}}
	{{- if .Parameter}}
	` + "`" + `parameter` + "`" + `=:parameter,
	{{- end}}
	{{- if .EjectType}}
	` + "`" + `eject_type` + "`" + `=:eject_type,
	{{- end}}
	{{- if .LinkName}}
	` + "`" + `link_name` + "`" + `=:link_name,
	{{- end}}
	{{- if .ApiMethod}}
	` + "`" + `api_method` + "`" + `=:api_method,
	{{- end}}
	{{- if .LinkType}}
	` + "`" + `link_type` + "`" + `=:link_type,
	{{- end}}
	{{- if .ApiUrl}}
	` + "`" + `api_url` + "`" + `=:api_url,
	{{- end}}
	{{- if .LinkChartId}}
	` + "`" + `link_chart_id` + "`" + `=:link_chart_id,
	{{- end}}
{{end}}

{{define "InsertReportLink"}}
INSERT INTO ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `jimu_report_link` + "`" + `
(
` + "`" + `id` + "`" + `,
` + "`" + `report_id` + "`" + `,
` + "`" + `parameter` + "`" + `,
` + "`" + `eject_type` + "`" + `,
` + "`" + `link_name` + "`" + `,
` + "`" + `api_method` + "`" + `,
` + "`" + `link_type` + "`" + `,
` + "`" + `api_url` + "`" + `,
` + "`" + `link_chart_id` + "`" + `)
VALUES (
	   :id,
	   :report_id,
	   :parameter,
	   :eject_type,
	   :link_name,
	   :api_method,
	   :link_type,
	   :api_url,
	   :link_chart_id)
{{end}}

{{define "UpdateReportLink"}}
UPDATE ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `jimu_report_link` + "`" + `
SET
	` + "`" + `report_id` + "`" + `=:report_id,
	` + "`" + `parameter` + "`" + `=:parameter,
	` + "`" + `eject_type` + "`" + `=:eject_type,
	` + "`" + `link_name` + "`" + `=:link_name,
	` + "`" + `api_method` + "`" + `=:api_method,
	` + "`" + `link_type` + "`" + `=:link_type,
	` + "`" + `api_url` + "`" + `=:api_url,
	` + "`" + `link_chart_id` + "`" + `=:link_chart_id
WHERE
    ` + "`" + `id` + "`" + ` =:id
{{end}}

{{define "UpdateReportLinkNoneZero"}}
UPDATE ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `jimu_report_link` + "`" + `
SET
    {{Eval "NoneZeroSet" . | TrimSuffix ","}}
WHERE
    ` + "`" + `id` + "`" + `=:id
{{end}}

{{define "UpsertReportLink"}}
INSERT INTO ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `jimu_report_link` + "`" + `
(
` + "`" + `id` + "`" + `,
` + "`" + `report_id` + "`" + `,
` + "`" + `parameter` + "`" + `,
` + "`" + `eject_type` + "`" + `,
` + "`" + `link_name` + "`" + `,
` + "`" + `api_method` + "`" + `,
` + "`" + `link_type` + "`" + `,
` + "`" + `api_url` + "`" + `,
` + "`" + `link_chart_id` + "`" + `)
VALUES (
        :id,
        :report_id,
        :parameter,
        :eject_type,
        :link_name,
        :api_method,
        :link_type,
        :api_url,
        :link_chart_id) ON DUPLICATE KEY
UPDATE
		` + "`" + `report_id` + "`" + `=:report_id,
		` + "`" + `parameter` + "`" + `=:parameter,
		` + "`" + `eject_type` + "`" + `=:eject_type,
		` + "`" + `link_name` + "`" + `=:link_name,
		` + "`" + `api_method` + "`" + `=:api_method,
		` + "`" + `link_type` + "`" + `=:link_type,
		` + "`" + `api_url` + "`" + `=:api_url,
		` + "`" + `link_chart_id` + "`" + `=:link_chart_id
{{end}}

{{define "UpsertReportLinkNoneZero"}}
INSERT INTO ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `jimu_report_link` + "`" + `
(
` + "`" + `id` + "`" + `,
` + "`" + `report_id` + "`" + `,
` + "`" + `parameter` + "`" + `,
` + "`" + `eject_type` + "`" + `,
` + "`" + `link_name` + "`" + `,
` + "`" + `api_method` + "`" + `,
` + "`" + `link_type` + "`" + `,
` + "`" + `api_url` + "`" + `,
` + "`" + `link_chart_id` + "`" + `)
VALUES (
        :id,
        :report_id,
        :parameter,
        :eject_type,
        :link_name,
        :api_method,
        :link_type,
        :api_url,
        :link_chart_id) ON DUPLICATE KEY
UPDATE
		{{Eval "NoneZeroSet" . | TrimSuffix ","}}
{{end}}

{{define "GetReportLink"}}
select *
from ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `jimu_report_link` + "`" + `
where ` + "`" + `id` + "`" + ` = ?
{{end}}

{{define "UpdateReportLinks"}}
UPDATE ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `jimu_report_link` + "`" + `
SET
	` + "`" + `report_id` + "`" + `=:report_id,
	` + "`" + `parameter` + "`" + `=:parameter,
	` + "`" + `eject_type` + "`" + `=:eject_type,
	` + "`" + `link_name` + "`" + `=:link_name,
	` + "`" + `api_method` + "`" + `=:api_method,
	` + "`" + `link_type` + "`" + `=:link_type,
	` + "`" + `api_url` + "`" + `=:api_url,
	` + "`" + `link_chart_id` + "`" + `=:link_chart_id
{{end}}

{{define "UpdateReportLinksNoneZero"}}
UPDATE ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `jimu_report_link` + "`" + `
SET
    {{Eval "NoneZeroSet" . | TrimSuffix ","}}
{{end}}

{{define "InsertIgnoreReportLink"}}
INSERT IGNORE INTO ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `jimu_report_link` + "`" + `
(
` + "`" + `id` + "`" + `,
` + "`" + `report_id` + "`" + `,
` + "`" + `parameter` + "`" + `,
` + "`" + `eject_type` + "`" + `,
` + "`" + `link_name` + "`" + `,
` + "`" + `api_method` + "`" + `,
` + "`" + `link_type` + "`" + `,
` + "`" + `api_url` + "`" + `,
` + "`" + `link_chart_id` + "`" + `)
VALUES (
	   :id,
	   :report_id,
	   :parameter,
	   :eject_type,
	   :link_name,
	   :api_method,
	   :link_type,
	   :api_url,
	   :link_chart_id)
{{end}}

{{define "UpdateClauseReportLink"}}
ON DUPLICATE KEY
UPDATE
		` + "`" + `report_id` + "`" + `=VALUES(report_id),
		` + "`" + `parameter` + "`" + `=VALUES(parameter),
		` + "`" + `eject_type` + "`" + `=VALUES(eject_type),
		` + "`" + `link_name` + "`" + `=VALUES(link_name),
		` + "`" + `api_method` + "`" + `=VALUES(api_method),
		` + "`" + `link_type` + "`" + `=VALUES(link_type),
		` + "`" + `api_url` + "`" + `=VALUES(api_url),
		` + "`" + `link_chart_id` + "`" + `=VALUES(link_chart_id)
{{end}}

{{define "UpdateClauseSelectReportLink"}}
ON DUPLICATE KEY
UPDATE
		{{- range $i, $co := .Columns}}
		{{- if $i}},{{end}}
		` + "`" + `{{$co}}` + "`" + `=VALUES({{$co}})
		{{- end }}
{{end}}`
