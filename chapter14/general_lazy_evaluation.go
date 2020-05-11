package main

import "fmt"

//练习14.12 通过使用14.12中工厂函数生成前10个斐波那契数

type Any interface {
}

//一个函数类型
type EvalFunc func(Any) (Any, Any)

func main() {
	//事件函数
	evenFunc := func(state Any) (Any, Any) {
		os := state.(int)
		ns := os + 2
		return os, ns
	}

	//懒加载，从 0 开始
	even := BuildLazyIntEvaluator(evenFunc, 0)

	for i := 0; i < 10; i++ {
		fmt.Printf("%vth even: %v\n", i, even())
	}
}

func BuildLazyEvaluator(evalFunc EvalFunc, initState Any) func() Any {
	//数据，放在通道里
	retValChan := make(chan Any)

	//异步执行的协程
	loopFunc := func() {
		var actState Any = initState

		var retVal Any

		for {
			//数据，进行处理，然后传入通道， 这是没有缓冲的，所以没消费的话，会被阻塞，来实现懒加载
			retVal, actState = evalFunc(actState)
			retValChan <- retVal
		}
	}

	//从通道获取数据的 函数
	retFunc := func() Any {
		return <-retValChan
	}

	//协程异步执行
	go loopFunc()

	return retFunc
}

func BuildLazyIntEvaluator(evalFunc EvalFunc, initState Any) func() int {
	ef := BuildLazyEvaluator(evalFunc, initState)
	return func() int {
		return ef().(int)
	}
}
