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

func (t Tree) Find(data int) bool {
	return t.find(data, t.root)
}

func (t Tree) find(data int, node *Node) bool {
	if node == nil {
		return false
	}
	var result bool
	switch {
	case data == node.key:
		result = true
	case data < node.key:
		result = t.find(data, node.lson)
	case data > node.key:
		result = t.find(data, node.rson)
	}
	return result
}

func (t Tree) Trace(up bool) []int {
	//  up - return tree in the ascending order
	// !up - return tree in the descending order
    return t.trace(t.root, up)
}    

func (t Tree) trace(node *Node, up bool) []int {
	if node == nil {
		return []int{}
	}
	// node != nil
	if up {  // ascending order
		return append ( append(t.trace(node.lson, true), node.key), t.trace(node.rson, true)...)
	} else { // descending order	
		return append ( append(t.trace(node.rson, false), node.key), t.trace(node.lson, false)...)
	}
}
    
func (t *Tree) Delete(data int) {
	t.root = t.delete(data, t.root)
}

func (t *Tree) delete(data int, node *Node) *Node {
	if node == nil {
		return nil
	}

	if data < node.key {
		node.lson = t.delete(data, node.lson)
	} else if data > node.key {
		node.rson = t.delete(data, node.rson)
	} else  // data == node.key
	if node.lson != nil && node.rson != nil {
		min := t.min(node.rson)
		node.rson = t.delete(min.key, node.rson) //Удалять не по ключу, а по поинтеру
		min.lson, min.rson = node.lson, node.rson
		node = min
	} else if node.lson == nil {
		node = node.rson
	} else { // node.rson == nil
		node = node.lson
	}
	return node
}

func (t *Tree) min(node *Node) *Node {
	for node.lson != nil {
		node = node.lson
	}
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

func (tree Tree) DraftDisplay() string {
	return tree.draftDisplay("", tree.root)
}

func (t Tree) draftDisplay(prefix string, node *Node) string {
	if node == nil {
		return ""
	}
	return t.draftDisplay("   "+prefix, node.rson) +
	       prefix + strconv.Itoa(node.key) + "\n" + 
	       t.draftDisplay("   "+prefix, node.lson)
}

func main() {
	data := []int{34, 45, 36, 7, 24, 2, 40, 27, 5, 3}
	tree := InitTree()
	for _, key := range data {
		tree.Insert(key)
	}
	fmt.Println(tree.DraftDisplay())
	fmt.Println(tree.TreeString())
	p := tree.root
	fmt.Println(*p)
	tree.Delete(34)
	fmt.Println(*p)
	fmt.Println(tree.Find(12))     // true
	fmt.Println(tree.Find(18))     // false
	fmt.Println(tree.Find(3))      // true
	fmt.Println(tree.Trace(true))  // [2 3 5 7 7 12 15 16 24 27 34 36 40 45 52]
	fmt.Println(tree.Trace(false)) // [52 45 40 36 34 27 24 16 15 12 7 7 5 3 2]
	for _, key := range data {
		if key%2 == 0 {
			tree.Delete(key)
		}
	}
	fmt.Println(tree.Trace(true))  //
	fmt.Println(tree.Trace(false)) //
	fmt.Println(tree.DraftDisplay())
	fmt.Println(tree.TreeString())
}
