package main

import (
	"os"
	"text/template"
)

func main() {
	//运行 Execute 产生的结果来自模板的输出，它包含静态文本，以及被 {{}} 包裹的称之为管道的文本。
	tEmpty := template.New("template test")
	//我们可以对管道数据的输出结果用 if-else-end 设置条件约束：如果管道是空的 则不会输出 Will not print
	tEmpty = template.Must(tEmpty.Parse("Empty pipeline if demo: {{if ``}} Will not print. {{end}}\n")) //empty pipeline following if
	tEmpty.Execute(os.Stdout, nil)

	//管道不是空的  会输出
	tWithValue := template.New("template test")
	tWithValue = template.Must(tWithValue.Parse("Non empty pipeline if demo: {{if `anything`}} Will print. {{end}}\n")) //non empty pipeline following if condition
	tWithValue.Execute(os.Stdout, nil)

	tIfElse := template.New("template test")
	tIfElse = template.Must(tIfElse.Parse("if-else demo: {{if `anything`}} Print IF part. {{else}} Print ELSE part.{{end}}\n")) //non empty pipeline following if condition
	tIfElse.Execute(os.Stdout, nil)

	t := template.New("test")
	t, _ = t.Parse("{{with `hello`}}{{.}}{{end}}!\n")
	t.Execute(os.Stdout, nil)

	//点号（.）可以在 Go 模板中使用：其值 {{.}} 被设置为当前管道的值。
	//with 语句将点号设为管道的值。如果管道是空的，那么不管 with-end 块之间有什么，
	// 都会被忽略。在被嵌套时，点号根据最近的作用域取得值。
	t, _ = t.Parse("{{with `hello`}}{{.}} {{with `Mary`}}{{.}}{{end}}{{end}}!\n")
	t.Execute(os.Stdout, nil)

	//可以在模板内为管道设置本地变量，变量名以 $ 符号作为前缀。
	// 变量名只能包含字母、数字和下划线
	t = template.New("test")
	t = template.Must(t.Parse("{{with $3 := `hello`}}{{$3}}{{end}}!\n"))
	t.Execute(os.Stdout, nil)

	t = template.Must(t.Parse("{{with $x3 := `hola`}}{{$x3}}{{end}}!\n"))
	t.Execute(os.Stdout, nil)

	t = template.Must(t.Parse("{{with $x_1 := `hey`}}{{$x_1}} {{.}} {{$x_1}}{{end}}!\n"))
	t.Execute(os.Stdout, nil)

	//range-end 结构格式为：{{range pipeline}} T1 {{else}} T0 {{end}}。
	//
	//range 被用于在集合上迭代：管道的值必须是数组、切片或 map。
	// 如果管道的值长度为零，点号的值不受影响，且执行 T0；
	// 否则，点号被设置为数组、切片或 map 内元素的值，并执行 T1。
	t = template.Must(t.Parse("{{range .}}{{.}}{{end}}"))

	s := []int{1, 2, 3, 4}
	t.Execute(os.Stdout, s)

	//一些可以在模板代码中使用的预定义函数，例如 printf 函数工作方式类似于 fmt.Sprintf：
	t = template.New("test")
	t = template.Must(t.Parse("{{with $x := `hello`}}{{printf `%s %s` $x `Mary`}}{{end}}!\n"))
	t.Execute(os.Stdout, nil)

}
