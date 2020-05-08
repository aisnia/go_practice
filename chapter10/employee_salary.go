package main

import (
	"container/list"
	"fmt"
)

type employee struct {
	salary float64
}

func (this employee) giveRaise(raise float64) float64 {
	return this.salary * (raise + 1)
}

//类型和 在它上面定义的方法 必须在同一个包内,  基本类型不能定义方法的原因

//func (p *list.List) Iter() {
//	// ...
//}

//但是可以用别名的方式
type myList struct {
	list list.List
}

//遍历链表
func (this *myList) iter() {
	for p := this.list.Front(); p != nil; p = p.Next() {
		fmt.Println(p.Value)
	}
}
func main() {
	e := employee{200}
	fmt.Println(e.giveRaise(0.2))

	//lst := new(list.List)
	//for _ = range lst.Iter() {
	//}
}
