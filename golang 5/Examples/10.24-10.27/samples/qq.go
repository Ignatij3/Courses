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
        len int
    }
)

func initList() list  {
	return list{nil, 0}
}

func (t *list) Add(k int) {
    (*t).len++
    p:= new(lmnt)
    (*p).x = k
    (*p).next = (*t).head
    (*t).head = p
    
    // (*t).head = &lmnt{k, (*t).head}  <==> 28-30
}

func (t list) Print() {
    for runner:= t.head; runner != nil; runner = (*runner).next {
	   fmt.Println((*runner).x)
    }   	
}	

func (t list) Max() int {
    max:= (*t.head).x
    for runner:= t.head; runner != nil; runner = (*runner).next {
	   if (*runner).x > max  { max = (*runner).x }
    }   	
    return max
}	

func main() {  
	f, err := os.Open("numbers.dat") 
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
    fmt.Println(s.Max())
}
