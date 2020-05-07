package main

import "fmt"

func main() {

	fmt.Println("before")
	greeting()
	fmt.Println("after")

	//Go 默认使用按值传递来传递参数 _抛弃不需要的值
	var x int
	x, _ = Multiply(1, 2, 3)
	fmt.Println(x)

	//长参数,求最大值
	fmt.Println(max(1, 3, 5, 12, 5, 123, 21))

	//defer追踪 推迟到函数返回之前 ,类似与java里面的 finally，用于释放资源
	fun1()

	f()
}
func f() {
	i := 0
	defer fmt.Println(i) //打印的是0,参数是会延迟的
	i++
	return
}

func fun1() {
	fmt.Println("fun1 before")
	defer fun2()
	fmt.Println("fun1 after")
}

func fun2() {
	fmt.Println("fun2")
}

func greeting() {
	fmt.Println("greeting")
}

//返回值在后面
func Multiply(a, b, c int) (int, string) {
	return a * b * c, ""
}
func max(a ...int) int {
	max := a[0]
	for i := 1; i < len(a); i++ {
		if max < a[i] {
			max = a[i]
		}
	}
	return max
}
