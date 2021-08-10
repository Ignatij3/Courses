package main

import "fmt"

type arr [1000000]int64

func UnlimitedRecursion(a arr) {
    fmt.Println(a[0])
    a[0]++
    UnlimitedRecursion(a)	
}

func main() {
	var c arr
	UnlimitedRecursion(c)
}
