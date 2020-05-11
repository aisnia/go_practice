package main

import (
	"fmt"
	"time"
)

func main() {
	stream := pump()
	//go suck(stream)
	suckFor(stream)
	time.Sleep(1e9)
}

//工厂 返回一个通道
func pump() chan int {
	ch := make(chan int)
	//不将通道作为参数传递给协程，而用函数来生成一个通道并返回（工厂角色）；
	// 函数内有个匿名函数被协程调用。
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func suck(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}

func suckFor(ch chan int) {
	go func() {
		//通过for 去遍历通道， 可以在里面用一个协程运行 匿名函数
		for v := range ch {
			fmt.Println(v)
		}
	}()

}
