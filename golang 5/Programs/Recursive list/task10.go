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

func (m *lmnt) CountX(k int) int {
	if m != nil {
		n := (*m).next.CountX(k)
		if (*m).x == k {return n + 1} else {return n}
	}
	return 0
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
	
	rand.Seed(time.Now().UnixNano())
	x = rand.Intn(55)
	//x = 25
	s.Print()
	fmt.Printf("\nКол-во вхождений числа %d в слайс - %v", x, (*s).head.CountX(x))
}
