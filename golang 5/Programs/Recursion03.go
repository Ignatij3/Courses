package main

type ( //4-е задание
	recur struct {
		next1 *booltr
		next2 *recur
	}
	booltr struct {
		trues *bool
	}
)

func main() {
	var (
		s *recur
		true1, true2 bool = true, true
	)
	s = new(recur)
	
	(*s).next1 = &booltr{&true1}
	
	(*s).next2 = &recur{(*s).next1, nil}
	(*(*s).next2).next1 = &booltr{&true1}
	
	(*(*s).next2).next2 = &recur{nil, nil}
	(*(*(*s).next2).next2).next1 = &booltr{&true2}
}
