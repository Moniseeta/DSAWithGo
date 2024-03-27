package main

import "fmt"

type Node struct {
	key   int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(i int) {
	if i < n.key {
		if n.Left == nil {
			n.Left = &Node{key: i}
		} else {
			n.Left.Insert(i)
		}
	} else if i > n.key {
		if n.Right == nil {
			n.Right = &Node{key: i}
		} else {
			n.Right.Insert(i)
		}
	}
}

func (n *Node) Search(key int) bool {
	if n.key == key {
		return true
	} else if key < n.key {
		if n.Left == nil {
			return false
		} else {
			n.Left.Search(key)
		}
	} else if key > n.key {
		if n.Right == nil {
			return false
		} else {
			n.Right.Search(key)
		}
	}
	return false
}

func minValueNode(node *Node) *Node {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}
func (n *Node) Delete(key int) {
	if n == nil {
		fmt.Println("key not found in tree")
		return
	}
	if n.key < key {
		if n.Right.key == key && n.Right.Right == nil && n.Right.Left == nil {
			n.Right = nil
			return
		}
		n.Right.Delete(key)
	} else if n.key > key {
		if n.Left.key == key && n.Left.Right == nil && n.Left.Left == nil {
			n.Left = nil
			return
		}
		n.Left.Delete(key)
	} else {
		if n.Left == nil {
			n.key = n.Right.key
			n.Left = n.Right.Left
			n.Right = n.Right.Right
			return
		} else if n.Right == nil {
			n.key = n.Left.key
			n.Left = n.Left.Left
			n.Right = n.Left.Right
			return
		}
		//We keep n.Left as is. Find the least value node in n.Right, since that will have value > n.left.key but less than
		// all values in n.Right, it will replace current n and we will delete that node from n.Right
		n.key = minValueNode(n.Right).key
		n.Right.Delete(n.key)
	}
}
func main() {
	tree := &Node{key: 100}
	tree.Insert(20)
	tree.Insert(200)
	tree.Insert(210)
	tree.Insert(120)
	tree.Insert(40)
	tree.Insert(26)
	tree.Insert(2)
	tree.Insert(99)
	tree.Insert(2)
	tree.Insert(207)

	fmt.Println(tree)

	fmt.Println(tree.Search(40))
	fmt.Println(tree.Search(30))

	tree.Delete(20)
	fmt.Println(tree.Search(20))

}
