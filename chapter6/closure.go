package main

import "fmt"

func main() {
	f()
	fmt.Println(f2())

	// 函数别名
	p2 := Add2()
	fmt.Printf("Call Add2 for 3 gives: %v\n", p2(3))
	//以2为参数的返回的 另一个函数 即里面那个 这时a已经赋值为 2了
	TwoAdder := Adder(2)
	//然后在调用该函数，a + b = 5
	fmt.Printf("The result is: %v\n", TwoAdder(3))
}
func f() {
	for i := 0; i < 4; i++ {
		//闭包，即匿名的函数，函数的地址给了g
		g := func(i int) {
			fmt.Printf("%d ", i)
		}
		g(i)
		//类型是函数，地址都是一样的
		fmt.Printf(" - g is of type %T and has value %v\n", g, g)
	}
}

//变量 ret 的值为 2，因为 ret++ 是在执行 return 1 语句后发生的。
//闭包可以改变函数的返回值
func f2() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}

//3传进来，然后到闭包里面作为参数，返回 b+2 即 3 + 2
func Add2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

//3传进来
func Adder(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}
