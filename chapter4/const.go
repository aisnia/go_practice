package main

import "fmt"

func main() {
	/*

	 */

	//	定义常量,可以省略类型，编译器会推断
	const PI = 3.14159
	//显示
	const s1 string = "abc"
	//隐式
	const s2 = "abc"

	//未指定类型常量会根据环境使用 来推断
	//int初值为0，5是常量被推断为int型，而且声明了变量后必须要使用，否则报错
	var n int
	n += 5
	fmt.Println(n)

	const c1 = 2 / 3
	//错误，常量是编译期就要确定的，自定义的函数无法用于常量的赋值
	//const c2 = getNum()

	//多种定义的方式
	//const Monday, Tuesday, Webnesday, Thursday, Friday, Saturday = 1, 2, 3, 4, 5, 6
	const (
		Monday, Tuesday, Webnesday = 1, 2, 3
		Thursday, Friday, Saturday = 4, 5, 6
	)

	//	枚举 性别
	const (
		MALE   = 0
		FEMALE = 1
		UNKNOW = 3
	)
	//  没有赋初值则与上面一行一样
	const (
		a int = 1
		b
		c string = "abc"
		d
		e
	)
	fmt.Printf("%T,%d\n", a, a)
	fmt.Printf("%T,%d\n", b, b)
	fmt.Printf("%T,%s\n", c, c)
	fmt.Printf("%T,%s\n", d, d)
	fmt.Printf("%T,%s\n", e, e)

}

func getNum() int {
	return 1
}
