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
	p := new(lmnt)
	(*p).x = k
	(*p).next = (*t).head
	(*t).head = p
	
	//(*t).head = &lmnt{k, (*t).head, g}  <==> 27-30
}

func (t *list) DeleteX(k data) {
	var run1, run2 *lmnt
	for run1 = (*t).head; run1 != nil ; run1 = (*run1).next {
		if (*run1).x == k {
			if run2 == nil {
				(*t).head = (*run1).next
			} else {
				(*run2).next = (*run1).next
			}
		} else {
			run2 = run1
		}
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
	x = data(20)
	
	fmt.Println("Old list:")
	s.Print()
	s.DeleteX(x)
	fmt.Println("\nNew list:")
	s.Print()
	fmt.Printf("\nDeleted number was %d", x)
	
}
