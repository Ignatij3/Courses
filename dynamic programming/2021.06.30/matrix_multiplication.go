package main

import (
	"fmt"
	"math"
)

var k = 9

func main() {
	n := []int{4, 7, 3, 5, 6, 2, 9, 3, 5, 2}
	a := make([][]int, k)
	steps := make([][]int, k)
	for i := 0; i < k; i++ {
		a[i] = make([]int, k+1)
		steps[i] = make([]int, k+1)
	}

	for i := 0; i < k; i++ {
		a[i][i+1] = 0
	}
	for i := 0; i < k-1; i++ {
		a[i][i+2] = n[i] * n[i+1] * n[i+2]
		steps[i][i+2] = n[i+1]
	}

	for delta := 3; delta <= k; delta++ {
		for i := 0; i <= k-delta; i++ {
			min := math.MaxInt64
			for j := i + 1; j < delta+i; j++ {
				temp := a[i][j] + n[i]*n[j]*n[delta+i] + a[j][delta+i]
				if temp < min {
					min = temp
					steps[i][delta+i] = n[j]
				}
			}
			a[i][delta+i] = min
		}
	}

	print_steps(steps)
}

func print_steps(steps [][]int) {
	print_with_pos(steps, 0, k)
	fmt.Println(steps[0][k])
}

func print_with_pos(steps [][]int, start, stop int) {
	if steps[start][stop]+1 < stop && steps[start][stop]-1 > start {
		print_with_pos(steps, start, steps[start][stop])
		print_with_pos(steps, steps[start][stop], stop)
	}
	fmt.Println(steps[start][stop])
}
