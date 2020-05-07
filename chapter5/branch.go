package main

import (
	"fmt"
	"strconv"
)

func main() {

	//if
	a := 10
	if a < 20 {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

	//初始值，只属于 if else条件语句里面
	if x := 10; x%2 == 0 {
		fmt.Println("ok")
	} else {
		fmt.Println("no")
	}
	//fmt.Println(x)  访问不到

	if s, err := strconv.Atoi("123"); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}

	//switch 变量 var1 可以是任何类型
	//每个case符合条件后，不用去执行下一个分支的代码默认break了。，可以加fallthrough,让他可以继续判断
	var num1 int = 100

	switch num1 {
	case 98, 99:
		fmt.Println("98或99")
	case 100:
		fmt.Println("100")
	default:
		fmt.Println("其他")
	}

	k := 6
	switch k {
	case 4:
		fmt.Println("was <= 4")
		fallthrough
	case 5:
		fmt.Println("was <= 5")
		fallthrough
	case 6:
		fmt.Println("was <= 6")
		fallthrough
	case 7:
		fmt.Println("was <= 7")
		fallthrough
	case 8:
		fmt.Println("was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}
	//6,7,8, default都会执行
}
