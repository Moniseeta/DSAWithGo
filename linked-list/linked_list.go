package main

import "fmt"

type linkedList struct {
	head   *node
	length int
}

type node struct {
	data int
	next *node
}

func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l linkedList) printList() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Println("")
}

func (l *linkedList) deleteWithValue(val int) {
	if l.length == 0 {
		return
	}

	if l.head.data == val {
		l.head = l.head.next
		l.length--
		return
	}

	prevToDelete := l.head
	for prevToDelete.next.data != val {
		prevToDelete = prevToDelete.next
		if prevToDelete.next.next == nil {
			return
		}
	}
	prevToDelete.next = prevToDelete.next.next
	l.length--
}
func main() {
	myList := linkedList{}
	node1 := &node{data: 48}
	node2 := &node{data: 14}
	node3 := &node{data: 16}
	node4 := &node{data: 43}
	node5 := &node{data: 17}
	node6 := &node{data: 19}
	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)
	myList.prepend(node5)
	myList.prepend(node6)
	fmt.Println(myList)

	myList.printList()
	myList.deleteWithValue(17)
	myList.printList()
}
