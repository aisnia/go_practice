package main

import (
	"fmt"
	"reflect"
)

type T struct {
	A int
	B string
}

func main() {
	t := T{23, "skidoo"}
	//直接对结构体进行反射，传入指针，并设置可以修改
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	//NumField属性的个数
	for i := 0; i < s.NumField(); i++ {
		//获取对应的属性
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
	//修改
	s.Field(0).SetInt(77)
	s.Field(1).SetString("Sunset Strip")
	fmt.Println("t is now", t)
}
