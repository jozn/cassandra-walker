package xc

import (
	"strings"
	"time"
    "github.com/gocql/gocql"
    "ms/sun/shared/helper"
)


{{range .Tables }}
	type {{ .TableNameGo }} struct{
		{{range .Columns }}
			{{- .ColumnNameGO }} {{ .TypeGo }} // {{ .ColumnName }}  {{ .Kind }}
		{{end}}
	}
/*
:= &xc.{{ .TableNameGo }} {
	{{- range .Columns }}
	{{ .ColumnNameGO }}: {{.TypeDefaultGo}},
	{{- end }}
*/
{{end}}

{{range .Tables }}
{{ $TableNameGo := .TableNameGo }}
func ({{.TableShortName}} *{{.TableNameGo}}) Save(session *gocql.Session) error {
	var cols []string
	var q []string
	var vals []interface{}

	{{range .Columns }}
		{{- if  eq .TypeGo "int" }}
			if {{.OutNameShorted}} != 0 {
				cols = append(cols, "{{.ColumnName}}")
				q = append(q, "?")
				vals = append(vals, {{.OutNameShorted}})
			}
		{{- else if eq .TypeGo "string"}}
			if {{.OutNameShorted}} != "" {
				cols = append(cols, "{{.ColumnName}}")
				q = append(q, "?")
				vals = append(vals, {{.OutNameShorted}})
			}
		{{- else}}
				cols = append(cols, "{{.ColumnName}}")
				q = append(q, "?")
				vals = append(vals, {{.OutNameShorted}})
		{{end}}
	{{end}}

	if len(cols) == 0 {
	    return errors.New("can not insert empty row.")
    }

	colOut := strings.Join(cols, ",")
	qOut := strings.Join(q, ",")
	cql := "insert into {{.TableSchemeOut}} (" + colOut + ") values (" + qOut + ") "

	err := session.Query(cql, vals... ).Exec()
	if err != nil {
		if LogTableCqlReq.{{ $TableNameGo }} {
			helper.XCLogErr(err)
		}
	}
	return err
}
{{end}}

// logs tables
type LogTableCql struct{
    {{range .Tables }}
    {{ .TableNameGo }} bool
    {{- end}}
}

var LogTableCqlReq = LogTableCql{
	{{- range .Tables }}
    {{ .TableNameGo }}: true ,
    {{- end}}
}
