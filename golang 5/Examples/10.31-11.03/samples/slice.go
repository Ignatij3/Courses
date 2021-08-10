package main

import 	(
    "fmt"
    "math"
)

const
    Infinity = math.MaxInt64
    
type   
    list []int

func initList() list  {
	return list{Infinity}
}

func (s list) Print()  {
    for i, x:= range s  {
        if i == len(s) -1  { break }
        fmt.Print(x, " ")
    }
    fmt.Println()
}	

func (s *list) Insert(num int)  {
	for i, x:= range *s  {
        if x >= num  { 
            *s = append(*s, 0)
            copy((*s)[i+1:], (*s)[i:])
            (*s)[i] = num
            break 
        }
    }
}

func main() {  
    l:= initList()
    l.Insert(2)
    l.Insert(5)
    l.Insert(4)
    l.Insert(1)
    l.Insert(2)
    l.Insert(8)
    l.Print()
}
