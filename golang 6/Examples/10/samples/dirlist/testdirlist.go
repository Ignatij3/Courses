package main

import (
	"fmt"
	"container/dirlist"
)

func main() {
	show:= func(v interface{}) {
		fmt.Printf("%5d", v.(int))
	}
	subj:= func (n int) func(v interface{}) bool {
		return 	func(v interface{}) bool {
					return v.(int) == n
				}
	}			

	l := dirlist.NewDirectedList()
	r, err:= l.Find(subj(8))		
	fmt.Println(r, err)	            	// <nil> no such element
	l.Traverse(show)					//
	fmt.Println("Traverse empty list")	// Traverse empty list  
	l.PushFront(7)
	l.Traverse(show)					//    7
	fmt.Println()
	l.PushBack(11)
	l.Traverse(show)					//    7   11 
	fmt.Println()
	l.PushBack(3)
	l.Traverse(show)					//    7   11    3
	fmt.Println()
	l.PushFront(5)
	l.PushFront(3)
	l.PushBack(8)
	l.Reverse(show)    					//    8    3   11    7    5    3
	fmt.Println()
	r, err = l.Find(subj(8))			//	&{<nil> 8} <nil>
	fmt.Println(r, err)	              
	r, err = l.Find(subj(7))        	//  &{0xc0000040f0 7} <nil>  
	fmt.Println(r, err)	                
	r, err = l.Find(subj(3))			//	&{0xc0000042b8 3} <nil>
	fmt.Println(r, err)	
	r, err = l.Find(subj(4))			//	<nil> no such element
	fmt.Println(r, err, "\n\n\n")	
	
	ls := dirlist.NewSortList(func(p, q *dirlist.Node) bool {
		return (*p).Value().(int) > (*q).Value().(int)
	})
	ls.Insert(2) 						//    2
	ls.Traverse(show)
	fmt.Println()
	ls.Insert(4)
	ls.Traverse(show)					//    4    2
	fmt.Println()
	ls.Insert(1)
	ls.Insert(8)
	ls.Insert(5)
	ls.Traverse(show)					//    8    5    4    2    1

	fmt.Println()
	ls.Reverse(show)					//    1    2    4    5    8
	fmt.Print("\n\n\n")

	ls = dirlist.NewSortList(func(p, q *dirlist.Node) bool {
		return (*p).Value().(int) < (*q).Value().(int)
	})
	ls.PushFront(3)
	ls.PushBack(8)
	ls.PushFront(9)
	ls.PushFront(1)
	ls.PushBack(3)
	ls.PushBack(0)
	ls.PushFront(5)
	ls.PushBack(2)
	ls.Traverse(show)		//    5    1    9    3    8    3    0    2
	fmt.Println()
	ls.Sort()
	ls.Traverse(show)		//    0    1    2    3    3    5    8    9
	fmt.Println()
}
