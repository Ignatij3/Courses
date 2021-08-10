package main

import 	(
	"container/list"
    "fmt"
    "math"
)

const  Infinity = math.MaxInt64

func main() {
	// Create a new list and put some Infinity in it.
	l := list.New()
	l.PushFront(Infinity)
	x := 1000000000
    for i:= 3; i>0; i = (i+6)%19 {
		fmt.Print(i, " ")
        for e := l.Front(); ; e = e.Next() {
            if e.Value.(int) >= i  {
                l.InsertBefore(i, e)
                break
            }    
        }
    }
    fmt.Println()
    l.Remove(l.Back())

	// Iterate through list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() {
        fmt.Print(e.Value, " ") 
	}
    fmt.Println()

//  Output:
//  3 9 15 2 8 14 1 7 13
//  1 2 3 7 8 9 13 14 15
}
