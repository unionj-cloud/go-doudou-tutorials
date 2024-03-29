/**
* Generated by go-doudou v2.0.4.
* You can edit it as your need.
 */
package dao

var squaretaskauditdaosql = `{{define "NoneZeroSet"}}
	{{- if .TaskId}}
	` + "`" + `task_id` + "`" + `=:task_id,
	{{- end}}
	{{- if .AuditStatus}}
	` + "`" + `audit_status` + "`" + `=:audit_status,
	{{- end}}
	{{- if .Reason}}
	` + "`" + `reason` + "`" + `=:reason,
	{{- end}}
	{{- if .CreateTime}}
	` + "`" + `create_time` + "`" + `=:create_time,
	{{- end}}
	{{- if .CreateBy}}
	` + "`" + `create_by` + "`" + `=:create_by,
	{{- end}}
	{{- if .CommitTime}}
	` + "`" + `commit_time` + "`" + `=:commit_time,
	{{- end}}
	{{- if .UpdateTime}}
	` + "`" + `update_time` + "`" + `=:update_time,
	{{- end}}
	{{- if .UpdateBy}}
	` + "`" + `update_by` + "`" + `=:update_by,
	{{- end}}
{{end}}

{{define "InsertSquareTaskAudit"}}
INSERT INTO ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `cad_square_task_audit` + "`" + `
(
` + "`" + `id` + "`" + `,
` + "`" + `task_id` + "`" + `,
` + "`" + `audit_status` + "`" + `,
` + "`" + `reason` + "`" + `,
` + "`" + `create_time` + "`" + `,
` + "`" + `create_by` + "`" + `,
` + "`" + `commit_time` + "`" + `,
` + "`" + `update_time` + "`" + `,
` + "`" + `update_by` + "`" + `)
VALUES (
	   :id,
	   :task_id,
	   :audit_status,
	   :reason,
	   :create_time,
	   :create_by,
	   :commit_time,
	   :update_time,
	   :update_by)
{{end}}

{{define "UpdateSquareTaskAudit"}}
UPDATE ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `cad_square_task_audit` + "`" + `
SET
	` + "`" + `task_id` + "`" + `=:task_id,
	` + "`" + `audit_status` + "`" + `=:audit_status,
	` + "`" + `reason` + "`" + `=:reason,
	` + "`" + `create_time` + "`" + `=:create_time,
	` + "`" + `create_by` + "`" + `=:create_by,
	` + "`" + `commit_time` + "`" + `=:commit_time,
	` + "`" + `update_time` + "`" + `=:update_time,
	` + "`" + `update_by` + "`" + `=:update_by
WHERE
    ` + "`" + `id` + "`" + ` =:id
{{end}}

{{define "UpdateSquareTaskAuditNoneZero"}}
UPDATE ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `cad_square_task_audit` + "`" + `
SET
    {{Eval "NoneZeroSet" . | TrimSuffix ","}}
WHERE
    ` + "`" + `id` + "`" + `=:id
{{end}}

{{define "UpsertSquareTaskAudit"}}
INSERT INTO ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `cad_square_task_audit` + "`" + `
(
` + "`" + `id` + "`" + `,
` + "`" + `task_id` + "`" + `,
` + "`" + `audit_status` + "`" + `,
` + "`" + `reason` + "`" + `,
` + "`" + `create_time` + "`" + `,
` + "`" + `create_by` + "`" + `,
` + "`" + `commit_time` + "`" + `,
` + "`" + `update_time` + "`" + `,
` + "`" + `update_by` + "`" + `)
VALUES (
        :id,
        :task_id,
        :audit_status,
        :reason,
        :create_time,
        :create_by,
        :commit_time,
        :update_time,
        :update_by) ON DUPLICATE KEY
UPDATE
		` + "`" + `task_id` + "`" + `=:task_id,
		` + "`" + `audit_status` + "`" + `=:audit_status,
		` + "`" + `reason` + "`" + `=:reason,
		` + "`" + `create_time` + "`" + `=:create_time,
		` + "`" + `create_by` + "`" + `=:create_by,
		` + "`" + `commit_time` + "`" + `=:commit_time,
		` + "`" + `update_time` + "`" + `=:update_time,
		` + "`" + `update_by` + "`" + `=:update_by
{{end}}

{{define "UpsertSquareTaskAuditNoneZero"}}
INSERT INTO ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `cad_square_task_audit` + "`" + `
(
` + "`" + `id` + "`" + `,
` + "`" + `task_id` + "`" + `,
` + "`" + `audit_status` + "`" + `,
` + "`" + `reason` + "`" + `,
` + "`" + `create_time` + "`" + `,
` + "`" + `create_by` + "`" + `,
` + "`" + `commit_time` + "`" + `,
` + "`" + `update_time` + "`" + `,
` + "`" + `update_by` + "`" + `)
VALUES (
        :id,
        :task_id,
        :audit_status,
        :reason,
        :create_time,
        :create_by,
        :commit_time,
        :update_time,
        :update_by) ON DUPLICATE KEY
UPDATE
		{{Eval "NoneZeroSet" . | TrimSuffix ","}}
{{end}}

{{define "GetSquareTaskAudit"}}
select *
from ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `cad_square_task_audit` + "`" + `
where ` + "`" + `id` + "`" + ` = ?
{{end}}

{{define "UpdateSquareTaskAudits"}}
UPDATE ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `cad_square_task_audit` + "`" + `
SET
	` + "`" + `task_id` + "`" + `=:task_id,
	` + "`" + `audit_status` + "`" + `=:audit_status,
	` + "`" + `reason` + "`" + `=:reason,
	` + "`" + `create_time` + "`" + `=:create_time,
	` + "`" + `create_by` + "`" + `=:create_by,
	` + "`" + `commit_time` + "`" + `=:commit_time,
	` + "`" + `update_time` + "`" + `=:update_time,
	` + "`" + `update_by` + "`" + `=:update_by
{{end}}

{{define "UpdateSquareTaskAuditsNoneZero"}}
UPDATE ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `cad_square_task_audit` + "`" + `
SET
    {{Eval "NoneZeroSet" . | TrimSuffix ","}}
{{end}}

{{define "InsertIgnoreSquareTaskAudit"}}
INSERT IGNORE INTO ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `cad_square_task_audit` + "`" + `
(
` + "`" + `id` + "`" + `,
` + "`" + `task_id` + "`" + `,
` + "`" + `audit_status` + "`" + `,
` + "`" + `reason` + "`" + `,
` + "`" + `create_time` + "`" + `,
` + "`" + `create_by` + "`" + `,
` + "`" + `commit_time` + "`" + `,
` + "`" + `update_time` + "`" + `,
` + "`" + `update_by` + "`" + `)
VALUES (
	   :id,
	   :task_id,
	   :audit_status,
	   :reason,
	   :create_time,
	   :create_by,
	   :commit_time,
	   :update_time,
	   :update_by)
{{end}}

{{define "UpdateClauseSquareTaskAudit"}}
ON DUPLICATE KEY
UPDATE
		` + "`" + `task_id` + "`" + `=VALUES(task_id),
		` + "`" + `audit_status` + "`" + `=VALUES(audit_status),
		` + "`" + `reason` + "`" + `=VALUES(reason),
		` + "`" + `create_time` + "`" + `=VALUES(create_time),
		` + "`" + `create_by` + "`" + `=VALUES(create_by),
		` + "`" + `commit_time` + "`" + `=VALUES(commit_time),
		` + "`" + `update_time` + "`" + `=VALUES(update_time),
		` + "`" + `update_by` + "`" + `=VALUES(update_by)
{{end}}

{{define "UpdateClauseSelectSquareTaskAudit"}}
ON DUPLICATE KEY
UPDATE
		{{- range $i, $co := .Columns}}
		{{- if $i}},{{end}}
		` + "`" + `{{$co}}` + "`" + `=VALUES({{$co}})
		{{- end }}
{{end}}`
