package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a int //8字节
	fmt.Println(unsafe.Sizeof(a))

}
