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

func (t *lmnt) First(k data) int {
	if t != nil {
		if t.x == k {return 0}
		return t.next.First(k) + 1
	}
	return -14
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
	//x = data(20)
	s.Print()
	fmt.Println()
	fmt.Println(s.head.First(x))
	fmt.Printf("Number was - %d", x)
}
