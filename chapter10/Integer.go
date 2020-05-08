package main

import "fmt"

type Integer1 int

func (this Integer1) get() int {
	return int(this)
}

type Integer2 struct {
	n int
}

func (this Integer2) get() int {
	return this.n
}

func main() {
	var i1 Integer1
	fmt.Println(i1.get())

	var i2 Integer2
	fmt.Println(i2.get())

}
