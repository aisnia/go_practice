// json.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

//结构体类型
type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func main() {
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}
	// fmt.Printf("%v: \n", vc) // {Jan Kersschot [0x126d2b80 0x126d2be0] none}:
	// JSON format: 序列号
	js, _ := json.Marshal(vc)
	//直接打印
	fmt.Printf("JSON format: %s \n", js)
	// using an encoder:
	file, _ := os.OpenFile("vcard.json", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	//json的编码器
	enc := json.NewEncoder(file)
	//将对象编码成json格式，然后写入文件
	err := enc.Encode(vc)
	if err != nil {
		log.Println("Error in encoding json")
	}

	decode()
}

func decode() {
	b := []byte(`{"Name": "Wednesday", "Age": 6, "Parents": ["Gomez", "Morticia"]}`)
	var f interface{}
	err := json.Unmarshal(b, &f)
	if err != nil {
		fmt.Println("error")
		return
	}
	fmt.Println(f)
	//map[Name:Wednesday Age:6 Parents:[Gomez Morticia]]  是一个map类型
	//断言，装换
	m := f.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", v)
		case int:
			fmt.Println(k, "is int", v)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "don't kown type", v)
		}
	}
}
