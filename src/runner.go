package src

import (
	_ "github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
	"ms/sun/shared/helper"
	"regexp"
)


func Run() {
	DB, err := sqlx.Connect("mysql", "root:123456@tcp(localhost:3306)/sun?charset=utf8mb4")
	DB.MapperFunc(func(s string) string { return s })
	DB = DB.Unsafe()
	helper.NoErr(err)

	//OutPutBuffer := &GenOut{}
	for _, db := range DATABASES {
		tables, err := My_LoadTables(DB, db, "BASE TABLE")
		helper.NoErr(err)
		OutPutBuffer.Tables = append(OutPutBuffer.Tables, tables...)
	}

	for _, table := range OutPutBuffer.Tables {
		table.Columns, _ = My_LoadTableColumns(DB, table.DataBase, table.TableName, table)
		table.Indexes, _ = MyTableIndexes(DB, table.DataBase, table.TableName, table)
	}

	for _, table := range OutPutBuffer.Tables {
		if table.NeedTrigger {
			OutPutBuffer.TablesTriggers = append(OutPutBuffer.TablesTriggers, table)
		}
		if table.PrimaryKey != nil {
		    table.XPrimaryKeyGoType = table.PrimaryKey.GoTypeOut
        }
	}

	build(OutPutBuffer)
	helper.PertyPrint(OutPutBuffer.Tables)

}
