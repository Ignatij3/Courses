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

func initList() list  {
	return list{nil}
}

func (t list) Print() {
    for runner:= t.head; runner != nil; runner = (*runner).next {
	   fmt.Println((*runner).x)
    }   	
}

func (t *list) Add(k int) {
	(*t).head = &lmnt{k, (*t).head}
}

func (t *lmnt) CopyList(r *lmnt) {
	if r != nil && t != nil {
		(*r).x = (*t).x
		t.next.CopyList(r.next)
	}
}

func main() {  
	var x int
	
	f, err := os.Open("../../Files/numbers.dat") 
	if err != nil {return}
	defer f.Close()
	
	s := initList()
	r := initList()
	for {
		_, err := fmt.Fscanf(f, "%d\n", &x)
		if err != nil {break}
		s.Add(x)
	}
	
	r.head = s.head
	s.head.CopyList(r.head)
	
	fmt.Println("Old list:")
	s.Print()
	fmt.Println("\nNew list:")
	r.Print()
}
