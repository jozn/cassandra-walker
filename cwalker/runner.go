package cwalker

import (
	"github.com/gocql/gocql"
)

var gen = &GenOut{}
var args *ConfigArgs

func Runner(arg *ConfigArgs) {
    args = arg
	for _, db := range arg.Keyspaces {
		// connect to the cluster
		//cluster := gocql.NewCluster("192.168.1.1", "192.168.1.2", "192.168.1.250")
		//cluster := gocql.NewCluster("192.168.1.250")
		cluster := gocql.NewCluster(arg.Host)
		cluster.Keyspace = db
		//cluster.Keyspace = "system"
		cluster.Consistency = gocql.One
		session, _ := cluster.CreateSession()
		defer session.Close()

		//loadTables("system_schema", gen, cluster)
		//loadTables("sunc", gen, cluster)
		tables := loadTables(db, cluster)
		//describeKeyspace("sunc", gen, cluster)
		//loadTables("system", gen, cluster)

        loadColumns(tables, cluster)

		for _, t := range tables {
			gen.TablesExtracted = append(gen.TablesExtracted, t)
		}

	}

    setTableParams(gen)

    PertyPrint(gen)

	//gen.Tables =gen.Tables[1:2]
	build(gen)

	for i := 0; i < 100; i++ {
		//insert2(session)
	}

}
