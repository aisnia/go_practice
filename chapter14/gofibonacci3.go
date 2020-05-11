package main

import (
	"fmt"
	"time"
)

func dup3(in <-chan int) (<-chan int, <-chan int, <-chan int) {
	a, b, c := make(chan int, 2), make(chan int, 2), make(chan int, 2)
	go func() {
		for {
			x := <-in
			a <- x
			b <- x
			c <- x
		}
	}()
	return a, b, c
}

//计算
func fib() <-chan int {
	x := make(chan int, 2)
	//三个通道，只要 x 有数据，就会传给他们，
	a, b, out := dup3(x)

	go func() {
		x <- 0
		x <- 1
		<-a
		for {
			//开始
			//a 1 1 2 3 5    //a,b每次都会消费的
			//b 0 1 1 2 3 5
			//c 0 1 1 2 3 5
			x <- <-a + <-b
		}
	}()
	//0 去除
	<-out
	return out
}

func main() {
	start := time.Now()
	x := fib()
	for i := 0; i < 10; i++ {
		fmt.Println(<-x)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}
