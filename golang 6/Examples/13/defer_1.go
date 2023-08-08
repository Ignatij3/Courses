package main

import "fmt"

func printValue(x int) {
	fmt.Println("value of parameter in deferred function =", x)
}

func main() {
	z := 5
	defer printValue(z)
	z = 10
	fmt.Println("value of parameter before deferred function call =", z)
}
//	value of parameter before deferred function call = 10
//	value of parameter in deferred function = 5
