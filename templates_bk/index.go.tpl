package {{ .PackageName}}

import (
    "strconv"
    "ms/sun/shared/base"
    "github.com/jmoiron/sqlx"
)
//todo this code can be used for multi secondery coulmn -- but this one interfier with secondery_template - in futer merege this two temple to one unifed that uses both primiry keys and othe seconder options and with reloadings
{{range .Tables}}

{{- $short := .ShortName}}
{{- $table := .}}
{{- $typ := .TableNameGo }}
{{- $_ := "" }}

{{range .Indexes}}
	{{ if .IsPrimary}}
		// {{ .FuncNameOut }} Generated from index '{{ .IndexName }}' -- retrieves a row from '{{ $table.TableSchemeOut }}' as a {{ $table.TableNameGo  }}.
		func {{ .FuncNameOut }}(db *sqlx.DB {{ goparamlist .Columns true true }}) ({{ if not .IsUnique }}[]{{ end }}*{{ $table.TableNameGo }}, error) {
			var err error

			const sqlstr = `SELECT * ` +
				`FROM {{ $table.TableSchemeOut }} ` +
				`WHERE {{ colnamesquery .Columns " AND " }}`

			XOLog(sqlstr{{ goparamlist .Columns true false }})

		{{- if .IsUnique }}
			{{ $short }} := {{ $table.TableNameGo }}{
			{{- if  $table.PrimaryKey }}
				_exists: true,
			{{ end -}}
			}

			err = db.Get(&{{ $short }},sqlstr{{ goparamlist .Columns true false }})
			if err != nil {
				XOLogErr(err)
				return nil, err
			}

			On{{ $table.TableNameGo }}_LoadOne(&{{ $short }})

			return &{{ $short }}, nil
		{{- else }}
			res := []*{{ $table.TableNameGo }}{}
			err = db.Select(&res,sqlstr{{ goparamlist .Columns true false }})
			if err != nil {
				XOLogErr(err)
				return res, err
			}
			On{{ $table.TableNameGo }}_LoadMany(res)

			return res, nil
		{{- end }}
		}
	{{end}}
{{end}}
{{end}}