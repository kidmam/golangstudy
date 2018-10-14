package main

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

//Person -
type Person3 struct {
	Name string
	Age  int
}

func main() {
	person := Person3{"Alex", 10}
	pbytes, _ := xml.Marshal(person)
	buff := bytes.NewBuffer(pbytes)

	// Request 객체 생성
	req, err := http.NewRequest("POST", "http://httpbin.org/post", buff)
	if err != nil {
		panic(err)
	}

	//Content-Type 헤더 추가
	req.Header.Add("Content-Type", "application/xml")

	// Client객체에서 Request 실행
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Response 체크.
	respBody, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		str := string(respBody)
		println(str)
	}
}
