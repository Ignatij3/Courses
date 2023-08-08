package binarytree

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// Algorithms
type BinaryTree interface {
	Empty() bool
	Left()  BinaryTree
	Right() BinaryTree
	Value() interface{}
}

type TraverseOrder uint

const (
	PreOrder	TraverseOrder = iota // NLR
	InOrder                 		 // LNR = ascending order
	PostOrder         		         // LRN
	ReversePreOrder   		         // NRL
	ReverseInOrder    		         // RNL = descending order
	ReversePostOrder  		         // RLN
)

func traverse(t BinaryTree, order TraverseOrder, visit func(r BinaryTree)) {
	switch order {
		case PreOrder: traverseNLR(t, visit)
		case InOrder:  traverseLNR(t, visit)
		case PostOrder: traverseLRN(t, visit)
		case ReversePreOrder: traverseNRL(t, visit)
		case ReverseInOrder: traverseRNL(t, visit)
		case ReversePostOrder: traverseRLN(t, visit)
	}
}

func traverseNLR(t BinaryTree, visit func(r BinaryTree)) {
	if t.Empty() { return }
	visit(t)
	traverseNLR(t.Left(), visit)
	traverseNLR(t.Right(), visit)
}	

func traverseLNR(t BinaryTree, visit func(r BinaryTree)) {
	if t.Empty() { return }
	traverseLNR(t.Left(), visit)
	visit(t)
	traverseLNR(t.Right(), visit)
}	

func traverseLRN(t BinaryTree, visit func(r BinaryTree)) {
	if t.Empty() { return }
	traverseLRN(t.Left(), visit)
	traverseLRN(t.Right(), visit)
	visit(t)
}	

func traverseNRL(t BinaryTree, visit func(r BinaryTree)) {
	if t.Empty() { return }
	visit(t)
	traverseNRL(t.Right(), visit)
	traverseNRL(t.Left(), visit)
}	

func traverseRNL(t BinaryTree, visit func(r BinaryTree)) {
	if t.Empty() { return }
	traverseRNL(t.Right(), visit)
	visit(t)
	traverseRNL(t.Left(), visit)
}	

func traverseRLN(t BinaryTree, visit func(r BinaryTree)) {
	if t.Empty() { return }
	traverseRLN(t.Right(), visit)
	traverseRLN(t.Left(), visit)
	visit(t)
}	

var ImageWidth int
func init() {
	ImageWidth = 5
}

func diagram(t BinaryTree, showRoot func (x interface{}) string) []string {
	if t.Empty() {
		return []string{""}
	}
	sL := diagram(t.Left(), showRoot)
	sR := diagram(t.Right(), showRoot)
	spaceL := strings.Repeat(" ", utf8.RuneCountInString(sL[0]))
	for len(sL) < len(sR) {
		sL = append(sL, spaceL)
	}
	spaceR := strings.Repeat(" ", utf8.RuneCountInString(sR[0]))
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

	s := []string{spaceL + fmt.Sprintf("%*s", ImageWidth, "|") + spaceR}
	s = append(s, sL[0] + showRoot(t.Value()) + sR[0])
	for i := 1; i < len(sL); i++ {
		s = append(s, sL[i]+strings.Repeat(" ", ImageWidth)+sR[i])
	}
	return s
}

 func rotatedDiagram( t BinaryTree, indent string,  show func(x interface{}) string) string {
	if t.Empty() {
		return ""
	}
	return rotatedDiagram(t.Right(), indent+"   ", show) + indent + show(t.Value()) + 
						  "\n" + rotatedDiagram(t.Left(), indent+"   ", show)
}

func search(t BinaryTree, v interface{}, less func(x, y interface{}) bool) bool {
	var search func(t BinaryTree) bool
	search = func(t BinaryTree) bool {
		if t.Empty() {
			return false
		}	
		if less(v, t.Value()) {
			return search(t.Left())
		}
		if less(t.Value(), v) {
			return search(t.Right())
		}	
		return true
	}
	return search(t)
}	
