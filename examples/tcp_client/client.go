package tcp_client

import (
	"GoCQLSockets/examples/config"
	"crypto/tls"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"
)

type Client struct {
	Socket net.Conn
	Data   chan []byte
	Error   chan bool
}

var GlobalClient *Client

func StartClientMode() *Client {
	GlobalClient = &Client{ Data: make(chan []byte),  Error: make(chan bool, 1)}
	fmt.Println("Starting tcp_client...")
	GlobalClient.reconnect()
	go GlobalClient.receive()
	return GlobalClient
}

func (client *Client) reconnect() {
	cert, err := tls.LoadX509KeyPair(config.Config.Client.Certs.Directory+config.Config.Client.Certs.Pem, config.Config.Client.Certs.Directory+config.Config.Client.Certs.Key)
	if err != nil {
		log.Fatal(err)
	}
	tlsConfig := tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}
	for {
		connection, err := tls.Dial("tcp", config.Config.Client.IPAddress+config.Config.Client.Port, &tlsConfig)

		if err != nil {
			fmt.Println(err)
			fmt.Println("Reconnecting tcp_client...")
			time.Sleep(time.Duration(rand.Intn(config.Config.Client.ReconnectTime)) * time.Second)
		} else {
			client.Socket = connection
			fmt.Println("tcp_client is connected")
			return
		}
	}

}

func (client *Client) receive() {
	for {
		message := make([]byte, 4096)
		length, err := client.Socket.Read(message)
		if err != nil {
			client.Socket.Close()
			client.reconnect()
			client.Error <- true
			continue
		}
		if length > 0 {
			client.Data <- message
		}
	}

}

func (client *Client) Write(message []byte) int {
	length, err := client.Socket.Write(message)
	if err != nil {
		return -1
	}

	return length
}
