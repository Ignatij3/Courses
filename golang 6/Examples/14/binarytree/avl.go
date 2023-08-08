package binarytree

type (
	AVLNode struct {
		value  interface{}
		lson   AVLTree
		rson   AVLTree
		height int
	}
	AVLTree struct {
		root *AVLNode
		less func(x, y interface{}) bool   // ?x<y
	}
)

//Adapter
type avlTree AVLTree

func (t avlTree) Left()  BinaryTree  { return avlTree(t.root.lson) }
func (t avlTree) Right() BinaryTree  { return avlTree(t.root.rson) }
func (t avlTree) Empty() bool        { return t.root == nil }
func (t avlTree) Value() interface{} { return t.root.value }


func (t AVLTree) Value() interface{} { return t.root.value }	
func (t AVLTree) Empty() bool { return t.root == nil }	
func (t AVLTree) Left() *AVLTree { return &t.root.lson }
func (t AVLTree) Right() *AVLTree { return &t.root.rson }	

func (t AVLTree) Height() int {
	if t.Empty() {
		return 0
	}
	return t.root.height
}	

func NewAVLTree(less func(x, y interface{}) bool) AVLTree {
	return AVLTree{less: less}  // root = nil
} 

func (t *AVLTree) Insert(x interface{}) {
	if avlTree(*t).Empty() {
		//new(Node)
		t.root = &AVLNode{value: x, lson: NewAVLTree(t.less), rson: NewAVLTree(t.less), height: 1}
		return
	} 
	
	if t.less(x, t.Value()) {
		t.root.lson.Insert(x)
		if t.root.lson.Height()-t.root.rson.Height() == 2 {
			if t.less(x, t.root.lson.Value()) { //left left
				t.singleRotateLeft()
			} else { //left right
				t.doubleRotateLeftRight()
			}
		}
	} else {
		t.root.rson.Insert(x)
		if t.root.rson.Height()-t.root.lson.Height() == 2 {
			if !t.less(x, t.root.rson.Value()) {
				t.singleRotateRight()
			} else {
				t.doubleRotateRightLeft()
			}
		}
	}
	t.updateHeight()
}

// left rotate a tree, and update tree's height as well as height of all subtrees
func (t *AVLTree) singleRotateLeft() {
	var left AVLTree
	if !t.Empty() {
		// turn left
		left = t.root.lson
		t.root.lson = left.root.rson
		left.root.rson = *t

		//update height
		t.updateHeight()
		left.updateHeight()

		*t = left
	}
}

// right rotate a tree, and update tree's height as well as height of all subtrees
func (t *AVLTree) singleRotateRight() {
	var right AVLTree
	if !t.Empty() {
		// turn right
		right = t.root.rson
		t.root.rson = right.root.lson
		right.root.lson = *t

		//update height
		t.updateHeight()
		right.updateHeight()

		*t = right
	}
}

// v = subtree root, vl = v's left child, vlr = vl's right child
// right rotate vl & vlr, left rotate v & v's left child
// return a new tree
func (t *AVLTree) doubleRotateLeftRight() {
	//right rotatel between vl & vlr
	t.root.lson.singleRotateRight()

	//left rotate between v and his left child
	t.singleRotateLeft()
}

// v = subtree root, vr = vr's right child, vrl = vr's left child
// left rotate vr & vrl, right rotate v & v's right child
// return a new tree
func (t *AVLTree) doubleRotateRightLeft() {
	//left rotatel between vr & vrl
	t.root.rson.singleRotateLeft()

	//right rotate between v and his left child
	t.singleRotateRight()
}

//recalculate the height of the node
func (t *AVLTree) updateHeight() {
	if t.Empty() {
		return
	}
	if t.root.lson.Height() > t.root.rson.Height() {
		t.root.height = t.root.lson.Height() + 1
	} else {
		t.root.height = t.root.rson.Height() + 1
	}
}

func (t *AVLTree) Delete(x interface{}) {
	if t.Empty() {
		return
	}	

	if t.less(x, t.Value()) || (!t.root.lson.Empty() && !t.less(t.root.lson.Value(), x)) { // <= - главное отличие удаления
		t.root.lson.Delete(x)
		if t.root.rson.Height() - t.root.lson.Height() == 2 {
			if t.root.rson.root.lson.Height() <= t.root.rson.root.rson.Height() {
				t.singleRotateRight()
			} else {
				t.doubleRotateRightLeft()
			}
		}
	} else if t.less(t.Value(), x) {
		t.root.rson.Delete(x)
		if t.root.lson.Height() -t.root.rson.Height() == 2 {
			if t.root.lson.root.rson.Height() <= t.root.lson.root.lson.Height() {
				t.singleRotateLeft()
			} else {
				t.doubleRotateLeftRight()
			}
		}
	} else if t.root.lson.Empty() {
		*t = t.root.rson
	} else if t.root.rson.Empty() {
		*t = t.root.lson
	} else
	/* !leftsubtree.Empty() && !rightsubtree.Empty() */ {
		min := t.root.rson.leftmost()
		t.root.rson.Delete(min.Value())
		min.root.lson, min.root.rson = t.root.lson, t.root.rson
		t.root = min.root
		if t.root.lson.Height() - t.root.rson.Height() == 2 {
			if t.root.lson.root.rson.Height() <= t.root.lson.root.lson.Height() {
				t.singleRotateLeft()
			} else {
				t.doubleRotateLeftRight()
			}
		}
	}
	t.updateHeight()
}

func (t AVLTree) leftmost() AVLTree {
	if t.Empty() {
		return t
	}
	for !t.root.lson.Empty() {
		t = t.root.lson
	}
	return t
}

func (t AVLTree) RotatedDiagram(indent string, show func (x interface{}) string) string {
	return rotatedDiagram(avlTree(t), indent,  show)
}

func (t AVLTree) Search(v interface{}) bool {
	return search(avlTree(t), v, t.less)	
}

func (t AVLTree) Traverse(order TraverseOrder, visit func(r AVLTree)) {
	traverse(avlTree(t), order, func (t BinaryTree) {
		visit(AVLTree(t.(avlTree)))
	})	 
}		

func (t AVLTree) Diagram (showRoot func (x interface{}) string) []string {
	return diagram( avlTree(t), showRoot)	
}	
