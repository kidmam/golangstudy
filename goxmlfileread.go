package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

//Member -
type Member6 struct {
	Name   string
	Age    int
	Active bool
}

//Members -
type Members struct {
	Member []Member6
}

func main() {
	// xml 파일 오픈
	fp, err := os.Open("/Users/netive/GoProject/src/github.com/kidmam/golangstudy/test.xml")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	// xml 파일 읽기
	data, err := ioutil.ReadAll(fp)

	// xml 디코딩
	var members Members
	xmlerr := xml.Unmarshal(data, &members)
	if xmlerr != nil {
		panic(xmlerr)
	}

	fmt.Println(members)
}
