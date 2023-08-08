package binarytree

type (
	BSNode struct {
		value interface{}
		size int
		lson  BSTree
		rson  BSTree
	}
	BSTree struct {
		root *BSNode
		less func(x, y interface{}) bool   // ?x<y
	}
)

//Adapter
type bsTree BSTree

func (t bsTree) Left()  BinaryTree  { return bsTree(t.root.lson) }
func (t bsTree) Right() BinaryTree  { return bsTree(t.root.rson) }
func (t bsTree) Empty() bool        { return t.root == nil }
func (t bsTree) Value() interface{} { return t.root.value }
func (t bsTree) Size()  int         { 
	if t.Empty() { return 0 }
	return t.root.size 
}	

func (t BSTree) Value() interface{} { return t.root.value }
func (t BSTree) Empty() bool { return t.root == nil }
func (t BSTree) Left()  *BSTree { return &t.root.lson }
func (t BSTree) Right() *BSTree { return &t.root.rson }	
func (t BSTree) Size()  int { return bsTree(t).Size() }	

func NewBSTree(less func(x, y interface{}) bool) BSTree {
	return BSTree{less: less}  // root = nil
} 

func (t *BSTree) Insert(newval interface{}) {
	if bsTree(*t).Empty() {
		t.root = &BSNode{value: newval, lson: NewBSTree(t.less), rson: NewBSTree(t.less), size:1 }
		return
	}
	if t.less(newval, t.Value()) {
		t.root.lson.Insert(newval)
	} else {
		t.root.rson.Insert(newval)
	}
	t.root.size++  //t.updateSize()
}

func (t *BSTree) updateSize() {
	if !t.Empty() {
		(*t).root.size = t.Left().Size() + t.Right().Size() + 1
	}
}	

func (t *BSTree) Delete(value interface{}) {
	if t.Empty() {
		return
	}
	switch {
	case t.less(value, t.Value()):
		// value "<" t.root.value
		t.root.lson.Delete(value)
	case t.less(t.Value(), value):
		// value ">" t.root.value:
		t.root.rson.Delete(value)
	default:
		// value "==" (*t.root).value:
		switch {
		case t.root.lson.Empty():
			*t = t.root.rson
		case t.root.rson.Empty():
			*t = t.root.lson
		default:
			// !t.root.lson.Empty() && t.root.rson.Empty()
			min := t.root.rson.leftmost()
			t.root.rson.Delete(min.root.value)
			min.root.lson, min.root.rson = t.root.lson, t.root.rson
			*t = min
		}
	}
	t.updateSize()
}

func (t BSTree) leftmost() BSTree {
	if t.Empty() {
		return t
	}
	for !t.root.lson.Empty() {
		t = t.root.lson
	}
	return t
}

func (t BSTree) RotatedDiagram(indent string, show func (x interface{}) string) string {
	return rotatedDiagram(bsTree(t), indent,  show)
}

func (t BSTree) Search(v interface{}) bool {
	return search(bsTree(t), v, t.less)	
}

func (t BSTree) Traverse(order TraverseOrder, visit func(r BSTree)) {
	traverse(bsTree(t), order, func (t BinaryTree) {
		visit(BSTree(t.(bsTree)))
	})	 
}		

func (t BSTree) Diagram (showRoot func (x interface{}) string) []string {
	return diagram( bsTree(t), showRoot)	
}	

func (t BSTree) FindKth(k int) interface{} {
	return findKth(bsTree(t), k)	
}		
