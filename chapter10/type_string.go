package main

import (
	"fmt"
	"strconv"
)

type T struct {
	a int
	b float32
	c string
}

//类似于Java的toString方法
func (this *T) String() string {
	return strconv.Itoa(this.a) +
		"/" + strconv.FormatFloat(float64(this.b), 'f', 6, 32) + "/" + this.c
}

func main() {
	t := &T{1, 1.1232141242131, "hello"}
	fmt.Println(t)
}
