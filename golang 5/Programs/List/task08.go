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
	p := new(lmnt)
	(*p).x = k
	(*p).next = (*t).head
	(*t).head = p
	
	//(*t).head = &lmnt{k, (*t).head, g}  <==> 27-30
}

func (t *list) ConcatList(t2 *list) {
	lmntL := (*t).head
	for ; (*lmntL).next != nil; lmntL = (*lmntL).next {}
	(*lmntL).next = (*t2).head
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
	s.ConcatList(&r)
	fmt.Println()
	fmt.Println("\nMerged lists:")
	s.Print()
	fmt.Println()
	fmt.Println("\nMerged lists:")
	r.Print()
}
