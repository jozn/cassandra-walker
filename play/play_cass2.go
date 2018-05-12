package main

import (
	"log"

	"github.com/gocql/gocql"
	"ms/sun/shared/helper"
)

type Gen struct {
	Tables []*Table
}

type Table struct {
	TableName string
	Keyspace  string
	Columns   []*Column
}

type TableOut struct {
	Table
	Columns        []*ColumnOut
	ShortName      string
	TableNameGo    string
	TableSchemeOut string
	Comment        string
}

type Column struct {
	ColumnName string
	Order      int
	TypeCql    string
}

type ColumnOut struct {
	ColumnName string
	Order      int
	TypeCql    string
}



func main() {
	// connect to the cluster
	//cluster := gocql.NewCluster("192.168.1.1", "192.168.1.2", "192.168.1.250")
	cluster := gocql.NewCluster("192.168.1.250")
	cluster.Keyspace = "sunc"
	cluster.Consistency = gocql.One
	session, _ := cluster.CreateSession()
	defer session.Close()

	gen := &Gen{}

	//loadTables("system_schema", gen, cluster)
	loadTables("sunc", gen, cluster)
	loadColumns(gen, cluster)
	helper.PertyPrint(gen)

	for i := 0; i < 100; i++ {
		insert2(session)
	}

}

func insert(session *gocql.Session) {
	err := session.Query("insert into sunc.admins( first_name, last_name,userid) values (?,?,?) ",
		helper.RandAlphaNumbericString(5),
		helper.RandAlphaNumbericString(5),
		helper.RandAlphaNumbericString(5)).Exec()
	helper.NoErr(err)
}

func insert2(session *gocql.Session) {
	err := session.Query("insert into sunc.admins(userid) values (?) ",
		helper.RandAlphaNumbericString(5)).Exec()
	helper.NoErr(err)
}
