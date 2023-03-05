/**
* Generated by go-doudou v2.0.4.
* You can edit it as your need.
 */
package dao

var squaretaskknowledgedaosql = `{{define "NoneZeroSet"}}
	{{- if .TaskId}}
	` + "`" + `task_id` + "`" + `=:task_id,
	{{- end}}
	{{- if .KnowledgeId}}
	` + "`" + `knowledge_id` + "`" + `=:knowledge_id,
	{{- end}}
	{{- if .CreateTime}}
	` + "`" + `create_time` + "`" + `=:create_time,
	{{- end}}
{{end}}

{{define "InsertSquareTaskKnowledge"}}
INSERT INTO ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `cad_square_task_knowledge` + "`" + `
(
` + "`" + `id` + "`" + `,
` + "`" + `task_id` + "`" + `,
` + "`" + `knowledge_id` + "`" + `,
` + "`" + `create_time` + "`" + `)
VALUES (
	   :id,
	   :task_id,
	   :knowledge_id,
	   :create_time)
{{end}}

{{define "UpdateSquareTaskKnowledge"}}
UPDATE ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `cad_square_task_knowledge` + "`" + `
SET
	` + "`" + `task_id` + "`" + `=:task_id,
	` + "`" + `knowledge_id` + "`" + `=:knowledge_id,
	` + "`" + `create_time` + "`" + `=:create_time
WHERE
    ` + "`" + `id` + "`" + ` =:id
{{end}}

{{define "UpdateSquareTaskKnowledgeNoneZero"}}
UPDATE ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `cad_square_task_knowledge` + "`" + `
SET
    {{Eval "NoneZeroSet" . | TrimSuffix ","}}
WHERE
    ` + "`" + `id` + "`" + `=:id
{{end}}

{{define "UpsertSquareTaskKnowledge"}}
INSERT INTO ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `cad_square_task_knowledge` + "`" + `
(
` + "`" + `id` + "`" + `,
` + "`" + `task_id` + "`" + `,
` + "`" + `knowledge_id` + "`" + `,
` + "`" + `create_time` + "`" + `)
VALUES (
        :id,
        :task_id,
        :knowledge_id,
        :create_time) ON DUPLICATE KEY
UPDATE
		` + "`" + `task_id` + "`" + `=:task_id,
		` + "`" + `knowledge_id` + "`" + `=:knowledge_id,
		` + "`" + `create_time` + "`" + `=:create_time
{{end}}

{{define "UpsertSquareTaskKnowledgeNoneZero"}}
INSERT INTO ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `cad_square_task_knowledge` + "`" + `
(
` + "`" + `id` + "`" + `,
` + "`" + `task_id` + "`" + `,
` + "`" + `knowledge_id` + "`" + `,
` + "`" + `create_time` + "`" + `)
VALUES (
        :id,
        :task_id,
        :knowledge_id,
        :create_time) ON DUPLICATE KEY
UPDATE
		{{Eval "NoneZeroSet" . | TrimSuffix ","}}
{{end}}

{{define "GetSquareTaskKnowledge"}}
select *
from ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `cad_square_task_knowledge` + "`" + `
where ` + "`" + `id` + "`" + ` = ?
{{end}}

{{define "UpdateSquareTaskKnowledges"}}
UPDATE ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `cad_square_task_knowledge` + "`" + `
SET
	` + "`" + `task_id` + "`" + `=:task_id,
	` + "`" + `knowledge_id` + "`" + `=:knowledge_id,
	` + "`" + `create_time` + "`" + `=:create_time
{{end}}

{{define "UpdateSquareTaskKnowledgesNoneZero"}}
UPDATE ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `cad_square_task_knowledge` + "`" + `
SET
    {{Eval "NoneZeroSet" . | TrimSuffix ","}}
{{end}}

{{define "InsertIgnoreSquareTaskKnowledge"}}
INSERT IGNORE INTO ` + "`" + `jeecg-boot` + "`" + `.` + "`" + `cad_square_task_knowledge` + "`" + `
(
` + "`" + `id` + "`" + `,
` + "`" + `task_id` + "`" + `,
` + "`" + `knowledge_id` + "`" + `,
` + "`" + `create_time` + "`" + `)
VALUES (
	   :id,
	   :task_id,
	   :knowledge_id,
	   :create_time)
{{end}}

{{define "UpdateClauseSquareTaskKnowledge"}}
ON DUPLICATE KEY
UPDATE
		` + "`" + `task_id` + "`" + `=VALUES(task_id),
		` + "`" + `knowledge_id` + "`" + `=VALUES(knowledge_id),
		` + "`" + `create_time` + "`" + `=VALUES(create_time)
{{end}}

{{define "UpdateClauseSelectSquareTaskKnowledge"}}
ON DUPLICATE KEY
UPDATE
		{{- range $i, $co := .Columns}}
		{{- if $i}},{{end}}
		` + "`" + `{{$co}}` + "`" + `=VALUES({{$co}})
		{{- end }}
{{end}}`