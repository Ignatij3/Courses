package main

import 	(
    "fmt"
    "errors"
    "math"
)

type (
    lmnt struct {
		n int
		next *lmnt
	}	
    ring struct {
		current *lmnt
		len int
	}	
)

func initRing(len int) ring {
	return ring {nil, 0}
}

func (r ring) Empty() bool  {
    return r.Length() == 0
}	

func (r ring) Length() int  {
    return r.len
}	

func (r ring) GetCurrentValue() int {
    if r.current == nil {
		return math.MinInt64
    } else {
        return (*r.current).n
    }    
}	

func (r *ring) RemoveNext () error {
    if r.Empty() {
        return errors.New("Invalid removing")
    }  
    if r.Length() > 1 {
       (*(*r).current).next = ((*(*r).current).next).next
    } else {
    // r.Length() == 1
		(*r).current = nil
    }
    (*r).len--
    return nil
}	

func (r *ring) PopNext () (int, error) {
    if r.Empty() {
        return 0, errors.New("Invalid removing")
    }  
    val:= (*(*(*r).current).next).n
    return val, (*r).RemoveNext()
}

func (r *ring) InsertFront (value int) {
    if r.Empty() {
        (*r).current = new(lmnt)
        (*(*r).current).n = value
        (*(*r).current).next = (*r).current
    } else {
		p:= &lmnt{value, (*(*r).current).next}
        (*(*r).current).next = p
    }
    (*r).len++
}	

func (r *ring) OneStep () error {
    if (*r).Empty() {
        return errors.New("Invalid step")
    }  else  {
        (*r).current = (*(*r).current).next
        return nil
    }    
}	

func main() {  
    var N, M int
    fmt.Print("Enter N: ")
    fmt.Scanf("%d\n", &N)
    fmt.Print("Enter M: ")
    fmt.Scanf("%d\n", &M)

    r:= initRing(N)
    for i:= N; i > 0; i-- {
        r.InsertFront(i)
    }    
    
    for i:= 0; i < N; i++ {
		for j:= 0; j < (M -1)% r.Length(); j++  { r.OneStep() }
		if tmp, err:= r.PopNext(); err == nil  {fmt.Println(tmp)}
    }    
}

