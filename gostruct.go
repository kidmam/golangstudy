package main

import "fmt"

type StructA struct {
	Name string
}

// 초기화 시에 하고 싶은 처리
func (self *StructA) SomeInitialize() {
	// Some initialize
}

// 생성자 함수를 정의
func NewStructA(name string) *StructA {
	structA := &StructA{Name: name}
	structA.SomeInitialize()
	return structA
}

func main() {
	//_ = NewStructA("name")

	structA := NewStructA("홍길동")

	fmt.Printf(structA.Name)
}
