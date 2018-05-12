package cwalker

import (
	"fmt"
	"github.com/jozn/protobuf/protoc-gen-go/generator"
)

type GenOut struct {
	TablesExtracted []*Table
	Tables          []*TableOut
}

type Table struct {
	TableName        string
	Keyspace         string
	Columns          []*Column
	PartitionColumns []*Column
	ClusterColumns   []*Column
}

type TableOut struct {
	Table
	Columns        []*ColumnOut
	TableShortName string
	TableNameGo    string
	TableSchemeOut string
	Comment        string
	OutColParams   string
}

type Column struct {
	ColumnName   string
	Kind         string
	Order        int
	TypeCql      string
	IsPartition  bool
	IsClustering bool
	IsRegular    bool //regular column types
}

type ColumnOut struct {
	Column
	ColumnNameGO   string
	OutNameShorted string
	TypeGo         string
	TypeDefaultGo  string
}

func setTableParams(gen *GenOut) {
	for _, table := range gen.TablesExtracted {
		t := &TableOut{
			Table:          *table,
			Comment:        fmt.Sprintf("table: %s", table.TableName),
			TableShortName: shortname(table.TableName),
			TableSchemeOut: table.Keyspace + "." + table.TableName,
			TableNameGo:    generator.CamelCase(table.TableName),
		}
		var outColParams = ""
		for _, col := range table.Columns {
			typGo, defGo := cqlTypesToGoType(col.TypeCql)
			c := &ColumnOut{
				Column:        *col,
				ColumnNameGO:  generator.CamelCase(col.ColumnName),
				TypeGo:        typGo,
				TypeDefaultGo: defGo,
			}
			c.OutNameShorted = fmt.Sprintf(" %s.%s", t.TableShortName, c.ColumnNameGO)
			t.Columns = append(t.Columns, c)
			outColParams += c.OutNameShorted + "," //fmt.Sprintf(" %s.%s,", t.TableShortName, c.ColumnNameGO)
		}

		t.OutColParams = outColParams[:len(outColParams)-1]
		gen.Tables = append(gen.Tables, t)
	}
}
