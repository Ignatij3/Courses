package sorted

type Element struct {
	Next *Element
	Value int
}

type SortedList struct {
	Head Element // sentinel list Element; 
		      // it's located at the top of the list
	Len  int  // current list length excluding sentinel
	less func(*Element, *Element) bool 
}

func NewSortedList (less func(*Element, *Element) bool ) SortedList { 
	return SortedList{less: less}
}

func (l *SortedList) Add(x int) {
	p := &(l.Head)
	v := &Element{Value:x}
	for i:= 0; i < l.Len; i++ {
		if l.less(v, (*p).Next) { break }
		p = (*p).Next
	}
	(*v).Next = (*p).Next
	p.Next =  v
	l.Len++
}
