package dirlist

import (
	"sort"
)	

type SortList struct {
	DirectedList
	less func(*Node, *Node) bool
}	

func NewSortList(less func(*Node, *Node) bool) SortList {
	return SortList{DirectedList: NewDirectedList(), less: less}
}

func (l *SortList) Insert(x interface{}) {
	if l.Empty() {	// empty list
		l.PushFront(x)
		return
	}
	v := &Node{value: x}
	if l.head == l.tail {	// 1-element list
		if l.less(v, l.head.next) {
			l.PushFront(x)
		} else {
			l.PushBack(x)
		}		
		return
	}
	p := &(l.head)
	var pNext *Node
	for {
		pNext = (*p).next
		if p == l.tail.next || l.less(v, pNext) {
			break
		}
		p = pNext
	}
	(*v).next = (*p).next
	if p ==  l.tail.next {
		l.tail.next = v
	}	
	(*p).next = v
}

func (l *SortList) Sort() {
	var s []*Node
	for p:= l.Head().next; p!= nil; p = (*p).next {
		s = append(s, p) 
	}
	
	sort.Slice(s, func(i, j int) bool { return l.less(s[i], s[j])})
	
	l.head.next = s[0]
	l.tail.next = s[len(s)-1]
	n:= s[0]
	for i:= 0; i < len(s)-1; i++ {
		(*n).next = s[i+1]
		n = s[i+1]
	}
	(*n).next = nil	
}	
