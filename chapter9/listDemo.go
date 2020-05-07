package main

import (
	"container/list"
	"fmt"
)

func main() {
	//返回的是 *List
	list := list.New()
	list.PushBack(101)
	list.PushBack(102)
	list.PushBack(103)

	for p := list.Front(); p != nil; p = p.Next() {
		fmt.Println(p.Value)
	}

}
