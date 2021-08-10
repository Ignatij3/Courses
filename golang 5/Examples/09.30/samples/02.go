package main

import "fmt"

func inc(val *int, delta int) {  
    *val += delta
}

func main() {  
    a := 12345
    fmt.Println("a =", a)	// a = 12345
    inc(&a, 5)
    fmt.Println("a =", a)	// a = 12350
    b := &a
    inc(b, -3)
    fmt.Println("a =", a)	// a = 12347
}

