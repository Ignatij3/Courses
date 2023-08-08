package sorted

import (
	"fmt"
	"order"
)

type (
	SortedCollection []order.Ordered
)

func (c *SortedCollection) Insert(x order.Ordered) {
	left, right := 0, len(*c)-1
	// the insertion location is searched using the binary search method.
	// invariant: (*c)[left].Before(x) && !(*c).[right].Before(x),
	// roughly speaking: (*c)[left] < x <= (*c)[right]
	(*c) = append((*c), x)
	if right < 0 || (*c)[right].Before(x) { //right < 0 <==> empty collection
		return
	}
	if !(*c)[0].Before(x) {
		copy((*c)[1:], (*c))
		(*c)[0] = x
		return
	}
	var center int
	for right-left > 1 {
		center = (left + right) / 2
		if (*c)[center].Before(x) {
			left = center
		} else {
			right = center
		}
	}
	copy((*c)[right+1:], (*c)[right:])
	(*c)[right] = x
}

func (c SortedCollection) Print() {
	for _, x := range c {
		fmt.Print(x.Show())
	}
}

type (
	Node struct {
		Key  order.Ordered
		Tail SortedList
	}
	SortedList struct {
		First *Node
	}
)

func NewSortedList() SortedList {
	return SortedList{} // <==> return SortedList{First:nil}
}

func (s SortedList) Empty() bool {
	return s.First == nil
}

func (s *SortedList) Insert(x order.Ordered) {
	if (*s).Empty() {
		(*s).First = &Node{Key: x, Tail: NewSortedList()}
		return
	}
	if !(*(*s).First).Key.Before(x) {
		*s = SortedList{&Node{Key: x, Tail: *s}}
		return
	}
	(*s.First).Tail.Insert(x)
}

func (s SortedList) Print() {
	if s.Empty() {
		return
	}
	fmt.Print((*s.First).Key.Show())
	(*s.First).Tail.Print()
}
