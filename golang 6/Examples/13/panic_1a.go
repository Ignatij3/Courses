package main

import (  
    "fmt"
)

func fullName(firstName *string, lastName *string) *string {  
    s:= *firstName + " " + *lastName
    fmt.Println("returned normally from fullName")
	return &s
}

func main() {  
    firstName, lastName := "Vasja", "Pupkin"
    fmt.Println(*fullName(&firstName, &lastName), "\n\n")
    fmt.Println(*fullName(&firstName, nil), "\n\n")
    fmt.Println("returned normally from main")
}

//	returned normally from fullName
//	Vasja Pupkin
//
//
//	panic: runtime error: invalid memory address or nil pointer dereference
//	[signal 0xc0000005 code=0x0 addr=0x0 pc=0x24c5a4]
//
//	goroutine 1 [running]:
//	main.fullName(0xc000077f10, 0x0)
//	        I:/Work. GO/work_VI/13/samples/panic_1a.go:8 +0x44
//	main.main()
//	        I:/Work. GO/work_VI/13/samples/panic_1a.go:16 +0xd5
//	exit status 2
