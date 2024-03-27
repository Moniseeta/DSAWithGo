package main

import "fmt"

const ArrayLength = 7

// Hash table struct
type Hashtable struct {
	array [ArrayLength]*bucket
}

type bucket struct {
	head *node
}

type node struct {
	next *node
	key  string
}

// Insert
func (h *Hashtable) Insert(word string) {
	index := hash(word)
	h.array[index].insert(word)
}

// Search in Hashtable
func (h *Hashtable) Search(word string) bool {
	index := hash(word)
	return h.array[index].search(word)
}

// Delete from Hashtable
func (h *Hashtable) Delete(word string) {
	index := hash(word)
	h.array[index].delete(word)
}

func hash(word string) (h int) {
	sum := 0
	for _, char := range word {
		sum += int(char)
	}
	h = sum % ArrayLength
	return h
}

func Init() *Hashtable {
	t := &Hashtable{}
	for i := range t.array {
		t.array[i] = &bucket{}
	}
	return t
}

// insert - we make the key being inserted as the new head of the linked list.
// this is because we want most of the time complexities to be O(1) for a hashttable
func (b *bucket) insert(k string) {
	if b.search(k) {
		return
	}
	newNode := &node{
		key: k,
	}
	newNode.next = b.head
	b.head = newNode
}

// search return true if key is found
func (b *bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

// delete key
func (b *bucket) delete(k string) {
	if b.head == nil {
		fmt.Printf("%s not found- cant be deleted \n", k)
		return
	}
	if b.head.key == k {
		b.head = b.head.next
		return
	}
	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.next.key == k {
			previousNode.next = previousNode.next.next
		}
		previousNode = previousNode.next
	}
	fmt.Println("KEY not found- cant be deleted")
}

func main() {
	hashTable := Init()
	list := []string{
		"ERIC",
		"KENNY",
		"KYLE",
		"TOKEN",
		"MOMO",
		"APPLE",
	}
	for _, v := range list {
		hashTable.Insert(v)
	}
	fmt.Println(hashTable.Search("MOMO"))
	fmt.Println(hashTable.Search("Momo"))
	hashTable.Delete("Momo")
	hashTable.Delete("MOMO")
	fmt.Println(hashTable.Search("MOMO"))
	/*testBucket := &bucket{}
	testBucket.insert("Momo")
	fmt.Println(testBucket.search("Momo"))
	testBucket.delete("Momo")
	fmt.Println(testBucket.search("Momo"))*/

}
