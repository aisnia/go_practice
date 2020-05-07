package main

import "fmt"

func main() {
	//将函数作为参数传递
	callback(1, Add)
}

func Add(a, b int) {
	//求和函数
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
}

func callback(y int, f func(int, int)) {
	//此时的f为 ADD
	f(y, 2)
}
