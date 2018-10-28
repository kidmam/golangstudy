package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		panic("에러: 2개 미만의 argument")
	}

	programName := os.Args[0:1]
	firstArg := os.Args[1:2]
	secondArg := os.Args[2:3]
	allArgs := os.Args[1:]

	fmt.Println(programName, firstArg, secondArg, allArgs)
}
