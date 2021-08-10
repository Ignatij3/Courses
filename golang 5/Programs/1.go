package main

import (
	"fmt"
)

func Hanoy(existence, shit, amen, cow19, prtr uint64) {
	if shit > existence {
		fmt.Println("нет в последовательности")
	} else if existence != shit {
		Hanoy(existence, amen + prtr, cow19, prtr, amen + prtr)
	} else if existence == shit {
		fmt.Println("есть в последовательности")
	}
}

func main() {
	Hanoy(47, 0, 1, 0, 1)
}
