package main

import (
	"errors"
	"fmt"
)

//创建一个错误类型
var errNotFound error = errors.New("Not found error")

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		//当场窗口返回
		return 0, errors.New("math - square root of negative number")
	}
	//fmt.Errorf的方式也行
	if f < 0 {
		return 0, fmt.Errorf("math: square root of negative number %g", f)
	}
	return f, nil
}

func main() {
	fmt.Printf("error : %v\n", errNotFound)

	if f, err := Sqrt(-1); err != nil {
		fmt.Println("Error : %s \n", err)
	} else {
		fmt.Println(f)
	}
}
