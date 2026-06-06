package bst

import "fmt"

type node struct {
	val   int
	left  *node
	right *node
}

type BST struct {
	root *node
}

func NewBST() *BST {
	return &BST{}
}

func (b *BST) Insert(val int) {
	b.root = insert(b.root, val)
}

func insert(n *node, val int) *node {
	if n == nil {
		fmt.Printf("inserted value %d\n", val)
		return &node{
			val:   val,
			left:  nil,
			right: nil,
		}
	}
	if val == n.val {
		fmt.Println("duplicate val")
	} else if val < n.val {
		n.left = insert(n.left, val)
	} else if val > n.val {
		n.right = insert(n.right, val)
	}
	return n
}

func (b *BST) Search(val int) bool {
	return b.root.search(val)
}

func (n *node) search(val int) bool {
	if n == nil {
		return false
	}

	if val == n.val {
		return true
	}

	if val < n.val {
		return n.left.search(val)
	} else {
		return n.right.search(val)
	}
}

func (b *BST) Inorder() {
	b.root.inorder()
}

func (n *node) inorder() {
	if n == nil {
		return
	}
	n.left.inorder()
	fmt.Printf("%v ", n.val)
	n.right.inorder()
}

func (b *BST) Postorder() {
	b.root.postorder()
}

func (n *node) postorder() {
	if n == nil {
		return
	}
	fmt.Printf("%v ", n.val)
	n.left.postorder()
	n.right.postorder()
}

func (b *BST) Preorder() {
	b.root.preorder()
}

func (n *node) preorder() {
	if n == nil {
		return
	}
	n.left.preorder()
	n.right.preorder()
	fmt.Printf("%v ", n.val)

}
