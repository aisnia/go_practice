package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	inputFile, _ := os.Open("goprogram")
	outputFile, _ := os.OpenFile("goprogramT", os.O_WRONLY|os.O_CREATE, 0666)
	defer inputFile.Close()
	defer outputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	outputWriter := bufio.NewWriter(outputFile)
	for {
		//每次读取一行
		inputString, _, readerError := inputReader.ReadLine()

		if readerError == io.EOF {
			fmt.Println("EOF")
			return
		}
		//第3到5的字节
		outputString := string(inputString[2:5]) + "\r\n"
		_, err := outputWriter.WriteString(outputString)
		//记得刷新缓冲
		outputWriter.Flush()
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	fmt.Println("Conversion done")
}
