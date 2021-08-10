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

func (t *list) CompareList(t2 *list) bool {
	var identical bool = true
	for run1, run2 := (*t).head, (*t2).head; run1 != nil && run2 != nil; run1, run2 = (*run1).next, (*run2).next {
		if (*run1).x != (*run2).x {identical = false}
	}
	return identical
}

func main() {  
	var x, neg data
	
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
			neg = data(rand.Intn(100))
			neg = -neg
			r.Add(neg)
		} else {
			r.Add(data(rand.Intn(100)))
		}
	}
	
	same := s.CompareList(&r)
	if same {
		fmt.Println("Lists are the same")
	} else {
		fmt.Println("Lists are different")
	}
}
