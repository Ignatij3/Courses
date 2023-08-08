package dirlist

import (
	"errors"
)

// interfaces and algorithms
type DirList interface {
	Head() Element
	Tail() Element
	Empty() bool
}

type Element interface {
	Next() Element
	Value() interface{}
}	

func Traverse(l DirList, f func(v interface{})) {
	if l.Empty() {return}
	p:= l.Head()
	var pNext Element	
	for {
		pNext = p.Next()
		f(pNext.Value()) 
		if pNext == l.Tail().Next() { return }
		p = pNext
	}
}

func Reverse(l DirList, f func(v interface{})) {	
	if l.Empty() {return}
	var reverse func(n Element) 
	reverse = func (n Element) {
		if n == l.Tail().Next() { return }
 		n = n.Next()
 		reverse(n)
		f(n.Value())
	}
	reverse (l.Head())	
}

func FirstThat(l DirList, f func(v interface{}) bool) (Element, error) {
	if l.Empty() {
		return l.Head(), errors.New("no such element")
	}
	p := l.Head()
	var pNext Element
	for p != l.Tail().Next() {
		pNext = p.Next()
		if f(pNext.Value()) { return pNext, nil }
		p = pNext
	}
	return p, errors.New("no such element")
}		
