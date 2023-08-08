package main

import (
	"fmt"
	"math/rand"
)

//==  Algorithm  ===========================================
type BinaryTree interface {
	Empty() bool
	Left() BinaryTree
	Right() BinaryTree
	Value() interface{}
}

func Traverse(t BinaryTree, visit func(r BinaryTree)) {
	var traverse func(t BinaryTree)
	traverse = func(t BinaryTree) {
		if !t.Empty() {
			traverse(t.Left())
			visit(t)
			traverse(t.Right())
		}
	}
	traverse(t)
}

//===  Adapter  ==========================================
type BTree BinTree

func (t BTree) Left() BinaryTree   { return BTree(t.root.lson) }
func (t BTree) Right() BinaryTree  { return BTree(t.root.rson) }
func (t BTree) Empty() bool        { return t.root == nil }
func (t BTree) Value() interface{} { return t.root.value }

//==  BinTree  ===========================================

type (
	Node struct {
		value interface{}
		lson  BinTree
		rson  BinTree
	}
	BinTree struct {
		root *Node
	}
)

func NewBinTree() BinTree {
	return BinTree{} // <==> return BinTree{root: nil}
}

func (t *BinTree) Insert(value interface{}) {
	if t.root == nil {
		t.root = &Node{value: value, lson: NewBinTree(), rson: NewBinTree()}
		return
	}
	if rand.Intn(2) == 0 {
		t.root.lson.Insert(value)
	} else {
		t.root.rson.Insert(value)
	}
}

//==  main  ===========================================
func main() {
	tree := NewBinTree()
	for _, val := range []int{34, 45, 36, 7, 24, 2, 40, 27, 5, 3} {
		tree.Insert(val)
	}
	fmt.Println()
	Traverse(BTree(tree), func(t BinaryTree) {
		fmt.Println(t.Value())
	})
	fmt.Println()
}
