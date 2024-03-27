package main

import "fmt"

type Queue struct {
	items []int
}

func (s *Queue) EnQueue(i int) {
	s.items = append(s.items, i)
}

func (s *Queue) DeQueue() (i int) {
	i = s.items[0]
	s.items = s.items[1:]
	return i
}

func main() {
	s := &Queue{}
	s.EnQueue(100)
	s.EnQueue(200)
	s.EnQueue(300)
	fmt.Println(s.items)
	s.DeQueue()
	fmt.Println(s.items)
}
