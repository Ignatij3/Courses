package main

import 	(
    "fmt"
    "errors"
    "os"
)

type stack []rune

func (ps *stack) Pop () (rune, error)   {
    if (*ps).Empty()  {
        return 0, errors.New("Pop from empty stack")
    }  else  {
        c:= (*ps)[len(*ps)-1]
        *ps = (*ps)[:len(*ps)-1]
        return c, nil
    }
}	

func (ps *stack) Push (r rune) {
    *ps = append(*ps, r)
}	

func (s stack) Empty() bool  {
    return len(s) == 0
}	

func check(brackets string) bool  {
    var s stack
    for _, b := range []rune(brackets)  {
        switch b  {
        case '(', '[', '{':  
            s.Push(b)
        case ')':  
            if top, err:= s.Pop(); err != nil || top != '('  { return false }
        case ']':  
            if top, err:= s.Pop(); err != nil || top != '['  { return false }
        case '}':  
            if top, err:= s.Pop(); err != nil || top != '{'  { return false }
        default:
            return false
        }    
    }	
    return s.Empty()
}	 

func main() {  
	f, err := os.Open("brackets.dat") 
    if err != nil {
        return
    }
    defer f.Close()
    var line string
    for  {
        _, err := fmt.Fscanf(f, "%s\n", &line)
        if err !=nil  { break }
        fmt.Println(line, check(line))
    }    
}
