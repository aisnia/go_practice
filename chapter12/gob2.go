package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

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

var content string

func encode() {

	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}
	// fmt.Printf("%v: \n", vc) // {Jan Kersschot [0x126d2b80 0x126d2be0] none}:
	// using an encoder:
	file, _ := os.OpenFile("vcard.gob", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	enc := gob.NewEncoder(file)
	err := enc.Encode(vc)
	if err != nil {
		log.Println("Error in encoding gob")
	}

}

//解码
func decode() {
	file, _ := os.OpenFile("vcard.gob", os.O_RDONLY, 0)
	dec := gob.NewDecoder(file)
	vc := new(VCard)
	err := dec.Decode(vc)
	fmt.Println(vc)
	if err != nil {
		fmt.Println("error")
	}
}

func main() {
	//encode() 编码
	decode()
}
