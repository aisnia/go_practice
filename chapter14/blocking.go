package main

import (
	"fmt"
)

func f1(in chan int) {
	fmt.Println(<-in)
}

func main() {
	//out := make(chan int)
	//解决方法 可以使用有缓冲的通道，比如这里容量为1，当缓冲没满的时候 往里面写数据是不会阻塞的
	out := make(chan int, 1)
	//会产生死锁，因为 out队列是 无缓冲的 所有需要有消费者 消费才能继续执行，于是一直会阻塞在这
	out <- 2

	go f1(out)

}
