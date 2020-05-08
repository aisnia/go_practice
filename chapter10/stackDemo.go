package main

import (
	"fmt"
	"strconv"
)

const LIMIT = 4

type Stack [LIMIT]int

//入栈，找到第一个0，非空的
func (stack *Stack) Push(n int) {
	for i, v := range stack {
		if v == 0 {
			stack[i] = n
			break
		}
	}
}

//出栈先进后出，从后面找
func (stack *Stack) Pop() int {
	v := 0
	for i := len(stack) - 1; i >= 0; i-- {
		if v = stack[i]; v != 0 {
			stack[i] = 0
			return v
		}
	}
	return 0
}
func (stack *Stack) String() string {
	str := ""
	for _, v := range stack {
		str += strconv.Itoa(v) + " "
	}
	return str
}

type Stack2 struct {
	index int
	arr   [LIMIT]int
}

func (stack *Stack2) Push(x int) {
	if stack.index < LIMIT {
		stack.arr[stack.index] = x
		stack.index++
	}
}

func (stack *Stack2) Pop() int {
	if stack.index > 0 {
		stack.index--
		res := stack.arr[stack.index]
		stack.arr[stack.index] = 0
		return res
	}
	return 0
}

func (stack *Stack2) String() string {
	str := ""
	for _, v := range stack.arr {
		str += strconv.Itoa(v) + " "
	}
	return "index :" + strconv.Itoa(stack.index) + "value :" + str
}

func main() {
	stack := new(Stack)

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	fmt.Println(stack)

	stack.Pop()
	fmt.Println(stack)

	stack2 := new(Stack2)
	stack2.Push(1)
	stack2.Push(2)
	stack2.Push(3)
	stack2.Push(4)
	fmt.Println(stack2)

	stack2.Pop()
	fmt.Println(stack2)

}
