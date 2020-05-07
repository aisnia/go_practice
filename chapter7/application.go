package main

import "fmt"

func main() {
	//7.13  分割字符串
	fmt.Println(spilt("123456789", 5))
	//7.14	反转
	fmt.Println(reverse("123"))
	//7.15
	fmt.Println(deal("12345666"))
	//7.16	排序
	fmt.Println(sort([]int{1, 3, 5, 7, 9, 2, 4, 6, 8, 0}))
	//7.17	对切片每个元素 执行f(i)
	f := func(x int) int {
		return x * 10
	}
	fmt.Println(f17(f, []int{1, 2, 3, 4, 5, 6}))
}

func spilt(str string, i int) (res1, res2 string) {
	b := []byte(str)
	return string(b[:i]), string(b[i:])

}

func reverse(str string) string {
	b := []byte(str)
	for i, j := 0, len(b)-1; i < j; {
		b[i], b[j] = b[j], b[i]
		i++
		j--
	}
	return string(b)
}

func deal(str string) string {
	b := []byte(str)
	res := make([]byte, len(b))
	res[0] = b[0]
	for i := 1; i < len(b); i++ {
		if b[i-1] != b[i] {
			res[i] = b[i-1]
		} else {
			res[i] = res[i-1]
		}
	}
	return string(res)
}

func sort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func f17(f func(x int) int, arr []int) []int {
	for i, _ := range arr {
		arr[i] = f(arr[i])
	}
	return arr
}
