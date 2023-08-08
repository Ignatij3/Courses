package main

import (
	"fmt"
	"math/rand"
	"time"
	"container/skiplist"
)

func main() {
	// Initialize random number generator.
	rand.Seed(time.Now().UTC().UnixNano())

	show:= func(v interface{}) {
		fmt.Printf("%5d", v.(int))
	}
	
	list := skiplist.NewSkipList(4, 0.7, func(p, q interface {}) bool {
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
