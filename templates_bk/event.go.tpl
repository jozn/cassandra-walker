package {{ .PackageName}}

import (
    "strconv"
    "ms/sun/shared/base"
)

{{range .Tables}}

{{- $short := .ShortName}}
{{- $table := .TableSchemeOut}}
{{- $typ := .TableNameGo }}
{{- $_ := "" }}

{{/* - * (Manually copy this to other location) */}}
//{{ .TableNameGo }} Events
{{if (eq .PrimaryKey.GoTypeOut "int") }}
func  On{{ .TableNameGo }}_AfterInsert{{$_}} (row *{{ .TableNameGo }}) {
	RowCache.Set("{{ .TableNameGo }}:"+strconv.Itoa(row.{{.PrimaryKey.ColumnNameCamel}}), row,time.Hour* 0)
}

func  On{{ .TableNameGo }}_AfterUpdate{{$_}} (row *{{ .TableNameGo }}) {
	RowCache.Set("{{ .TableNameGo }}:"+strconv.Itoa(row.{{.PrimaryKey.ColumnNameCamel}}), row,time.Hour* 0)
}

func  On{{ .TableNameGo }}_AfterDelete{{$_}} (row *{{ .TableNameGo }}) {
	RowCache.Delete("{{ .TableNameGo }}:"+strconv.Itoa(row.{{.PrimaryKey.ColumnNameCamel}}))
}

func  On{{ .TableNameGo }}_LoadOne{{$_}} (row *{{ .TableNameGo }}) {
	RowCache.Set("{{ .TableNameGo }}:"+strconv.Itoa(row.{{.PrimaryKey.ColumnNameCamel}}), row,time.Hour* 0)
}

func  On{{ .TableNameGo }}_LoadMany{{$_}} (rows []*{{ .TableNameGo }}) {
	for _, row:= range rows {
		RowCache.Set("{{ .TableNameGo }}:"+strconv.Itoa(row.{{.PrimaryKey.ColumnNameCamel}}), row,time.Hour* 0)
	}
}
{{else if ( eq .PrimaryKey.GoTypeOut "string" ) }}
func  On{{ .TableNameGo }}_AfterInsert{{$_}} (row *{{ .TableNameGo }}) {
	RowCache.Set("{{ .TableNameGo }}:"+row.{{.PrimaryKey.ColumnNameCamel}}, row,time.Hour* 0)
}

func  On{{ .TableNameGo }}_AfterUpdate{{$_}} (row *{{ .TableNameGo }}) {
	RowCache.Set("{{ .TableNameGo }}:"+row.{{.PrimaryKey.ColumnNameCamel}}, row,time.Hour* 0)
}

func  On{{ .TableNameGo }}_AfterDelete{{$_}} (row *{{ .TableNameGo }}) {
	RowCache.Delete("{{ .TableNameGo }}:"+row.{{.PrimaryKey.ColumnNameCamel}})
}

func  On{{ .TableNameGo }}_LoadOne{{$_}} (row *{{ .TableNameGo }}) {
	RowCache.Set("{{ .TableNameGo }}:"+row.{{.PrimaryKey.ColumnNameCamel}}, row,time.Hour* 0)
}

func  On{{ .TableNameGo }}_LoadMany{{$_}} (rows []*{{ .TableNameGo }}) {
	for _, row:= range rows {
		RowCache.Set("{{ .TableNameGo }}:"+row.{{.PrimaryKey.ColumnNameCamel}}, row,time.Hour* 0)
	}
}
{{end}}

{{end}}