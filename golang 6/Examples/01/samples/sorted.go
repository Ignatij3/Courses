package sorted

import "fmt"

type Ordered interface {
	Before(b Ordered) bool
}

type SortedCollection []Ordered

func (c *SortedCollection) Insert(x Ordered) {
	*c = append(*c, x)
	i := len(*c) - 1
	for i > 0 && x.Before((*c)[i-1]) {
		i--
	}
	copy((*c)[i+1:], (*c)[i:])
	(*c)[i] = x
}

func (c SortedCollection) Print() {
	fmt.Println(c)
}
