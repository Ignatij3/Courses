package main

import 	(
    "fmt"
    "errors"
    "os"
)

func Pop (pstack *[]rune) (rune, error)   {
    if Empty(*pstack)  {
        return 0, errors.New("")
    }  else  {
        c:= (*pstack)[len(*pstack)-1]
        *pstack = (*pstack)[:len(*pstack)-1]
        return c, nil
    }
}	

func Push (pstack *[]rune, r rune) {
    *pstack = append(*pstack, r)
}	

func Empty(stack []rune) bool  {
    return len(stack) == 0
}	

func check(brackets string) bool  {
    var stack []rune
    for _, b := range []rune(brackets)  {
        switch b  {
        case '(', '[', '{':  
            Push(&stack, b)
        case ')':  
            if top, err:= Pop(&stack); err != nil || top != '('  { return false }
        case ']':  
            if top, err:= Pop(&stack); err != nil || top != '['  { return false }
        case '}':  
            if top, err:= Pop(&stack); err != nil || top != '{'  { return false }
        default:
            return false
        }    
    }	
    return Empty(stack)
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
