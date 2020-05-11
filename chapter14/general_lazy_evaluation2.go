package main

import (
	"fmt"
)

type Any interface {
}

type Fi func(Any, Any) (Any, Any)

func main() {
	evenFunc := func(p1, p2 Any) (Any, Any) {
		return p2, p1.(int) + p2.(int)
	}

	even := BuildLazyIntEvaluator2(evenFunc, 0, 1)

	for i := 0; i < 10; i++ {
		fmt.Printf("%vth even: %v\n", i, even())
	}

}

func BuildLazyEvaluator2(fi Fi, p1, p2 Any) func() Any {
	ch := make(chan Any)

	loopFunc := func() {
		var t1 Any = p1
		var t2 Any = p2
		for {
			t1, t2 = fi(t1, t2)
			ch <- t1
		}
	}

	retFunc := func() Any {
		return <-ch
	}

	go loopFunc()

	return retFunc
}

func BuildLazyIntEvaluator2(evalFunc Fi, p1, p2 Any) func() int {
	ef := BuildLazyEvaluator2(evalFunc, p1, p2)
	return func() int {
		return ef().(int)
	}
}
