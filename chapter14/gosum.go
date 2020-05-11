package main

import "fmt"

func sum(x, y int, ch chan int) {
	ch <- x + y
}
func main() {
	ch := make(chan int)
	go sum(1, 2, ch)
	fmt.Println(<-ch)
}
