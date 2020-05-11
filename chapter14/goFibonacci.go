package main

import (
	"fmt"
	"os"
)

func fibonacci(n int, ch chan int) {
	for p1, p2 := 0, 1; p1 < n; {
		ch <- p1
		t := p1 + p2
		p1 = p2
		p2 = t
	}
	close(ch)
}

func fibonacci2(n int, ch chan int, quit chan bool) {
	for p1, p2 := 0, 1; p1 < n; {
		ch <- p1
		t := p1 + p2
		p1 = p2
		p2 = t
	}
	quit <- true
}

func main() {
	ch := make(chan int)

	//第一张 close
	//go fibonacci(30,ch)

	//for{
	//	if i,ok := <-ch;ok{
	//		fmt.Println(i)
	//	}
	//}

	quit := make(chan bool)

	go fibonacci2(30, ch, quit)
	//第二种select
	for {
		select {
		case v := <-ch:
			fmt.Println(v)
		case <-quit:
			os.Exit(0)
		}
	}

}
