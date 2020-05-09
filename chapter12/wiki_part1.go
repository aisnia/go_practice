package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func (this *Page) save() {
	file, err := os.OpenFile(this.Title, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("创建文件失败")
		return
	}
	defer file.Close()

	file.Write(this.Body)
}

func (this *Page) load(title string) {
	ioutil.ReadFile(title)

	this.Title = title
	var err error
	this.Body, err = ioutil.ReadFile(title)
	if err != nil {
		fmt.Println("读取失败")
		return
	}
}

func main() {
	title := "a.txt"
	body := []byte("hello,World!")

	page := Page{title, body}

	page.save()

	newpage := new(Page)

	newpage.load(title)

	fmt.Println(newpage.Title)
	fmt.Println(string(newpage.Body))

}
