package main

import (
	"fmt"

	"./median"
)

func main() {
	set := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("expected: %d\n", 5)
	med, _ := median.QuickSearch(set)
	fmt.Printf("s1-qs: %d\n", med)
	med, _ = median.MedianOfMedians(set)
	fmt.Printf("s1-mm: %d\n\n", med)

	set = []int{48, 49, 50, 51, 52, 43, 44, 45, 46, 47, 38, 39, 40, 41, 42, 33, 34, 35, 36, 37, 28, 29, 30, 31, 32, 23, 21, 22, 13, 14, 15, 16, 17, 8, 9, 10, 53, 54, 24, 25, 26, 27, 18, 19, 20}
	fmt.Printf("expected: %d\n", 32)
	med, _ = median.QuickSearch(set)
	fmt.Printf("s2-qs: %d\n", med)
	med, _ = median.MedianOfMedians(set)
	fmt.Printf("s2-mm: %d\n\n", med)

	set = []int{1, 2, 3, 4, 5, 1000, 8, 9, 99}
	fmt.Printf("expected: %d\n", 5)
	med, _ = median.QuickSearch(set)
	fmt.Printf("s2-qs: %d\n", med)
	med, _ = median.MedianOfMedians(set)
	fmt.Printf("s2-mm: %d\n\n", med)

	set = []int{43, -32, 58, 57, 52, 78, -15, 61, 42, 0, -40, -67, -72, -79, 87, -28, -33, 71, 60, -19, -98, -45, 40, -40, 70, 12, 20, -50, -68, 57, 59, 58, 67, 16, 51, -99, -76, 6, -70, -84, 21, -60, -36, 80, -19, -1, 14, -54, 41, -8, 1, -32, 95, 28, -50, -94, -42, -98, 17, 43, -11, 40, 19, 24, -2, 26, 53, 14, 53, -96, -43, 22, -94, -39, 78, 89, 82, 64, 66, -43, -78, 88, -11, -72, 68, -25, -68, 7, 16, 16, 7, 40, 64, 66, -61, -65, 15, -38, -5, -61, 25, 84, 84, -39, 15, -61, 100, -7, 87, -97, 6, -3, 70, 73, 47, 2, -86, -35, -98, 36}
	fmt.Printf("expected: %d\n", 7)
	med, _ = median.QuickSearch(set)
	fmt.Printf("s2-qs: %d\n", med)
	med, _ = median.MedianOfMedians(set)
	fmt.Printf("s2-mm: %d\n\n", med)
}
