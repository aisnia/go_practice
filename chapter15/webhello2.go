package main

import (
	"fmt"
	"net/http"
	"strings"
)

//编写一个网页服务器监听端口 9999，有如下处理函数：
//
//当请求 http://localhost:9999/hello/Name 时，响应：hello Name（Name 需是一个合法的姓，比如 Chris 或者 Madeleine）
//
//当请求 http://localhost:9999/shouthello/Name 时，响应：hello NAME

func myHandler(w http.ResponseWriter, req *http.Request) {
	s := req.URL.Path[1:]
	strs := strings.Split(s, "/")
	if len(strs) != 2 {
		fmt.Fprint(w, "Error")
		return
	}
	switch strs[0] {
	case "hello":
		fmt.Fprintf(w, "hello %s", strs[1])
	case "shouthello":
		fmt.Fprintf(w, "hello %s", strings.ToUpper(strs[1]))
	default:
		fmt.Fprint(w, "Error")
	}
}

func main() {

	http.Handle("/", http.HandlerFunc(myHandler))
	err := http.ListenAndServe("localhost:9999", nil)

	if err != nil {
		fmt.Printf("Error")
	}
}
