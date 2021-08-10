package main

import (
	"fmt"
	"binheap"
)

func HeapSort(a []binheap.Tdata) []binheap.Tdata{
	var (
		b binheap.LocatorBinaryHeap
		t binheap.Lmnt
	)	
	b.Init(a)
	result:= make([]binheap.Tdata, len(a), len(a)) 
	for k:= len(a) - 1; k >= 0; k-- {
		t = b.ExtractMax()
		result[k] = t.Value
	}
	return result
}	
	
func main() {
	a:= []binheap.Tdata{2,5,7,2,4,9,1,6}
	fmt.Println(a)				// [2 5 7 2 4 9 1 6]
	fmt.Println(HeapSort(a))	// [1 2 2 4 5 6 7 9]

	a = []binheap.Tdata{3,7,2}
	var b binheap.LocatorBinaryHeap
	b.Init(a)
	fmt.Println(b)	// {[{1 7} {0 3} {2 2}] [1 0 2]}
	b.Delete(1)
	fmt.Println(b)	// {[{0 3} {2 2}] [0 -1 1]}
	b.Add(binheap.Lmnt{1,17})
	fmt.Println(b)	// {[{1 17} {2 2} {0 3}] [2 0 1]}
	b.Add(binheap.Lmnt{1,7})
	fmt.Println(b)	// {[{1 17} {2 2} {0 3}] [2 0 1]}
}	

