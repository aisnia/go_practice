package main

import "fmt"

func main() {

	//切面动态的数组，长度可变，引用类型
	// len长度即元素个数，  cap容量
	var arr1 [6]int              //0,1,2,3,4,5   这是数组
	var slice1 []int = arr1[2:5] //定义切片，取数组 2到5 左闭右开  【2 3 4】 可以省略:前或后代表从0开始，  到结尾
	for i := 0; i < 6; i++ {
		arr1[i] = i
	}
	fmt.Println(slice1)
	fmt.Println("length:", len(slice1))
	fmt.Println("cap:", cap(slice1))

	//cap为5的切片
	var slice2 []int = make([]int, 5)
	for i := 0; i < len(slice2); i++ {
		slice2[i] = i * 5
	}
	fmt.Println(slice2)
	fmt.Println("length:", len(slice2))
	fmt.Println("cap:", cap(slice2))

	//var p *[]int = new([]int) // *p == nil; with len and cap 0 没有被初始化，只是分配了内存
	p1 := new([]int)
	fmt.Println(p1) //&[]
	//fmt.Println("length:",len(p1))
	//fmt.Println("cap:",cap(p1))
	p2 := make([]int, 0) //初始化了，指向空的数组
	fmt.Println(p2)
	fmt.Println("length:", len(p2))
	fmt.Println("cap:", cap(p2))

	var v []int = make([]int, 10, 50) //长度为10，容量为50的切片
	fmt.Println(v)
	fmt.Println("length:", len(v))
	fmt.Println("cap:", cap(v))

	//练习7.7求和
	arrF := []float32{2.2, 2.3, 2.4, 2.1}
	fmt.Println(p77(arrF))

	//7.8
	arrMax := []int{1, 3, 5, 7, 9}
	fmt.Println("max", max(arrMax))
	arrmin := []int{1, 3, 5, 7, 9}
	fmt.Println("min:", min(arrmin))

}

//求和
func p77(arrF []float32) (res float32) {
	for _, v := range arrF {
		res += v
	}
	return
}

func max(arr []int) (max int) {
	max = arr[0]
	for _, v := range arr {
		if max < v {
			max = v
		}
	}
	return
}

func min(arr []int) (min int) {
	min = arr[0]
	for _, v := range arr {
		if min > v {
			min = v
		}
	}
	return
}
