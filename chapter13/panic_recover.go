// panic_recover.go
package main

import (
	"fmt"
)

func badCall() {
	//通常将 panic 转换成 error 来告诉调用方为何出错，然后调用者通过recover获取error对象来获取错误的信息
	//panic(error)
	panic("bad end")
}

func test() {
	//defer 修饰的一个方法，在程序中断前会执行，然后就可以在里面 调用recover()返回一个错误，做出我们的处理，类似于catch，让程序可以继续运行
	// 只要在这个方法里面定义了defer修饰的方法，那么就可以处理panic抛出的错误，如果没有panic则 recover返回的是nil
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Panicing %s\r\n", e)
		}
	}()

	badCall()
	fmt.Printf("After bad call\r\n") // <-- wordt niet bereikt
}

func main() {
	fmt.Printf("Calling test\r\n")
	test()
	fmt.Printf("Test completed\r\n")
}
