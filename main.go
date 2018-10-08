package main

import (
	"GoCQLSockets/client"
	"GoCQLSockets/server/database"
	"GoCQLSockets/server/sockets"
	"flag"
	"strings"
)

func main() {
	flagMode := flag.String("mode", "server", "start in client or server mode")
	flag.Parse()
	if strings.ToLower(*flagMode) == "server" {
		CassandraSession := cassandra.Session
		defer CassandraSession.Close()
		sockets.StartServerMode()
	} else {
		client.StartClientMode()
	}

}
