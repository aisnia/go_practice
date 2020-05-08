package main

import "fmt"

type Simpler interface {
	Get() int
	Set(x int)
}

type Simple struct {
	x int
}

func (this *Simple) Get() int {
	return this.x
}
func (this *Simple) Set(x int) {
	this.x = x
}

func f(simpler Simpler) {
	simpler.Set(1)
	fmt.Println(simpler.Get())
}

type RSimple struct {
	x int
}

func (this *RSimple) Get() int {
	return this.x
}
func (this *RSimple) Set(x int) {
	this.x = x
}

func fi(simpler []Simpler) {
	for i, v := range simpler {
		switch v.(type) {
		case *Simple:
			fmt.Printf("The %d is *Simple", i)
		case *RSimple:
			fmt.Printf("The %d is *RSimple", i)
		default:
			fmt.Println("not type")
		}
	}
}

//11.9
func gI(simpler interface{}) bool {
	if _, ok := simpler.(Simple); ok {
		return ok
	}
	return false
}
func main() {
	s := new(Simple)
	f(s)

	r := new(RSimple)
	r.Set(2)
	simpers := []Simpler{s, r}
	for _, v := range simpers {
		fmt.Printf("%T\n", v)
	}
	fi(simpers)
	fmt.Println(gI(s))

	/*if v,ok := r.(Simpler) ;ok{
		fmt.Printf("v implements Get(): %d\n", v.Get()) // note: sv, not v
	}*/
}
