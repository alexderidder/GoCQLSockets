package connector

import (
	"GoCQLSockets/server/config"
	"bytes"
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"strconv"
)

type ClientManager struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

type Error struct {
	ErrorMessage string
}

type Client struct {
	socket net.Conn
	data   chan []byte
}

func StartServerMode() {

	fmt.Println("Starting server...")
	cert, err := tls.LoadX509KeyPair(config.Config.Server.Certs.Directory+config.Config.Server.Certs.Pem, config.Config.Server.Certs.Directory+config.Config.Server.Certs.Key)

	if err != nil {
		log.Fatal(err)
	}

	tlsConfig := tls.Config{Certificates: []tls.Certificate{cert}, ClientAuth: tls.RequireAnyClientCert}
	tlsConfig.Rand = rand.Reader
	listener, err := tls.Listen("tcp", config.Config.Server.IPAddress+config.Config.Server.Port, &tlsConfig)

	if err != nil {
		fmt.Println(err)
	}
	manager := ClientManager{
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
	go manager.start()
	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		client := &Client{socket: connection, data: make(chan []byte)}
		manager.register <- client
		go manager.receive(client)
		go manager.send(client)
	}
}

func (manager *ClientManager) start() {
	for {
		select {
		case connection := <-manager.register:
			manager.clients[connection] = true
			fmt.Println("Added new connection!")
		case connection := <-manager.unregister:
			if _, ok := manager.clients[connection]; ok {
				close(connection.data)
				delete(manager.clients, connection)
				fmt.Println("A connection has terminated!")
			}

		}
	}
}

func (manager *ClientManager) receive(client *Client) {
	counter := 0
	for {
		message := make([]byte, 4096)
		length, err := client.socket.Read(message)
		if err != nil {
			manager.unregister <- client
			client.socket.Close()
			break
		}
		if length > 0 {
			counter++
			message = bytes.Trim(message, "\x00")
			fmt.Println(string(message) + " : " + strconv.Itoa(counter))
			client.data <- []byte(strconv.Itoa(counter))
		}
	}
}

func (manager *ClientManager) send(client *Client) {
	defer client.socket.Close()
	for {
		select {
		case message, ok := <-client.data:
			if !ok {
				return
			} else {
				client.socket.Write(append([]byte("Received:"), message...))
			}
		}
	}
}
