package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	s, _ := in.ReadString('S')

	a, b, c := 0, 0, 0
	bytes := []byte(s)

	for _, v := range bytes {
		switch {
		case v == '\n':
			c++ //
			b++ //换行单词也+1
		case v == ' ':
			b++
			fallthrough
		case v != '\r':
			a++
		}
	}
	fmt.Printf("字符个数：%d, 单词个数%d, 行数 %d\n", a, b, c)

}
