package tcp_client

import (
	"GoCQLSockets/examples/config"
	"bytes"
	"crypto/tls"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"
)

type Iets struct {
	Socket net.Conn
	Data   chan []byte
}

var Client *Iets

func StartClientMode() {
	Client = &Iets{}
	fmt.Println("Starting tcp_client...")
	reconnect()
	go Client.receive()
}

func reconnect() {
	for {

		cert, err := tls.LoadX509KeyPair(config.Config.Client.Certs.Directory+config.Config.Client.Certs.Pem, config.Config.Client.Certs.Directory+config.Config.Client.Certs.Key)
		if err != nil {
			log.Fatal(err)
		}
		tlsConfig := tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}
		connection, err := tls.Dial("tcp", config.Config.Client.IPAddress+config.Config.Client.Port, &tlsConfig)

		if err != nil {
			fmt.Println(err)
			fmt.Println("Reconnecting tcp_client...")
			time.Sleep(time.Duration(rand.Intn(config.Config.Client.ReconnectTime)) * time.Second)
		} else {
			Client.Socket = connection
			return
		}
	}

}

func (client *Iets) receive() {
	for {
		message := make([]byte, 4096)
		length, err := client.Socket.Read(message)
		if err != nil {
			client.Socket.Close()
			reconnect()
			continue
		}
		if length > 0 {
			message = bytes.Trim(message, "\x00")
			fmt.Println(string(message))
		}
	}

}

func (client *Iets) Write(message []byte) int {
	length, err := client.Socket.Write(message)
	if err != nil {
		return -1
	}
	return length
}
