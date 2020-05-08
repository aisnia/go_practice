package main

import (
	"fmt"
	"math"
)

//定义一个接口Shaper，只有一个Area方法
type Shaper interface {
	Area() float32
}

//具体的struct  Square
type Square struct {
	side float32
}

//Square 的一个方法，其中实现了 接口的Area方法
func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

//矩形
type Rectangle struct {
	length, width float32
}

//也实现了 Area方法
func (r Rectangle) Area() float32 {
	return r.length * r.width
}

//圆
type Circle struct {
	redio float32
}

func (this *Circle) Area() float32 {
	return math.Pi * this.redio * this.redio
}

func main() {
	sq1 := new(Square)
	sq1.side = 5

	//接口，可以直接用接口的引用指向 Square 然后调用它的Area方法
	var areaIntf Shaper
	areaIntf = sq1
	//其他赋值方式都可以
	//areaIntf := Shaper(sq1)
	//areaIntf := sq1
	fmt.Printf("The square has area: %f\n", areaIntf.Area())

	//注意，必须要实现接口的方法，才能调用
	//而且如果 Shaper 有另外一个方法 Perimeter()，
	//但是Square 没有实现它，即使没有人在 Square 实例上调用这个方法，编译器也会给出上面同样的错误。

	r := Rectangle{3, 5}
	q := &Square{5}

	var s Shaper
	c := &Circle{2.0}
	s = c
	fmt.Println(s.Area())

	//都是接口Shapes的实现是 Shapes类型,从而实现多态
	shapes := []Shaper{r, q, s}
	for i, v := range shapes {
		fmt.Println("Shaper is ", shapes[i], " Area is ", v.Area())
	}

}
