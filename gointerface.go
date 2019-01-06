package main

import (
	"fmt"
	. "github.com/kidmam/golangstudy/util"
)

func showArea(shapes ...Shape) {
	for _, s := range shapes {
		a := s.Area() //인터페이스 메서드 호출
		println(a)
	}
}

// 인터페이스 정의
type SomeBehivor interface {
	DoSomething(arg string) string
}

// 구조체 정의
type StructB struct {
}

// 인터페이스 구현
func (self *StructB) DoSomething(arg string) string {
	return arg
}

func main() {
	r := RectNew{10., 20.}
	c := Circle{10}

	showArea(r, c)

	// 인터페이스 타입 변수에 저장할 수 있다
	var behivor SomeBehivor
	behivor = &StructB{}
	fmt.Println(behivor.DoSomething("A"))
}
