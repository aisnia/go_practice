package main

import (
	"fmt"
	"math/rand"
	"time"
)

//类型别名
type TZ int

var tz TZ = 5

func main() {
	//bool类型 true,false 逻辑运算 && ||  初值为false
	//数字类型
	//1.整数   初值 0
	//int8(-128,127) int16 int32 int64 分别代表 8 16 32 64位的整数
	//无符号
	//uint8(0 , 256)，uint16, uint32, uint64

	//浮点类型 初值 0.0
	//float32，float64
	//complex64  32位的实数与虚数  complex128  64位

	//类型之间不能混合使用，不能隐身转换
	var a int
	var b int32
	a = 15
	//b = a + a  编译错误
	b = b + 5

	//但是可以显示转换

	b = int32(a)

	//格式化输出,和C类似
	fmt.Printf("%d ,   %d\n", a, b)

	//复数

	var c1 complex64 = 5 + 10i
	fmt.Printf("value: %v\n", c1)

	//随机数
	r()

	//字符 特殊的整形 uint8的, byte是uint8别名
	var ch byte = 65
	fmt.Printf("%c", ch) //A

}

func r() {
	for i := 1; i < 10; i++ {

		a := rand.Int()
		fmt.Printf("%d / ", a)
	}
	fmt.Println()
	for i := 0; i < 5; i++ {
		//0 - 左闭右开 伪随机
		r := rand.Intn(8)
		fmt.Printf("%d /", r)
	}

	fmt.Println()
	//随机种子，当前时间
	timens := int64(time.Now().Nanosecond())
	rand.Seed(timens)
	for i := 0; i < 10; i++ {
		fmt.Printf("%2.2f /", 100*rand.Float32())
	}
}
