package main

import (
	"math/rand"
	"time"
	"fmt"
)

func FillSlice(a *[]int) {
	rand.Seed(time.Now().UnixNano())
	for n := 0; n < 31; n++ {
		switch rand.Intn(3) {
			case 0:
				*a = append(*a, -rand.Intn(50))
			default:
				*a = append(*a, rand.Intn(50))
		}
	}
}

func HeapSort(a []int) {
	var pivot int = len(a) - 1
	for ; pivot > 0; pivot-- {
		if a[(pivot - 1) / 2] < a[pivot] {
			for nPivot := pivot; nPivot > 0 && a[(nPivot - 1) / 2] < a[nPivot]; nPivot = (nPivot - 1) / 2 {
				a[(nPivot - 1) / 2], a[nPivot] = a[nPivot], a[(nPivot - 1) / 2]
				for nPivot2 := pivot; (nPivot2 * 2) + 1 < len(a); {
					if len(a) % 2 == 0 && (nPivot2 * 2) + 2 == len(a) {
						if a[nPivot2] < a[(nPivot2 * 2) + 1] {a[(nPivot2 * 2) + 1], a[nPivot2] = a[nPivot2], a[(nPivot2 * 2) + 1]}
						break
					} else if a[(nPivot2 * 2) + 1] >= a[(nPivot2 * 2) + 2] && a[nPivot2] < a[(nPivot2 * 2) + 1] {
						a[(nPivot2 * 2) + 1], a[nPivot2] = a[nPivot2], a[(nPivot2 * 2) + 1]
						nPivot2 = (nPivot2 * 2) + 1
					} else if a[(nPivot2 * 2) + 2] > a[(nPivot2 * 2) + 1] && a[nPivot2] < a[(nPivot2 * 2) + 2] {
						a[(nPivot2 * 2) + 2], a[nPivot2] = a[nPivot2], a[(nPivot2 * 2) + 2]
						nPivot2 = (nPivot2 * 2) + 2
					} else {break}
				}
			}
		}
	}
}

func main() {
	var a []int
	FillSlice(&a)
	HeapSort(a)
	fmt.Println(a)
}
