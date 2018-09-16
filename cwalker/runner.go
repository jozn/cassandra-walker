package cwalker

import (
	"github.com/gocql/gocql"
)

var gen = &GenOut{}
var args *ConfigArgs

func Runner(arg *ConfigArgs) {
    args = arg
    gen.Package = args.Package

	for _, db := range arg.Keyspaces {
		// connect to the cluster
		cluster := gocql.NewCluster(arg.Host)
		cluster.Keyspace = db
		cluster.Consistency = gocql.One
		session, err := cluster.CreateSession()
		NoErr(err)
		defer session.Close()

		tables := loadTables(db, cluster)

        loadColumns(tables, cluster)

		for _, t := range tables {
			gen.TablesExtracted = append(gen.TablesExtracted, t)
		}
	}

    setTableParams(gen)

    PertyPrint(gen)

	build(gen)
}
