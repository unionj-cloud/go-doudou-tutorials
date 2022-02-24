package dao

var userdaosql = `{{define "NoneZeroSet"}}
	{{- if .Username}}
	` + "`" + `username` + "`" + `=:username,
	{{- end}}
	{{- if .Password}}
	` + "`" + `password` + "`" + `=:password,
	{{- end}}
	{{- if .Name}}
	` + "`" + `name` + "`" + `=:name,
	{{- end}}
	{{- if .Phone}}
	` + "`" + `phone` + "`" + `=:phone,
	{{- end}}
	{{- if .Dept}}
	` + "`" + `dept` + "`" + `=:dept,
	{{- end}}
	{{- if .Avatar}}
	` + "`" + `avatar` + "`" + `=:avatar,
	{{- end}}
	{{- if .DeleteAt}}
	` + "`" + `delete_at` + "`" + `=:delete_at,
	{{- end}}
{{end}}

{{define "InsertUser"}}
INSERT INTO ` + "`" + `tutorial` + "`" + `.` + "`" + `t_user` + "`" + `
(
` + "`" + `id` + "`" + `,
` + "`" + `username` + "`" + `,
` + "`" + `password` + "`" + `,
` + "`" + `name` + "`" + `,
` + "`" + `phone` + "`" + `,
` + "`" + `dept` + "`" + `,
` + "`" + `avatar` + "`" + `,
` + "`" + `delete_at` + "`" + `)
VALUES (
	   :id,
	   :username,
	   :password,
	   :name,
	   :phone,
	   :dept,
	   :avatar,
	   :delete_at)
{{end}}

{{define "UpdateUser"}}
UPDATE ` + "`" + `tutorial` + "`" + `.` + "`" + `t_user` + "`" + `
SET
	` + "`" + `username` + "`" + `=:username,
	` + "`" + `password` + "`" + `=:password,
	` + "`" + `name` + "`" + `=:name,
	` + "`" + `phone` + "`" + `=:phone,
	` + "`" + `dept` + "`" + `=:dept,
	` + "`" + `avatar` + "`" + `=:avatar,
	` + "`" + `delete_at` + "`" + `=:delete_at
WHERE
    ` + "`" + `id` + "`" + ` =:id
{{end}}

{{define "UpdateUserNoneZero"}}
UPDATE ` + "`" + `tutorial` + "`" + `.` + "`" + `t_user` + "`" + `
SET
    {{Eval "NoneZeroSet" . | TrimSuffix ","}}
WHERE
    ` + "`" + `id` + "`" + `=:id
{{end}}

{{define "UpsertUser"}}
INSERT INTO ` + "`" + `tutorial` + "`" + `.` + "`" + `t_user` + "`" + `
(
` + "`" + `id` + "`" + `,
` + "`" + `username` + "`" + `,
` + "`" + `password` + "`" + `,
` + "`" + `name` + "`" + `,
` + "`" + `phone` + "`" + `,
` + "`" + `dept` + "`" + `,
` + "`" + `avatar` + "`" + `,
` + "`" + `delete_at` + "`" + `)
VALUES (
        :id,
        :username,
        :password,
        :name,
        :phone,
        :dept,
        :avatar,
        :delete_at) ON DUPLICATE KEY
UPDATE
		` + "`" + `username` + "`" + `=:username,
		` + "`" + `password` + "`" + `=:password,
		` + "`" + `name` + "`" + `=:name,
		` + "`" + `phone` + "`" + `=:phone,
		` + "`" + `dept` + "`" + `=:dept,
		` + "`" + `avatar` + "`" + `=:avatar,
		` + "`" + `delete_at` + "`" + `=:delete_at
{{end}}

{{define "UpsertUserNoneZero"}}
INSERT INTO ` + "`" + `tutorial` + "`" + `.` + "`" + `t_user` + "`" + `
(
` + "`" + `id` + "`" + `,
` + "`" + `username` + "`" + `,
` + "`" + `password` + "`" + `,
` + "`" + `name` + "`" + `,
` + "`" + `phone` + "`" + `,
` + "`" + `dept` + "`" + `,
` + "`" + `avatar` + "`" + `,
` + "`" + `delete_at` + "`" + `)
VALUES (
        :id,
        :username,
        :password,
        :name,
        :phone,
        :dept,
        :avatar,
        :delete_at) ON DUPLICATE KEY
UPDATE
		{{Eval "NoneZeroSet" . | TrimSuffix ","}}
{{end}}

{{define "GetUser"}}
select *
from ` + "`" + `tutorial` + "`" + `.` + "`" + `t_user` + "`" + `
where ` + "`" + `id` + "`" + ` = ?
{{end}}

{{define "UpdateUsers"}}
UPDATE ` + "`" + `tutorial` + "`" + `.` + "`" + `t_user` + "`" + `
SET
	` + "`" + `username` + "`" + `=:username,
	` + "`" + `password` + "`" + `=:password,
	` + "`" + `name` + "`" + `=:name,
	` + "`" + `phone` + "`" + `=:phone,
	` + "`" + `dept` + "`" + `=:dept,
	` + "`" + `avatar` + "`" + `=:avatar,
	` + "`" + `delete_at` + "`" + `=:delete_at
{{end}}

{{define "UpdateUsersNoneZero"}}
UPDATE ` + "`" + `tutorial` + "`" + `.` + "`" + `t_user` + "`" + `
SET
    {{Eval "NoneZeroSet" . | TrimSuffix ","}}
{{end}}`
