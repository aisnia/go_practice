package main

import (
	"fmt"
	"math"
)

//全局变量
//大写开头，public的
var PI = 3.1415926

//小写 private的
var str = "hello"

func main() {
	//局部变量定义了就要使用
	//声明变量 var name type
	//var a, b int  //都是int型

	//分开声明
	//var a int
	//var b bool
	//var str string

	//或者
	var (
		a   int
		b   bool
		str string
	)
	//都有初值
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(str)

	//类型推断 int型
	var x = 1
	var y = "abc"
	fmt.Println(x)
	fmt.Println(y)

	//省略var,也可以自动类型转换
	//但是不能用于 全局变量的声明与赋值
	n := 2
	s := "bcd"
	fmt.Println(n)
	fmt.Println(s)

	//多变量声明  int可以不加
	var name1, name2, name3 int = 1, 2, 3
	// name1,name2,name3 := 1,2,3
	fmt.Println(name1, name2, name3)

	fmt.Println(PI)
	fmt.Println(str)

}

//最先执行，类似于构造方法，但是在初始化后面，类似Java
func init() {
	fmt.Println("init函数")
	PI = 4 * math.Atan(1)

}
