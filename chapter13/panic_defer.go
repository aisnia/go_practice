// panic_defer.go
package main

import "fmt"

//我的猜想 Calling g.    Printing in g 0, ...1, ...2 ,...3,
// i等于4时，   Panicking!   Defer in g 4,    Recovered in f 4

//中间的 Defer in g 4 错了      原因：painc之后 defer的语句也不会执行 所以没有打印 Defer in g 4
//其次 前面每个递归g函数时，都没有panic发生错误，都会调用想要的defer 有 Defer in g 3，。。。2，1，0)
func main() {
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
