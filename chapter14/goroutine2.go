package main

import (
	"fmt"
	"time"
)

func main() {
	//创建一个string 类型的通道
	ch := make(chan string)

	//去除一个则会 panic 因为runtime会检查所有协程 释放
	go sendData(ch)
	go getData(ch)

	//休息 10的9次方 纳秒  即1秒 必须等到两个协程完成
	time.Sleep(1e9)
}

func sendData(ch chan string) {
	//向通道发生数据
	ch <- "Hello"
	ch <- "World"
	ch <- "!"
}

func getData(ch chan string) {
	var str string
	for {
		//从通道获取数据，当通道ch 数据空了 就不会输出了，而且main协程消亡，这个也会消亡的
		str = <-ch
		fmt.Printf("%s", str)
	}
}
