package main

import (
	"fmt"
	"time"
)

func GreetingNight() {
	fmt.Println("Good Night")
}

func IsEvening() bool {
	t := time.Now()
	h := t.Hour()
	return h > 18
}
