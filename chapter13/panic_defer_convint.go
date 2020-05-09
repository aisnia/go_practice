package main

import (
	"fmt"
	"math"
)

func ConvertInt64ToInt(x int64) int {
	if x <= math.MaxInt32 && x >= math.MinInt32 {
		return int(x)
	}
	panic(fmt.Sprintf("%g is out of the int32 range", x))
}

func IntFromInt64(x int64) (i int, err error) {
	defer func() {
		if r := recover(); r != nil {
			//包装成error返回
			err = fmt.Errorf("%v", r)
		}
	}()
	res := ConvertInt64ToInt(x)
	return res, nil
}
func main() {
	if i, err := IntFromInt64(math.MaxInt32 + 1); err != nil {
		fmt.Println("Error :", err)
	} else {
		fmt.Printf("convert ok %d is %T", i, i)
	}

}
