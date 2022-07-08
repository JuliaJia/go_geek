package main

import (
	//"bufio"
	"fmt"
	"net"
	"protocol"
)

func client_tcp_fix_length(conn net.Conn) {
	fmt.Println("client, fix length")
	sendByte := make([]byte, 1024)
	sendMsg := "{\"这是第一条信息\":啦啦啦,\"这是第二条信息\",嘿嘿嘿}"
	for i := 0; i < 1000; i++ {
		tempByte := []byte(sendMsg)
		for j := 0; j < len(tempByte) && j < 1024; j++ {
			sendByte[j] = tempByte[j]
		}
		_, err := conn.Write(sendByte)
		if err != nil {
			fmt.Println(err, ",err index=", i)
			return
		}
		fmt.Println("发送一次信息！")
	}
}

func client_tcp_delimiter(conn net.Conn) {
	fmt.Println("client, delimiter based")
	var sendMsgs string
	sendMsg := "{\"这是第一条信息\":啦啦啦,\"这是第二条信息\",嘿嘿嘿}\n"
	for i := 0; i < 1000; i++ {
		sendMsgs += sendMsg
		_, err := conn.Write([]byte(sendMsgs))
		if err != nil {
			fmt.Println(err, ",err index=", i)
			return
		}
		fmt.Println("发送一次信息！")
	}
}

func client_tcp_frame_decoder(conn net.Conn) {
	fmt.Println("client, length field based frame decoder")
	for i := 0; i < 1000; i++ {
		sendMsg := "{\"这是第一条信息\":啦啦啦,\"这是第二条信息\",嘿嘿嘿}"
		_, err := conn.Write(protocol.Packet([]byte(sendMsg)))
		if err != nil {
			fmt.Println(err, ",err index=", i)
			return
		}
		fmt.Println("发送一次信息！")
	}
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:7777")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	//client_tcp_delimiter(conn)
	client_tcp_delimiter(conn)
}
