package xconst

const (
{{range .Tables}}
	{{.TableNameGo}}_Table = "{{.TableName}}"
	{{.TableNameGo}}_TableGo = "{{.TableNameGo}}"
{{- end}}
)
{{range .Tables}}
	var {{.TableNameGo}} = 	struct {
		{{range .Columns}}
			{{.ColumnNameCamel}} string
		{{- end}}
	}{
		{{range .Columns}}
        	{{.ColumnNameCamel}}: "{{.ColumnName}}",
        {{- end}}
	}
{{end}}

