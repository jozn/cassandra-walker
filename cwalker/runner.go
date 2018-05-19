package cwalker

import (
    "github.com/gocql/gocql"
    "ms/sun/shared/helper"
)

var gen = &GenOut{}
func Runner() {

    for _, db := range DATABASES {
        // connect to the cluster
        //cluster := gocql.NewCluster("192.168.1.1", "192.168.1.2", "192.168.1.250")
        //cluster := gocql.NewCluster("192.168.1.250")
        cluster := gocql.NewCluster("127.0.0.1")
        cluster.Keyspace = db
        //cluster.Keyspace = "system"
        cluster.Consistency = gocql.One
        session, _ := cluster.CreateSession()
        defer session.Close()



        //loadTables("system_schema", gen, cluster)
        //loadTables("sunc", gen, cluster)
        loadTables(db, gen, cluster)
        //describeKeyspace("sunc", gen, cluster)
        //loadTables("system", gen, cluster)
        loadColumns(gen, cluster)
        setTableParams(gen)
    }

    helper.PertyPrint(gen)

    //gen.Tables =gen.Tables[1:2]
    build(gen)

    for i := 0; i < 100; i++ {
        //insert2(session)
    }

}

