package main

import (
	"fmt"
	"time"
)

func producer(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i * 10
	}
}

func consumer(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}

func main() {
	ch := make(chan int)

	go producer(ch)

	go consumer(ch)

	time.Sleep(1e9)
}
