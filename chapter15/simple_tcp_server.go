// Simple multi-thread/multi-core TCP server.
package main

import (
	"flag"
	"fmt"
	"net"
	"syscall"
)

const maxRead = 25

func main() {
	//flag 包解析命令行参数
	flag.Parse()
	//如果不是两个参数 报错
	if flag.NArg() != 2 {
		panic("usage: host port")
	}
	//Sprintf 格式话字符串，将两个参数 传入，作为ip和端口
	hostAndPort := fmt.Sprintf("%s:%s", flag.Arg(0), flag.Arg(1))
	//初始化服务器，返回监听
	listener := initServer(hostAndPort)
	for {
		//监听连接，然后新建一个协程 进行处理
		conn, err := listener.Accept()
		checkError(err, "Accept: ")
		go connectionHandler(conn)
	}
}
func initServer(hostAndPort string) net.Listener {
	// 解析得到 ip 与 端口， 可以看看格式是否正确
	serverAddr, err := net.ResolveTCPAddr("tcp", hostAndPort)
	checkError(err, "Resolving address:port failed: '"+hostAndPort+"'")
	//进行监听，返回listener
	listener, err := net.ListenTCP("tcp", serverAddr)
	checkError(err, "ListenTCP: ")
	println("Listening to: ", listener.Addr().String())
	return listener
}

//请求的处理
func connectionHandler(conn net.Conn) {
	//获取 地址
	connFrom := conn.RemoteAddr().String()
	println("Connection from: ", connFrom)
	//写回去个响应，代表连接成功
	sayHello(conn)
	for {
		var ibuf []byte = make([]byte, maxRead+1)
		//读取数据
		length, err := conn.Read(ibuf[0:maxRead])
		ibuf[maxRead] = 0 // to prevent overflow
		switch err {
		case nil:
			//没有错误，则发送消息
			handleMsg(length, err, ibuf)
		case syscall.EAGAIN: // 资源暂时不可以，稍后重试
			continue
		default:
			//关闭连接，goto跳出循环
			goto DISCONNECT
		}
	}
DISCONNECT:
	err := conn.Close()
	println("Closed connection: ", connFrom)
	checkError(err, "Close: ")
}
func sayHello(to net.Conn) {
	obuf := []byte{'L', 'e', 't', '\'', 's', ' ', 'G', 'O', '!', '\n'}
	wrote, err := to.Write(obuf)
	checkError(err, "Write: wrote "+string(wrote)+" bytes.")
}
func handleMsg(length int, err error, msg []byte) {
	if length > 0 {
		print("<", length, ":")
		for i := 0; ; i++ {
			if msg[i] == 0 {
				break
			}
			fmt.Printf("%c", msg[i])
		}
		print(">")
	}
}
func checkError(error error, info string) {
	if error != nil {
		panic("ERROR: " + info + " " + error.Error()) // terminate program
	}
}
