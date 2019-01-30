package main

import (
	"fmt"
	"os"
)

func main() {
	readBytes(10)
}

func readBytes(amount int) {
	pwd, _ := os.Getwd()
	f, _ := os.Open(pwd + "/awesome.txt")

	defer fmt.Println("File closed")
	defer f.Close()
	defer fmt.Println("Going to close file")

	bytes := make([]byte, amount)
	count, _ := f.Read(bytes)

	fmt.Printf("%d bytes: '%s'\n", count, string(bytes))
}
