package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

func (s *Stack) Pop() (i int) {
	l := len(s.items) - 1
	i = s.items[l]
	s.items = s.items[:l]
	return i
}

func main() {
	s := &Stack{}
	s.Push(100)
	s.Push(200)
	s.Push(300)
	fmt.Println(s.items)
	s.Pop()
	fmt.Println(s.items)
}
