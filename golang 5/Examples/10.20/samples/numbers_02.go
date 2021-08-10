package main

import 	(
    "fmt"
    "errors"
    "os"
    "math"
)

type lmnt struct {
    n int
    next *lmnt
}

type queue struct {
    head *lmnt
    tail *lmnt
}

func initQueue() queue  {
    return queue{nil, nil}
}

func (q *queue) Add (n int) {
	if (*q).Empty()  {
        (*q).tail = &lmnt{n, nil}
        (*q).head = (*q).tail
    } else {
        (*(*q).tail).next = &lmnt{n, nil}
        (*q).tail = (*(*q).tail).next
    }
}	

func (q *queue) Remove () error   {
    if (*q).Empty()  {
        return errors.New("Attempt to remove from empty queue")
    }  else  {
        (*q).head = (*(*q).head).next
        return nil
    }    
}	

func (q queue) Max () int {
    res := math.MinInt64
    runner:= q.head
    for runner != nil  {
        if (*runner).n > res  { res = (*runner).n }
        runner = (*runner).next
    }  		
	return res	
}	

func (q queue) Empty() bool  {
    return q.head == nil
}	

func main() {  
	f, err := os.Open("numbers.dat") 
    if err != nil {
        return
    }
    defer f.Close()
    var k int
    _, err = fmt.Fscanf(f, "%d\n", &k)
    if err != nil  { return }
    var x int
    q:= initQueue()
    for i:= 0; i < k; i++  {
        _, err := fmt.Fscanf(f, "%d\n", &x)
        if err !=nil  { break }
        q.Add(x)
    }    
    for {
        fmt.Println(q.Max())
        _, err := fmt.Fscanf(f, "%d\n", &x)
        if err != nil  { break }
        q.Add(x)
        q.Remove()
    }    
}
