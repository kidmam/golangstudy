package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Open("C:\\temp\\1.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	println(f.Name())
}
