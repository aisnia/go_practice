package main

var a = "G"

var b string

func main() {
	n()
	m()
	n()
	print("\n------------\n")

	b = "G" //我猜是 G，O，G, 对了，因为每个函数有自己的作用域，共享全局变量
	print(a)
	f1()
}

func f1() {
	b := "O"
	print(b)
	f2()
}
func f2() {
	print(b)
}

func n() {
	print(a)
}

func m() {
	a := "O"
	//a = "O"  //这样才是赋值， G O O
	print(a)
}

//我的解答    G，O，O
//实际是   G O G
//发现问题是   a:= "O"  这是个赋值语句，是一个局部变量
