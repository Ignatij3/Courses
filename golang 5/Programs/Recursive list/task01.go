package main

import 	(
    "fmt"
    "os"
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

func (t *list) Add(k int) {
    (*t).head = &lmnt{k, (*t).head}
}

func (t list) Print() {
    for runner:= t.head; runner != nil; runner = (*runner).next {
	   fmt.Println((*runner).x)
    }   	
}	

func (t *lmnt) Max() int {
    if (*t).next != nil {
		if (*t).x > (*t).next.Max() {
			return (*t).x
		} else {
			return (*t).next.Max()
		}
	}
	return -1
}	

func main() {  
	f, err := os.Open("../../Files/numbers.dat") 
    if err != nil {
        return
    }
    defer f.Close()
    s:= initList()
    var x int
    for {
        _, err := fmt.Fscanf(f, "%d\n", &x)
        if err != nil  { break }
        s.Add(x)
    }    
    s.Print()
    fmt.Println("\nMax", s.head.Max())
}
