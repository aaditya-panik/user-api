package cassandra

import (
	"github.com/gocql/gocql"
	"log"
)

var Session *gocql.Session

func init() {
	var err error

	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "userapi"
	Session, err = cluster.CreateSession()
	if err != nil {
		log.Print(err.Error())
		panic(err)
	}
	log.Print("## Cassandra Initialized")
}
