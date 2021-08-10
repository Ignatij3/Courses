package main

import 	(
    "fmt"
    "errors"
    "os"
)

type  (
    node struct {  
        bracket rune
        next *node
    }
    stack *node
)    

func Pop (ps *stack) (rune, error)   {
    if Empty(*ps)  {
        return 0, errors.New("")
    }  else  {
        c:= (*(*ps)).bracket
        *ps = (*(*ps)).next
        return c, nil
    }
}	

func Push (ps *stack, r rune) {
	p:= new(node)
	(*p).bracket = r
	(*p).next = *ps
	*ps = p
}	

func Empty(s stack) bool  {
    return s == nil
}	

func check(brackets string) bool  {
    var s stack
    for _, b := range []rune(brackets)  {
        switch b  {
        case '(', '[', '{':  
            Push(&s, b)
        case ')':  
            if top, err:= Pop(&s); err != nil || top != '('  { return false }
        case ']':  
            if top, err:= Pop(&s); err != nil || top != '['  { return false }
        case '}':  
            if top, err:= Pop(&s); err != nil || top != '{'  { return false }
        default:
            return false
        }    
    }	
    return Empty(s)
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
