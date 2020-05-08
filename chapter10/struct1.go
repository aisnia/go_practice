package main

import (
	"fmt"
	"strings"
)

type struct1 struct {
	i1  int
	f1  float32
	str string
}

func main() {
	ms := new(struct1)
	ms.i1 = 10
	ms.f1 = 1.1
	ms.str = "3"
	//快速定义
	ms2 := struct1{10, 1.1, "123"}
	ms3 := struct1{i1: 10, f1: 1.1, str: "123"}
	fmt.Println(ms)
	fmt.Println(ms2)
	fmt.Println(ms3)
}
