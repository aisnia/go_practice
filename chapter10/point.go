package main

import (
	"fmt"
	"math"
)

type Point struct {
	X int
	Y int
}
type Polar struct {
	X int
	Y int
	Z int
}

func Abs(point *Point) float64 {
	return math.Sqrt(float64(point.X*point.X + point.Y*point.Y))

}
func main() {
	person := &Point{X: 3, Y: 4}
	fmt.Println(Abs(person))
}
