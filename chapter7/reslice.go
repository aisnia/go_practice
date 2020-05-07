package main

import "fmt"

func main() {
	var ar = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var a = ar[5:7] //{5,6} - len(a) 2    cap(a)   5
	fmt.Println(a, " len", len(a), "cap", cap(a))
	a = a[0:4] //重新分片   5,6,7,8  数组还是同一个ar
	fmt.Println(a, " len", len(a), "cap", cap(a))

	//切片追加 复制
	Myappend()
}

func Myappend() {
	slFrom := []int{1, 2, 3}

	//创建大小为10的 []int切片
	slTo := make([]int, 10)

	//复制，返回拷贝元素的个数
	n := copy(slTo, slFrom)
	fmt.Println(slTo)
	fmt.Printf("Copied %d elements\n", n) // n == 3
	//append向切片后面追加,如果sl3的容量不足，则会分配新的切片
	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3)

	// 7.9 扩容
	fmt.Println(Expansion(slFrom, 2))

	//7.10 ,过滤选出偶数的
	fn := func(int2 int) bool {
		return int2%2 == 0
	}
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(Filter(s, fn))

	//7.11
	from := []int{1, 2, 3}
	to := []int{0, 4, 5, 6}
	fmt.Println(InsertStringSlice(from, 1, to))

	//7.12
	arr := []int{1, 2, 3, 5, 6, 7, 4}
	fmt.Println(RemoveStringSlice(arr, 3, 5))

}

func Expansion(s []int, factor int) (res []int) {
	res = make([]int, len(s)*factor)
	copy(res, s)
	return
}
func Filter(s []int, fn func(int2 int) bool) (res []int) {
	for _, v := range s {
		if fn(v) {
			res = append(res, v)
		}
	}
	return
}

func InsertStringSlice(from []int, index int, to []int) []int {
	/*//t := to[:index]  不能这样对原理的 to会有影响
	var  t []int = make([]int,0,len(from) + len(to))
	for i:=0;i < index;i++ {
		t = append(t, to[i])
	}
	for _,v:=range from {
		t = append(t, v)
	}
	for i:=index;i < len(to);i++ {
		t = append(t, to[i])
	}
	return t*/

	//更加简便的方式，多使用几个append
	return append(to[:index], append(from, to[index:]...)...)
}
func RemoveStringSlice(arr []int, start int, end int) []int {
	return append(arr[:start], arr[end+1:]...)
}
