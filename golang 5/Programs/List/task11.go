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

func (t *list) Truncate(k int) {
	/*for lmntL := t.head; lmntL != nil ; lmntL = (*lmntL).next {
		if lmntL.x == k{found = true; p = lmntL.pos}
	}*/
	listk := (*t).head
	for i := 0; i < k; i++ {
		listk = (*listk).next
	}
	(*listk).next = nil
}

func main() {  
	var x, lenS int
	
	f, err := os.Open("../../Files/numbers.dat") 
	if err != nil {return}
	defer f.Close()
	
	s := initList()
	for ; ; lenS++ {
		_, err := fmt.Fscanf(f, "%d\n", &x)
		if err != nil {break}
		s.Add(x)
	}
	
	rand.Seed(time.Now().UnixNano())
	x = rand.Intn(20)
	//x = 26
	
	if x >= lenS {
		fmt.Printf("Number is too large (%d)", x)
	} else {
		fmt.Println("Before:")
		s.Print()
		s.Truncate(x)
		fmt.Println("\nAfter:")
		s.Print()
		fmt.Println("\nCut position -", x)
	}
}
