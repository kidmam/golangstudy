package main

import "fmt"

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func traverse(t *Tree) {
	if t == nil {
		return
	}
	traverse(t.Left)
	fmt.Println(t.Value, " ")
	traverse(t.Right)
}

func main() {

}
