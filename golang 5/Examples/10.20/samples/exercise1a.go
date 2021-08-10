package main


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
	(*s).next = new(trio)
	(*(*s).next).n = 9
	(*(*s).next).pc = new(rune)
	*((*(*s).next).pc) = 'W'
	(*(*s).next).next = s
}	

