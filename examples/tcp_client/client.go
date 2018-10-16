package tcp_client

import (
	"GoCQLSockets/config"
	"bytes"
	"fmt"
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
	connection, error := net.Dial("tcp", config.Config.Client.IPAddress+config.Config.Client.Port)
	if error != nil {
		fmt.Println(error)
		reconnect()
	}else{
		Client.Socket = connection
	}
	go Client.receive()
}

func reconnect(){
	for {
		fmt.Println("Reconnecting tcp_client...")
		connection, error := net.Dial("tcp", config.Config.Client.IPAddress+config.Config.Client.Port)
		if error != nil {
			fmt.Println(error)
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

func (client *Iets) Write(message []byte) int{
		length, err := client.Socket.Write(message)
		if err != nil{
			return -1
		}
		return length
}