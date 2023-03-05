/**
* Generated by go-doudou v2.0.4.
* You can edit it as your need.
 */
package dao

var squarecommitdaosql = `{{define "NoneZeroSet"}}
	{{- if .DelFlag}}
	` + "`" + `del_flag` + "`" + `=:del_flag,
	{{- end}}
	{{- if .TaskId}}
	` + "`" + `task_id` + "`" + `=:task_id,
	{{- end}}
	{{- if .Version}}
	` + "`" + `version` + "`" + `=:version,
	{{- end}}
	{{- if .Input}}
	` + "`" + `input` + "`" + `=:input,
	{{- end}}
	{{- if .Output}}
	` + "`" + `output` + "`" + `=:output,
	{{- end}}
	{{- if .IsBest}}
	` + "`" + `is_best` + "`" + `=:is_best,
	{{- end}}
	{{- if .IsPublished}}
	` + "`" + `is_published` + "`" + `=:is_published,
	{{- end}}
	{{- if .Remark}}
	` + "`" + `remark` + "`" + `=:remark,
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
	{{- if .AuditStatus}}
	` + "`" + `audit_status` + "`" + `=:audit_status,
	{{- end}}
	{{- if .IsRead}}
	` + "`" + `is_read` + "`" + `=:is_read,
	{{- end}}
	{{- if .FilePath}}
	` + "`" + `file_path` + "`" + `=:file_path,
	{{- end}}
{{end}}

{{define "InsertSquareCommit"}}
INSERT INTO ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `cad_square_commit` + "`" + `
(
` + "`" + `id` + "`" + `,
` + "`" + `del_flag` + "`" + `,
` + "`" + `task_id` + "`" + `,
` + "`" + `version` + "`" + `,
` + "`" + `input` + "`" + `,
` + "`" + `output` + "`" + `,
` + "`" + `is_best` + "`" + `,
` + "`" + `is_published` + "`" + `,
` + "`" + `remark` + "`" + `,
` + "`" + `create_time` + "`" + `,
` + "`" + `create_by` + "`" + `,
` + "`" + `update_time` + "`" + `,
` + "`" + `update_by` + "`" + `,
` + "`" + `sys_org_code` + "`" + `,
` + "`" + `audit_status` + "`" + `,
` + "`" + `is_read` + "`" + `,
` + "`" + `file_path` + "`" + `)
VALUES (
	   :id,
	   :del_flag,
	   :task_id,
	   :version,
	   :input,
	   :output,
	   :is_best,
	   :is_published,
	   :remark,
	   :create_time,
	   :create_by,
	   :update_time,
	   :update_by,
	   :sys_org_code,
	   :audit_status,
	   :is_read,
	   :file_path)
{{end}}

{{define "UpdateSquareCommit"}}
UPDATE ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `cad_square_commit` + "`" + `
SET
	` + "`" + `del_flag` + "`" + `=:del_flag,
	` + "`" + `task_id` + "`" + `=:task_id,
	` + "`" + `version` + "`" + `=:version,
	` + "`" + `input` + "`" + `=:input,
	` + "`" + `output` + "`" + `=:output,
	` + "`" + `is_best` + "`" + `=:is_best,
	` + "`" + `is_published` + "`" + `=:is_published,
	` + "`" + `remark` + "`" + `=:remark,
	` + "`" + `create_time` + "`" + `=:create_time,
	` + "`" + `create_by` + "`" + `=:create_by,
	` + "`" + `update_time` + "`" + `=:update_time,
	` + "`" + `update_by` + "`" + `=:update_by,
	` + "`" + `sys_org_code` + "`" + `=:sys_org_code,
	` + "`" + `audit_status` + "`" + `=:audit_status,
	` + "`" + `is_read` + "`" + `=:is_read,
	` + "`" + `file_path` + "`" + `=:file_path
WHERE
    ` + "`" + `id` + "`" + ` =:id
{{end}}

{{define "UpdateSquareCommitNoneZero"}}
UPDATE ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `cad_square_commit` + "`" + `
SET
    {{Eval "NoneZeroSet" . | TrimSuffix ","}}
WHERE
    ` + "`" + `id` + "`" + `=:id
{{end}}

{{define "UpsertSquareCommit"}}
INSERT INTO ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `cad_square_commit` + "`" + `
(
` + "`" + `id` + "`" + `,
` + "`" + `del_flag` + "`" + `,
` + "`" + `task_id` + "`" + `,
` + "`" + `version` + "`" + `,
` + "`" + `input` + "`" + `,
` + "`" + `output` + "`" + `,
` + "`" + `is_best` + "`" + `,
` + "`" + `is_published` + "`" + `,
` + "`" + `remark` + "`" + `,
` + "`" + `create_time` + "`" + `,
` + "`" + `create_by` + "`" + `,
` + "`" + `update_time` + "`" + `,
` + "`" + `update_by` + "`" + `,
` + "`" + `sys_org_code` + "`" + `,
` + "`" + `audit_status` + "`" + `,
` + "`" + `is_read` + "`" + `,
` + "`" + `file_path` + "`" + `)
VALUES (
        :id,
        :del_flag,
        :task_id,
        :version,
        :input,
        :output,
        :is_best,
        :is_published,
        :remark,
        :create_time,
        :create_by,
        :update_time,
        :update_by,
        :sys_org_code,
        :audit_status,
        :is_read,
        :file_path) ON DUPLICATE KEY
UPDATE
		` + "`" + `del_flag` + "`" + `=:del_flag,
		` + "`" + `task_id` + "`" + `=:task_id,
		` + "`" + `version` + "`" + `=:version,
		` + "`" + `input` + "`" + `=:input,
		` + "`" + `output` + "`" + `=:output,
		` + "`" + `is_best` + "`" + `=:is_best,
		` + "`" + `is_published` + "`" + `=:is_published,
		` + "`" + `remark` + "`" + `=:remark,
		` + "`" + `create_time` + "`" + `=:create_time,
		` + "`" + `create_by` + "`" + `=:create_by,
		` + "`" + `update_time` + "`" + `=:update_time,
		` + "`" + `update_by` + "`" + `=:update_by,
		` + "`" + `sys_org_code` + "`" + `=:sys_org_code,
		` + "`" + `audit_status` + "`" + `=:audit_status,
		` + "`" + `is_read` + "`" + `=:is_read,
		` + "`" + `file_path` + "`" + `=:file_path
{{end}}

{{define "UpsertSquareCommitNoneZero"}}
INSERT INTO ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `cad_square_commit` + "`" + `
(
` + "`" + `id` + "`" + `,
` + "`" + `del_flag` + "`" + `,
` + "`" + `task_id` + "`" + `,
` + "`" + `version` + "`" + `,
` + "`" + `input` + "`" + `,
` + "`" + `output` + "`" + `,
` + "`" + `is_best` + "`" + `,
` + "`" + `is_published` + "`" + `,
` + "`" + `remark` + "`" + `,
` + "`" + `create_time` + "`" + `,
` + "`" + `create_by` + "`" + `,
` + "`" + `update_time` + "`" + `,
` + "`" + `update_by` + "`" + `,
` + "`" + `sys_org_code` + "`" + `,
` + "`" + `audit_status` + "`" + `,
` + "`" + `is_read` + "`" + `,
` + "`" + `file_path` + "`" + `)
VALUES (
        :id,
        :del_flag,
        :task_id,
        :version,
        :input,
        :output,
        :is_best,
        :is_published,
        :remark,
        :create_time,
        :create_by,
        :update_time,
        :update_by,
        :sys_org_code,
        :audit_status,
        :is_read,
        :file_path) ON DUPLICATE KEY
UPDATE
		{{Eval "NoneZeroSet" . | TrimSuffix ","}}
{{end}}

{{define "GetSquareCommit"}}
select *
from ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `cad_square_commit` + "`" + `
where ` + "`" + `id` + "`" + ` = ?
{{end}}

{{define "UpdateSquareCommits"}}
UPDATE ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `cad_square_commit` + "`" + `
SET
	` + "`" + `del_flag` + "`" + `=:del_flag,
	` + "`" + `task_id` + "`" + `=:task_id,
	` + "`" + `version` + "`" + `=:version,
	` + "`" + `input` + "`" + `=:input,
	` + "`" + `output` + "`" + `=:output,
	` + "`" + `is_best` + "`" + `=:is_best,
	` + "`" + `is_published` + "`" + `=:is_published,
	` + "`" + `remark` + "`" + `=:remark,
	` + "`" + `create_time` + "`" + `=:create_time,
	` + "`" + `create_by` + "`" + `=:create_by,
	` + "`" + `update_time` + "`" + `=:update_time,
	` + "`" + `update_by` + "`" + `=:update_by,
	` + "`" + `sys_org_code` + "`" + `=:sys_org_code,
	` + "`" + `audit_status` + "`" + `=:audit_status,
	` + "`" + `is_read` + "`" + `=:is_read,
	` + "`" + `file_path` + "`" + `=:file_path
{{end}}

{{define "UpdateSquareCommitsNoneZero"}}
UPDATE ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `cad_square_commit` + "`" + `
SET
    {{Eval "NoneZeroSet" . | TrimSuffix ","}}
{{end}}

{{define "InsertIgnoreSquareCommit"}}
INSERT IGNORE INTO ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `cad_square_commit` + "`" + `
(
` + "`" + `id` + "`" + `,
` + "`" + `del_flag` + "`" + `,
` + "`" + `task_id` + "`" + `,
` + "`" + `version` + "`" + `,
` + "`" + `input` + "`" + `,
` + "`" + `output` + "`" + `,
` + "`" + `is_best` + "`" + `,
` + "`" + `is_published` + "`" + `,
` + "`" + `remark` + "`" + `,
` + "`" + `create_time` + "`" + `,
` + "`" + `create_by` + "`" + `,
` + "`" + `update_time` + "`" + `,
` + "`" + `update_by` + "`" + `,
` + "`" + `sys_org_code` + "`" + `,
` + "`" + `audit_status` + "`" + `,
` + "`" + `is_read` + "`" + `,
` + "`" + `file_path` + "`" + `)
VALUES (
	   :id,
	   :del_flag,
	   :task_id,
	   :version,
	   :input,
	   :output,
	   :is_best,
	   :is_published,
	   :remark,
	   :create_time,
	   :create_by,
	   :update_time,
	   :update_by,
	   :sys_org_code,
	   :audit_status,
	   :is_read,
	   :file_path)
{{end}}

{{define "UpdateClauseSquareCommit"}}
ON DUPLICATE KEY
UPDATE
		` + "`" + `del_flag` + "`" + `=VALUES(del_flag),
		` + "`" + `task_id` + "`" + `=VALUES(task_id),
		` + "`" + `version` + "`" + `=VALUES(version),
		` + "`" + `input` + "`" + `=VALUES(input),
		` + "`" + `output` + "`" + `=VALUES(output),
		` + "`" + `is_best` + "`" + `=VALUES(is_best),
		` + "`" + `is_published` + "`" + `=VALUES(is_published),
		` + "`" + `remark` + "`" + `=VALUES(remark),
		` + "`" + `create_time` + "`" + `=VALUES(create_time),
		` + "`" + `create_by` + "`" + `=VALUES(create_by),
		` + "`" + `update_time` + "`" + `=VALUES(update_time),
		` + "`" + `update_by` + "`" + `=VALUES(update_by),
		` + "`" + `sys_org_code` + "`" + `=VALUES(sys_org_code),
		` + "`" + `audit_status` + "`" + `=VALUES(audit_status),
		` + "`" + `is_read` + "`" + `=VALUES(is_read),
		` + "`" + `file_path` + "`" + `=VALUES(file_path)
{{end}}

{{define "UpdateClauseSelectSquareCommit"}}
ON DUPLICATE KEY
UPDATE
		{{- range $i, $co := .Columns}}
		{{- if $i}},{{end}}
		` + "`" + `{{$co}}` + "`" + `=VALUES({{$co}})
		{{- end }}
{{end}}`