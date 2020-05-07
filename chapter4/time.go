package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Now()
	fmt.Println(t)

	fmt.Printf("%d/%d/%d\n", t.Year(), t.Month(), t.Day()) //2020/5/6

	t = time.Now().UTC()
	fmt.Println(t)
	fmt.Println(time.Now())

	s := t.Format("20200512")
	fmt.Println(t, "->", s)

	//指针
	var i1 = 5
	fmt.Printf("值：%d,地址：%p\n", i1, &i1) //5
	//p是执行i1的指针
	var p *int = &i1
	*p = 100
	//一个指针变量可以指向任何一个值的内存地址
	fmt.Println(i1) //100

	//其他类型如 string，或对象也是一样的
	//指针的一个高级应用是你可以传递一个变量的引用（如函数的参数），这样不会传递变量的拷贝。
	//是引用传递

}
