package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	myVersion := runtime.Version()
	fmt.Println("Version: ", myVersion)

	major := strings.Split(myVersion, ".")[0]
	fmt.Println("major Version: ", major)

	minor := strings.Split(myVersion, ".")[1]
	fmt.Println("minor minor: ", minor)

	m1, _ := strconv.Atoi(string(major))
	fmt.Println("m1: ", m1)

	m2, _ := strconv.Atoi(minor)
	fmt.Println("m2: ", m2)

	if m1 == 1 && m2 < 8 {
		fmt.Println("Need Go Version 1.8 or higher!")
		return
	}

	fmt.Println("You are using Go Version 1.8 or higher!")
}
