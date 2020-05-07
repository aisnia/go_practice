package main

import (
	"fmt"
	"sort"
)

var (
	barVal = map[string]int{
		"alpha": 34, "bravo": 56, "charlie": 23,
		"delta": 87, "echo": 56, "foxtrot": 12,
		"golf": 34, "hotel": 16, "indio": 87,
		"juliet": 65, "kili": 43, "lima": 98}
)

func main() {
	sortMap()
	reserve()
}

func sortMap() {
	//没排序
	fmt.Println("unsorted:")
	for k, v := range barVal {
		fmt.Printf("Key: %v, Value: %v / ", k, v)
	}
	//创建一个切片用来保存 键
	keys := make([]string, len(barVal))
	i := 0
	for k, _ := range barVal {
		keys[i] = k
		i++
	}
	//对于键进行排序
	sort.Strings(keys)

	fmt.Println()
	fmt.Println("sorted:")
	//然后遍历键，获取值
	for _, k := range keys {
		fmt.Printf("Key: %v, Value: %v / ", k, barVal[k])
	}
}

func reserve() {
	res := make(map[int]string, len(barVal))

	for k, v := range barVal {
		res[v] = k
	}
}
