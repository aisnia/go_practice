package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	//打开链接
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		//无法创建连接
		fmt.Println("Error Dialing", err.Error())
		return
	}

	inputReader := bufio.NewReader(os.Stdin)

	fmt.Println("First,what is you name?")

	//输入连接名
	clientName, _ := inputReader.ReadString('\n')

	//去除前后回车
	trimmedClient := strings.Trim(clientName, "\n")

	for {
		fmt.Println("What to send to the server? Type Q to quit :")
		input, _ := inputReader.ReadString('\n')

		trimmedInput := strings.Trim(input, "\n")

		if trimmedInput == "Q" {
			return
		}

		//写入给服务器
		_, err = conn.Write([]byte(trimmedClient + " says: " + trimmedInput))

	}
}
