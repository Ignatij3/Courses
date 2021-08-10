package main

import 	(
	"fmt"
	"os"
	"math/rand"
	"time"
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

func (m *lmnt) SubstX(k, c int) {
	if m != nil {
		if (*m).x == k {(*m).x = c}
		(*m).next.SubstX(k, c)
	}
}

func main() {  
	var x, y int
	
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
	x, y = rand.Intn(55), rand.Intn(10000)
	//x, y = 7, 549
	fmt.Printf("Number %d will be replaced with %d\n", x, y)
	fmt.Println("Before:")
	s.Print()
	(*s).head.SubstX(x, y)
	fmt.Println("\nAfter:")
	s.Print()
}
