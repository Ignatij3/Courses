package main

import (  
    "fmt"
    "runtime/debug"
)

func recoverFullName() {  
    if r := recover(); r != nil {
        fmt.Println("recovered from \n\t", r)
        debug.PrintStack()
    }
}

func fullName(firstName *string, lastName *string) {  
    defer recoverFullName()
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
//	recovered from
//		     runtime error: last name cannot be nil
//	goroutine 1 [running]:
//	runtime/debug.Stack()
//          C:/Program Files/Go/src/runtime/debug/stack.go:24 +0x65
//	runtime/debug.PrintStack()
//          C:/Program Files/Go/src/runtime/debug/stack.go:16 +0x19
//	main.recoverFullName()
//          I:/Work. GO/work_VI/13/samples/recover_1b.go:11 +0x79
//	panic({0x646b80, 0x678620})
//          C:/Program Files/Go/src/runtime/panic.go:1038 +0x215
//	main.fullName(0x1b1dc120a28, 0x60)
//          I:/Work. GO/work_VI/13/samples/recover_1b.go:21 +0x152
//	main.main()
//          I:/Work. GO/work_VI/13/samples/recover_1b.go:30 +0x8c
//	returned normally from main
//	deferred call in main
