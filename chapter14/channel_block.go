package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	go pump(ch1)

	//从通道获取数据，如果没有则会阻塞
	//fmt.Println(<- ch1)
	//main结束，其他协程也结束，只输出1

	go suck(ch1)

	//一秒钟
	time.Sleep(1e9)
}

//生产者
func pump(ch chan int) {
	for i := 1; ; i++ {
		ch <- i
	}
}

//消费者
func suck(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}
