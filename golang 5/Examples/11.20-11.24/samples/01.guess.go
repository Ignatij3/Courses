package main

import "fmt"

func guess(left, right int)  {
	var answer string
    c:= (left + right)/2
    fmt.Println("my shot: ", c)
    fmt.Scanln(&answer) 
    switch answer {
	case "<": 
	    guess(left, c-1)   	
    case ">":
        guess(c+1, right)
    case "=":
        fmt.Println("Congratulations!")
        return
    default:
        fmt.Println("Incorrect answer")
        fmt.Println("Enter \"<\" if your secret number <", c)    	
        fmt.Println("Enter \">\" if your secret number >", c)    	
        fmt.Println("Enter \"=\" if your secret number is", c)
        guess(left, right)
    }
}	

func main()  {
    guess(1, 100)
}
