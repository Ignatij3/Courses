package bst

import (
	"fmt"
	"order"
	"strings"
)

type (
	Node struct {
		Value order.Key
		Lson  Tree
		Rson  Tree
	}
	Tree struct {
		Root *Node
	}
)

func NewTree() Tree {
	return Tree{Root: nil} // <==> return Tree{}
}

func (t Tree) Empty() bool {
	return t.Root == nil
}

func (t *Tree) Insert(value order.Key) {
	if (*t).Empty() {
		(*t).Root = &Node{Value: value, Lson: NewTree(), Rson: NewTree()}
		return
	}
	if value.Before((*t.Root).Value) {
		// value "<"" (*t.Root).Value
		(*t.Root).Lson.Insert(value)
	} else {
		// Value "!<" (*t.Root).Value {
		(*t.Root).Rson.Insert(value)
	}
}

func (t Tree) Find(n order.Ordered) bool {
	switch {
	case t.Empty():
		return false
	case n.Before((*t.Root).Value):
		//	n before (*t.Root).Value
		return (*t.Root).Lson.Find(n)
	case (*t.Root).Value.Before(n):
		//	n after (*t.Root).Value
		return (*t.Root).Rson.Find(n)
	default:
		// n is equivalent to (*t.Root).Value
		return true
	}
}

func (t Tree) Traversal(dir int, f func(x order.Key)) {
	if t.Empty() {
		return
	}
	switch dir {
	case order.PreOrder: // NLR
		f((*t.Root).Value)
		(*t.Root).Lson.Traversal(dir, f)
		(*t.Root).Rson.Traversal(dir, f)
	case order.InOrder: // LNR
		(*t.Root).Lson.Traversal(dir, f)
		f((*t.Root).Value)
		(*t.Root).Rson.Traversal(dir, f)
	case order.PostOrder: // LRN
		(*t.Root).Lson.Traversal(dir, f)
		(*t.Root).Rson.Traversal(dir, f)
		f((*t.Root).Value)
	case order.ReversePreOrder: // NRL
		f((*t.Root).Value)
		(*t.Root).Rson.Traversal(dir, f)
		(*t.Root).Lson.Traversal(dir, f)
	case order.ReverseInOrder: // RNL
		(*t.Root).Rson.Traversal(dir, f)
		f((*t.Root).Value)
		(*t.Root).Lson.Traversal(dir, f)
	case order.ReversePostOrder: // RLN
		(*t.Root).Rson.Traversal(dir, f)
		(*t.Root).Lson.Traversal(dir, f)
		f((*t.Root).Value)
	}
}

func (t *Tree) Delete(value order.Ordered) {
	if t.Empty() {
		return
	}
	switch {
	case value.Before((*t.Root).Value):
		// value "<" (*t.Root).Value
		(*t.Root).Lson.Delete(value)
	case (*t.Root).Value.Before(value):
		// value ">" (*t.Root).Value:
		(*t.Root).Rson.Delete(value)
	default:
		// value "==" (*t.Root).Value:
		switch {
		case (*t.Root).Lson.Empty():
			*t = (*t.Root).Rson
		case (*t.Root).Rson.Empty():
			*t = (*t.Root).Lson
		default:
			// !(*t.Root).Lson.Empty() && (*t.Root).Rson.Empty()
			min := (*t.Root).Rson.leftmost()
			(*t.Root).Rson.Delete((*min.Root).Value)
			(*min.Root).Lson, (*min.Root).Rson = (*t.Root).Lson, (*t.Root).Rson
			*t = min
		}
	}
}

func (t Tree) leftmost() Tree {
	if t.Empty() {
		return NewTree()
	}
	for !(*t.Root).Lson.Empty() {
		t = (*t.Root).Lson
	}
	return t
}

func (t Tree) DraftDisplay(indent string) string {
	if t.Empty() {
		return ""
	}
	return (*t.Root).Rson.DraftDisplay(indent+"   ") +
		indent + (*t.Root).Value.Show() + "\n" +
		(*t.Root).Lson.DraftDisplay(indent+"   ")
}

func (t Tree) Diagram() []string {
	if t.Empty() {
		return []string{""}
	}
	sL := (*t.Root).Lson.Diagram()
	sR := (*t.Root).Rson.Diagram()
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
	s = append(s, sL[0]+(*t.Root).Value.Show()+sR[0])
	for i := 1; i < len(sL); i++ {
		s = append(s, sL[i]+strings.Repeat(" ", order.ImageWidth)+sR[i])
	}
	return s
}

func stringLen(s string) int {
	return len([]rune(s))
}
