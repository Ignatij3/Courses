package main

import "fmt"

type trio struct {
    n int
    pc *rune
    next *trio
}

func main() {  
	var s *trio
	s = new(trio)
	(*s).n = 7
	(*s).pc = new(rune)
	*(*s).pc = 'R'

	(*s).next = &trio{9, nil, s}
	(*(*s).next).pc = new(rune)
	*(*(*s).next).pc = 'W'
	fmt.Println(*s)
}	

