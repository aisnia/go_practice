package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	CopyFile("b.txt", "a.txt")
	fmt.Println("Copy Done")
}
func CopyFile(target, source string) {
	s, err := os.Open(source)

	if err != nil {
		return
	}
	defer s.Close()

	t, err := os.Create(target)
	if err != nil {
		return
	}
	defer t.Close()

	io.Copy(t, s)
}
