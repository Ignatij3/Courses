package main

import 	(
    "fmt"
)

type  ( 
    list struct {
        head *lmnt
    }
    lmnt struct {
        x int
        next *lmnt
    }
)

func initList() list  {
	return list{nil}
}

func (s list) Empty() bool  {
	return s.head == nil
}

func (s list) Len() int  {
    if s.Empty()  {
        return 0
    }  else  {
        return 1 + list{(*s.head).next}.Len()
	}	
}	

func (s list) PrintThere()  {
    if s.Empty()  {
		fmt.Println()
    }  else  {
        fmt.Print((*s.head).x, " ")
        list{(*s.head).next}.PrintThere()
	}	
}	

func (s list) PrintBack()  {
    if !s.Empty()  {
        list{(*s.head).next}.PrintBack()
        fmt.Print((*s.head).x, " ")
	}
}	

func (s *list) Add(n int)  {
    (*s).head = &lmnt{x: n, next: (*s).head}
}

func (s list) AddBack(n int) list {
    if s.Empty() {
        return list{head: &lmnt{x: n, next: nil}}
    } else { 
        tail:= list{(*s.head).next}.AddBack(n) 
        (*s.head).next = tail.head 
        return s
    }		
}

func main() {  
    l:= initList()
    for _, x := range []int{2, 5, 4, 1, 2, 8} {
        l.Add(x)
    }    
    fmt.Println(l.Len())    // 6
    l.PrintThere()          // 8 2 1 4 5 2
    l.PrintBack()           // 2 5 4 1 2 8
    fmt.Println()
    l = initList()
    for _, x := range []int{2, 5, 4, 1, 2, 8} {
        l = l.AddBack(x)
    }    
    fmt.Println(l.Len())    // 6
    l.PrintThere()          // 2 5 4 1 2 8
    l.PrintBack()           // 8 2 1 4 5 2
    fmt.Println()
}
