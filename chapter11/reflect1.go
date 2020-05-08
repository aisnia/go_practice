package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	//reflect.TypeOf(x) 获取x的类型
	fmt.Println("type:", reflect.TypeOf(x))
	//反射获取值
	v := reflect.ValueOf(x)
	fmt.Println("value:", v)
	//值v 有相关的方法, Type,Kind,Float
	fmt.Println("type:", v.Type())
	fmt.Println("kind:", v.Kind())
	fmt.Println("value:", v.Float())
	//接口值
	fmt.Println(v.Interface())
	fmt.Printf("value is %5.2e\n", v.Interface())
	//
	y := v.Interface().(float64)
	fmt.Println(y)
}
