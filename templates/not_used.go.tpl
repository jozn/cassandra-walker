
/*
//not for now -- but keep this
func MassInsert_{{.TableNameGo}}(rows []*{{.TableNameGo}}, session *gocql.Session) error {
    if len(rows) == 0 {
        return errors.New("rows slice should not be empty - inserted nothing")
    }
    var err error
    ln := len(rows)
    insVals := helper.SqlManyDollars( {{len .Columns }} ,len(rows),true)
    
    sqlstr := "INSERT INTO {{.TableSchemeOut}} (" +
       " {{ .ColumnNamesParams }} " +
        ") VALUES " + insVals

    // run query
    vals := make([]interface{}, 0, ln*5) //5 fields

    for _, row := range rows {
    	{{- range .Columns}}
    		vals = append(vals, row.{{.ColumnNameGO}})
    	{{- end}}
    }

    if LogTableCqlReq.{{.TableNameGo}} {
        helper.XCLog(" MassInsert len = ", ln, sqlstr ,vals)
    }
    err = session.Query(sqlstr, vals...).Exec()
    if err != nil {
        helper.XCLogErr(err)
        return err
    }

    return nil
}
*/
