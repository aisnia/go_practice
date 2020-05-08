package main

import (
	"os"
	"strconv"
)

//字符串接口
type Stringer interface {
	String() string
}

type Celsius float64

//温度。float64 保留1位小数
func (c Celsius) String() string {
	return strconv.FormatFloat(float64(c), 'f', 1, 64) + " °C"
}

//day
type Day int

var dayName = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

func (day Day) String() string {
	return dayName[day]
}

//打印
func print(args ...interface{}) {
	for i, arg := range args {
		if i > 0 {
			os.Stdout.WriteString(" ")
		} //空格隔开
		switch a := arg.(type) { // type switch 根据类型输出String
		case Stringer:
			os.Stdout.WriteString(a.String())
		case int:
			os.Stdout.WriteString(strconv.Itoa(a))
		case string:
			os.Stdout.WriteString(a)
		// more types
		default:
			os.Stdout.WriteString("???")
		}
	}
}

func main() {
	print(Day(1), "was", Celsius(18.36)) // Tuesday was 18.4 °C
}
