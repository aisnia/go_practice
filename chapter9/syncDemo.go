package main

import (
	"fmt"
	"math"
	"math/big"
	"sync"
)

type Info struct {
	mu sync.Mutex
	// ... other fields, e.g.: Str string
}

func Update(info *Info) {
	//上锁
	info.mu.Lock()

	//do somethings
	//解锁
	info.mu.Unlock()
}

func main() {
	//big.NewInt(1)   整数
	//big.NewFloat(1)  浮点数
	i := big.NewRat(1, 2)
	//分数 1/2

	fmt.Println(i)
	//大数
	im := big.NewInt(math.MaxInt64)
	in := im
	io := big.NewInt(1956)
	ip := big.NewInt(1)
	//就是 im = io + ip  ,三个操作数
	im.Add(io, ip)
	fmt.Printf("result is %v\n", im)
	ip.Mul(im, in).Add(ip, im).Div(ip, im)
	fmt.Printf("result is %v\n", ip)
}
