package main

import (
	"fmt"
)

func IsB(n, last, prelast, preprelast uint64) bool {
	if last <= n {return last == n || IsB(n, last + preprelast, last, prelast)}
	return false
}

func main() {
	fmt.Println(IsB(47, 0, 1, 1))
}
