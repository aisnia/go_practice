package main

import "fmt"

type Triangle struct {
	h float32
	d float32
}
type AreaInterface interface {
	Area() float32
}

func (this *Triangle) Area() float32 {
	return 0.5 * this.d * this.h
}

type Square struct {
	side float32
}
type PeriInterface interface {
	Perimeter() float32
}

func (this *Square) Perimeter() float32 {
	return 4 * this.side
}

func main() {
	s := Square{1.0}
	t := Triangle{2, 3}
	fmt.Println(s.Perimeter())

	fmt.Println(t.Area())
}
