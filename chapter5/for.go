package main

import (
	"fmt"
)

func main() {
	//1 到 5
	for i := 1; i < 5; i++ {
		fmt.Printf("%d \t", i)
	}

	str := "abc"
	fmt.Println(len(str))

	for i := 0; i < len(str); i++ {
		fmt.Printf("index : %d, char : %c\t", i, str[i])
	}
	fmt.Println()
	name := "小强"
	fmt.Println(len(name))

	// loop.go
	for i := 1; i <= 15; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	i := 1
	//只使用 goto
loop:
	fmt.Printf("%d ", i)
	if i < 15 {
		i++
		goto loop
	}
	fmt.Println("-------嵌套--------")

	//嵌套
	for i := 1; i <= 6; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("G")
		}
		fmt.Println()
	}
	fmt.Println("--------拼接-------")
	s := ""
	for i := 1; i <= 6; i++ {
		s += "G"
		fmt.Println(s)
	}
	fmt.Println("--------按位补码-------")
	for i := 0; i <= 10; i++ {
		fmt.Printf("i : %d, bit : %b\n", i, ^i)

	}

	fmt.Println("------Fizz-Buzz-----------")
	for i := 1; i <= 100; i++ {
		switch {
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		case i%15 == 0:
			fmt.Println("FizzBuzz")
		default:
			fmt.Println(i)

		}
	}
	fmt.Println("--------使用 * 符号打印宽为 20，高为 10 的矩形。 -----")
	for i := 0; i < 10; i++ {
		for j := 0; j < 20; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	s = "abcdef"
	for pos, char := range s {
		fmt.Printf("%d,%c\n", pos, char)
	}

	//break,continue 跳出，继续 java类似

	//goto 跳到标签处
	i = 0
flag:
	fmt.Print(i)
	i++
	if i < 5 {
		goto flag
	}

}
