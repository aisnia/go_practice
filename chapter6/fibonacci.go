package main

import "fmt"

func main() {
	fmt.Println(fibonacci(5)) // 1 1 2 3 5 8
	fmt.Println(p64(5))

	//10打印到1
	p65(10)

	fmt.Println("-----------------")
	//阶乘
	fmt.Println(p66(30))

}
func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}

func p64(n int) (index, res int) {
	if n <= 1 {
		index, res = 1, 1
	} else {
		index = n
		_, res = p64(n - 1)
		_, t := p64(n - 2)
		res += t
	}
	return
}
func p65(n int) {
	if n < 1 {
		return
	}
	fmt.Println(n)
	p65(n - 1)
}

func p66(n int) uint64 {
	if n == 0 {
		return 1
	}
	pre := p66(n - 1)
	fmt.Println(pre)
	return pre * uint64(n)

}
