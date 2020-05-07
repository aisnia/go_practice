package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	//目标字符串
	searchIn := "John: 2578.34 William: 4567.23 Steve: 5632.18"
	pat := "[0-9]+.[0-9]+" //正则

	f := func(s string) string {
		v, _ := strconv.ParseFloat(s, 32) //转为float32

		//返回格式化的 float   目标v，  f浮点, 2 保留两位数，   32位的float
		return strconv.FormatFloat(v*2, 'f', 2, 32)
	}

	//是否有匹配   目标字符串转为[]byte
	if ok, _ := regexp.Match(pat, []byte(searchIn)); ok {
		fmt.Println("Match Found!")
	}

	//编译？
	re, _ := regexp.Compile(pat)
	//将匹配到的部分替换为"##.#"
	str := re.ReplaceAllString(searchIn, "##.#")
	fmt.Println(str)
	//参数为函数时  所有匹配的部分通过函数进行运算后返回
	str2 := re.ReplaceAllStringFunc(searchIn, f)
	fmt.Println(str2)
}
