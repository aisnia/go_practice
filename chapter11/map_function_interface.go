package main

import "fmt"

type obj interface {
}

func main() {
	//做出对应的处理   参数是 空接口，可接受任意类型
	mf := func(o obj) obj {
		switch o.(type) {
		case int:
			return o.(int) * 2
		case string:
			return o.(string) + o.(string)
		}
		return o
	}

	isl := []obj{0, 1, 2, 3, 4, 5}

	//进行处理 传入我们定义的函数
	res1 := mapFunc(mf, isl)
	for _, v := range res1 {
		fmt.Println(v)
	}
	println()

	ssl := []obj{"0", "1", "2", "3", "4", "5"}
	res2 := mapFunc(mf, ssl)
	for _, v := range res2 {
		fmt.Println(v)
	}
}

//不定数量的 与切面差不多
func mapFunc2(mf func(obj) obj, list ...obj) []obj {
	result := make([]obj, len(list))

	for ix, v := range list {
		result[ix] = mf(v)
	}
	return result
}

func mapFunc(mf func(obj) obj, list []obj) []obj {
	result := make([]obj, len(list))

	for ix, v := range list {
		result[ix] = mf(v)
	}
	return result
}
