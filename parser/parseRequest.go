package parser

import (
	"encoding/binary"
	"fmt"
)

func ParseHeader(request []byte) (uint32, uint32, uint32, uint32) {
	result := make([]uint32, 4)
	for i := 0; i < 4; i++ {
		result[i] = ByteToInt(request, i*4)
	}

	return result[0], result[1], result[2], result[3]
}

func ParseOpCode(opCode uint32, message []byte) []byte {
	switch opCode {

	case 300:
		{
			return parseCounter(message)
		}

	default:
		fmt.Println("opcode unknown")
	}
	return []byte{1, 1, 1, 1}
}


func MakeHeader(messageLength, requestID, responseID, opCode uint32) []byte {
	var requestHeader []byte
	//Request headers
	variable := make([]byte, 4)

	binary.LittleEndian.PutUint32(variable, messageLength)
	requestHeader = append(requestHeader, variable...)

	binary.LittleEndian.PutUint32(variable, requestID)
	requestHeader = append(requestHeader, variable...)

	binary.LittleEndian.PutUint32(variable, responseID)
	requestHeader = append(requestHeader, variable...)

	binary.LittleEndian.PutUint32(variable, opCode)
	requestHeader = append(requestHeader, variable...)
	return requestHeader
}

func ByteToInt(request []byte, beginIndex int) uint32 {
	var result uint32
	result |= uint32(request[beginIndex])
	beginIndex++
	result |= uint32(request[beginIndex]) << 8
	beginIndex++
	result |= uint32(request[beginIndex]) << 16
	beginIndex++
	result |= uint32(request[beginIndex]) << 24
	return result
}
