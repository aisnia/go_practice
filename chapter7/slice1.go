package main

import "fmt"

func main() {
	//数组
	var arr1 [5]int
	//下标方式遍历
	for i := 1; i < len(arr1); i++ {
		arr1[i] = i * 2
	}
	//range遍历
	for i, v := range arr1 {
		fmt.Printf("index is %d, value is %d\n", i, v)
	}
	fmt.Println("-----------------")
	// [...]任意长度
	a := [...]string{"a", "b", "c", "d"}
	fmt.Println(len(a))
	//v可以省略不写 或者用 _
	for i, v := range a {
		fmt.Println("Array item", i, "is", a[i], "=", v)
	}

}
