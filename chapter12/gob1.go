package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type P struct {
	X, Y, Z int
	Name    string
}

type Q struct {
	X, Y *int32
	Name string
}

func main() {
	var network bytes.Buffer

	//编码器，用了写入，序列化我们的network 为 gob
	enc := gob.NewEncoder(&network)
	//解码器 相反
	dec := gob.NewDecoder(&network)

	err := enc.Encode(P{3, 4, 5, "P"})
	if err != nil {
		fmt.Println("error", err)
	}

	var q Q
	err = dec.Decode(&q)

	if err != nil {
		fmt.Println("decode error", err)
	}

	fmt.Printf("%q : {%d,%d}\n", q.Name, *q.X, *q.Y)

}
