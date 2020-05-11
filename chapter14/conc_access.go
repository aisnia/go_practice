package main

import (
	"fmt"
	"strconv"
)

//person 类  chF用来存放 匿名函数的 通道
type Person struct {
	Name   string
	salary float64
	chF    chan func()
}

//当创建一个 person对象，会用一个协程 一直循环的执行，放在 chF中的函数
//改和读取salary的方法会通过将一个匿名函数写入chF通道中，然后让backend()按顺序执行以达到其目的。
func NewPerson(name string, salary float64) *Person {
	p := &Person{name, salary, make(chan func())}
	go p.backend()
	return p
}
func (p *Person) backend() {
	for f := range p.chF {
		f()
	}
}

// 通过匿名函数设置 salary
func (p *Person) SetSalary(sal float64) {
	p.chF <- func() { p.salary = sal }
}

// 通过匿名函数获取
func (p *Person) Salary() float64 {
	fChan := make(chan float64)
	p.chF <- func() { fChan <- p.salary }
	return <-fChan
}

func (p *Person) String() string {
	return "Person - name is: " + p.Name + " - salary is: " + strconv.FormatFloat(p.Salary(), 'f', 2, 64)
}

func main() {
	bs := NewPerson("Smith Bill", 2500.5)
	fmt.Println(bs)
	bs.SetSalary(4000.25)
	fmt.Println("Salary changed:")
	fmt.Println(bs)
}
