package examples

import (
	"GoCQLSockets/examples/tcp_client"
	"GoCQLSockets/parser"
	"encoding/binary"
	"fmt"
	"time"
)

func Example1GET() {
	tcp_client.StartClientMode()
	defer tcp_client.GlobalClient.Socket.Close()

	for {
		updateOrReceiveCounterByFlag(10, 300)
		time.Sleep(time.Second)
	}

}

func Example1Increase() {
	tcp_client.StartClientMode()
	defer tcp_client.GlobalClient.Socket.Close()

	for {
		updateOrReceiveCounterByFlag(1, 300)
		time.Sleep(time.Second*2)
	}
}

func Example1Decrease() {
	tcp_client.StartClientMode()
	defer tcp_client.GlobalClient.Socket.Close()


	for {
		updateOrReceiveCounterByFlag(2, 300)
		time.Sleep(time.Second*5)
	}
}


func updateOrReceiveCounterByFlag(flags, opCode uint32){
	var requestHeader = parser.MakeHeader(20, 1, 0, opCode)

	//OP_QUERY
	opCodeRequest := make([]byte, 4)
	binary.LittleEndian.PutUint32(opCodeRequest, flags)
	request := append(requestHeader, opCodeRequest...)
	tcp_client.GlobalClient.Write(request)

	select{
		case response, ok := <-tcp_client.GlobalClient.Data :
			if !ok {
				return
			} else {

				requestLengthResponse, _, responseID, opCodeResponse := parser.ParseHeader(response)
				if 1 == opCodeResponse && responseID == 1 {
					//Trim response
					response = response[16:]
					response = response[:requestLengthResponse]
					fmt.Print("Flag: ")
					fmt.Println(parser.ByteToInt(response, 0))
					if flags == 10 {
						flags := parser.ByteToInt(response, 4)
						fmt.Print("Counter: ")
						fmt.Println(flags)
					} else if flags == 100 {
						fmt.Print("Error: ")
					}

				}
			}
		case <-tcp_client.GlobalClient.Error :
			fmt.Println("Cancel 'wait for response', because server error")
			return
	}

}
