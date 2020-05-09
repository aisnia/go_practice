package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//与标准输入os.Stdin 一样是 *os.File类型的，表示一个文件描述符， 打开input.txt 只读
	inputFile, inputError := os.Open("input.txt")

	if inputError != nil {
		fmt.Printf("读取文件失败")
		fmt.Println(inputError)
		return
	}
	//释放资源， 退出之前
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	for {
		//结束标记符为 /n 逐行的读取   注意windows 的 回车是 \r\n,但是不需要关心
		str, error := inputReader.ReadString('\n')
		fmt.Println("the input string is", str)
		//读取完成到文件末尾
		if error == io.EOF {
			return
		}
	}

}
