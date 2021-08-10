package main

import 	(
    "fmt"
    "math"
)

const  (
    PlusInfinity = math.MaxInt64
    MinusInfinity = math.MinInt64
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
	return list{&lmnt{MinusInfinity, &lmnt{PlusInfinity, nil}}}
}

func (s list) Print()  {
    for runner:= (*s.head).next; (*runner).x < PlusInfinity; runner = (runner).next  {
        fmt.Print((*runner).x, " ")
    }
    fmt.Println()
}	

func (s *list) Insert1(num int)  {
    runner:= (*s).head
	for (*(*runner).next).x < num  {
        runner = (*runner).next
    }
    (*runner).next = &lmnt{num, (*runner).next}
}

func (s *list) Insert2(num int)  {
    runner:= (*s).head
    runner2:= (*runner).next
	for (*runner2).x < num  {
        runner, runner2 = runner2, (*runner2).next
    }
    (*runner).next = &lmnt{num, runner2}
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
