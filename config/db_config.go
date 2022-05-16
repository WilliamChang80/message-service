package config

import (
	"fmt"
	"log"

	"github.com/gocql/gocql"
)

func InitDb() *gocql.Session {
	cluster1 := fmt.Sprintf("%s:%d", GetDbUrl(), GetDbPort())
	cluster := gocql.NewCluster(cluster1)
	cluster.Keyspace = GetDbKeyspace()
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()

	if err != nil {
		log.Fatal("Failed to connect cluster")
	}

	return session
}
