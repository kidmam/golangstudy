package main

import (
	"encoding/json"
	"fmt"
)

//Member -
type Member struct {
	Name   string
	Age    int
	Active bool
}

func main() {

	// Go 데이타
	mem := Member{"Alex", 10, true}

	// JSON 인코딩
	jsonBytes, err := json.Marshal(mem)
	if err != nil {
		panic(err)
	}

	// JSON 바이트를 문자열로 변경
	jsonString := string(jsonBytes)

	fmt.Println(jsonString)
}
