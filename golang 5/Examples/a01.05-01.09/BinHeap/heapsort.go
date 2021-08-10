package main

import (
	"fmt"
	"binheap"
)

func HeapSort(a []binheap.Tdata) {
	var bheap binheap.BinaryHeap
	bheap.Init(a)
	for k:= len(a) - 1; k >= 0; k-- {
		a[k] = bheap.ExtractMax()
	}
}	
	
func main() {
	a:= []binheap.Tdata{2,5,7,2,4,9,1,6}
	fmt.Println(a)	// [2 5 7 2 4 9 1 6]
	HeapSort(a)
	fmt.Println(a)	// [1 2 2 4 5 6 7 9]
}	
