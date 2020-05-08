package main

import "fmt"

type inner struct {
	int1 int
	int2 int
}

type outer struct {
	b   int
	c   int
	int //匿名的 通过类型访问
	//int 不能有两个相同类型的
	inner //嵌套的结构体，一种继承
}
type A struct{ a int }
type B struct{ a, b int }

type C struct {
	A
	B
}

var c C //A，B中的 a 在同一层，命名冲突

type D struct {
	B
	b float32
}

var d D //D的b 在外层，会覆盖B的d

func main() {
	outer := new(outer)
	outer.b = 1
	outer.c = 2
	outer.int = 3
	outer.inner = inner{4, 5}

	fmt.Println(outer)

	a := A{1}
	b := B{2, 3}
	c = C{a, b}
	//fmt.Println(c.a) //报错
	d = D{b, 1.2}
	fmt.Println(d.b) //输出的是1.2 就是外层的  用于方法或字段的重载
}
