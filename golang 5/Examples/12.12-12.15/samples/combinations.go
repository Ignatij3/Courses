package main

import "fmt"

var N, K int 

var P []int //глобальный "массив"

func Cikl(CurrentPos int, FirstItem int) {
	// CurrentPos - степень вложенности цикла,
	// FirstItem - стартовое число очередного цикла
	for i := FirstItem; i <= N-(K-CurrentPos); i++ {
		P[CurrentPos-1] = i
		// Цикл вложенности K - последний.
		if CurrentPos == K {
			// Печатаем комбинацию (сочетание) 
			fmt.Println(P)
		} else {
			Cikl(CurrentPos+1, i+1)
		}
	}
}

func main() {
	for {
		fmt.Print("Enter N: ")
		fmt.Scanln(&N)
		fmt.Print("Enter K: ")
		fmt.Scanln(&K)
		if K>0 && N>=K { break }
	}
	// Алгоритм: реализуем цикл вложенности K
	// for i1:= 1; i <= N-(K-1) {
	//     for i2:= i1+1; i <=  N-(K-2) {
	//         for i3:= i2+1; i <=  N-(K-3) {
	//           . . .      . . .      . . .
	//             for iK:= i(K-1)+1; To N-(K-K) {
	//                 Print (i1, i2, ..., iK)
	//             }
	//           . . .      . . .      . . .
	//         }
	//     }
	// }
	P = make([]int, K, K)
	Cikl(1, 1)
}
