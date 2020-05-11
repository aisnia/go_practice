package main

import (
	"fmt"
	"net/http"
)

type Hello struct{}

//自定义结构体，实现http.Request的方法
func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

func main() {
	var h Hello
	//默认是监听 / 所有请求
	http.ListenAndServe("localhost:4000", h)
}
