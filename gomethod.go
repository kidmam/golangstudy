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

// 포인터 Receiver
func (r *Rect) area2() int {
	r.width++
	return r.width * r.height
}

func main() {
	rect := Rect{10, 20}
	area := rect.area() //메서드 호출
	fmt.Println(area)

	area2 := rect.area2()      //메서드 호출
	println(rect.width, area2) // 11 220 출력
}
