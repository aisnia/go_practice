package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"os"
)

func main() {
	fileName := "MyFile.gz"
	var r *bufio.Reader
	fi, err := os.Open(fileName)
	if err != nil {
		fmt.Println(os.Stderr, "%v,Can't %s: error: %s\n", os.Args[0], fileName, err)
		os.Exit(0)
	}

	defer fi.Close()

	//通过 gzip 来读取 压缩文件
	fz, err := gzip.NewReader(fi)
	if err != nil {
		r = bufio.NewReader(fi)
	} else {
		r = bufio.NewReader(fz)
	}

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			fmt.Println("已读完文件")
			os.Exit(0)
		}
		fmt.Println(line)
	}

}
