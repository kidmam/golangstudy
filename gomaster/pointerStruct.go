package main

import "fmt"

type myStructure struct {
	Name    string
	Surname string
	Height  int32
}

func createStructure(n, s string, h int32) *myStructure {
	if h > 300 {
		h = 0
	}
	return &myStructure{n, s, h}
}

func retStructure(n, s string, h int32) myStructure {
	if h > 300 {
		h = 0
	}
	return myStructure{n, s, h}
}

func main() {
	s1 := createStructure("길동", "홍", 123)
	s2 := retStructure("길동", "홍", 123)

	fmt.Println(s1)
	fmt.Println((*s1).Name)

	fmt.Println(s2)
	fmt.Println(s2.Name)

	pS := new(myStructure)
	fmt.Println(pS)

	sP := new([]myStructure)
	fmt.Println(sP)

}
