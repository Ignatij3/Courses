package main 

import (
	"testing"
	"math/rand"
	"math"
)

type data []int

const (
	MinData = math.MinInt32
	LEN = 90000000
)

var slice data

func init() {
	for n := 0; n < LEN; n++ {
		switch rand.Intn(2) {
			case 0:
				slice = append(slice, -rand.Intn(LEN / 2))
			case 1:
				slice = append(slice, rand.Intn(LEN / 2))
		}
	}
}

func (b data) pushDown(place int) {
	if place >= len(b) || place < 0 { 
		return
	}	
	x := b[place]
	for  {
		if 2*place + 1 >= len(b) {
			break
		}
		maxson := 2*place + 1
		rson:= maxson + 1
		if rson < len(b) && b[rson] > b[maxson] {
			maxson = rson
		}
		if b[maxson] <= x {
			break
		}
		b[place] = b[maxson]	 	
		place = maxson
	}
	b[place] = x	
}	

func (b data) ExtractMax() int {
	if len(b) > 0 {
		max:= b[0]
		b[0] = b[len(b)-1]
		b = b[:len(b)-1]
		b.pushDown(0)
		return max
	} else {
		return MinData
	}
}

func (a data) HeapSort() {
	for k := len(a) - 1; k >= 0; k-- {
		a[k] = a.ExtractMax()
	}
}	

func (a data) MyHeapSort() {
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

func BenchmarkMyHeapSort(b *testing.B) {
	for i:= 0; i < b.N; i++  {
        slice.MyHeapSort()
    }
}

func BenchmarkHeapSort(b *testing.B) {
	for i:= 0; i < b.N; i++  {
        slice.HeapSort()
    }
}



