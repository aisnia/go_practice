package main

import "fmt"

type Base struct {
	id int
}

func (this *Base) Id() int {
	return this.id
}

func (this *Base) setId(id int) {
	this.id = id
}

type person struct {
	Base      //继承关系
	FirstName string
	LastName  string
}

type Employee struct {
	person //继承
	salary float64
}

func main() {
	e := new(Employee)
	e.setId(1)
	fmt.Println(e.Id())
}
