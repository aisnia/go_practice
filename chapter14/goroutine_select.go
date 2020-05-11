package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go pump1(ch1)
	go pump2(ch2)
	go suck(ch1, ch2)
	time.Sleep(1e9)

}

func pump1(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

func pump2(ch chan int) {
	for i := 0; ; i++ {
		ch <- i * 2
	}
}

func suck(ch1, ch2 chan int) {
	for {
		//死循环里面轮询
		select {
		case v := <-ch1:
			fmt.Println("channel1", v)
		case v := <-ch2:
			fmt.Println("channel2", v)
			//default:
			//做一些其他的事情,可以不阻塞select
		}
	}
}
