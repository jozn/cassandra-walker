package xc

import (
	"strings"
	"time"
    "github.com/gocql/gocql"
    "github.com/scylladb/gocqlx"
    "ms/sun/shared/helper"
)

//////////////// Constants //////////////////

type whereClause struct {
    condition string
    args      []interface{}
}

func whereClusesToSql(wheres []whereClause, whereSep string) (string, []interface{}) {
	var wheresArr []string
	for _, w := range wheres {
		wheresArr = append(wheresArr, w.condition)
	}
	wheresStr := strings.Join(wheresArr, whereSep)

	var args []interface{}
	for _, w := range wheres {
		args = append(args, w.args...)
	}
	return wheresStr, args
}

//////////////// End of Constants ///////////////

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
func (u *{{ $selectorType}} ) Limit(limit int) *{{ $selectorType}} {
    u.limit = limit
    return u
}

func (u *{{ $selectorType}}) allowFiltering()  *{{ $selectorType}} {
    u.allowFilter = true
    return u
}


func New{{.TableNameGo}}_Selector() *{{ $selectorType}} {
    u := {{ $selectorType}} {}
    return &u
}

//each select columns
{{ range .Columns }}
func (u *{{ $selectorType}}) Select_{{ .ColumnNameGO }}() *{{ $selectorType}} {
    u.selectCol = append(u.selectCol, "{{.ColumnName}}" )
    return u
}
//each column orders //just ints
func (u *{{ $selectorType}})  OrderBy_{{ .ColumnNameGO }}_Desc() *{{ $selectorType}} {
    u.orderBy = append(u.orderBy, " {{.ColumnName}} DESC")
    return u
}

func (u *{{ $selectorType}}) OrderBy_{{ .ColumnNameGO }}_Asc() *{{ $selectorType}} {
    u.orderBy = append(u.orderBy, " {{.ColumnName}} ASC")
    return u
}

{{ end }}

//////////////////// just Deleter
//each column delete
{{ range .Columns }}
func (u *{{ $deleterType}}) Delete_{{ .ColumnNameGO }}() *{{ $deleterType}} {
    u.deleteCol = append(u.deleteCol, "{{.ColumnName}}")
    return u
}
{{ end }}
//////////////////// End of just Deleter

//////////////////// just Updater
//each column delete
{{ range .Columns }}
	{{if (eq .TypeGo "int")}}
		func (u *{{ $updaterType}}) {{ .ColumnNameGO }}(newVal int)  *{{ $updaterType}} {
		    u.updates["{{.ColumnName}}"] = newVal
		     return u
		}
	{{else if (eq .TypeGo "string")}}
		func (u *{{ $updaterType}}) {{ .ColumnNameGO }}(newVal string) *{{ $updaterType}} {
		    u.updates["{{.ColumnName}}"] = newVal
		     return u
		}
	{{else if (eq .TypeGo "[]byte")}}
		func (u *{{ $updaterType}}) {{ .ColumnNameGO }}(newVal []byte) *{{ $updaterType}} {
		    u.updates["{{.ColumnName}}"] = newVal
		     return u
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

/////////////////////////////////////// Start of select //////////////////
func (u *{{ $selectorType }}) _toSql() (string,[]interface{}) {

	sqlWheres, whereArgs := whereClusesToSql(u.wheres, "")
	selectCols := "*"
	if len(u.selectCol) > 0 {
		selectCols = strings.Join(u.selectCol, ", ")
	}
	sqlstr := "SELECT " + selectCols + " FROM {{ $table.TableSchemeOut }}"

	if len(strings.Trim(sqlWheres, " ")) > 0 { //2 for safty
		sqlstr += " WHERE " + sqlWheres
	}

	if len(u.orderBy) > 0 {
		orders := strings.Join(u.orderBy, ", ")
		sqlstr += " ORDER BY " + orders
	}

	if u.limit != 0 {
		sqlstr += " LIMIT " + strconv.Itoa(u.limit)
	}
	
	return sqlstr, whereArgs
}

func (u *{{$selectorType}}) GetRow (session *gocql.Session) (*{{ $table.TableNameGo }},error) {
	var err error

	u.limit = 1
	sqlstr, whereArgs := u._toSql()

	if LogTableCqlReq.{{.TableNameGo}} {
		helper.XCLog(sqlstr,whereArgs )
	}

	query := session.Query(sqlstr, whereArgs...)
	var row *{{ $table.TableNameGo }}
	//by Sqlx
	// err = gocqlx.Get(row ,query)
	rows,err := {{.TableNameGo}}_Iter(query.Iter(),1)
	if err != nil {
		if LogTableCqlReq.{{.TableNameGo}} {
            helper.XCLogErr(err)
        }
		return nil, err
	}
	if len(rows) == 0 {
	    return nil,errors.New("empty rows")
    }else {
        row = rows[0]
    }

	//row._exists = true

	//On{{ .TableNameGo}}_LoadOne(row)

	return row, nil
}

func (u *{{$selectorType}}) GetRows (session *gocql.Session) ([]*{{ $table.TableNameGo }},error) {
	var err error

	sqlstr, whereArgs := u._toSql()

	if LogTableCqlReq.{{.TableNameGo}} {
		helper.XCLog(sqlstr,whereArgs )
	}

	query := session.Query(sqlstr, whereArgs...)

	rows,err := {{.TableNameGo}}_Iter(query.Iter(),-1)
	if err != nil {
		if LogTableCqlReq.{{.TableNameGo}} {
            helper.XCLogErr(err)
        }
		return rows, err
	}


	// for i:=0;i< len(rows);i++ {
	// 	rows[i]._exists = true
	// }

	// On{{ .TableNameGo}}_LoadMany(rows)

	return rows, nil
}

func (u *{{$updaterType}}) Update(session *gocql.Session) ( error) {
    var err error

    var updateArgs []interface{}
    var sqlUpdateArr []string
    for up, newVal := range u.updates {
        sqlUpdateArr = append(sqlUpdateArr, up)
        updateArgs = append(updateArgs, newVal)
    }
    sqlUpdate := strings.Join(sqlUpdateArr, ",")

    sqlWheres, whereArgs := whereClusesToSql(u.wheres, "")

    var allArgs []interface{}
    allArgs = append(allArgs, updateArgs...)
    allArgs = append(allArgs, whereArgs...)

    sqlstr := `UPDATE {{.TableSchemeOut}} SET ` + sqlUpdate

    if len(strings.Trim(sqlWheres, " ")) > 0 { 
        sqlstr += " WHERE " + sqlWheres
    }
    if LogTableCqlReq.{{.TableNameGo}} {
        helper.XCLog(sqlstr,allArgs)
    }
    err = session.Query(sqlstr, allArgs).Exec()
    if err != nil {
        helper.XCLogErr(err)
        return  err
    }

    return nil
}

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


//////////////////// Iternations //////////////////
{{- range .Tables }} 
func {{ .TableNameGo }}_Iter(iter *gocql.Iter, limit int) ([]*{{ .TableNameGo }}, error) {
	var rows []*{{ .TableNameGo }}
	if limit < 1 {
		limit = 1e6
	}

	for i := 0; i < limit; i++ {
		m := make(map[string]interface{}, 10)
		row := &{{ .TableNameGo }}{}
		if iter.MapScan(m) {

			{{range .Columns }}
				if val, ok := m["{{.ColumnName}}"]; ok {
					row.{{.ColumnNameGO}} = val.({{.TypeGo}})
				}
			{{ end }}

			rows = append(rows, row)
		} else {
			break
		}
	}
	err := iter.Close()

	return rows, err
}

{{- end }}