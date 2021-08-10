package main

import 	(
    "fmt"
    "errors"
    "math"
)

type (
    lmnt int
    ring struct {
		row []lmnt
		currentPos int
	}	
)

func initRing(len int) ring {
	var r ring
	r.row = make([]lmnt, len, len)
	r.currentPos = 0
	return r
}

func (r ring) Empty() bool  {
    return r.Length() == 0
}	

func (r ring) Length() int  {
    return len(r.row)
}	

func (r ring) GetCurrent() lmnt {
    if r.currentPos >= r.Length() || r.currentPos < 0 {
		return math.MinInt64
    } else {
        return r.row[r.currentPos]
    }    
}	

func (r *ring) RemoveCurrent () error {
    if r.currentPos >= r.Length() || r.currentPos < 0 {
        return errors.New("Invalid removing")
    }  else  {
        (*r).row = Delete((*r).row, (*r).currentPos)
        return nil
    }
}	

func Delete(s []lmnt, i int) []lmnt {
    return s[:i+copy(s[i:], s[i+1:])]
}    

func (r *ring) Forward (step int) error {
    if (*r).Empty() {
        return errors.New("Invalid step")
    }  else  {
        (*r).currentPos = ((*r).currentPos + step) % (*r).Length()
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
    for i, _ := range r.row  {
        r.row[i] = lmnt(i+1)
    }    

    for i:= 0; i < N; i++ {
		r.Forward(M-1)
		fmt.Println(r.GetCurrent())
		r.RemoveCurrent()
    }    
}

