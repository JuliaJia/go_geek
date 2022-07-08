package protocol

import (
	"bytes"
	"encoding/binary"
)

const (
	Ch   = "www.baidu.com"
	Chl  = 13
	Csdl = 4
)

// 封包
func Packet(message []byte) []byte {
	// 头部信息 + body长度 + 消息
	return append(append([]byte(Ch), IntToBytes(len(message))...), message...)
}

// 解包
func Unpack(buffer []byte, readerChannel chan []byte) []byte {
	length := len(buffer)

	var i int
	for i = 0; i < length; i++ {
		if length < i+Chl+Csdl {
			break
		}
		if string(buffer[i:i+Chl]) == Ch {
			ml := BytesToInt(buffer[i+Chl : i+Chl+Csdl])
			if length < i+Chl+Csdl+ml {
				break
			}
			data := buffer[i+Chl+Csdl : i+Chl+Csdl+ml] // 信息
			readerChannel <- data

			i += Chl + Csdl + ml - 1 // end index
		}
	}

	if i == length {
		return make([]byte, 0)
	}
	return buffer[i:] // return message
}

// 整形转换成字节 int32 4个字节
func IntToBytes(n int) []byte {
	x := int32(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

// 字节转换成整形
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}
