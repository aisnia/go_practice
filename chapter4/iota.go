package main

import "fmt"

func main() {
	/*
		iota:特殊的常量，可以被编译期修改
		没当我们定义一个 const， iota初始值为 0
		每当定义一个常量 则会自动累加1
		直到下一个const出现，从0再开始
	*/
	const (
		a = iota //0
		b = iota //1
		c = iota //2
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	//重新记为0
	const (
		d = iota //0
		e        //1,与上面一样是 iota
	)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println("----------------")
	const (
		A = iota  //0
		B         //1
		C         //2
		D = "abc" //abc   但是 iota 还是会累加 为 3
		E         //abc  iota = 4
		F = 100   //100    iota = 5
		G         // 100   iota = 6
		H = iota  //iota 为7了
		I         // 8
	)

	const (
		J = iota //遇到const从0开始   0
	)

	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(F)
	fmt.Println(G)
	fmt.Println(H)
	fmt.Println(I)
	fmt.Println(J)
}
