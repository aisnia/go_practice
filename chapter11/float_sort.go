package main

import (
	"fmt"
	"math/rand"
	"time"
)

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

type Float64Array [25]float64

func (this *Float64Array) Len() int { return len(this) }
func (this *Float64Array) Less(i, j int) bool {
	return this[i] < this[j] || isNaN(this[i]) && !isNaN(this[j])
}
func (this *Float64Array) Swap(i, j int) { this[i], this[j] = this[j], this[i] }
func isNaN(f float64) bool {

	return f != f
}
func SortFloat64s(a [25]float64) {
	f := Float64Array(a)
	s := &f
	Sort(s)
}
func isSortFloat64s(a [25]float64) bool {
	f := Float64Array(a)
	s := &f
	return IsSorted(s)
}

func NewFloat64Array() [25]float64 {
	res := new([25]float64)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < cap(res); i++ {
		res[i] = rand.Float64() * 100
	}
	return *res
}

func main() {
	f := NewFloat64Array()
	fmt.Println(f)
	fmt.Println(isSortFloat64s(f))
	SortFloat64s(f)
	fmt.Println(f)
	fmt.Println(isSortFloat64s(f))

}
