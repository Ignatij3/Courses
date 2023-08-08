package main

import (  
    "fmt"
)

func appendStr() func(string) string {  
    t := "Hello"
    c := func(s string) string {
        t = t + " " + s
        return t
    }
    return c
}

func main() {
    a := appendStr()
    b := appendStr()
    fmt.Println(a("World"))		// Hello World
    fmt.Println(b("Everyone"))	// Hello Everyone
    fmt.Println(a("Gopher"))	// Hello World Gopher
    fmt.Println(b(":)"))		// Hello Everyone :)
    
    for i:= 0; i<5; i++ {
			fmt.Println(a("!"))
	}	
	//	Hello World Gopher !
	//	Hello World Gopher ! !
	//	Hello World Gopher ! ! !
	//	Hello World Gopher ! ! ! !
	//	Hello World Gopher ! ! ! ! !
}
