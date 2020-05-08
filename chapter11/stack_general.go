package main

import "fmt"

type Stack struct {
	index int
	data  []interface{}
}

func (this *Stack) Len() int {
	return this.index
}
func (this *Stack) IsEmpty() bool {
	return this.index == 0
}
func (this *Stack) Push(x interface{}) {
	this.data[this.index] = x
	this.index++

}
func (this *Stack) Pop() (interface{}, error) {
	if this.index > 0 {
		this.index--
		return this.data[this.index], nil
	}
	return nil, nil
}
func main() {
	str := "1"
	i := 2
	f := 3.4

	stack := new(Stack)
	//new只是分配空间，注意要初始化
	stack.data = make([]interface{}, 3)
	stack.Push(str)
	stack.Push(i)
	stack.Push(f)
	fmt.Println(stack.Len())
	for !stack.IsEmpty() {
		fmt.Println(stack.Pop())
	}
}
