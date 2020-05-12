package main

import (
	"fmt"
	"time"
)

var values = [5]int{10, 11, 12, 13, 14}

func main() {
	// 版本A:
	for ix := range values { // ix是索引值
		func() {
			fmt.Print(ix, " ")
		}() // 调用闭包打印每个索引值
	}
	fmt.Println()
	// 版本B: 和A版本类似，但是通过调用闭包作为一个协程  4 4 4 4 4
	// 每个都是单体的值，因为协程可能在循环结束后还没有开始执行，而此时ix值是4。
	for ix := range values {
		go func() {
			fmt.Print(ix, " ")
		}()
	}
	time.Sleep(5e9)
	fmt.Println()
	// 版本C: 正确的处理方式
	//每次传入闭包 都赋值   并将每个协程的ix放置在栈中 顺序取决于那个协程先启动
	for ix := range values {
		go func(ix interface{}) {
			fmt.Print(ix, " ")
		}(ix)
	}
	time.Sleep(5e9)
	fmt.Println()
	// 版本D: 输出值:
	for ix := range values {
		//每次循环的 都会声明创建一个 val，是不共享的
		val := values[ix]
		go func() {
			fmt.Print(val, " ")
		}()
	}
	time.Sleep(1e9)
}

//输出：
//
//0 1 2 3 4
//
//4 4 4 4 4
//
//1 0 3 4 2
//
//10 11 12 13 14
