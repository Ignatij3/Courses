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

func (t *lmnt) ConcatList(t2 *lmnt) {
	if t.next == nil {t.next = t2} else {t.next.ConcatList(t2)}
}

func main() {  
	var x, neg int
	
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
	r := initList()
	
	for k := 0; k < 10; k++ {
		if rand.Intn(4) == 0 {
			neg = rand.Intn(100)
			neg = -neg
			r.Add(neg)
		} else {
			r.Add(rand.Intn(100))
		}
	}
	fmt.Println()
	fmt.Println("Old list:")
	s.Print()
	fmt.Println()
	fmt.Println("\nNew list:")
	r.Print()
	s.head.ConcatList(r.head)
	fmt.Println()
	fmt.Println("\nMerged lists:")
	s.Print()
}
