package main

type ( //3-е задание
	recur struct {
		next1 *adrnum
		next2 *recur
	}
	adrnum struct {
		next *adrnum
		num int
	}
)

func main() {
	var s *recur
	
	s = new(recur)
	
	(*s).next1 = &adrnum{nil, 11}
	(*(*s).next1).next = &adrnum{nil, 12}
	(*(*(*s).next1).next).next = &adrnum{nil, 13}
	
	(*s).next2 = &recur{nil, nil}
	
	(*(*s).next2).next1 = &adrnum{nil, 21}
	(*(*(*s).next2).next1).next = &adrnum{nil, 22}
	
	(*(*s).next2).next2 = &recur{nil, nil}
	
	(*(*(*s).next2).next2).next2 = &recur{nil, nil}
	
	(*(*(*s).next2).next2).next1 = &adrnum{nil, 41}
}
