package main

import (
	"GoCQLSockets/examples"
	"GoCQLSockets/server/database/connector"
	"GoCQLSockets/server/tcp_server/connector"
	"flag"
	"fmt"
	"strings"
)

func main() {
	flagMode := flag.String("mode", "", "start in tcp_client or server mode")
	flag.Parse()

	switch strings.ToLower(*flagMode) {
	case "server":
		go cassandra.StartCassandra()
		CassandraSession := cassandra.Session
		defer CassandraSession.Close()
		connector.StartServerMode()
	case "example1_get":
		examples.Example1GET()
	case "example1_inc":
		examples.Example1Increase()
	case "example1_dec":
		examples.Example1Decrease()
	default:
		fmt.Println("Mode isnt specified or doesn't exists")
	}
}
