package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Starting the Server")
	//创建listener
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error Listening", err.Error())
		return
	}

	//监听来着客户端的链接
	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("Error Accept", err.Error())
			return
		}
		//另起一个协程 进行处理
		go doServer(conn)
	}
}

func doServer(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		//读取数据 到buf中
		len, err := conn.Read(buf)

		if err != nil {
			fmt.Println("Error reading", err.Error())
			return
		}

		fmt.Printf("received data:\n%v", string(buf[:len]))
	}
}
