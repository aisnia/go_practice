package main

import (
	"fmt"
	"reflect"
)

type TagType struct { // tags 标签，类似于Java中的注解
	field1 bool   "An important answer"
	field2 string "The name of the thing"
	field3 int    "How much there are"
}

func main() {
	tt := TagType{true, "Barak Obama", 1}
	for i := 0; i < 3; i++ {
		refTag(tt, i)
	}
}

func refTag(tt TagType, ix int) {
	//反射
	ttType := reflect.TypeOf(tt)
	//获取对应的属性
	ixField := ttType.Field(ix)
	//打印属性的的标签
	fmt.Printf("%v\n", ixField.Tag)
	fmt.Printf("%v\n", ixField)
}
