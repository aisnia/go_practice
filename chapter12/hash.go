package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io"
)

func sha() {
	//创建一个 hash.Hash对象
	hasher := sha1.New()

	//将数据添加到 散列中
	io.WriteString(hasher, "test")

	b := []byte{}
	//sum 向当前散列增加 b
	fmt.Printf("Result is %x\n", hasher.Sum(b))
	fmt.Printf("Result is %d\n", hasher.Sum(b))

	//将hash 重置为初始状态
	hasher.Reset()
	data := []byte("We shall overcome")
	//直接写入data
	n, err := hasher.Write(data)

	if n != len(data) || err != nil {
		fmt.Println("error")
	}

	checksum := hasher.Sum(b)

	fmt.Printf("result: %x\n", checksum)
}
func hash1() {
	hasher := md5.New()
	b := []byte{}
	io.WriteString(hasher, "test")
	fmt.Printf("Result is %x\n", hasher.Sum(b))
	fmt.Printf("Result is %d\n", hasher.Sum(b))
}
func main() {
	sha()
	hash1()
}
