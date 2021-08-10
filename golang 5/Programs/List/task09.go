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

func (t *list) First(k data) int {
	var p int
	for lmntL := (*t).head; lmntL != nil ; lmntL = (*lmntL).next {
		if lmntL.x == k {return p}
		p++
	}
	return -1
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
	x = data(rand.Intn(57))
	x = data(20)
	s.Print()
	fmt.Println()
	fmt.Println(s.First(x))
	fmt.Printf("Number was - %d", x)
}
