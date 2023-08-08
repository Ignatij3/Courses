package sorted

type Element struct {
	next  *Element
	Value interface{}
}

type List struct {
	head Element // sentinel list Element;
	// it's located at the top of the list
	len  int // current list length excluding sentinel
	less func(*Element, *Element) bool
}

func NewList(less func(*Element, *Element) bool) List {
	return List{less: less}
}

func (l *List) Add(x interface{}) {
	p := &(l.head)
	v := &Element{Value: x}
	for i := 0; i < l.len; i++ {
		if l.less(v, (*p).next) {
			break
		}
		p = (*p).next
	}
	(*v).next = (*p).next
	p.next = v
	l.len++
}

func (l List) Do(f func(v interface{})) {
	p := l.head
	for i := 0; i < l.len; i++ {
		p = *(p.next)
		f(p.Value)
	}
}

func (l List) Len() int {
	return l.len
}
