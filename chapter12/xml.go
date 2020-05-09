// xml.go
package main

import (
	"encoding/xml"
	"fmt"
	"strings"
)

var t, token xml.Token
var err error

func main() {
	//xml字符串
	input := "<Person><FirstName>Laura</FirstName><LastName>Lynn</LastName></Person>"
	inputReader := strings.NewReader(input)

	//xml 解码器
	p := xml.NewDecoder(inputReader)

	for t, err = p.Token(); err == nil; t, err = p.Token() {
		switch token := t.(type) {
		case xml.StartElement: //开始
			name := token.Name.Local
			fmt.Printf("Token name: %s\n", name)
			for _, attr := range token.Attr {
				attrName := attr.Name.Local
				attrValue := attr.Value
				fmt.Printf("An attribute is: %s %s\n", attrName, attrValue)
				// ...
			}
		case xml.EndElement:
			fmt.Println("End of token")
		case xml.CharData: //文本
			content := string([]byte(token))
			fmt.Printf("This is the content: %v\n", content)
			// ...
		default:
			// ...
		}
	}
}
