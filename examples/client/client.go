package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gocql/gocql"
	"net"
	"time"
)

type Client struct {
	socket net.Conn
	data   chan []byte
}

type Measurements struct {
	Measurements []float32 `json:"measurements"`
	CreationDate time.Time `json:"creation_date"`
}
type Error struct {
	ErrorMessage string
}

type Request struct {
	Type      int64      `json:"name"`
	ID        gocql.UUID `json:"id"`
	BeginDate time.Time  `json:"begin_date"`
	EndDate   time.Time  `json:"end_date"`
}

func StartClientMode() {
	fmt.Println("Starting client...")
	connection, error := net.Dial("tcp", "localhost:12345")
	if error != nil {
		fmt.Println(error)
	}
	client := &Client{socket: connection}
	go client.receive()
	defer connection.Close()
	requestMeasurementData(connection)
}

func (client *Client) receive() {
	for {
		message := make([]byte, 4096)
		length, err := client.socket.Read(message)
		if err != nil {
			client.socket.Close()
			break
		}
		if length > 0 {
			message = bytes.Trim(message, "\x00")
			fmt.Println("RECEIVED: " + string(message))
			//doSomethingWithReceivedData(message)
		}
	}
}

func requestMeasurementData(connection net.Conn) {
	for {
		beginDate := time.Date(
			2018, 1, 1, 0, 0, 0, 0, time.UTC)
		endDate := time.Date(
			2018, 8, 1, 0, 0, 0, 0, time.UTC)

		id, err := gocql.ParseUUID("6623c850-c008-11e8-a355-529269fb1459")
		if err != nil {
			fmt.Println(err.Error())
		} else {
			request := Request{ID: id, Type: 0, BeginDate: beginDate, EndDate: endDate}
			result, err := json.Marshal(request)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				connection.Write(result)
				fmt.Print("Sent: ")
				fmt.Println(result)
			}
		}
		time.Sleep(time.Second * 10)
	}

}

//func doSomethingWithReceivedData(message []byte){
//var errorMessage Error
//err := json.Unmarshal(message, &errorMessage)
//if err == nil && errorMessage.ErrorMessage != ""{
//	//can do something with error
//}else{
//	var measurement Measurements
//	err := json.Unmarshal(message, &measurement)
//	if err != nil{
//		fmt.Println(err)
//	}
//	//else {
//	//	Can do something with measuremnt
//	//}
//}
//}
