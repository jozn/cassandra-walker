package cwalker

import (
    "github.com/gocql/gocql"
    "ms/sun/shared/helper"
)

var gen = &GenOut{}
func Runner() {
    // connect to the cluster
    //cluster := gocql.NewCluster("192.168.1.1", "192.168.1.2", "192.168.1.250")
    cluster := gocql.NewCluster("192.168.1.250")
    cluster.Keyspace = "sunc"
    cluster.Consistency = gocql.One
    session, _ := cluster.CreateSession()
    defer session.Close()



    //loadTables("system_schema", gen, cluster)
    loadTables("sunc", gen, cluster)
    loadColumns(gen, cluster)
    setTableParams(gen)
    helper.PertyPrint(gen)

    build(gen)

    for i := 0; i < 100; i++ {
        //insert2(session)
    }

}

