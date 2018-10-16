package main

import (
	"encoding/xml"
	"fmt"
)

//Member -
type Member5 struct {
	Name   string
	Age    int
	Active bool
}

func main() {
	// 테스트용 데이타
	xmlBytes, _ := xml.Marshal(Member5{"Tim", 1, true})

	// XML 디코딩
	var mem Member5
	err := xml.Unmarshal(xmlBytes, &mem)
	if err != nil {
		panic(err)
	}

	// mem 구조체 필드 엑세스
	fmt.Println(mem.Name, mem.Age, mem.Active)
}
