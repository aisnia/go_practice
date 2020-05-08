package main

import (
	"fmt"
	"time"
)

type Address struct {
	country  string
	province string
	city     string
}

type VCard struct {
	name      string
	address   *Address
	birthday  time.Time
	imagePath string
}

func main() {
	card := new(VCard)
	card.name = "xiaoqiang"
	card.address = &Address{country: "中国", province: "湖南省", city: "株洲市"}
	card.birthday = time.Now()
	card.imagePath = "/xiaoqiang.jpg"

	fmt.Println(card)
}
