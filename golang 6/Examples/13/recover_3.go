package main

import (  
    "fmt"
    "time"
)

func do1(a, b int) {  
    defer func () {  
		if r := recover(); r != nil {
			fmt.Println("recovered in do1:", r)
		}
	} ()
	mult(a, b)
	div(a, b)
}

func do2(a, b int) {  
    defer func () {  
		if r := recover(); r != nil {
			fmt.Println("recovered in do2:", r)
		}
	} ()
	go mult(a, b)
	go div(a, b)
}

func mult(x, y int) {  
	fmt.Print("mult: ")  
	time.Sleep(100 * time.Millisecond)
    fmt.Printf ("%d * %d = %d\n", x, y, x * y)
}

func div(x, y int) {  
	fmt.Print("div: ")  
	time.Sleep(150 * time.Millisecond)
    fmt.Printf ("%d / %d = %d\n", x, y, x / y)
}

func main() {  
	do1(21, 4)
	do1(15, 0)
    fmt.Println("normal resumption\n\n")
	do2(21, 4)
	time.Sleep(500 * time.Millisecond)
	do2(15, 0)
	time.Sleep(500 * time.Millisecond)
    fmt.Println("normally returned from main")
}

//	mult: 21 * 4 = 84
//	div: 21 / 4 = 5
//	mult: 15 * 0 = 0
//	div: recovered in do1: runtime error: integer divide by zero
//	normal resumption
//
//
//	div: mult: 21 * 4 = 84
//	21 / 4 = 5
//	div: mult: 15 * 0 = 0
//	panic: runtime error: integer divide by zero
//
//	goroutine 19 [running]:
//	main.div(0xf, 0x0)
//	        I:/Work. GO/work_VI/13/samples/recover_2a.go:37 +0x167
//	created by main.do2
//	        I:/Work. GO/work_VI/13/samples/recover_2a.go:25 +0xab
//	exit status 2
