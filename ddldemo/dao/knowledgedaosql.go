/**
* Generated by go-doudou v2.0.4.
* You can edit it as your need.
 */
package dao

var knowledgedaosql = `{{define "NoneZeroSet"}}
	{{- if .DelFlag}}
	` + "`" + `del_flag` + "`" + `=:del_flag,
	{{- end}}
	{{- if .Title}}
	` + "`" + `title` + "`" + `=:title,
	{{- end}}
	{{- if .Content}}
	` + "`" + `content` + "`" + `=:content,
	{{- end}}
	{{- if .ViewCount}}
	` + "`" + `view_count` + "`" + `=:view_count,
	{{- end}}
	{{- if .Catalog}}
	` + "`" + `catalog` + "`" + `=:catalog,
	{{- end}}
	{{- if .Tag}}
	` + "`" + `tag` + "`" + `=:tag,
	{{- end}}
	{{- if .Status}}
	` + "`" + `status` + "`" + `=:status,
	{{- end}}
	{{- if .LastReleaseVersionId}}
	` + "`" + `last_release_version_id` + "`" + `=:last_release_version_id,
	{{- end}}
	{{- if .LastReleaseTime}}
	` + "`" + `last_release_time` + "`" + `=:last_release_time,
	{{- end}}
	{{- if .LastReleaseBy}}
	` + "`" + `last_release_by` + "`" + `=:last_release_by,
	{{- end}}
	{{- if .CreateTime}}
	` + "`" + `create_time` + "`" + `=:create_time,
	{{- end}}
	{{- if .CreateBy}}
	` + "`" + `create_by` + "`" + `=:create_by,
	{{- end}}
	{{- if .UpdateTime}}
	` + "`" + `update_time` + "`" + `=:update_time,
	{{- end}}
	{{- if .UpdateBy}}
	` + "`" + `update_by` + "`" + `=:update_by,
	{{- end}}
	{{- if .SysOrgCode}}
	` + "`" + `sys_org_code` + "`" + `=:sys_org_code,
	{{- end}}
	{{- if .ReferenceDocuments}}
	` + "`" + `reference_documents` + "`" + `=:reference_documents,
	{{- end}}
	{{- if .SortOrder}}
	` + "`" + `sort_order` + "`" + `=:sort_order,
	{{- end}}
	{{- if .ActualTitles}}
	` + "`" + `actual_titles` + "`" + `=:actual_titles,
	{{- end}}
{{end}}

{{define "InsertKnowledge"}}
INSERT INTO ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `cad_knowledge` + "`" + `
(
` + "`" + `id` + "`" + `,
` + "`" + `del_flag` + "`" + `,
` + "`" + `title` + "`" + `,
` + "`" + `content` + "`" + `,
` + "`" + `view_count` + "`" + `,
` + "`" + `catalog` + "`" + `,
` + "`" + `tag` + "`" + `,
` + "`" + `status` + "`" + `,
` + "`" + `last_release_version_id` + "`" + `,
` + "`" + `last_release_time` + "`" + `,
` + "`" + `last_release_by` + "`" + `,
` + "`" + `create_time` + "`" + `,
` + "`" + `create_by` + "`" + `,
` + "`" + `update_time` + "`" + `,
` + "`" + `update_by` + "`" + `,
` + "`" + `sys_org_code` + "`" + `,
` + "`" + `reference_documents` + "`" + `,
` + "`" + `sort_order` + "`" + `,
` + "`" + `actual_titles` + "`" + `)
VALUES (
	   :id,
	   :del_flag,
	   :title,
	   :content,
	   :view_count,
	   :catalog,
	   :tag,
	   :status,
	   :last_release_version_id,
	   :last_release_time,
	   :last_release_by,
	   :create_time,
	   :create_by,
	   :update_time,
	   :update_by,
	   :sys_org_code,
	   :reference_documents,
	   :sort_order,
	   :actual_titles)
{{end}}

{{define "UpdateKnowledge"}}
UPDATE ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `cad_knowledge` + "`" + `
SET
	` + "`" + `del_flag` + "`" + `=:del_flag,
	` + "`" + `title` + "`" + `=:title,
	` + "`" + `content` + "`" + `=:content,
	` + "`" + `view_count` + "`" + `=:view_count,
	` + "`" + `catalog` + "`" + `=:catalog,
	` + "`" + `tag` + "`" + `=:tag,
	` + "`" + `status` + "`" + `=:status,
	` + "`" + `last_release_version_id` + "`" + `=:last_release_version_id,
	` + "`" + `last_release_time` + "`" + `=:last_release_time,
	` + "`" + `last_release_by` + "`" + `=:last_release_by,
	` + "`" + `create_time` + "`" + `=:create_time,
	` + "`" + `create_by` + "`" + `=:create_by,
	` + "`" + `update_time` + "`" + `=:update_time,
	` + "`" + `update_by` + "`" + `=:update_by,
	` + "`" + `sys_org_code` + "`" + `=:sys_org_code,
	` + "`" + `reference_documents` + "`" + `=:reference_documents,
	` + "`" + `sort_order` + "`" + `=:sort_order,
	` + "`" + `actual_titles` + "`" + `=:actual_titles
WHERE
    ` + "`" + `id` + "`" + ` =:id
{{end}}

{{define "UpdateKnowledgeNoneZero"}}
UPDATE ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `cad_knowledge` + "`" + `
SET
    {{Eval "NoneZeroSet" . | TrimSuffix ","}}
WHERE
    ` + "`" + `id` + "`" + `=:id
{{end}}

{{define "UpsertKnowledge"}}
INSERT INTO ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `cad_knowledge` + "`" + `
(
` + "`" + `id` + "`" + `,
` + "`" + `del_flag` + "`" + `,
` + "`" + `title` + "`" + `,
` + "`" + `content` + "`" + `,
` + "`" + `view_count` + "`" + `,
` + "`" + `catalog` + "`" + `,
` + "`" + `tag` + "`" + `,
` + "`" + `status` + "`" + `,
` + "`" + `last_release_version_id` + "`" + `,
` + "`" + `last_release_time` + "`" + `,
` + "`" + `last_release_by` + "`" + `,
` + "`" + `create_time` + "`" + `,
` + "`" + `create_by` + "`" + `,
` + "`" + `update_time` + "`" + `,
` + "`" + `update_by` + "`" + `,
` + "`" + `sys_org_code` + "`" + `,
` + "`" + `reference_documents` + "`" + `,
` + "`" + `sort_order` + "`" + `,
` + "`" + `actual_titles` + "`" + `)
VALUES (
        :id,
        :del_flag,
        :title,
        :content,
        :view_count,
        :catalog,
        :tag,
        :status,
        :last_release_version_id,
        :last_release_time,
        :last_release_by,
        :create_time,
        :create_by,
        :update_time,
        :update_by,
        :sys_org_code,
        :reference_documents,
        :sort_order,
        :actual_titles) ON DUPLICATE KEY
UPDATE
		` + "`" + `del_flag` + "`" + `=:del_flag,
		` + "`" + `title` + "`" + `=:title,
		` + "`" + `content` + "`" + `=:content,
		` + "`" + `view_count` + "`" + `=:view_count,
		` + "`" + `catalog` + "`" + `=:catalog,
		` + "`" + `tag` + "`" + `=:tag,
		` + "`" + `status` + "`" + `=:status,
		` + "`" + `last_release_version_id` + "`" + `=:last_release_version_id,
		` + "`" + `last_release_time` + "`" + `=:last_release_time,
		` + "`" + `last_release_by` + "`" + `=:last_release_by,
		` + "`" + `create_time` + "`" + `=:create_time,
		` + "`" + `create_by` + "`" + `=:create_by,
		` + "`" + `update_time` + "`" + `=:update_time,
		` + "`" + `update_by` + "`" + `=:update_by,
		` + "`" + `sys_org_code` + "`" + `=:sys_org_code,
		` + "`" + `reference_documents` + "`" + `=:reference_documents,
		` + "`" + `sort_order` + "`" + `=:sort_order,
		` + "`" + `actual_titles` + "`" + `=:actual_titles
{{end}}

{{define "UpsertKnowledgeNoneZero"}}
INSERT INTO ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `cad_knowledge` + "`" + `
(
` + "`" + `id` + "`" + `,
` + "`" + `del_flag` + "`" + `,
` + "`" + `title` + "`" + `,
` + "`" + `content` + "`" + `,
` + "`" + `view_count` + "`" + `,
` + "`" + `catalog` + "`" + `,
` + "`" + `tag` + "`" + `,
` + "`" + `status` + "`" + `,
` + "`" + `last_release_version_id` + "`" + `,
` + "`" + `last_release_time` + "`" + `,
` + "`" + `last_release_by` + "`" + `,
` + "`" + `create_time` + "`" + `,
` + "`" + `create_by` + "`" + `,
` + "`" + `update_time` + "`" + `,
` + "`" + `update_by` + "`" + `,
` + "`" + `sys_org_code` + "`" + `,
` + "`" + `reference_documents` + "`" + `,
` + "`" + `sort_order` + "`" + `,
` + "`" + `actual_titles` + "`" + `)
VALUES (
        :id,
        :del_flag,
        :title,
        :content,
        :view_count,
        :catalog,
        :tag,
        :status,
        :last_release_version_id,
        :last_release_time,
        :last_release_by,
        :create_time,
        :create_by,
        :update_time,
        :update_by,
        :sys_org_code,
        :reference_documents,
        :sort_order,
        :actual_titles) ON DUPLICATE KEY
UPDATE
		{{Eval "NoneZeroSet" . | TrimSuffix ","}}
{{end}}

{{define "GetKnowledge"}}
select *
from ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `cad_knowledge` + "`" + `
where ` + "`" + `id` + "`" + ` = ?
{{end}}

{{define "UpdateKnowledges"}}
UPDATE ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `cad_knowledge` + "`" + `
SET
	` + "`" + `del_flag` + "`" + `=:del_flag,
	` + "`" + `title` + "`" + `=:title,
	` + "`" + `content` + "`" + `=:content,
	` + "`" + `view_count` + "`" + `=:view_count,
	` + "`" + `catalog` + "`" + `=:catalog,
	` + "`" + `tag` + "`" + `=:tag,
	` + "`" + `status` + "`" + `=:status,
	` + "`" + `last_release_version_id` + "`" + `=:last_release_version_id,
	` + "`" + `last_release_time` + "`" + `=:last_release_time,
	` + "`" + `last_release_by` + "`" + `=:last_release_by,
	` + "`" + `create_time` + "`" + `=:create_time,
	` + "`" + `create_by` + "`" + `=:create_by,
	` + "`" + `update_time` + "`" + `=:update_time,
	` + "`" + `update_by` + "`" + `=:update_by,
	` + "`" + `sys_org_code` + "`" + `=:sys_org_code,
	` + "`" + `reference_documents` + "`" + `=:reference_documents,
	` + "`" + `sort_order` + "`" + `=:sort_order,
	` + "`" + `actual_titles` + "`" + `=:actual_titles
{{end}}

{{define "UpdateKnowledgesNoneZero"}}
UPDATE ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `cad_knowledge` + "`" + `
SET
    {{Eval "NoneZeroSet" . | TrimSuffix ","}}
{{end}}

{{define "InsertIgnoreKnowledge"}}
INSERT IGNORE INTO ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `cad_knowledge` + "`" + `
(
` + "`" + `id` + "`" + `,
` + "`" + `del_flag` + "`" + `,
` + "`" + `title` + "`" + `,
` + "`" + `content` + "`" + `,
` + "`" + `view_count` + "`" + `,
` + "`" + `catalog` + "`" + `,
` + "`" + `tag` + "`" + `,
` + "`" + `status` + "`" + `,
` + "`" + `last_release_version_id` + "`" + `,
` + "`" + `last_release_time` + "`" + `,
` + "`" + `last_release_by` + "`" + `,
` + "`" + `create_time` + "`" + `,
` + "`" + `create_by` + "`" + `,
` + "`" + `update_time` + "`" + `,
` + "`" + `update_by` + "`" + `,
` + "`" + `sys_org_code` + "`" + `,
` + "`" + `reference_documents` + "`" + `,
` + "`" + `sort_order` + "`" + `,
` + "`" + `actual_titles` + "`" + `)
VALUES (
	   :id,
	   :del_flag,
	   :title,
	   :content,
	   :view_count,
	   :catalog,
	   :tag,
	   :status,
	   :last_release_version_id,
	   :last_release_time,
	   :last_release_by,
	   :create_time,
	   :create_by,
	   :update_time,
	   :update_by,
	   :sys_org_code,
	   :reference_documents,
	   :sort_order,
	   :actual_titles)
{{end}}

{{define "UpdateClauseKnowledge"}}
ON DUPLICATE KEY
UPDATE
		` + "`" + `del_flag` + "`" + `=VALUES(del_flag),
		` + "`" + `title` + "`" + `=VALUES(title),
		` + "`" + `content` + "`" + `=VALUES(content),
		` + "`" + `view_count` + "`" + `=VALUES(view_count),
		` + "`" + `catalog` + "`" + `=VALUES(catalog),
		` + "`" + `tag` + "`" + `=VALUES(tag),
		` + "`" + `status` + "`" + `=VALUES(status),
		` + "`" + `last_release_version_id` + "`" + `=VALUES(last_release_version_id),
		` + "`" + `last_release_time` + "`" + `=VALUES(last_release_time),
		` + "`" + `last_release_by` + "`" + `=VALUES(last_release_by),
		` + "`" + `create_time` + "`" + `=VALUES(create_time),
		` + "`" + `create_by` + "`" + `=VALUES(create_by),
		` + "`" + `update_time` + "`" + `=VALUES(update_time),
		` + "`" + `update_by` + "`" + `=VALUES(update_by),
		` + "`" + `sys_org_code` + "`" + `=VALUES(sys_org_code),
		` + "`" + `reference_documents` + "`" + `=VALUES(reference_documents),
		` + "`" + `sort_order` + "`" + `=VALUES(sort_order),
		` + "`" + `actual_titles` + "`" + `=VALUES(actual_titles)
{{end}}

{{define "UpdateClauseSelectKnowledge"}}
ON DUPLICATE KEY
UPDATE
		{{- range $i, $co := .Columns}}
		{{- if $i}},{{end}}
		` + "`" + `{{$co}}` + "`" + `=VALUES({{$co}})
		{{- end }}
{{end}}`
