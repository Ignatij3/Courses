package main

import 	(
	"fmt"
	"os"
	"math/rand"
	"time"
)

type (
	data int
	lmnt struct {
		x data
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

func (t *list) Add(k data) {
	(*t).head = &lmnt{k, (*t).head}
}

func (t *lmnt) DeleteX(k data, t2 *lmnt) {
	if t != nil {
		if (*t).x == k {(*t2).next = (*t).next}
		(*t).next.DeleteX(k, t)
	}
}

func main() {  
	var x data
	
	f, err := os.Open("../../Files/numbers.dat") 
	if err != nil {return}
	defer f.Close()
	
	s := initList()
	for {
		_, err := fmt.Fscanf(f, "%d\n", &x)
		if err != nil {break}
		s.Add(x)
	}
	
	rand.Seed(time.Now().UnixNano())
	x = data(rand.Intn(60))
	//x = data(20)
	
	fmt.Println("Old list:")
	s.Print()
	s.head.next.DeleteX(x, s.head)
	fmt.Println("\nNew list:")
	s.Print()
	fmt.Printf("\nDeleted number was %d", x)
	
}
