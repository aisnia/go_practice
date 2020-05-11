package main

import (
	"fmt"
	"strconv"
)

//请求，内置了通道
type Request struct {
	a, b   int
	replyc chan int
}

func (this Request) String() string {
	return strconv.Itoa(<-this.replyc)
}

type binOp func(a, b int) int

func run(op binOp, req *Request) {
	req.replyc <- op(req.a, req.b)
}

//轮训两个通道 service是请求，  quit是否退出
func server(op binOp, service chan *Request, quit chan bool) {
	for {
		select {
		case req := <-service:
			go run(op, req)
		case <-quit:
			return
		}
	}
}

//开启服务器，设置好对应的通道，并且 go service
func startServer(op binOp) (service chan *Request, quit chan bool) {
	service = make(chan *Request)
	quit = make(chan bool)
	go server(op, service, quit)
	return service, quit
}

func main() {
	//传入我们处理操作的函数，这里是相加
	adder, quit := startServer(func(a, b int) int { return a + b })
	//模拟 100个请求
	const N = 100
	var reqs [N]Request

	//初始化 然后传个服务器的通道
	for i := 0; i < N; i++ {
		req := &reqs[i]
		req.a = i
		req.b = i + N
		req.replyc = make(chan int)
		adder <- req
	}
	// checks:检查每个通道处理后的结果  i + i + N 即 2 * i+N
	for i := N - 1; i >= 0; i-- { // doesn't matter what order
		fmt.Println(reqs[i])
		//	if <-reqs[i].replyc != N+2*i {
		//		fmt.Println("fail at", i)
		//	} else {
		//		fmt.Println("Request ", i, " is ok!")
		//	}
	}
	quit <- true
	fmt.Println("done")
}
