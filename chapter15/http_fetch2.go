package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

//练习 15.4：扩展 http_fetch.go 使之可以从控制台读取url，使用 12.1节 学到的接收控制台输入的方法（http_fetch2.go）

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	str, err := inputReader.ReadString('\n')

	checkError(err)
	//除去 /n
	str = str[:len(str)-1]

	req, err := http.Get(str)
	checkError(err)

	data, err := ioutil.ReadAll(req.Body)

	checkError(err)
	fmt.Printf("Got: %q", string(data))

}

func checkError(err error) {
	if err != nil {
		log.Fatal("Error", err)
	}
}
