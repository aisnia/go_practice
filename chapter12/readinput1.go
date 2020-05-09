package main

import (
	"fmt"
)

var (
	firstName, lastName, s string
	i                      int
	f                      float32
	input                  = "56.12 / 5212 / Go"
	format                 = "%f / %d /%s"
)

func main() {
	fmt.Println("请输入你的名字")
	fmt.Scanln(&firstName, &lastName)
	//类似c的输入
	//fmt.Scanf("%s %s", &firstName, &lastName)
	fmt.Println("Hi", firstName, lastName)

	//Scanf 从标准输入读取
	//Sscanf 从字符串 input中读取
	fmt.Sscanf(input, format, &f, &i, &s)

	fmt.Println(f, i, s)

}
