package main

import 	(
    "fmt"
    "errors"
    "os"
    "math"
)

type queue []int

func initQueue() queue  {
	return make([]int, 0, 0)
}

func (q *queue) Add (n int) {
    *q = append(*q, n)
}	

func (q queue) Max () int {
    res := math.MinInt64
    for _, x := range q  {
		if x > res  { res = x }
	}
	return res	
}	


func (q *queue) Remove () error   {
    if (*q).Empty()  {
        return errors.New("Attempt to remove from empty queue")
    }  else  {
        *q = (*q)[1:]
        return nil
    }
}	


func (q queue) Empty() bool  {
    return len(q) == 0
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
