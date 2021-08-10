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
	p := new(lmnt)
	(*p).x = k
	(*p).next = (*t).head
	(*t).head = p
	
	//(*t).head = &lmnt{k, (*t).head, g}  <==> 27-30
}

func (t *list) CopyList() list {
	r := initList()
	slice := make([]*lmnt, 0)
	
	for runner:= t.head; runner != nil; runner = (*runner).next {
	   slice = append(slice, runner)
    }
	
	for s := len(slice) - 1; s >= 0; s-- {
		r.Add((*slice[s]).x)
	}
	return r
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
	
	r := s.CopyList()
	
	fmt.Println("Old list:")
	s.Print()
	fmt.Println("\nNew list:")
	r.Print()
}
