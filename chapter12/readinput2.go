package main

import (
	"bufio"
	"fmt"
	"os"
)

var inputReader *bufio.Reader
var input2 string
var err error

func main() {
	//将标准输入流  包一层缓存
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("输入")
	//遇到 换行停止  也可以是其他字符
	input2, err := inputReader.ReadString('\n')
	if err == nil {
		fmt.Println(input2)
	}

}
