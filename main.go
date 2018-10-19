package main

import (
	"GoCQLSockets/examples"
	"GoCQLSockets/server/database/connector"
	"GoCQLSockets/server/tcp_server/connector"
	"flag"
	"strings"
)

func main() {
	flagMode := flag.String("mode", "server", "start in tcp_client or server mode")
	flag.Parse()
	if strings.ToLower(*flagMode) == "server" {
		go cassandra.StartCassandra()
		CassandraSession := cassandra.Session
		defer CassandraSession.Close()
		connector.StartServerMode()
	} else {
		examples.Example()
	}

}
