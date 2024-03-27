package main

import (
	"fmt"
)

type MaxHeap struct {
	array []int
}

func (h *MaxHeap) swap(index1, index2 int) {
	h.array[index1], h.array[index2] = h.array[index2], h.array[index1]
}

func parentIndex(index int) (pIndex int) {
	return (index - 1) / 2
}

func rightChildIndex(index int) (rIndex int) {
	return (index * 2) + 2
}

func leftChildIndex(index int) (lIndex int) {
	return (index * 2) + 1
}

func (h *MaxHeap) heapifyUp(index int) {
	//l := len(h.array) -1
	if index != 0 && h.array[parentIndex(index)] < h.array[index] {
		h.swap(parentIndex(index), index)
		h.heapifyUp(parentIndex(index))
	}
}

func (h *MaxHeap) heapifyDown(index int) {
	// Concept is - when we extract the largest elem of a max heap, we move the last elem of array to root position
	// Then we compare it with the larger of its two children
	// To find which child it should be compared to - we check if it has children
	// If not - nothing needs to be done
	// If yes - we check if left child is lesser than or equal to the last index of array
	// If equal - it does not have a right child, also left child is leaf node - so compare and swap if necessary
	// If less - compare left and right child - the larger is the child that needs to be compared/swapped with current index
	// Once swapped - we know that the highest value is in root position - but we need to heapify the tree with the swapped child
	// So this loop continues till the leaf nodes are reached
	lastIndex := len(h.array) - 1
	childToCompare := 0
	l, r := leftChildIndex(index), rightChildIndex(index)

	for l <= lastIndex {
		if l == lastIndex {
			childToCompare = l
		} else if h.array[l] > h.array[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = leftChildIndex(index), rightChildIndex(index)
		} else {
			return
		}
	}
}

func (h *MaxHeap) Add(element int) {
	h.array = append(h.array, element)
	h.heapifyUp(len(h.array) - 1)
}

func (h *MaxHeap) GetLargest() (element int) {
	if len(h.array) == 0 {
		fmt.Println("Cannot extract from empty array")
	}
	element = h.array[0]
	indx := len(h.array) - 1
	h.array[0] = h.array[indx]
	h.array = h.array[:indx]
	h.heapifyDown(0)
	return element
}

func main() {

	arr := []int{10, 40, 50, 50, 40, 15, 100}

	m := &MaxHeap{}
	for _, v := range arr {
		m.Add(v)
	}
	fmt.Printf("Heapified : %v ", m.array)

	fmt.Printf("largest : %d", m.GetLargest())
	fmt.Printf("Heapified : %v ", m.array)
}
