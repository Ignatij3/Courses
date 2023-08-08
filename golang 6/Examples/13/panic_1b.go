package main

import (  
    "fmt"
)

func slicePanic() {  
    n := []int{5, 7, 4}
    fmt.Println(n[4])
    fmt.Println("normally returned from slicePanic")
}
func main() {  
    slicePanic()
    fmt.Println("normally returned from main")
}

//	panic: runtime error: index out of range [4] with length 3
//
//	goroutine 1 [running]:
//	main.slicePanic()
//	        I:/Work. GO/work_VI/13/samples/panic_1b.go:9 +0x1d
//	main.main()
//	        I:/Work. GO/work_VI/13/samples/panic_1b.go:13 +0x19
//	exit status 2
