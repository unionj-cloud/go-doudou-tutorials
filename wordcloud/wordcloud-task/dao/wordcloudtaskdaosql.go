package dao

var wordcloudtaskdaosql = `{{define "NoneZeroSet"}}
	{{- if .SrcUrl}}
	` + "`" + `src_url` + "`" + `=:src_url,
	{{- end}}
	{{- if .ImgUrl}}
	` + "`" + `img_url` + "`" + `=:img_url,
	{{- end}}
	{{- if .Lang}}
	` + "`" + `lang` + "`" + `=:lang,
	{{- end}}
	{{- if .Top}}
	` + "`" + `top` + "`" + `=:top,
	{{- end}}
	{{- if .Status}}
	` + "`" + `status` + "`" + `=:status,
	{{- end}}
	{{- if .Error}}
	` + "`" + `error` + "`" + `=:error,
	{{- end}}
	{{- if .UserId}}
	` + "`" + `user_id` + "`" + `=:user_id,
	{{- end}}
	{{- if .DeleteAt}}
	` + "`" + `delete_at` + "`" + `=:delete_at,
	{{- end}}
{{end}}

{{define "InsertWordCloudTask"}}
INSERT INTO ` + "`" + `tutorial` + "`" + `.` + "`" + `t_word_cloud_task` + "`" + `
(
` + "`" + `id` + "`" + `,
` + "`" + `src_url` + "`" + `,
` + "`" + `img_url` + "`" + `,
` + "`" + `lang` + "`" + `,
` + "`" + `top` + "`" + `,
` + "`" + `status` + "`" + `,
` + "`" + `error` + "`" + `,
` + "`" + `user_id` + "`" + `,
` + "`" + `delete_at` + "`" + `)
VALUES (
	   :id,
	   :src_url,
	   :img_url,
	   :lang,
	   :top,
	   :status,
	   :error,
	   :user_id,
	   :delete_at)
{{end}}

{{define "UpdateWordCloudTask"}}
UPDATE ` + "`" + `tutorial` + "`" + `.` + "`" + `t_word_cloud_task` + "`" + `
SET
	` + "`" + `src_url` + "`" + `=:src_url,
	` + "`" + `img_url` + "`" + `=:img_url,
	` + "`" + `lang` + "`" + `=:lang,
	` + "`" + `top` + "`" + `=:top,
	` + "`" + `status` + "`" + `=:status,
	` + "`" + `error` + "`" + `=:error,
	` + "`" + `user_id` + "`" + `=:user_id,
	` + "`" + `delete_at` + "`" + `=:delete_at
WHERE
    ` + "`" + `id` + "`" + ` =:id
{{end}}

{{define "UpdateWordCloudTaskNoneZero"}}
UPDATE ` + "`" + `tutorial` + "`" + `.` + "`" + `t_word_cloud_task` + "`" + `
SET
    {{Eval "NoneZeroSet" . | TrimSuffix ","}}
WHERE
    ` + "`" + `id` + "`" + `=:id
{{end}}

{{define "UpsertWordCloudTask"}}
INSERT INTO ` + "`" + `tutorial` + "`" + `.` + "`" + `t_word_cloud_task` + "`" + `
(
` + "`" + `id` + "`" + `,
` + "`" + `src_url` + "`" + `,
` + "`" + `img_url` + "`" + `,
` + "`" + `lang` + "`" + `,
` + "`" + `top` + "`" + `,
` + "`" + `status` + "`" + `,
` + "`" + `error` + "`" + `,
` + "`" + `user_id` + "`" + `,
` + "`" + `delete_at` + "`" + `)
VALUES (
        :id,
        :src_url,
        :img_url,
        :lang,
        :top,
        :status,
        :error,
        :user_id,
        :delete_at) ON DUPLICATE KEY
UPDATE
		` + "`" + `src_url` + "`" + `=:src_url,
		` + "`" + `img_url` + "`" + `=:img_url,
		` + "`" + `lang` + "`" + `=:lang,
		` + "`" + `top` + "`" + `=:top,
		` + "`" + `status` + "`" + `=:status,
		` + "`" + `error` + "`" + `=:error,
		` + "`" + `user_id` + "`" + `=:user_id,
		` + "`" + `delete_at` + "`" + `=:delete_at
{{end}}

{{define "UpsertWordCloudTaskNoneZero"}}
INSERT INTO ` + "`" + `tutorial` + "`" + `.` + "`" + `t_word_cloud_task` + "`" + `
(
` + "`" + `id` + "`" + `,
` + "`" + `src_url` + "`" + `,
` + "`" + `img_url` + "`" + `,
` + "`" + `lang` + "`" + `,
` + "`" + `top` + "`" + `,
` + "`" + `status` + "`" + `,
` + "`" + `error` + "`" + `,
` + "`" + `user_id` + "`" + `,
` + "`" + `delete_at` + "`" + `)
VALUES (
        :id,
        :src_url,
        :img_url,
        :lang,
        :top,
        :status,
        :error,
        :user_id,
        :delete_at) ON DUPLICATE KEY
UPDATE
		{{Eval "NoneZeroSet" . | TrimSuffix ","}}
{{end}}

{{define "GetWordCloudTask"}}
select *
from ` + "`" + `tutorial` + "`" + `.` + "`" + `t_word_cloud_task` + "`" + `
where ` + "`" + `id` + "`" + ` = ?
{{end}}

{{define "UpdateWordCloudTasks"}}
UPDATE ` + "`" + `tutorial` + "`" + `.` + "`" + `t_word_cloud_task` + "`" + `
SET
	` + "`" + `src_url` + "`" + `=:src_url,
	` + "`" + `img_url` + "`" + `=:img_url,
	` + "`" + `lang` + "`" + `=:lang,
	` + "`" + `top` + "`" + `=:top,
	` + "`" + `status` + "`" + `=:status,
	` + "`" + `error` + "`" + `=:error,
	` + "`" + `user_id` + "`" + `=:user_id,
	` + "`" + `delete_at` + "`" + `=:delete_at
{{end}}

{{define "UpdateWordCloudTasksNoneZero"}}
UPDATE ` + "`" + `tutorial` + "`" + `.` + "`" + `t_word_cloud_task` + "`" + `
SET
    {{Eval "NoneZeroSet" . | TrimSuffix ","}}
{{end}}`
