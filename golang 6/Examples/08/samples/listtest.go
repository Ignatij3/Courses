package main

import (
	"container/list"
	"fmt"
)

func main() {
	var l list.List
	fmt.Printf("%v\n", l) // {{<nil> <nil> <nil> <nil>} 0}
	l.Init()
	fmt.Printf("%v\n", l) // {{0xc000074480 0xc000074480 <nil> <nil>} 0}
	ll := list.New()
	fmt.Printf("%v\n", *ll) // {{0xc000074510 0xc000074510 <nil> <nil>} 0}
	e := l.PushFront(777)
	fmt.Printf("%v\n", l)  // {{0xc000074570 0xc000074570 <nil> <nil>} 1}
	fmt.Printf("%v\n", *e) // {0xc000074480 0xc000074480 0xc000074480 777}
	l.PushFront(1234)
	fmt.Println(*l.Front(), *l.Back())
	// {0xc000074570 0xc000074480 0xc000074480 1234} {0xc000074480 0xc000074600 0xc000074480 777}
}
