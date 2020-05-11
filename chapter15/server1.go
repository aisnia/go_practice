package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

//增加一个检查错误的函数 checkError(error)；讨论如下方案的利弊：为什么这个重构可能并没有那么理想？看看在 示例15.14 中它是如何被解决的
//使客户端可以通过发送一条命令 SH 来关闭服务器
//让服务器可以保存已经连接的客户端列表（他们的名字）；当客户端发送 WHO 指令的时候，服务器将显示如下列表：
var users map[string]int

func main() {
	fmt.Println("server is starting.....")
	users = make(map[string]int)
	listener, err := net.Listen("tcp", "localhost:8080")
	checkError(err)

	for {
		conn, err := listener.Accept()
		checkError(err)
		go doserver(conn)
	}
}

func doserver(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		len, err := conn.Read(buf)
		checkError(err)
		//第一次发过来的是 name

		data := string(buf[:len])

		//关闭服务器
		if strings.Contains(data, ": SH") {
			os.Exit(0)
		}

		if strings.Contains(data, ": WHO") {
			DisplayList()
		}

		index := strings.Index(data, "says")

		name := data[:index-1]

		users[name] = 1

		fmt.Printf("Received data: %v", data)
	}
}

func DisplayList() {
	fmt.Println("This is the client list: 1:active, 0=inactive")
	for name, i := range users {
		fmt.Printf("User %s is %d\n", name, i)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error ", err.Error())
		//抛出异常，终止程序
		panic(err)
	}
}
