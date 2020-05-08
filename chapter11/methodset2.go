package main

import (
	"fmt"
)

//list 切片
type List []int

//len方法
func (l List) Len() int {
	return len(l)
}

//追加
func (l *List) Append(val int) {
	*l = append(*l, val)
}

type Appender interface {
	Append(int)
}

//普通函数，添加start到end 传的都是接口
func CountInto(a Appender, start, end int) {
	for i := start; i <= end; i++ {
		a.Append(i)
	}
}

type Lener interface {
	Len() int
}

func LongEnough(l Lener) bool {
	return l.Len()*10 > 42
}

func main() {
	//值类型
	var lst List
	// 不能将 List 作为参数，传递给Appender   因为 List实现的Append 接受的参数是 l *List  指针 代表是*List实现了Append
	CountInto(&lst, 1, 10)
	if LongEnough(lst) { //Len 接受参数是 List 代表是 值类型List 实现了Len， 可以传递指针和值
		fmt.Printf("- lst is long enough\n")
	}

	// 指针类型
	plst := new(List)
	CountInto(plst, 1, 10)
	if LongEnough(plst) {
		// VALID: a *List can be dereferenced for the receiver
		fmt.Printf("- plst is long enough\n")
	}
}
