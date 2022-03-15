package dao

var invalidtokendaosql = `{{define "NoneZeroSet"}}
	{{- if .Token}}
	` + "`" + `token` + "`" + `=:token,
	{{- end}}
	{{- if .ExpireAt}}
	` + "`" + `expire_at` + "`" + `=:expire_at,
	{{- end}}
	{{- if .DeleteAt}}
	` + "`" + `delete_at` + "`" + `=:delete_at,
	{{- end}}
{{end}}

{{define "InsertInvalidToken"}}
INSERT INTO ` + "`" + `tutorial` + "`" + `.` + "`" + `t_invalid_token` + "`" + `
(
` + "`" + `id` + "`" + `,
` + "`" + `token` + "`" + `,
` + "`" + `expire_at` + "`" + `,
` + "`" + `delete_at` + "`" + `)
VALUES (
	   :id,
	   :token,
	   :expire_at,
	   :delete_at)
{{end}}

{{define "UpdateInvalidToken"}}
UPDATE ` + "`" + `tutorial` + "`" + `.` + "`" + `t_invalid_token` + "`" + `
SET
	` + "`" + `token` + "`" + `=:token,
	` + "`" + `expire_at` + "`" + `=:expire_at,
	` + "`" + `delete_at` + "`" + `=:delete_at
WHERE
    ` + "`" + `id` + "`" + ` =:id
{{end}}

{{define "UpdateInvalidTokenNoneZero"}}
UPDATE ` + "`" + `tutorial` + "`" + `.` + "`" + `t_invalid_token` + "`" + `
SET
    {{Eval "NoneZeroSet" . | TrimSuffix ","}}
WHERE
    ` + "`" + `id` + "`" + `=:id
{{end}}

{{define "UpsertInvalidToken"}}
INSERT INTO ` + "`" + `tutorial` + "`" + `.` + "`" + `t_invalid_token` + "`" + `
(
` + "`" + `id` + "`" + `,
` + "`" + `token` + "`" + `,
` + "`" + `expire_at` + "`" + `,
` + "`" + `delete_at` + "`" + `)
VALUES (
        :id,
        :token,
        :expire_at,
        :delete_at) ON DUPLICATE KEY
UPDATE
		` + "`" + `token` + "`" + `=:token,
		` + "`" + `expire_at` + "`" + `=:expire_at,
		` + "`" + `delete_at` + "`" + `=:delete_at
{{end}}

{{define "UpsertInvalidTokenNoneZero"}}
INSERT INTO ` + "`" + `tutorial` + "`" + `.` + "`" + `t_invalid_token` + "`" + `
(
` + "`" + `id` + "`" + `,
` + "`" + `token` + "`" + `,
` + "`" + `expire_at` + "`" + `,
` + "`" + `delete_at` + "`" + `)
VALUES (
        :id,
        :token,
        :expire_at,
        :delete_at) ON DUPLICATE KEY
UPDATE
		{{Eval "NoneZeroSet" . | TrimSuffix ","}}
{{end}}

{{define "GetInvalidToken"}}
select *
from ` + "`" + `tutorial` + "`" + `.` + "`" + `t_invalid_token` + "`" + `
where ` + "`" + `id` + "`" + ` = ?
{{end}}

{{define "UpdateInvalidTokens"}}
UPDATE ` + "`" + `tutorial` + "`" + `.` + "`" + `t_invalid_token` + "`" + `
SET
	` + "`" + `token` + "`" + `=:token,
	` + "`" + `expire_at` + "`" + `=:expire_at,
	` + "`" + `delete_at` + "`" + `=:delete_at
{{end}}

{{define "UpdateInvalidTokensNoneZero"}}
UPDATE ` + "`" + `tutorial` + "`" + `.` + "`" + `t_invalid_token` + "`" + `
SET
    {{Eval "NoneZeroSet" . | TrimSuffix ","}}
{{end}}`
