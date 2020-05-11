package main

import "fmt"

func generate(ch chan int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

func filter1(in, out chan int, prime int) {
	for {

		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}
func main() {
	ch := make(chan int) // 输入通道
	go generate(ch)      // 一个协程生产数据 从2开始。。。。
	for {
		prime := <-ch //获取生产的数据 比如2
		fmt.Print(prime, " ")
		ch1 := make(chan int)      //输出通道 不能整除 prime放入，2的数
		go filter1(ch, ch1, prime) //过滤
		ch = ch1                   //输入通道又变成 ch1了即把整除prime的数据，剔除了， 进而在剔除后面的 比如 3整除的，剔除剩下的都是质数了，就输出 即prime
	}
}
