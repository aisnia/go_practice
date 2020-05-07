package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//字符串,go使用的都是UTF-8编码标识Unicode文本
	var str string = "abc"
	//前缀
	fmt.Printf("%t\n", strings.HasPrefix(str, "ab")) //true

	//包含关系,子串
	//strings.Contains(s, substr string) bool

	fmt.Println(strings.Contains(str, "c")) //true

	//strings.Index(s, str string) int  出现的第一个索引  LastIndex最后一个从后面开始

	var s string = "hello,world"

	fmt.Printf("%d\n", strings.Index(s, "ll"))
	fmt.Printf("%d\n", strings.LastIndex(s, "or"))

	//字符串替换  strings.Replace(str, old, new, n) string
	//统计字符串出现次数  strings.Count(s, str string) int str在s出现的次数
	var manyG = "gggggggggg"

	fmt.Printf("%d\n", strings.Count(manyG, "g"))
	fmt.Printf("%s\n", strings.Replace(manyG, "gg", "ab", -1))

	// 重复字符串  strings.Repeat(s, count int) string 重复s count次
	fmt.Println(strings.Repeat("a", 5)) //aaaaa

	// 修改字符串大小写 strings.ToLower(s) string 小写，  strings.ToUpper(s) string 大写

	// 修剪字符串,去除两边的字符
	fmt.Println(strings.TrimSpace("   hi   "))  //hi
	fmt.Println(strings.Trim("abcabcab", "ab")) //cabc

	//分割字符串

	fmt.Println(strings.Split("hello world go", " "))

	//拼接 slice 到字符串  strings.Join(sl []string, sep string) string
	//向数组s1中添加 字符串

	//从字符串中读取内容  strings.NewReader(str) 用于生成一个 Reader 并读取字符串中的内容

	//字符串与其它类型的转换 strconv
	var x int
	x, _ = strconv.Atoi("12312")
	fmt.Println(x)

	x += 5
	fmt.Println(strconv.Itoa(x))
}
