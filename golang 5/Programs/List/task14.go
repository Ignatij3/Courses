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

func (t *list) Reverse() { //Собрать в слайс и переставить 1 - n-1, 2, n-2...
	slice := make([]*lmnt, 0)
	for runner:= t.head; runner != nil; runner = (*runner).next {
	   slice = append(slice, runner)
    }
    
	var r1, r2 int = 0, len(slice) - 1
    for r1 < r2 {
		(*slice[r1]).x, (*slice[r2]).x = (*slice[r2]).x, (*slice[r1]).x
		r1++
		r2--
	}
	fmt.Println((*t).head)
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
	s.Reverse()
	fmt.Println("\nNew list:")
	s.Print()
}
