package cwalker

import (
	"fmt"
	"strings"

	"github.com/jozn/protobuf/protoc-gen-go/generator"
)

type GenOut struct {
	TablesExtracted   []*Table
	Tables            []*TableOut
	KeyspacesDescribe []string
	Package           string
}

type Table struct {
	TableName string
	Keyspace  string
	Columns   []*Column
	//PartitionColumns []*Column
	//ClusterColumns   []*Column
}

type TableOut struct {
	Table
	Columns          []*ColumnOut
	PartitionColumns []*ColumnOut
	ClusterColumns   []*ColumnOut
	TableShortName   string
	TableNameGo      string
	TableSchemeOut   string
	Comment          string
	OutColParams     string
	PrefixHidden     string  //hide ex: Table_Selector in docs
	GenOut           *GenOut //we need this for package refrencing, could be done better, but good
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
	TypeGoOriginal string
	TypeDefaultGo  string
	WhereModifiers []WhereModifier
}

type WhereModifier struct {
	Suffix    string
	Prefix    string
	Condition string
	AndOr     string
	FuncName  string
}

type WhereModifierIns struct {
	Suffix   string
	Prefix   string
	AndOr    string
	FuncName string
}

func setTableParams(gen *GenOut) {
	for _, table := range gen.TablesExtracted {
		t := &TableOut{
			Table:          *table,
			Comment:        fmt.Sprintf("table: %s", table.TableName),
			TableShortName: shortname(table.TableName),
			TableSchemeOut: table.Keyspace + "." + table.TableName,
			TableNameGo:    generator.CamelCase(table.TableName),
			PrefixHidden:   "",
			GenOut:         gen,
		}
		if args.Minimize {
			t.PrefixHidden = "__"
		}
		var outColParams = ""
		for _, col := range table.Columns {
			typGo, typOrg, defGo := cqlTypesToGoType(col.TypeCql)
			c := &ColumnOut{
				Column:         *col,
				ColumnNameGO:   generator.CamelCase(col.ColumnName),
				TypeGo:         typGo,
				TypeGoOriginal: typOrg,
				TypeDefaultGo:  defGo,
			}
			c.OutNameShorted = fmt.Sprintf(" %s.%s", t.TableShortName, c.ColumnNameGO)
			t.Columns = append(t.Columns, c)
			if c.IsPartition {
				t.PartitionColumns = append(t.PartitionColumns, c)
			}

			if col.IsClustering {
				t.ClusterColumns = append(t.ClusterColumns, c)
			}

			outColParams += c.OutNameShorted + "," //fmt.Sprintf(" %s.%s,", t.TableShortName, c.ColumnNameGO)
			c.WhereModifiers = c.GetModifiers()
		}

		t.OutColParams = outColParams[:len(outColParams)-1]
		gen.Tables = append(gen.Tables, t)
	}
}

func (c *ColumnOut) GetModifiers() (res []WhereModifier) {
	add := func(m WhereModifier) {
		if len(m.AndOr) > 0 {
			m.FuncName = m.AndOr + "_" + c.ColumnNameGO + m.Suffix
		} else {
			m.FuncName = c.ColumnNameGO + m.Suffix
		}
		res = append(res, m)
	}
	eqAdd := func(filter, andOr string) {
		//sufix := filter + andOr
		add(WhereModifier{"_Eq" + filter, andOr, "=", andOr, ""})
	}

	notEqs := func(filter, andOr string) {
		sufix := filter //+ andOr
		and := andOr
		add(WhereModifier{"_LT" + sufix, and, "<", andOr, ""})
		add(WhereModifier{"_LE" + sufix, and, "<=", andOr, ""})
		add(WhereModifier{"_GT" + sufix, and, ">", andOr, ""})
		add(WhereModifier{"_GE" + sufix, and, ">=", andOr, ""})
	}
	const filter = "_FILTERING"
	for _, andOr := range []string{"", "And", "Or"} {
		if c.TypeGo == "int" || c.TypeGo == "int64" {
			filter := "_Filtering"
			if c.IsPartition {
				eqAdd("", andOr)
				notEqs(filter, andOr)
			}
			if c.IsClustering {
				eqAdd("", andOr)
				notEqs("", andOr)
			}
			if c.IsRegular {
				eqAdd(filter, andOr)
				notEqs(filter, andOr)
			}
		}
		if c.TypeGo == "string" {
			if c.IsPartition {
				eqAdd("", andOr)
			}
			if c.IsClustering {
				eqAdd("", andOr)
			}
			if c.IsRegular {
				eqAdd(filter, andOr)
			}
		}
	}

	return
}

func (c *ColumnOut) GetModifiersIns() (res []WhereModifierIns) {
	add := func(m WhereModifierIns) {
		if len(m.AndOr) > 0 {
			m.FuncName = m.AndOr + "_" + c.ColumnNameGO + m.Suffix
		} else {
			m.FuncName = c.ColumnNameGO + m.Suffix
		}
		res = append(res, m)
	}
	inAdd := func(filter, andOr string) {
		add(WhereModifierIns{"_In" + filter, andOr, andOr, ""})
	}

	const filter = "_FILTERING"

	for _, andOr := range []string{"", "And", "Or"} {
		if c.TypeGo == "int" {
			if c.IsPartition {
				inAdd("", andOr)
			}
			if c.IsClustering {
				inAdd("", andOr)
			}
			if c.IsRegular {
				inAdd(filter, andOr)
			}
		}
		if c.TypeGo == "string" {
			if c.IsPartition {
				inAdd("", andOr)
			}
			if c.IsClustering {
				inAdd("", andOr)
			}
			if c.IsRegular {
				inAdd(filter, andOr)
			}
		}
	}

	return
}

func (table *TableOut) ColumnNamesParams() string {
	var arr []string
	for _, t := range table.Columns {
		arr = append(arr, t.ColumnName)
	}
	return strings.Join(arr, ",")
}
