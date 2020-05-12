package main

import (
	"fmt"
	"os"
	"text/template"
)

type Person struct {
	Name string
	//	nonExportedAgeField string
	NonExportedAgeField string
}

func main() {
	//创建一个 template，名称是hello
	t := template.New("hello")
	//Parse 方法通过解析模板定义字符串，生成模板的内部表示。  {{.Name}} 结构体某个字段名，被替换
	//{{html .Name}} 或 {{.Name |html}} 可以进行过滤 即转义 例如把 > 替换为 &gt;
	//必须是导出的 才能解析用到模板
	t, _ = t.Parse("hello {{.Name |html}}!  Field {{.NonExportedAgeField }}")
	//Parse 方法通过解析模板定义字符串，生成模板的内部表示。
	//t.ParseFiles("字符串文件名")

	//验证模板格式 不对就好 panic
	//template.Must(t.Parse("hello {{.Name |html}}!  Field {{.NonExportedAgeField }}"))

	p := Person{Name: "<Mary>", NonExportedAgeField: "31  "}
	//Execute,将模板与数据结构整合，然后写入 第一个参数 io.Writer
	if err := t.Execute(os.Stdout, p); err != nil {
		fmt.Println("There was an error:", err.Error())
	}
}
