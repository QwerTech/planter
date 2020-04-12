package main

const entryTmpl = `
table("{{ .Name }}", "{{ .Comment.String }}") {
{{- range .Columns }}
  {{- if .IsPrimaryKey }}
  <b>primary_key({{ .Name }})</b> {{ .DataType }}{{- if .Comment.Valid }} : {{ .Comment.String }}{{- end }}
  {{- end }}
{{- end }}
  --
{{- range .Columns }}
  {{- if not .IsPrimaryKey }}
  {{ if .HasIndex }}index_column {{ end}}{{ if .HasFk }}fk {{ end}}{{ if not .NotNull }}<i>{{- end }}<b>{{ .Name }}</b>{{ if not .NotNull }}</i>{{- end }} {{ .DataType }} {{- if .Comment.Valid }} : {{ .Comment.String }}{{- end }}
  {{- end }}
{{- end }}
}
`

const relationTmpl = `
{{ if .IsOneToOne }} {{ .SourceTableName }} ||-|| {{ .TargetTableName }}{{else}} {{ .SourceTableName }} }-- {{ .TargetTableName }}{{end}}
`
const headerTmpl = `
skinparam linetype ortho
!define table(x,y) class x << (T,#FFAAAA) y >>
!define primary_key(x) <&key> x
!define foreign_key(x) <&link-intact> x
!define index_column <&info>
!define fk <&action-undo>
`
const footerTmpl = `
legend
|= |= Legend |
| <b>primary_key(column)</b> | Primary key |
| fk <b>column</b> | Column has FK constraint |
| index_column <b>column</b> | Column has an index |
| <b><i>column</i></b> | Nullable column |
endlegend
`
