package main

import "fmt"

func main() {
	//运行时的错误，即运行时执行 panic ，程序会崩溃并且抛出一个runtime.Error的接口值
	fmt.Println("Starting the program")
	panic("A severe error occurred: stopping the program!")
	fmt.Println("Ending the program")
}
