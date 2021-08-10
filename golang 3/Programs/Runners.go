package main

import  (
	"fmt"
)

func hasMet(distA, distB, track, i int)  bool{
	return (distA == track && distB == 0 && i > 0) || 
			   (distA == 0 && distB == track && i > 0)  || 
			   (distA == track && distB == track && i > 0) || 
			   (distA == 0 && distB == 0 && i > 0);
}


func Race(a, b, track int) (int, int, int) {
	var (
		collisions, iterations1, iterations2, distA, distB int
	)
	distA, distB = 0, 0
	
	for i := 0; i >= 0; i++ {
		if (hasMet(distA, distB, track, i))  {break}
		//fmt.Println(collisions, iterations1, iterations2, distA, distB, track)
		distA -= a
		distB += b
		if distA <= 0 {
			distA = track + distA
			iterations1++
		}
		if distB >= track {
			distB = (distB - track)
			iterations2++
		}
		if distA == distB {collisions++}
	}
	return collisions, iterations1, iterations2
}

func main() {
	var (
		runA, runB, track int
	)
	fmt.Print("Введите скорость первого участника: ")
	fmt.Scan(&runA)
	fmt.Print("Введите скорость второго участника: ")
	fmt.Scan(&runB)
	fmt.Print("Введите длину трассы: ")
	fmt.Scan(&track)
	collisions, iterations1, iterations2 := Race(runA, runB, track)
	fmt.Println("Они встретятся на старте через", iterations1, "кругов(а) для первого бегуна и через", iterations2, "кругов(а) для второго бегуна, но перед этим столкнутся", collisions, "раз(а)")
	
}
