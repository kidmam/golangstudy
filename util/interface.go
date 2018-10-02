package util

import "math"

type Shape interface {
	Area() float64
	Perimeter() float64
}

//Rect 정의
type RectNew struct {
	Width, Height float64
}

//Circle 정의
type Circle struct {
	Radius float64
}

//Rect 타입에 대한 Shape 인터페이스 구현
func (r RectNew) Area() float64 { return r.Width * r.Height }
func (r RectNew) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

//Circle 타입에 대한 Shape 인터페이스 구현
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}
