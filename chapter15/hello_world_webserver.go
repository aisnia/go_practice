package main

import (
	"fmt"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside HelleServer handler")
	//获取url中的参数 如 localhost:8080/world 则 URL.Path就是 /world 然后取出 1后面的
	//Fprintf 向 w Writer中 写入数据  可以写html
	fmt.Fprintf(w, "<h1>hello,%s<h1>", req.URL.Path[1:])
	//fmt.Fprintf(w, "Hello,"+req.URL.Path[1:])

}

func main() {
	http.HandleFunc("/", HelloServer)
	//这种方式也行，将我们定义的方法，转为 http.HandlerFunc的方法
	//http.Handle("/",http.HandlerFunc(HelloServer))
	//监听地址
	err := http.ListenAndServe("localhost:8080", nil)

	if err != nil {
		log.Fatal("ListenAndServer: ", err.Error())
	}
}
