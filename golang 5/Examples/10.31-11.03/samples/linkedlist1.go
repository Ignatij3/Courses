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

func (s list) Print()  {
    for runner:= s.head; runner != nil; runner = (runner).next  {
        fmt.Print((*runner).x, " ")
    }
    fmt.Println()
}	

func (s *list) Insert1(num int)  {
    runner:= (*s).head
    if runner == nil  {
        (*s).head = &lmnt{num, nil}
        return
    }
    if num <= (*runner).x {
        (*s).head = &lmnt{num, runner}
        return
    }    
    for ((*runner).next!=nil) && ((*(*runner).next).x < num)  {
        runner = (*runner).next
    }
    (*runner).next = &lmnt{num, (*runner).next}
}

func (s *list) Insert2(num int)  {
	var runner, runner2 *lmnt
    runner2 = (*s).head
    if runner2 == nil  {
        (*s).head = &lmnt{num, nil}
        return
    }
    if num <= (*runner2).x {
        (*s).head = &lmnt{num, runner2}
        return
    }    
	for (runner2 != nil) && ((*runner2).x < num)  { 
	  runner, runner2 = runner2, (*runner2).next
	}
    (*runner).next = &lmnt{num, (*runner).next}
}

func main() {  
    l:= initList()
    l.Insert2(2)
    l.Insert2(5)
    l.Insert2(4)
    l.Insert2(1)
    l.Insert2(2)
    l.Insert2(8)
    l.Print()
}
