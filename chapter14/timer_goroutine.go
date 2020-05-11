package main

import (
	"fmt"
	"time"
)

func main() {
	//time.Tick() 函数声明为 Tick(d Duration) <-chan Time，
	// 当你想返回一个通道而不必关闭它的时候这个函数非常有用：它以 d 为周期给返回的通道发送时间，d是纳秒数。
	tick := time.Tick(1e8)
	//在 Duration d 之后，当前时间被  发到返回的通道；所以它和 NewTimer(d).C 是等价的；它类似 Tick()，但是 After() 只发送一次时间。
	boom := time.After(5e8)
	for {
		//select 是伪随机
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(5e7)
		}
	}
}
