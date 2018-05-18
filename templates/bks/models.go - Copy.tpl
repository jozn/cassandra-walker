package x

{{range . }}
	{{- if .Comment -}}
		// {{ .Comment }}
	{{- else -}}
		// {{ .TableName }} '{{ .TableNameGo }}'.
	{{- end }}
	type {{ .TableNameGo }} struct {
		{{- range .Columns }}
			{{ .ColumnNameCamel }} {{ .GoTypeOut }} {{ .StructTagOut }} {{ ms_col_comment_raw .Comment }}
		{{- end }}
		{{- if .PrimaryKey }}
			{{/* // xox fields */}}
			_exists, _deleted bool
		{{ end -}}
	}
	/*
:= &x.{{ .TableNameGo }} {
	{{- range .Columns }}
	{{ .ColumnNameCamel }}: {{.GoDefaultOut}},
	{{- end }}
	*/
{{end}}

///////////////// Skip Loging Tables ////////////////
type LogTableSql struct{
    {{range . }}
    {{ .TableNameGo }} bool
    {{- end}}
}

var LogTableSqlReq = LogTableSql{
	{{range . }}
    {{ .TableNameGo }}: true ,
    {{- end}}
}
