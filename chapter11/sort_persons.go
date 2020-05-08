package main

import (
	"fmt"
	"strings"
)

type Person struct {
	firstName string
	lastName  string
}

func (this *Person) String() string {
	return this.firstName + " " + this.lastName
}

type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

//冒泡排序
func Sort(data Sorter) {
	for pass := 1; pass < data.Len(); pass++ {
		for i := 0; i < data.Len()-pass; i++ {
			if data.Less(i+1, i) {
				data.Swap(i, i+1)
			}
		}
	}
}

//是否排序了
func IsSorted(data Sorter) bool {
	n := data.Len()
	for i := n - 1; i > 0; i-- {
		if data.Less(i, i-1) {
			return false
		}
	}
	return true
}

type Persons struct {
	data []Person
}

func (this *Persons) Len() int { return len(this.data) }
func (this *Persons) Less(i, j int) bool {
	return strings.Compare(this.data[i].firstName+this.data[i].lastName, this.data[j].firstName+this.data[j].lastName) < 0
}
func (this *Persons) Swap(i, j int) {
	this.data[i], this.data[j] = this.data[j], this.data[i]
}
func (this *Persons) String() string {
	str := ""
	for _, v := range this.data {
		str += v.String()
	}
	return str
}
func SortPersons(a Persons) { Sort(&a) }
func main() {
	p := []Person{
		{"d", "e"},
		{"a", "b"},
		{"a", "c"},
		{"b", "c"},
	}
	persons := Persons{p}
	fmt.Println(persons)
	SortPersons(persons)
	fmt.Println(persons)

}
