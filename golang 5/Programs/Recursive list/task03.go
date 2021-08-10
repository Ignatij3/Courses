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

func (t *lmnt) CompareList(t2 *lmnt) bool {
	if t != nil && t2 != nil {
		if (*t).x != (*t2).x || !t.next.CompareList(t2.next) {return false} else {return true}
	}
	return true
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
	
	same := s.head.CompareList(r.head)
	if same {
		fmt.Println("Lists are the same")
	} else {
		fmt.Println("Lists are different")
	}
}
