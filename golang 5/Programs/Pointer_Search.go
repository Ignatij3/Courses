package main

import (
	"fmt"
	"strconv"
)

type (
	Node struct {
		key  int
		lson *Node
		rson *Node
	}
	Tree struct {
		root *Node
	}
)

func InitTree() Tree {
	return Tree{root: nil}
}

func (t *Tree) Insert(data int) {
	t.root = t.insert(data, t.root)
}

func (t *Tree) insert(data int, node *Node) *Node {
	if node == nil {
		return &Node{key: data}
	}
	if data < node.key {
		node.lson = t.insert(data, node.lson)
	} else {
		// data >= node.key {
		node.rson = t.insert(data, node.rson)
	}
	return node
}

func (t *Tree) min(node *Node) *Node {
	for node.lson != nil && node.lson.key != node.key {node = node.lson}
	return node
}

func (t *Tree) Delete(data int) {
	t.root = t.delete(data, t.root)
}

func (t *Tree) delete(data int, node *Node) *Node {
	if node == nil {return nil}
	
	if data < node.key {
		node.lson = t.delete(data, node.lson)
	} else if data > node.key {
		node.rson = t.delete(data, node.rson)
	} else if node.lson != nil && node.rson != nil {
		min := t.min(node.rson)
		node.rson = t.delete(min.key, node.rson)
		min.lson, min.rson = node.lson, node.rson
		node = min
	} else if node.lson == nil {node = node.rson} else {node = node.lson}
	return node
}

func (tree Tree) TreeString() string {
	return tree.treeString("", true, "", tree.root)
}

func (tree Tree) treeString(prefix string, top bool, str string, node *Node) string {
	if node == nil {
		return ""
	}
	var temp string
	if node.rson != nil {
		if top {
			temp = prefix + "│   "
		} else {
			temp = prefix + "    "
		}
		str = tree.treeString(temp, false, str, node.rson)
	}
	str += prefix
	if top {
		str += "└──"
	} else {
		str += "┌──"
	}
	str += " " + strconv.Itoa(node.key) + "\n"
	if node.lson != nil {
		if top {
			temp = prefix + "    "
		} else {
			temp = prefix + "│   "
		}
		str = tree.treeString(temp, true, str, node.lson)
	}
	return str
}

func main() {
	//data := []int{34, 45, 36, 36, 35, 33, 36, 40, 42, 41, 7, 24, 2, 27, 5, 3}
	data := []int{1, 10, 6, 6, 6, 9, 5, 2}
	tree := InitTree()
	for _, key := range data {tree.Insert(key)}
	p := tree.root
	fmt.Println(tree.TreeString())
	fmt.Println(*p)
	//tree.Delete(36)
	tree.Delete(6)
	fmt.Println(*p)
	fmt.Println(tree.TreeString())
}
