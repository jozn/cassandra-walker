package xc

import (
	"strings"
	"time"
    "github.com/gocql/gocql"
    "ms/sun/shared/helper"
)


type whereClause struct {
    condition string
    args      []interface{}
}

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

////////////////////////////////////////// Query seletor updater and deleter /////////////////////////

{{range .Tables }}


{{- $deleterType := printf "__%s_Deleter" .TableNameGo}}
{{- $updaterType := printf "__%s_Updater" .TableNameGo}}
{{- $selectorType := printf "__%s_Selector" .TableNameGo}}

type {{ $selectorType}} struct {
    wheres      []whereClause
    selectCol   []string
    orderBy     []string //" order by id desc //for ints
    limit       int
    allowFilter bool
}

type {{ $updaterType }} struct {
    wheres   []whereClause
    updates  map[string]interface{}
}

type {{ $deleterType }} struct {
    wheres      []whereClause
    deleteCol   []string

}

//////////////////// just Selector
func (u *{{ $selectorType}} ) Limit(limit int) {
    u.limit = limit
}

func (u *{{ $selectorType}}) allowFiltering() {
    u.allowFilter = true
}


func New{{.TableNameGo}}_Selector() *{{ $selectorType}} {
    u := {{ $selectorType}} {}
    return &u
}

//each select columns
{{ range .Columns }}
func (u *{{ $selectorType}}) Select_{{ .ColumnNameGO }}() {
    u.selectCol = append(u.selectCol, "{{.ColumnName}}" )
}
//each column orders //just ints
func (u *{{ $selectorType}})  {{ .ColumnNameGO }}_Desc() {
    u.orderBy = append(u.selectCol, " {{.ColumnName}} DESC")
}

func (u *{{ $selectorType}}) {{ .ColumnNameGO }}_Asc() {
    u.orderBy = append(u.orderBy, " {{.ColumnName}} ASC")
}

{{ end }}

//////////////////// just Deleter
//each column delete
{{ range .Columns }}
func (u *{{ $deleterType}}) Delete_{{ .ColumnNameGO }}() {
    u.deleteCol = append(u.deleteCol, "{{.ColumnName}}")
}
{{ end }}
//////////////////// End of just Deleter

//////////////////// just Updater
//each column delete
{{ range .Columns }}
	{{if (eq .TypeGo "int")}}
		func (u *{{ $updaterType}}) {{ .ColumnNameGO }}(newVal int) {
		    u.updates["{{.ColumnName}}"] = newVal
		}
	{{else if (eq .TypeGo "string")}}
		func (u *{{ $updaterType}}) {{ .ColumnNameGO }}(newVal string) {
		    u.updates["{{.ColumnName}}"] = newVal
		}
	{{else if (eq .TypeGo "[]byte")}}
		func (u *{{ $updaterType}}) {{ .ColumnNameGO }}(newVal []byte) {
		    u.updates["{{.ColumnName}}"] = newVal
		}		
	{{ end }}
{{ end }}
//////////////////// End just Updater
{{$table := . }}
{{ range (ms_to_slice $deleterType $updaterType $selectorType) }}
	{{ $operationType := . }}
	{{ range $table.Columns }}
		{{ $col := . }}
		{{ with .GetModifiers }}
			{{ range . }}
				//{{.}}
				{{ if (or (eq $col.TypeGo "int" ) (eq $col.TypeGo "string" ) ) }}
					func (d *{{ $operationType }}) {{ .FuncName }} (val {{$col.TypeGo}}) *{{$operationType}} {
					    w := whereClause{}
					    var insWhere []interface{}
					    insWhere = append(insWhere,val)
					    w.args = insWhere
					    w.condition = "{{.AndOr}} {{ $col.ColumnName }} {{.Condition}} ? "
					    d.wheres = append(d.wheres, w)

						return d
					}
				{{end}}
			{{end}}
		{{end }}
	{{ end }}
{{ end }}
///////////////////////////// start of where cluases

{{end }}//end of table range 
///////

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
