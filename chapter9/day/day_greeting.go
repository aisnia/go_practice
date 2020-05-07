package main

import (
	"fmt"
	"time"
)

func GreetingDay() {
	fmt.Println("Good Day")
}

func IsAfternoon() bool {
	t := time.Now()
	day := t.Hour()
	return day < 18
}
