package main

import "fmt"

func hello() *string {  
    s := "hello, world"
    return &s
}

func main() {  
    fmt.Println(*hello())	// hello, world
    hi := hello()
    fmt.Printf("Тип переменной hi %T,\n", hi)		// Тип переменной hi *string,
    fmt.Printf("она содержит указатель на строку \"%s\"\n", *hi)		
								// она содержит указатель на строку "hello, world"
}
								
