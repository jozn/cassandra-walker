package main

import (
	"fmt"
	"github.com/gocql/gocql"
	"ms/cassandra_walker/play/out"
)

func main() {
	cluster := gocql.NewCluster("192.168.1.250")
	cluster.Keyspace = "sunc"
	cluster.Consistency = gocql.One
	session, _ := cluster.CreateSession()
	defer session.Close()

	for j := 0; j < 3; j++ {
		a := &xc.Admins{
			FirstName: "hamid",
			Userid:    fmt.Sprintf("%d", j),
		}
		a.Save(session)
	}
}
