package avl

import (
	"fmt"
	"order"
	"strings"
)

type (
	Node struct {
		Value  order.Key
		Lson   AVLtree
		Rson   AVLtree
		height int
	}
	AVLtree struct {
		root *Node
	}
)

func NewAVLtree() AVLtree {
	return AVLtree{root: nil}
}

func (t AVLtree) Empty() bool {
	return t.root == nil
}

func (t AVLtree) Traversal(dir int, f func(x order.Key)) {
	if t.Empty() {
		return
	}
	switch dir {
	case order.PreOrder: // NLR
		f((*t.root).Value)
		(*t.root).Lson.Traversal(dir, f)
		(*t.root).Rson.Traversal(dir, f)
	case order.InOrder: // LNR
		(*t.root).Lson.Traversal(dir, f)
		f((*t.root).Value)
		(*t.root).Rson.Traversal(dir, f)
	case order.PostOrder: // LRN
		(*t.root).Lson.Traversal(dir, f)
		(*t.root).Rson.Traversal(dir, f)
		f((*t.root).Value)
	case order.ReversePreOrder: // NRL
		f((*t.root).Value)
		(*t.root).Rson.Traversal(dir, f)
		(*t.root).Lson.Traversal(dir, f)
	case order.ReverseInOrder: // RNL
		(*t.root).Rson.Traversal(dir, f)
		f((*t.root).Value)
		(*t.root).Lson.Traversal(dir, f)
	case order.ReversePostOrder: // RLN
		(*t.root).Rson.Traversal(dir, f)
		(*t.root).Lson.Traversal(dir, f)
		f((*t.root).Value)
	}
}

func (t AVLtree) Find(n order.Ordered) (node *Node, ok bool) {
// Здесь мы никак не используем x.Show(), так что для поиска нам
// достаточно иметь ключ интерфейса order.Ordered, а не order.Key
	switch {
	case t.Empty():
		return nil, false
	case n.Before((*t.root).Value):
		//	n before (*t.Root).Value
		return (*t.root).Lson.Find(n)
	case (*t.root).Value.Before(n):
		//	n after (*t.Root).Value
		return (*t.root).Rson.Find(n)
	default:
		// n is equivalent to (*t.Root).Value
		return t.root, true
	}
}

//insert an x to AVLtree
func (t *AVLtree) Insert(x order.Key) {
	if (*t).Empty() {
		//new(Node)
		(*t).root = &Node{
			Value:  x,
			height: 1,
		}
	} else if x.Before((*(*t).root).Value) {
		(*(*t).root).Lson.Insert(x)
		if (*(*t).root).Lson.treeHeight()-(*(*t).root).Rson.treeHeight() == 2 {
			if x.Before((*(*(*t).root).Lson.root).Value) { //left left
				(*t).singleRotateLeft()
			} else { //left right
				(*t).doubleRotateLeftRight()
			}
		}
	} else /* !x.Before((*(*t).root).Value) */ {
		(*t).root.Rson.Insert(x)
		if (*(*t).root).Rson.treeHeight()-(*(*t).root).Lson.treeHeight() == 2 {
			if !x.Before((*(*(*t).root).Rson.root).Value) {
				(*t).singleRotateRight()
			} else {
				(*t).doubleRotateRightLeft()
			}
		}
	}
	(*t).updateHeight()
}

//delete an x in AVLtree
func (t *AVLtree) Delete(x order.Ordered) {
// И здесь мы никак не используем и не будем использовать в дальнейшем x.Show(), так что
// можно считать, что мы удаляем ключ именно интерфейса order.Ordered, а не order.Key
	if !(*t).Empty() {
		if x.Before((*(*t).root).Value) || (!(*(*t).root).Lson.Empty() && !(*(*(*t).root).Lson.root).Value.Before(x)) { // <= - главное отличие удаления
			(*(*t).root).Lson.Delete(x)
			if (*(*t).root).Rson.treeHeight()-(*(*t).root).Lson.treeHeight() == 2 {
				if (*(*(*t).root).Rson.root).Lson.treeHeight() <= (*(*(*t).root).Rson.root).Rson.treeHeight() {
					(*t).singleRotateRight()
				} else {
					(*t).doubleRotateRightLeft()
				}
			}
		} else if (*(*t).root).Value.Before(x) {
			(*(*t).root).Rson.Delete(x)
			if (*(*t).root).Lson.treeHeight()-(*(*t).root).Rson.treeHeight() == 2 {
				if (*(*(*t).root).Lson.root).Rson.treeHeight() <= (*(*(*t).root).Lson.root).Lson.treeHeight() {
					(*t).singleRotateLeft()
				} else {
					(*t).doubleRotateLeftRight()
				}
			}
		} else if (*(*t).root).Lson.Empty() {
			*t = (*(*t).root).Rson
		} else if (*(*t).root).Rson.Empty() {
			*t = (*(*t).root).Lson
		} else
		/* !leftsubtree.Empty() && !rightsubtree.Empty() */ {
			min := (*(*t).root).Rson.leftmost()
			(*(*t).root).Rson.Delete((*min).Value)
			(*min).Lson, (*min).Rson = (*(*t).root).Lson, (*(*t).root).Rson
			(*t).root = min
			if (*(*t).root).Lson.treeHeight()-(*(*t).root).Rson.treeHeight() == 2 {
				if (*(*(*t).root).Lson.root).Rson.treeHeight() <= (*(*(*t).root).Lson.root).Lson.treeHeight() {
					(*t).singleRotateLeft()
				} else {
					(*t).doubleRotateLeftRight()
				}
			}
		}
	}
	(*t).updateHeight()
}

//find leftmost min elem in the subtree
func (t AVLtree) leftmost() *Node {
	if t.Empty() {
		return nil
	}
	for !(*t.root).Lson.Empty() {
		t = (*t.root).Lson
	}
	return t.root
}

// left rotate a tree, and update tree's height as well as height of all subtrees
func (t *AVLtree) singleRotateLeft() {
	var left AVLtree
	if !(*t).Empty() {
		// turn left
		left = (*(*t).root).Lson
		(*(*t).root).Lson = (*left.root).Rson
		(*left.root).Rson = *t

		//update height
		(*t).updateHeight()
		left.updateHeight()

		*t = left
	}
}

// right rotate a tree, and update tree's height as well as height of all subtrees
func (t *AVLtree) singleRotateRight() {
	var right AVLtree
	if !(*t).Empty() {
		// turn right
		right = (*(*t).root).Rson
		(*(*t).root).Rson = (*right.root).Lson
		(*right.root).Lson = *t

		//update height
		(*t).updateHeight()
		right.updateHeight()

		*t = right
	}
}

// v = subtree root, vl = v's left child, vlr = vl's right child
// right rotate vl & vlr, left rotate v & v's left child
// return a new tree
func (t *AVLtree) doubleRotateLeftRight() {
	//right rotatel between vl & vlr
	(*(*t).root).Lson.singleRotateRight()

	//left rotate between v and his left child
	(*t).singleRotateLeft()
}

// v = subtree root, vr = vr's right child, vrl = vr's left child
// left rotate vr & vrl, right rotate v & v's right child
// return a new tree
func (t *AVLtree) doubleRotateRightLeft() {
	//left rotatel between vr & vrl
	(*(*t).root).Rson.singleRotateLeft()

	//right rotate between v and his left child
	(*t).singleRotateRight()
}

//return the height of the node
func (t AVLtree) treeHeight() int {
	if t.Empty() {
		return 0
	} else {
		lh, rh := (*t.root).Lson.treeHeight(), (*t.root).Rson.treeHeight()
		if rh > lh {
			return rh + 1
		} else {
			return lh + 1
		}
	}
}

//recalculate the height of the node
func (t *AVLtree) updateHeight() {
	if (*t).Empty() {
		return
	}
	if (*(*t).root).Lson.treeHeight() > (*(*t).root).Rson.treeHeight() {
		(*(*t).root).height = (*(*t).root).Lson.treeHeight() + 1
	} else {
		(*(*t).root).height = (*(*t).root).Rson.treeHeight() + 1
	}
}

func (t AVLtree) Diagram() []string {
	if t.Empty() {
		return []string{""}
	}
	sL := (*t.root).Lson.Diagram()
	sR := (*t.root).Rson.Diagram()
	spaceL := strings.Repeat(" ", stringLen(sL[0]))
	for len(sL) < len(sR) {
		sL = append(sL, spaceL)
	}
	spaceR := strings.Repeat(" ", stringLen(sR[0]))
	for len(sR) < len(sL) {
		sR = append(sR, spaceR)
	}
	r := []rune(sL[0])
	ch := ' '
	for i := range r {
		if r[i] == '|' {
			r[i] = '┌'
			ch = '─'
		} else {
			r[i] = ch
		}
	}
	sL[0] = string(r)
	r = []rune(sR[0])
	ch = '─'
	for i := range r {
		if r[i] == '|' {
			r[i] = '┐'
			ch = ' '
		} else {
			r[i] = ch
		}
	}
	sR[0] = string(r)

	s := []string{spaceL + fmt.Sprintf("%*s", order.ImageWidth, "|") + spaceR}
	s = append(s, sL[0]+(*t.root).Value.Show()+sR[0])
	for i := 1; i < len(sL); i++ {
		s = append(s, sL[i]+strings.Repeat(" ", order.ImageWidth)+sR[i])
	}
	return s
}

func stringLen(s string) int {
	return len([]rune(s))
}
