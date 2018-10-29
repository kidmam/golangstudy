package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(element interface{}) {
	*h = append(*h, element.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	element := old[n-1]
	*h = old[0 : n-1]
	return element
}

/* IntHeap 인터페이스 구현 여기에 추가 */

func main() {
	h := &IntHeap{2, 1, 7}
	heap.Init(h)
	fmt.Println(*h) // 1 2 7

	heap.Push(h, 4)
	heap.Push(h, 10)

	fmt.Println(*h) // 1 2 7 4 10

	// 1 2 4 7 10
	for h.Len() > 0 {
		m := heap.Pop(h)
		fmt.Print(m, " ")
	}
}
