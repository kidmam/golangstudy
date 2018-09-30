package main

import "fmt"

type person struct {
	name string
	age  int
}

type dict struct {
	data map[int]string
}

func newDict() *dict {
	d := dict{}
	d.data = map[int]string{}
	return &d
}

func main() {
	p := person{}
	p.name = "Lee"
	p.age = 10

	var p1 person
	p1 = person{"Bob", 20}
	p2 := person{name: "Sean", age: 50}

	fmt.Println(p)
	fmt.Println(p1)
	fmt.Println(p2)

	p3 := new(person)
	p3.name = "Lee" // p가 포인터라도 . 을 사용한다
	fmt.Println(*p3)

	dic := newDict() // 생성자 호출
	dic.data[1] = "A"
	fmt.Println(dic.data[1])

}
