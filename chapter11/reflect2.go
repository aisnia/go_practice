package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	v := reflect.ValueOf(x) //反射值
	// setting a value:
	// v.SetFloat(3.1415) // 不可设置
	fmt.Println("type of v:", v.Type())          //类型获取的是对应的 地址  float64
	fmt.Println("settability of v:", v.CanSet()) //CanSet判断是否设置
	v = reflect.ValueOf(&x)                      // 传递地址来 获取反射的值
	fmt.Println("type of v:", v.Type())          //类型获取的是对应的 地址  *float64
	fmt.Println("settability of v:", v.CanSet()) //
	v = v.Elem()                                 //间接的使用指针就可以设置了
	fmt.Println("The Elem of v is: ", v)
	fmt.Println("settability of v:", v.CanSet())
	v.SetFloat(3.1415) // this works!
	fmt.Println(v.Interface())
	fmt.Println(v)
}
