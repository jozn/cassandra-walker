package {{ .PackageName}}

import (
    "strconv"
    "ms/sun/shared/base"
    "errors"
)
//TODO: WE MUST separate int from string to not let empty string "" from preloading or loading and inserting into caches
{{range .Tables}}
	{{- $short := .ShortName}}
	{{- $table := .TableSchemeOut}}
	{{- $model := .TableNameGo -}}

	{{range .Indexes}}
		{{if (and (eq (len .Columns) 1) (not .IsPrimary) ) }}

		{{- $col := (index .Columns 0) -}}//field
		{{- $colType := $col.GoTypeOut -}}//field
		{{- $indexName := (printf "%s_%s" $model  .IndexName) -}}//field
		{{$param := (printf "%s" $col.ColumnNameCamel) }}
		///// Generated from index '{{ .IndexName }}'.
		func (c _StoreImpl) {{ $model }}_By{{$col.ColumnName}} ({{$param}} {{$colType}}) (*{{ $model }},bool){
			o ,ok :=RowCacheIndex.Get("{{ $indexName }}:"+fmt.Sprintf("%v",{{$param}}))
			if ok {
				if obj, ok := o.(*{{ $model }});ok{
					return obj, true
				}
			}

			row, err := New{{ $model }}_Selector().{{$col.ColumnName}}_Eq({{$param}}).GetRow(base.DB)
			if err == nil{
		        RowCacheIndex.Set("{{ $indexName }}:"+fmt.Sprintf("%v",row.{{$param}}), row,0)
		        return row, true
		    }

			XOLogErr(err)
			return nil, false
		}

		func (c _StoreImpl) {{ $model }}_By{{$col.ColumnName}}_JustCache ({{$param}} {{$colType}}) (*{{ $model }},bool){
			o ,ok :=RowCacheIndex.Get("{{ $indexName }}:"+fmt.Sprintf("%v",{{$param}}))
			if ok {
				if obj, ok := o.(*{{ $model }});ok{
					return obj, true
				}
			}

			XOLogErr(errors.New("_JustCache is empty for secondry index " + "{{ $indexName }}:"+fmt.Sprintf("%v",{{$param}})))
			return nil, false
		}

		{{$param := (printf "%ss" $col.ColumnName) }}
		func (c _StoreImpl) PreLoad{{ $model }}_By{{$col.ColumnName}}s ({{$param}} []{{$colType}}) {
			not_cached := make([]{{$colType}},0,len({{$param}}))

			for _,id := range {{$param}} {
				_ ,ok :=RowCacheIndex.Get("{{ $indexName }}:"+fmt.Sprintf("%v",id))
				if !ok {
					not_cached = append(not_cached,id)
				}
			}

			if len(not_cached) > 0 {
				rows, err := New{{ $model }}_Selector().{{$col.ColumnName}}_In(not_cached).GetRows(base.DB)
				if err == nil{
		            for _, row := range rows {
		                RowCacheIndex.Set("{{ $indexName }}:"+fmt.Sprintf("%v",row.{{$col.ColumnName}}), row,0)
		            }
		        }
			}
		}
		{{else}}
		// {{$model}} - {{.IndexName}}

		{{end}}
	{{end}}
{{end}}