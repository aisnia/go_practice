package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//文件名，  标志， 权限 读填0，而写权限填 0666
	//文件不存在则自动创建，  然后可以有写的权限
	outFile, err := os.OpenFile("output2.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("打开文件错误")
		return
	}

	defer outFile.Close()

	outputWriter := bufio.NewWriter(outFile)
	str := "hello,world\n"
	for i := 0; i < 10; i++ {
		outputWriter.WriteString(str)
	}

	outputWriter.Flush()

}
