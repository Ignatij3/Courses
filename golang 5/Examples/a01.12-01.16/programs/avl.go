package main
 
import (
	"fmt"
	"strconv"
)

type  (
    Node struct {
        key int
        lson *Node
        rson *Node
        height int
    }
    Tree struct {
        root *Node
    }
)

func InitTree() Tree {  
    return Tree{root: nil}
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
    
func (t *Tree) Find(key int) *Node {
	return t.find(key, t.root)
}

func (t *Tree) find(x int, node *Node) *Node {
	if node != nil {
		if x < node.key {
			return t.find(x, node.lson)
		} else if x > node.key {
			return t.find(x, node.rson)
		} else if x == node.key {
			return node
		}
	} 
	return nil		
}
 
//find min elem in the tree
func (t *Tree) MinNode() *Node {
	return t.minNode(t.root)
}
 
func (t *Tree) minNode(node *Node) *Node {
	if node != nil {
		if node.lson != nil {
			return t.minNode(node.lson)
		} else {return node}
	} else {return nil}
}

//insert an x to AvlTree
func (t *Tree) Insert(x int) {
	t.root = t.insert(t.root, x)
}

func (t *Tree) insert(node *Node, x int) *Node {
	if node == nil {
		//new(Node)
		node = &Node { 
			key : x,			
			height : 1,
		}
	} else if x < node.key {
		node.lson = t.insert(node.lson, x)
		if t.nodeHeight(node.lson) - t.nodeHeight(node.rson) == 2 {
			if x < node.lson.key { //left left
				node = t.singleRotateLeft(node)
			} else { //left right
				node = t.doubleRotateLeftRight(node)
			}
		}
	} else if x > node.key {
		node.rson = t.insert(node.rson, x)
		if t.nodeHeight(node.rson) - t.nodeHeight(node.lson) == 2 {
			if x >= node.rson.key {
				node = t.singleRotateRight(node)
			} else {
				node = t.doubleRotateRightLeft(node)
			}
		}
	} 
	t.updateHeight(node) 
	return node
}
 
//delete an x in AvlTree
func (t *Tree) Delete(x int) {
	t.root = t.delete(t.root, x)
}

func (t *Tree) delete(node *Node, x int) *Node {	
	if node != nil {
		if x < node.key {
			node.lson = t.delete(node.lson, x)	
			if t.nodeHeight(node.rson) - t.nodeHeight(node.lson) == 2 {
				if t.nodeHeight(node.rson.lson) <= t.nodeHeight(node.rson.rson) {
					node = t.singleRotateRight(node)
				} else {
					node = t.doubleRotateRightLeft(node)
				}			
			}
		} else 
		if x > node.key {
			node.rson = t.delete(node.rson, x)
			if t.nodeHeight(node.lson) - t.nodeHeight(node.rson) == 2 {
				if t.nodeHeight(node.lson.rson) <= t.nodeHeight(node.lson.lson) {
					node = t.singleRotateLeft(node)
				} else {
					node = t.doubleRotateLeftRight(node)
				}
			}
		} else {
		// x == node.key
			if node.lson != nil && node.rson != nil {
				min:= t.minNode(node.rson)
				node.rson = t.delete(node.rson, min.key)
				min.lson, min.rson = node.lson, node.rson
				node = min
				if t.nodeHeight(node.lson) - t.nodeHeight(node.rson) == 2 {
					if t.nodeHeight(node.lson.rson) <= t.nodeHeight(node.lson.lson) {
						node = t.singleRotateLeft(node)
					} else {
						node = t.doubleRotateLeftRight(node)
					}
				}			
			} else {			
				if node.lson == nil {
					node = node.rson
				} else 
				if node.rson == nil {
					node = node.lson
				}			
			}		
		}
	}
	t.updateHeight(node) 
	return node
}
 
// left rotate a tree, and update node's height
// return the new root 
func (t *Tree) singleRotateLeft(node *Node) *Node {
	var left *Node
	if node != nil {
		// turn left
		left = node.lson
		node.lson = left.rson
		left.rson = node
		
		//update height
		t.updateHeight(node) 
		t.updateHeight(left) 
 
		node = left
	}
 
	return node
}
 
// right rotate a tree, and update node's height
// return the new root
func (t *Tree) singleRotateRight(node *Node) *Node {
	var right *Node
	if node != nil {
		//turn right
		right = node.rson
		node.rson = right.lson
		right.lson = node
		
		//update height
		t.updateHeight(node) 
		t.updateHeight(right) 
 
		node = right
	}
	return node
}

// v = subtree root, vl = v's left child, vlr = vl's right child
// right rotate vl & vlr, left rotate v & v's left child
// return a new tree
func (t *Tree) doubleRotateLeftRight(v *Node) *Node  {
	//right rotatel between vl & vlr
	v.lson = t.singleRotateRight(v.lson)
 
	//left rotate between v and his left child
	return t.singleRotateLeft(v)
}
 
// v = subtree root, vr = vr's right child, vrl = vr's left child
// left rotate vr & vrl, right rotate v & v's right child
// return a new tree
func (t *Tree) doubleRotateRightLeft(v *Node) *Node  {
	//left rotatel between vr & vrl
	v.rson = t.singleRotateLeft(v.rson)
 
	//right rotate between v and his left child
	return t.singleRotateRight(v)
}

//return the height of the node
func (t *Tree) nodeHeight(node *Node) int {
	if node == nil {return 0} else {return node.height}
}

//recalculate the height of the node
func (t *Tree) updateHeight(node *Node) {
	if node == nil {
		return
	}	
	if t.nodeHeight(node.lson) > t.nodeHeight(node.rson) {
		node.height = t.nodeHeight(node.lson) + 1
	} else {
		node.height = t.nodeHeight(node.rson) + 1
	}	
}
 
func (t Tree) TreeString() string {
	return t.treeString("", true, "", t.root)
}

func (t Tree) treeString(prefix string, top bool, str string, node *Node) string {
	if (node == nil) {
		return ""
	}
	var temp string
	if (node.rson != nil) {
		if top {
			temp = prefix + "│   "
		} else {
			temp = prefix + "    "	
		}	
		str = t.treeString(temp, false, str, node.rson);
	}
	str += prefix
	if top {
		str += "└──"
	} else {
		str += "┌──"
	}	
	str += " " + strconv.Itoa(node.key) + "/" + strconv.Itoa(node.height) + "\n";
	if (node.lson != nil) {
		if top {
			temp = prefix + "    "
		} else {
			temp = prefix + "│   "	
		}	
		str = t.treeString(temp, true, str, node.lson);
	}
	return str
}


func main() {
    data:= []int{34, 45, 36, 8, 24, 2, 40, 40, 40, 27, 5, 3, 52, 7, 16, 12, 15}
    tree:= InitTree()
    for _, key := range data {tree.Insert(key)}
	fmt.Println(tree.TreeString())
	tree.Delete(24)
	fmt.Println(tree.TreeString())
	tree.Delete(34)
	fmt.Println(tree.TreeString())
	tree.Delete(36)
	fmt.Println(tree.TreeString())
} 
