package main

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

//Person -
type Person2 struct {
	Name string
	Age  int
}

func main() {
	person := Person2{"Alex", 10}
	pbytes, _ := xml.Marshal(person)
	buff := bytes.NewBuffer(pbytes)
	resp, err := http.Post("http://httpbin.org/post", "application/json", buff)

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
