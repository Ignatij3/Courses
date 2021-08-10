package main

import  (
    "fmt"
    "errors"
    "os"
)    

type lmnt struct {
    bracket rune
    next *lmnt
}

type stack struct {
    head *lmnt
}

func (s *stack) Push(bracket rune) {
    s.head = &lmnt{bracket, s.head}
}

func (s *stack) Pop() (rune, error) {
    if s.head == nil {
        return 0, errors.New("List is empty")
    }
    c:= s.head.bracket
    s.head = s.head.next
    return c, nil
}

func (s stack) Empty() bool  {
    return s.head == nil
}	

func initStack() stack {
    return stack{nil}
}

func check(brackets string) bool  {
    s := initStack()
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
