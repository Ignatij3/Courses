package dirlist


// exported types and adapters

// Adapters
type node struct {
	Node
}

func (e node) Next() Element {
	return node{*e.next}
}

func (e node) Value() interface{} {
	return e.value
}

type directedList struct {
	DirectedList
}

func (l directedList) Head() Element {
	return node{l.head}
}

func (l directedList) Tail() Element {
	return node{l.tail}
}

func (l directedList) Empty() bool {
	return l.DirectedList.Empty()
}

func NewdirectedList() directedList {
	l:= directedList{NewDirectedList()}
	l.tail = l.head
	return l
}

func (l directedList) Traverse(f func(v interface{})) {
	Traverse(l, f)
}

func (l directedList) Reverse(f func(v interface{})) {
	Reverse(l, f)
}

func (l directedList) Find(f func(v interface{}) bool ) (p *Node, err error) {
	pp, e := FirstThat(l, f);
	if e != nil { 
		err = e
		return 
	}
	pe:= pp.(node).Node
	p, err = &pe, e
	return
}

// Exported types
type Node struct {
	next  *Node
	value interface{}
}

func (e Node) Next() Node {
	return *e.next
}
func (e Node) Value() interface{} {
	return e.value
}

type DirectedList struct {
	head Node // sentinel First List Node;
	tail Node // sentinel Last List Node;
}

func (l DirectedList) Head() Node {
	return l.head
}
func (l DirectedList) Tail() Node {
	return l.tail
}

func (l DirectedList) Empty() bool {
	return l.tail.next == nil
}

func NewDirectedList() DirectedList {
	return DirectedList{head: Node{}, tail: Node{}}
}

func (l *DirectedList) PushFront(x interface{}) {
	if l.Empty() {
		l.head.next = &Node{value: x}
		l.tail.next = l.head.next
		return
	}
	l.head.next = &Node{value: x, next: l.head.next}
}

func (l *DirectedList) PushBack(x interface{}) {
	if l.Empty() {
		l.head.next = &Node{value: x}
		l.tail.next = l.head.next
		return
	}
	l.tail.next.next = &Node{value: x}
	l.tail.next = l.tail.next.next
}

func (l DirectedList) Traverse(f func(v interface{})) {
	directedList{l}.Traverse(f)
}

func (l DirectedList) Reverse(f func(v interface{})) {
	directedList{l}.Reverse(f)
}

func (l DirectedList) Find(f func(v interface{}) bool ) (*Node, error) {
	return directedList{l}.Find(f)
}
