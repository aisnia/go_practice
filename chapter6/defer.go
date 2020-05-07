package main

import "fmt"

func trace(s string)   { fmt.Println("entering:", s) }
func untrace(s string) { fmt.Println("leaving:", s) }

func main() {
	b()
}
func a() {
	trace("a")          //3 entering: a
	defer untrace("a")  //入栈
	fmt.Println("in a") //4 in a
	//a先出栈，5 出栈 leaving: a
}

func b() {
	trace("b")          //1    entering: b
	defer untrace("b")  //入栈，先进后出
	fmt.Println("in b") // in b
	a()                 //调用a

	//b出栈, 6 leaving: b
}
