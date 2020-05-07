package main

import "fmt"

/*
	go 中数组是一种值类型，可以使用new 来创建，

*/
func f(a [3]int) {
	fmt.Println(a)
}
func fp(a *[3]int) {
	fmt.Println(a)
}

func main() {
	var arr1 = new([5]int) // *[5]int 指针类型   引用传递
	var arr2 [5]int        //  [5]int值类型	值传递
	fmt.Printf("%T\n", arr1)
	fmt.Printf("%T\n", arr2)
	//传递的是引用指针
	deal1(arr1)
	fmt.Println(arr1)
	//传递的是拷贝的副本
	deal2(arr2)
	fmt.Println(arr2)

	var ar [3]int
	f(ar)   // passes a copy of ar
	fp(&ar) // passes a pointer to ar
}
func deal1(arr *[5]int) {
	arr[4] = 4
}
func deal2(arr [5]int) {
	arr[4] = 4
}
