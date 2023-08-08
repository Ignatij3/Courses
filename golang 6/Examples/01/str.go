package main

import "fmt"

type gh struct {
	a int
	b rune
}

func main() {
	var myst gh = gh{3, 'd'}
	fmt.Println(myst)
}
