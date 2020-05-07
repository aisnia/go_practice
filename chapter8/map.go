package main

import "fmt"

func main() {
	//声明
	var mapList map[string]int

	mapList = map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println(mapList["a"])
	fmt.Println(mapList["b"])
	fmt.Println(mapList["c"])

	mapCreate := make(map[string]float32)
	mapCreate["k1"] = 1.0
	mapCreate["k2"] = 2.0
	mapCreate["k3"] = 3.0

	//不要使用 new，永远用 make 来构造 map
	//newMap := new(map[string]int)
	//报错它是 *map[string]int的类型，不能直接使用
	//newMap["a"] = 1

	//判断某个 key是否存在
	if _, ok := mapCreate["k4"]; ok {
		//ok 为bool类型，即键是否存在
		fmt.Println("键存在")
	}

	//删除,某个键
	delete(mapCreate, "k1")

	//for range  遍历键与值，可以_省略某个
	for k, v := range mapCreate {
		fmt.Println("k : ", k, " v : ", v)
	}

	test()

}

func test() {
	// Version A:
	items := make([]map[int]int, 5)
	for i := range items {
		//对于切片的每一项，都初始化创建一个map，并1与2映射
		items[i] = make(map[int]int, 1)
		items[i][1] = 2
	}
	fmt.Printf("Version A: Value of items: %v\n", items)

	// Version B: NOT GOOD!
	items2 := make([]map[int]int, 5)
	for _, item := range items2 {
		//这里的切片，返回的是item是 一个拷贝的值而已，对原来的没有影响
		item = make(map[int]int, 1) // item is only a copy of the slice element.
		item[1] = 2                 // This 'item' will be lost on the next iteration.
	}
	fmt.Printf("Version B: Value of items: %v\n", items2)
}
