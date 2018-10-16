package main

import (
	"GoCQLSockets/examples"
	"GoCQLSockets/server/database"
	"GoCQLSockets/server/tcp_server"
	"flag"
	"strings"
)

func main() {
	flagMode := flag.String("mode", "server", "start in tcp_client or server mode")
	flag.Parse()
	if strings.ToLower(*flagMode) == "server" {
		cassandra.StartCassandra()
		CassandraSession := cassandra.Session
		defer CassandraSession.Close()
		tcp_server.StartServerMode()
	} else {
		examples.Example()
	}

}
