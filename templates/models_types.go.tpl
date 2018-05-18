package xc

import (
    "errors"
	"ms/sun/shared/helper"
	"strconv"
	"strings"

	"github.com/gocql/gocql"
)

{{range .Tables }}
	type {{ .TableNameGo }} struct{
		{{range .Columns }}
			{{- .ColumnNameGO }} {{ .TypeGo }} // {{ .ColumnName }}  {{ .Kind }}
		{{end}}

		_exists, _deleted bool
	}
/*
:= &xc.{{ .TableNameGo }} {
	{{- range .Columns }}
	{{ .ColumnNameGO }}: {{.TypeDefaultGo}},
	{{- end }}
*/

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