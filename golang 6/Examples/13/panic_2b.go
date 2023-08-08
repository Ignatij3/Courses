package main

import "fmt"

func fullName(firstName *string, lastName *string) {  
    defer fmt.Println("deferred call in fullName")
    if firstName == nil {
        panic("runtime error: first name cannot be nil")
    }
    if lastName == nil {
        panic("runtime error: last name cannot be nil")
    }
    fmt.Printf("%s %s\n", *firstName, *lastName)
    fmt.Println("returned normally from fullName")
}

func main() {  
    defer fmt.Println("deferred call in main")
    firstName := "Vasja"
    fullName(&firstName, nil)
    fmt.Println("returned normally from main")
}
/*
deferred call in fullName
deferred call in main
panic: runtime error: last name cannot be nil

goroutine 1 [running]:
main.fullName(0x1ec1f0a0108, 0x60)
        I:/Work. GO/work_VI/13/samples/panic_2b.go:11 +0x197
main.main()
        I:/Work. GO/work_VI/13/samples/panic_2b.go:20 +0x8c
exit status 2
*/
