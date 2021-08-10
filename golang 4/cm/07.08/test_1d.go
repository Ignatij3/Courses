package main

import "fmt"

// convert types: takes an int and returns a corresponding string value
type convert func(int) string

// function int2str implements convert, returning x as string
func int2str(x int) string {
	return fmt.Sprintf("%v", x)
}

// quote passes x to some convert func and returns quoted string
func quote(fn convert, x int) string {
	return fmt.Sprintf("%q", fn(x))
}

func main() {
	var result string
	
	result = int2str(54321)
	fmt.Println(result)    // 54321

	result = quote(int2str, 12345)
	fmt.Println(result)    // "12345"

	result = quote(func(x int) string { return fmt.Sprintf("%b", x) }, 341)
	fmt.Println(result)    // "101010101"

	trashFunc := func(x int) string { return "trash" }
	result = quote(trashFunc, 314159)
	fmt.Println(result)    // "trash"
}
