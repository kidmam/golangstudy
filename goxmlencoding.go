package main

import (
	"encoding/xml"
	"fmt"
)

//Member -
type Member3 struct {
	Name   string
	Age    int
	Active bool
}

func main() {
	mem := Member3{"Alex", 10, true}

	xmlBytes, err := xml.Marshal(mem)
	if err != nil {
		panic(err)
	}

	xmlString := string(xmlBytes)
	fmt.Println(xmlString)
}
