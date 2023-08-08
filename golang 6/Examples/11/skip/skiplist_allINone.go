package main

import (
	"fmt"
	"math/rand"
	"time"
	"errors"
)

type SkipListNode struct {
	next  []*SkipListNode
	value interface{}
}

func (u *SkipListNode) Value() interface{} {
	return u.value
}	

func (u *SkipListNode) Next() *SkipListNode {
	return u.next[0]
}	

func (u *SkipListNode) level() int {
	return len(u.next)-1
}	

type SkipList struct {
	head      *SkipListNode
	maxLevel  int
	topLevel  int
	p         float64
	less      func (interface{}, interface{}) bool
}

func NewSkipList(maxLevel int, p float64, less func (interface{}, interface{}) bool ) SkipList {
	return SkipList{head:&SkipListNode{next:make([]*SkipListNode, maxLevel+1, maxLevel+1)}, maxLevel:maxLevel, topLevel:0, p:p, less:less}
}

func (t SkipList) before(u *SkipListNode, v *SkipListNode) bool {  // u << v
	if v == nil {return true}	// v is tail node
	if u == nil {return false}	// u is tail node
	return t.less(u.Value(), v.Value())
} 
		
func (t SkipList) Empty() bool {
	return t.head.Next() == nil
}

func (t *SkipList) newNodeLevel() int {
	res:= 0
	for rand.Float64() < t.p && res <= t.topLevel && res < t.maxLevel {
		res++
	}
	if res > t.topLevel {
		t.topLevel = res
	}
	return res
}	

func (t *SkipList) Insert(v interface{}) {
	l:= t.newNodeLevel()
	node:= &SkipListNode{value: v, next: make([]*SkipListNode, l+1, l+1) }
	
	u:= t.head
	for level:= node.level(); level>=0; level-- {
		for {
			if t.before((*u).next[level], node) {
				u = (*u).next[level]
			} else {
				(*node).next[level] = (*u).next[level]
				(*u).next[level] = node
				break
			}
		}		 
	}	
}	

func (t SkipList) Remove(v interface{}) {
	node:= &SkipListNode{value:v}
	u:= t.head
	var uNext *SkipListNode
	for level:= t.topLevel; level>=0; level-- {
		for {
			uNext = (*u).next[level]
			if t.before(uNext, node) {
				u = uNext
			} else { 
				if !t.before(node, uNext) {
					(*u).next[level] = (*uNext).next[level]
				} 
				break	// jump to the previous level
			}	
		}		 
	}	
}	

func (t SkipList) Find(v interface{}) (*SkipListNode, error) {
	u:= t.head
	var uNext *SkipListNode
	node:= &SkipListNode{value:v}
	for level:= t.topLevel; level>=0; level-- {
		for {
			uNext = (*u).next[level]
			if t.before(uNext, node) {
				u = uNext
			} else 
			if t.before(node, uNext) {
				break	// jump to the previous level
			} else {
				return uNext, nil 
			}
		}		 
	}	
	return nil, errors.New("no such element")
}	

func (t SkipList) Traverse(f func(v interface{})) {
	p:= t.head
	for {
		p = p.Next()
		if p == nil { break }
		f(p.Value()) 
	}
}

func main() {
	// Initialize random number generator.
	rand.Seed(time.Now().UTC().UnixNano())

	show:= func(v interface{}) {
		fmt.Printf("%5d", v.(int))
	}
	
	list := NewSkipList(4, 0.7, func(p, q interface {}) bool {
		return p.(int) > q.(int)
	})
	
	list.Insert(2) 			//    2
	list.Traverse(show)
	fmt.Println() 
	list.Insert(4)
	list.Traverse(show)		//    4    2
	fmt.Println()
	list.Insert(1)
	list.Insert(8)
	list.Insert(5)
	list.Traverse(show)		//    8    5    4    2    1
	fmt.Println()
	
	p, err:= list.Find(8)	
	fmt.Println(p, err )	//	&{[0xc00006e5a0 0xc00006e5a0] 8} <nil>
	p, err = list.Find(5)
	fmt.Println(p, err ) 	//	&{[0xc00006e510 0xc00006e510 0xc00006e510 <nil>] 5} <nil>
	p, err = list.Find(1)
	fmt.Println(p, err )	//	&{[<nil> <nil> <nil>] 1} <nil>
	p, err = list.Find(10)
	fmt.Println(p, err )	//	<nil> no such element
	p, err = list.Find(7)
	fmt.Println(p, err )	//	<nil> no such element
	p, err = list.Find(-3)
	fmt.Println(p, err )	//	<nil> no such element
	
 	list.Remove(5)			
	list.Traverse(show)		//    8    4    2    1
	fmt.Println()
	list.Remove(8)
	list.Traverse(show)		//    4    2    1
	fmt.Println()
	list.Remove(1)
	list.Traverse(show)		//    4    2
	fmt.Println()
}
