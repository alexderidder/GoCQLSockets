package cassandra

import (
	"GoCQLSockets/config"
	"fmt"
	"github.com/gocql/gocql"
)

var Session *gocql.Session

var cluster *gocql.ClusterConfig

func StartCassandra() {
	clusterConfig()
	connect()
}

func connect(){
	var err error
	Session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	fmt.Println("cassandra connection")
}

func clusterConfig(){
	var ipAddresses []string
	for _, value :=  range config.Config.Database.Clusters{
		ipAddresses = append(ipAddresses, value.IPAddress)
	}
	cluster = gocql.NewCluster(ipAddresses...)
	cluster.Keyspace = config.Config.Database.Keyspace
}