package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*

文件 products.txt 的内容如下：

	"The ABC of Go";25.5;1500
	"Functional Programming with Go";56;280
	"Go for It";45.9;356
	"The Go Way";55;500

每行的第一个字段为 title，第二个字段为 price，第三个字段为 quantity。
内容的格式基本与 示例 12.3c 的相同，除了分隔符改成了分号。
请读取出文件的内容，创建一个结构用于存取一行的数据，然后使用结构的切片，并把数据打印出来。
*/

type Product struct {
	title    string
	price    float64
	quantity int
}

func main() {

	products := make([]Product, 0, 4)

	file, err := os.Open("products.txt")
	if err != nil {
		fmt.Println("文件无法打开", err)
	}
	reader := bufio.NewReader(file)
	file.Close()
	for {
		str, e := reader.ReadString('\n')

		strs := strings.Split(str, ";")
		title := strs[0]
		price, _ := strconv.ParseFloat(strs[1], 64)
		//出现了问题，str[2]后面带有了换行，无法转换，需要去除换行\n
		s := strings.Replace(strs[2], "\n", "", -1)
		q, _ := strconv.Atoi(s)
		fmt.Print(strs[2])
		produce := Product{title, price, q}

		products = append(products, produce)
		if e == io.EOF {
			break
		}
	}

	for i, v := range products {
		fmt.Printf("第%d个 title : %s, price : %.1f,quantity : %d\n", i+1, v.title, v.price, v.quantity)
	}
}
