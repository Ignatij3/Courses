package main

import 	(
    "fmt"
    "errors"
    "os"
    "math"
)

type (
    lmnt struct {
		val int
		pos int
	}	
    deque []lmnt
)

func initDeque() deque  {
	return make([]lmnt, 0, 0)
}

func (dq *deque) PushLeft (w lmnt) {
    *dq = append([]lmnt{w}, (*dq)...)
}	

func (dq *deque) PushRight (w lmnt) {
    *dq = append(*dq, w)
}	

func (dq deque) GetLeft() lmnt {
    if dq.Empty()  {
		return lmnt{math.MinInt64, -1}
    } else {
        return dq[0]
    }    
}	

func (dq deque) GetRight() lmnt {
    if dq.Empty()  {
		return lmnt{math.MinInt64, -1}
    } else {
        return dq[len(dq)-1]
    }    
}	

func (dq *deque) RemoveLeft () error {
    if (*dq).Empty()  {
        return errors.New("Attempt to remove from empty deque")
    }  else  {
        *dq = (*dq)[1:]
        return nil
    }
}	

func (dq *deque) RemoveRight () error {
    if (*dq).Empty()  {
        return errors.New("Attempt to remove from empty deque")
    } else {
        *dq = (*dq)[:len(*dq)-1]
        return nil
    }
}	

func (dq deque) Empty() bool  {
    return len(dq) == 0
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
    win:= initDeque()
    for i:= 0; i < k; i++  {
        _, err := fmt.Fscanf(f, "%d\n", &x)
        if err !=nil  { break }

        for !win.Empty() &&  win.GetRight().val <= x {
            win.RemoveRight()
        }
        win.PushRight(lmnt{x, i})
    }    
    fmt.Println(win.GetLeft().val)

    for i:= k; ; i++ {
        _, err := fmt.Fscanf(f, "%d\n", &x)
        if err != nil  { break }
        
        if win.GetLeft().pos == i-k {
            win.RemoveLeft()
		}	
        
        for !win.Empty() &&  win.GetRight().val <= x {
            win.RemoveRight()
        }
        win.PushRight(lmnt{x, i})

        fmt.Println(win.GetLeft().val)
    }    
}
