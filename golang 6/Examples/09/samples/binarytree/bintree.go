package binarytree

// BinTree
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

//Adapter
type bTree BinTree

func (t bTree) Left() BinaryTree   { return bTree(t.root.lson) }
func (t bTree) Right() BinaryTree  { return bTree(t.root.rson) }
func (t bTree) Empty() bool        { return t.root == nil }
func (t bTree) Value() interface{} { return t.root.value }

func NewBinTree() BinTree {
	return BinTree{}
}

func (t BinTree) Value() interface{} { return bTree(t).Value() }
func (t BinTree) Empty() bool        { return bTree(t).Empty() }
func (t BinTree) Left() *BinTree     { return &t.root.lson }
func (t BinTree) Right() *BinTree    { return &t.root.rson }

func (t *BinTree) Insert(newval interface{}, toLeft func(newval, rootval interface{}) bool) {
	if bTree(*t).Empty() {
		t.root = &Node{value: newval, lson: NewBinTree(), rson: NewBinTree()}
		return
	}
	if toLeft(newval, bTree(*t).Value()) {
		t.root.lson.Insert(newval, toLeft)
	} else {
		t.root.rson.Insert(newval, toLeft)
	}
}

func (t BinTree) RotatedDiagram(indent string, show func(x interface{}) string) string {
	return rotatedDiagram(bTree(t), indent, show)
}

func (t BinTree) Traverse(order TraverseOrder, visit func(r BinTree)) {
	traverse(bTree(t), order, func(t BinaryTree) {
		visit(BinTree(t.(bTree)))
	})
}

func (t BinTree) Diagram(showRoot func(x interface{}) string) []string {
	return diagram(bTree(t), showRoot)
}
