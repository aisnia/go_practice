package main

import "fmt"

type Two struct {
	a int
	b int
}

//属于 struct的发发发，其中前面的相当于 java里面的this
func (this *Two) AddThem() int {
	return this.a + this.b
}

//加上额外的数
func (this *Two) AddToParam(param int) int {
	return this.a + this.b + param
}

type IntArr []int

func (this IntArr) sum() (res int) {
	for _, v := range this {
		res += v
	}
	return
}

func main() {
	t := new(Two)
	t.a = 1
	t.b = 1
	fmt.Println(t.AddThem())

	fmt.Println(t.AddToParam(1))

	//非结构体也行,为类型 起别名
	arr := IntArr{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(arr.sum())
}
