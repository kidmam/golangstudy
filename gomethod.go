package main

import "fmt"

//Rect - struct 정의
type Rect struct {
	width, height int
}

//Rect의 area() 메소드
func (r Rect) area() int {
	return r.width * r.height
}

func main() {
	rect := Rect{10, 20}
	area := rect.area() //메서드 호출
	fmt.Println(area)
}
