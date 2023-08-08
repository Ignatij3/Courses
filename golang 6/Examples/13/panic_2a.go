package main

import "fmt"

func fullName(firstName *string, lastName *string) *string {  
    if firstName == nil {
        panic("runtime error: first name cannot be nil")
    }
    if lastName == nil {
        panic("runtime error: last name cannot be nil")
    }
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
//	panic: runtime error: last name cannot be nil
//
//	goroutine 1 [running]:
//	main.fullName(0x616700, 0xc000006018)
//	        I:/Work. GO/go_VI_semester_repa/13/samples/panic_2a.go:10 +0xf4
//	main.main()
//	        I:/Work. GO/go_VI_semester_repa/13/samples/panic_2a.go:20 +0xd5
//	exit status 2
