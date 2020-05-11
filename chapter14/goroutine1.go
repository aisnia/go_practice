package main

import (
	"fmt"
	"time"
)

func main() {

	//go每次都开启一个独立的协程 去 处理函数单元
	//当 main() 函数返回的时候，程序退出：它不会等待任何其他非 main 协程的结束。而导致其他协程也 消亡
	//服务器一般用一个死循环
	fmt.Println("In main()")
	go longWait()
	go shortWait()
	fmt.Println("About to sleep in main()")
	// sleep works with a Duration in nanoseconds (ns) !
	time.Sleep(10 * 1e9)
	fmt.Println("At the end of main()")
}

func longWait() {
	fmt.Println("Beginning longWait()")
	time.Sleep(5 * 1e9) // sleep for 5 seconds
	fmt.Println("End of longWait()")
}

func shortWait() {
	fmt.Println("Beginning shortWait()")
	time.Sleep(2 * 1e9) // sleep for 2 seconds
	fmt.Println("End of shortWait()")
}
