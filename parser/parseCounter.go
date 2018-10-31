package parser

import "encoding/binary"

var counter uint32 = 0

func parseCounter(request []byte) []byte {
	variable := make([]byte, 4)

	switch ByteToInt(request, 0) {

	case 1:
		counter++
		binary.LittleEndian.PutUint32(variable, 200)
	case 2:
		counter--
		binary.LittleEndian.PutUint32(variable, 200)
	case 10:
		binary.LittleEndian.PutUint32(variable, 200)
		result := make([]byte, 4)
		binary.LittleEndian.PutUint32(result, counter)
		variable = append(variable, result...)
	case 20:
		counter = 0
		binary.LittleEndian.PutUint32(variable, 200)
	default:
		binary.LittleEndian.PutUint32(variable, 300)
	}

	return variable
}