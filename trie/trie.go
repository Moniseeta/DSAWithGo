package main

import "fmt"

// AlphabetSize is the number of possible chars in a trie
const AlphabetSize = 26

// Node struct
type Node struct {
	children [26]*Node
	isEnd    bool
}

// Trie struct
type Trie struct {
	root *Node
}

// InitTrie is for initializing the Trie
func InitTrie() *Trie {
	result := &Trie{root: &Node{}}
	return result
}

// Insert is to add a word to Trie
func (t *Trie) Insert(w string) {
	root := t.root
	for _, val := range w {
		idx := int(val) - 'a'
		if root.children[idx] == nil {
			root.children[idx] = &Node{}
		}
		root = root.children[idx]
	}
	root.isEnd = true
}

// Search is to search a word to Trie
func (t *Trie) Search(w string) bool {
	root := t.root
	for _, val := range w {
		idx := int(val) - 'a'
		if root.children[idx] == nil {
			return false
		}
		root = root.children[idx]
	}
	if root.isEnd == true {
		return true
	}
	return false
}

func main() {
	t := InitTrie()
	words := []string{
		"ape",
		"apple",
		"apply",
		"application",
		"bat",
		"batman",
		"bold",
	}

	for _, word := range words {
		t.Insert(word)
	}
	fmt.Println(t.Search("application"))
	fmt.Println(t.Search("applicant"))
	fmt.Println(t.Search("bald"))
}
