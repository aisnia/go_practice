package main

//10.14
import "fmt"

type Day int

var days = [7]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

const (
	MO Day = iota
	TU Day = iota
	WE Day = iota
	TH Day = iota
	FR Day = iota
	SA Day = iota
	SU Day = iota
)

func (this Day) String() string {
	return days[this]
}
func main() {
	fmt.Println(MO)

}
