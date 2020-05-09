package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	inputFile := "input.txt"
	outputFile := "output.txt"

	//通过 io 工具 读写文件，返回的 buf是 []byte
	buf, err := ioutil.ReadFile(inputFile)

	if err != nil {
		fmt.Println("文件读取失败:", err)
	}

	fmt.Println(string(buf))

	//写入文件
	err = ioutil.WriteFile(outputFile, buf, 0644)

	if err != nil {
		panic(err.Error())
	}

	//带有缓冲的读写, 每次读取buffer大小的字节数， n为 读取的字节数
	//buffer := make([]byte,1024)
	//n ,e:= inputReader.Read(buffer)
	//
	//if e != nil {
	//
	//}

	//if(n == 0) {
	//
	//	return
	//}

}
