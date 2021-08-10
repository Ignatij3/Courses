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

func (t *list) CountX(k int) int {
	var count int
	for lmntL := (*t).head; lmntL != nil ; lmntL = (*lmntL).next {
		if lmntL.x == k{count++}
	}
	return count
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
	//x = 7
	s.Print()
	fmt.Printf("\nКол-во вхождений числа %d в слайс - %v", x, s.CountX(x))
}
