package main

import (
	"fmt"
	"os"
)

func tel(ch chan int) {
	for i := 0; i < 15; i++ {
		ch <- i
	}
	//死锁解决方法，关闭通道
	close(ch)
}

//另一种解决办法，
func tel2(ch chan int, quit chan bool) {
	for i := 0; i < 15; i++ {
		ch <- i
	}
	quit <- true
}

//练习 14.7

func main() {
	ch := make(chan int)
	//go tel(ch)
	//第一种 会死锁
	//for{
	//	fmt.Println(<-ch)
	//}

	//第二种：close
	//for{
	//	//ok代表通道 是否关闭
	//	if i,ok := <- ch;ok{
	//		fmt.Println(i)
	//	}
	//}

	//第三种select
	quit := make(chan bool)
	go tel2(ch, quit)

	for {
		select {
		case v := <-ch:
			fmt.Println(v)
		case <-quit:
			os.Exit(0)
		}
	}

}
