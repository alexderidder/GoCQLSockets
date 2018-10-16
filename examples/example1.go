package examples

import (
	"GoCQLSockets/examples/tcp_client"
	"time"
)

func Example() {
	tcp_client.StartClientMode()
	defer tcp_client.Client.Socket.Close()
	for {
		tcp_client.Client.Write([]byte("HALLO"))
		time.Sleep(time.Second)
	}

}
