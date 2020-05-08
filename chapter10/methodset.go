package main

import (
	"fmt"
)

type List []int

//传递的是值，拷贝  竟然都会改变，因为是切片传的是 引用
func (l List) change0() { l[0] = 3; return }
func (l List) Len() int { return len(l) }

//传递的是指针
func (l *List) Append(val int) { *l = append(*l, val) }

//指针 或 值都可以调用方法 原因是会自动转化  (&lst).Append(1)

type T struct {
	a int
}

//这个 clear没有用  说明还是值传递，是拷贝，  然后切片类型特殊
func (this T) clear1() {
	this.a = 0
}
func (this *T) clear2() {
	this.a = 0
}
func main() {
	// 值
	var lst List
	lst.Append(1)
	lst.change0()
	fmt.Printf("%v (len: %d)\n", lst, lst.Len()) // [3] (len: 1)

	// 指针
	plst := new(List)
	plst.Append(2)
	plst.change0()
	fmt.Printf("%v (len: %d)\n", plst, plst.Len()) // &[3] (len: 1)

	var t1 T
	t1.a = 1
	t1.clear1()
	fmt.Println(t1.a) //1
	t1.a = 1
	t1.clear2()
	fmt.Println(t1.a) //0

}
