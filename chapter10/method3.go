package main

import (
	"fmt"
	"math"
)

//点类型
type point struct {
	x, y float64
}

//点 有的方法
func (p *point) Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

//内嵌Point，外层类型 继承 了内嵌的这些方法
type namedPoint struct {
	point //这是内嵌  namedPoint 可以直接使用 point的方法
	//p point //这是组合了，通过 n.p.fun()去使用
	//Point //多继承
	name string
}

//可以覆盖即 重写
func (n *namedPoint) Abs() float64 {
	return n.point.Abs() * 100.
}

func main() {
	n := &namedPoint{point{3, 4}, "Pythagoras"}
	//namedPoint 直接可以有 Abs方法
	fmt.Println(n.Abs()) // 打印5
}
