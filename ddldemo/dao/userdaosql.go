package dao

var userdaosql = `{{define "NoneZeroSet"}}
	{{- if .Name}}
	` + "`" + `name` + "`" + `=:name,
	{{- end}}
	{{- if .Phone}}
	` + "`" + `phone` + "`" + `=:phone,
	{{- end}}
	{{- if .Age}}
	` + "`" + `age` + "`" + `=:age,
	{{- end}}
	{{- if .No}}
	` + "`" + `no` + "`" + `=:no,
	{{- end}}
	{{- if .School}}
	` + "`" + `school` + "`" + `=:school,
	{{- end}}
	{{- if .IsStudent}}
	` + "`" + `is_student` + "`" + `=:is_student,
	{{- end}}
	{{- if .DeleteAt}}
	` + "`" + `delete_at` + "`" + `=:delete_at,
	{{- end}}
	{{- if .AvgScore}}
	` + "`" + `avg_score` + "`" + `=:avg_score,
	{{- end}}
	{{- if .Hobby}}
	` + "`" + `hobby` + "`" + `=:hobby,
	{{- end}}
{{end}}

{{define "InsertUser"}}
INSERT INTO ` + "`" + `test` + "`" + `.` + "`" + `ddl_user` + "`" + `
(
` + "`" + `id` + "`" + `,
` + "`" + `name` + "`" + `,
` + "`" + `phone` + "`" + `,
` + "`" + `age` + "`" + `,
` + "`" + `no` + "`" + `,
` + "`" + `school` + "`" + `,
` + "`" + `is_student` + "`" + `,
` + "`" + `delete_at` + "`" + `,
` + "`" + `avg_score` + "`" + `,
` + "`" + `hobby` + "`" + `)
VALUES (
	   :id,
	   :name,
	   :phone,
	   :age,
	   :no,
	   :school,
	   :is_student,
	   :delete_at,
	   :avg_score,
	   :hobby)
{{end}}

{{define "UpdateUser"}}
UPDATE ` + "`" + `test` + "`" + `.` + "`" + `ddl_user` + "`" + `
SET
	` + "`" + `name` + "`" + `=:name,
	` + "`" + `phone` + "`" + `=:phone,
	` + "`" + `age` + "`" + `=:age,
	` + "`" + `no` + "`" + `=:no,
	` + "`" + `school` + "`" + `=:school,
	` + "`" + `is_student` + "`" + `=:is_student,
	` + "`" + `delete_at` + "`" + `=:delete_at,
	` + "`" + `avg_score` + "`" + `=:avg_score,
	` + "`" + `hobby` + "`" + `=:hobby
WHERE
    ` + "`" + `id` + "`" + ` =:id
{{end}}

{{define "UpdateUserNoneZero"}}
UPDATE ` + "`" + `test` + "`" + `.` + "`" + `ddl_user` + "`" + `
SET
    {{Eval "NoneZeroSet" . | TrimSuffix ","}}
WHERE
    ` + "`" + `id` + "`" + `=:id
{{end}}

{{define "UpsertUser"}}
INSERT INTO ` + "`" + `test` + "`" + `.` + "`" + `ddl_user` + "`" + `
(
` + "`" + `id` + "`" + `,
` + "`" + `name` + "`" + `,
` + "`" + `phone` + "`" + `,
` + "`" + `age` + "`" + `,
` + "`" + `no` + "`" + `,
` + "`" + `school` + "`" + `,
` + "`" + `is_student` + "`" + `,
` + "`" + `delete_at` + "`" + `,
` + "`" + `avg_score` + "`" + `,
` + "`" + `hobby` + "`" + `)
VALUES (
        :id,
        :name,
        :phone,
        :age,
        :no,
        :school,
        :is_student,
        :delete_at,
        :avg_score,
        :hobby) ON DUPLICATE KEY
UPDATE
		` + "`" + `name` + "`" + `=:name,
		` + "`" + `phone` + "`" + `=:phone,
		` + "`" + `age` + "`" + `=:age,
		` + "`" + `no` + "`" + `=:no,
		` + "`" + `school` + "`" + `=:school,
		` + "`" + `is_student` + "`" + `=:is_student,
		` + "`" + `delete_at` + "`" + `=:delete_at,
		` + "`" + `avg_score` + "`" + `=:avg_score,
		` + "`" + `hobby` + "`" + `=:hobby
{{end}}

{{define "UpsertUserNoneZero"}}
INSERT INTO ` + "`" + `test` + "`" + `.` + "`" + `ddl_user` + "`" + `
(
` + "`" + `id` + "`" + `,
` + "`" + `name` + "`" + `,
` + "`" + `phone` + "`" + `,
` + "`" + `age` + "`" + `,
` + "`" + `no` + "`" + `,
` + "`" + `school` + "`" + `,
` + "`" + `is_student` + "`" + `,
` + "`" + `delete_at` + "`" + `,
` + "`" + `avg_score` + "`" + `,
` + "`" + `hobby` + "`" + `)
VALUES (
        :id,
        :name,
        :phone,
        :age,
        :no,
        :school,
        :is_student,
        :delete_at,
        :avg_score,
        :hobby) ON DUPLICATE KEY
UPDATE
		{{Eval "NoneZeroSet" . | TrimSuffix ","}}
{{end}}

{{define "GetUser"}}
select *
from ` + "`" + `test` + "`" + `.` + "`" + `ddl_user` + "`" + `
where ` + "`" + `id` + "`" + ` = ?
{{end}}

{{define "UpdateUsers"}}
UPDATE ` + "`" + `test` + "`" + `.` + "`" + `ddl_user` + "`" + `
SET
	` + "`" + `name` + "`" + `=:name,
	` + "`" + `phone` + "`" + `=:phone,
	` + "`" + `age` + "`" + `=:age,
	` + "`" + `no` + "`" + `=:no,
	` + "`" + `school` + "`" + `=:school,
	` + "`" + `is_student` + "`" + `=:is_student,
	` + "`" + `delete_at` + "`" + `=:delete_at,
	` + "`" + `avg_score` + "`" + `=:avg_score,
	` + "`" + `hobby` + "`" + `=:hobby
{{end}}

{{define "UpdateUsersNoneZero"}}
UPDATE ` + "`" + `test` + "`" + `.` + "`" + `ddl_user` + "`" + `
SET
    {{Eval "NoneZeroSet" . | TrimSuffix ","}}
{{end}}`
