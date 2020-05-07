package main

import (
	"fmt"
	"time"
)

func main() {
	res := 0
	start := time.Now()
	for i := 0; i < 40; i++ {
		res = fibonacci(i)
		//res = fibonacci_buffer(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, res)
	}
	end := time.Now()
	//end 减去 start
	delta := end.Sub(start)
	fmt.Printf("time : %d", delta)

}

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}

const LIM = 41

//记忆备忘录
var fibs [LIM]int

func fibonacci_buffer(n int) (res int) {
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	fibs[n] = res
	return
}
