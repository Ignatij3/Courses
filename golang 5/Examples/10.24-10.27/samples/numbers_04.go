package main

import 	(
    "fmt"
    "errors"
    "os"
    "math"
)

type (
    data struct {
		val int
		pos int
    }
    lmnt struct {
        d data
        l *lmnt
        r *lmnt
	}	
    deque struct {
        left *lmnt
        right *lmnt
    }
)

func initDeque() deque  {
	return deque{nil, nil}
}

func (dq *deque) PushLeft (d data) {
	if (*dq).Empty()  {
        (*dq).left = &lmnt{d, nil, nil}
        (*dq).right = (*dq).left
    } else {
        (*(*dq).left).l = &lmnt{d, nil, (*dq).left}
        (*dq).left = (*(*dq).left).l
    }
}	

func (dq *deque) PushRight (d data) {
	if (*dq).Empty()  {
        (*dq).left = &lmnt{d, nil, nil}
        (*dq).right = (*dq).left
    } else {
        (*(*dq).right).r = &lmnt{d, (*dq).right, nil}
        (*dq).right = (*(*dq).right).r
    }
}	

func (dq deque) GetLeft() data {
    if dq.Empty()  {
		return data{math.MinInt64, -1}
    } else {
        return (*dq.left).d
    }    
}	

func (dq deque) GetRight() data {
    if dq.Empty()  {
		return data{math.MinInt64, -1}
    } else {
        return (*dq.right).d
    }    
}	

func (dq *deque) RemoveLeft () error {
    if (*dq).Empty()  {
        return errors.New("Attempt to remove from empty deque")
    }
    if (*dq).left == (*dq).right  {
		(*dq).left, (*dq).right = nil, nil
	} else {	
        (*dq).left = (*(*dq).left).r
    }
    return nil
}	

func (dq *deque) RemoveRight () error {
    if (*dq).Empty()  {
        return errors.New("Attempt to remove from empty deque")
    } 
    if (*dq).left == (*dq).right  {
		(*dq).left, (*dq).right = nil, nil
	} else {	
        (*dq).right = (*(*dq).right).l
    }
    return nil
}	

func (dq deque) Empty() bool  {
    return dq.left==nil
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
        win.PushRight(data{x, i})
    }    
    fmt.Println(win.GetLeft().val)

    for i:= k; ; i++ {
        _, err := fmt.Fscanf(f, "%d\n", &x)
        if err != nil  {fmt.Println(err); break}
        if win.GetLeft().pos == i-k {
            win.RemoveLeft()
		}	
        
        for !win.Empty() &&  win.GetRight().val <= x {
            win.RemoveRight()
        }
        win.PushRight(data{x, i})

        fmt.Println(win.GetLeft().val)
    }    
}
