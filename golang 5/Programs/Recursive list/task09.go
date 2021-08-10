package main

import 	(
	"fmt"
	"os"
)

type (
	lmnt struct {
		x int
		next *lmnt
	}	
	list struct {
		head *lmnt
	}
)

func initList() *list  {
	return &list{nil}
}

func (t list) Print() {
    for runner:= t.head; runner != nil; runner = (*runner).next {
	   fmt.Println((*runner).x)
    }
}

func (t *list) Add(k int) {
	(*t).head = &lmnt{k, (*t).head}
}

func (p0 *lmnt) Reverse(p1, p2 *lmnt) *lmnt {
	(*p1).next = p0
	if p2 != nil {
		return p1.Reverse(p2, (*p2).next)
	}
	return p1
}

func main() {
	var x int
	
	f, err := os.Open("../../Files/numbers.dat")
	if err != nil {return}
	defer f.Close()
	
	s := initList()
	for {
		_, err := fmt.Fscanf(f, "%d\n", &x)
		if err != nil {break}
		s.Add(x)
	}
	
	fmt.Println("Old list:")
	s.Print()
	
	start := (*s).head
	(*start).next, start = nil, start.Reverse((*start).next, (*(*start).next).next)
	(*s).head = start
	
	fmt.Println("\nNew list:")
	s.Print()
}
