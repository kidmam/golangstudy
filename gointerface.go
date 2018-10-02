package main

import (
	. "github.com/kidmam/golangstudy/util"
)

func showArea(shapes ...Shape) {
	for _, s := range shapes {
		a := s.Area() //인터페이스 메서드 호출
		println(a)
	}
}

func main() {
	r := RectNew{10., 20.}
	c := Circle{10}

	showArea(r, c)
}
