package main

type recur struct { //2-е задание
	pi *float64
	next *recur
}

func main() {
	var s *recur
	s = new(recur)
	
	(*s).pi = new(float64)
	*(*s).pi = 3.14
	(*s).next = &recur{(*s).pi, nil}
	
	(*(*s).next).next = &recur{(*s).pi, s}
}
