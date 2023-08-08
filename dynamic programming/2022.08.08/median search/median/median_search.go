package median

import (
	"math"
	"sort"
)

const (
	divnLen  = 5 // division length
	infinity = 1<<63 - 1
)

// QuickSelect searches for the median of the array using quick select
// It returns median and it's place in the array
func QuickSearch(arr []int) (int, int) {
	if len(arr) == 0 {
		return 0, 0
	}
	return quickSearch(append([]int{}, arr...), len(arr)/2, len(arr)-1)
}

func quickSearch(arr []int, medianPos int, pivotPos int) (int, int) {
	if len(arr) == 1 {
		return arr[0], 0
	}

	pivot := arr[pivotPos]
	place := 0

	for i := range arr {
		if arr[i] < pivot {
			arr[i], arr[place] = arr[place], arr[i]
			place++
		}
	}

	arr[place], arr[pivotPos] = arr[pivotPos], arr[place]

	if medianPos < place {
		return quickSearch(arr[:place], medianPos, place-1)
	} else if medianPos > place {
		return quickSearch(arr[place+1:], medianPos-(place+1), len(arr)-(place+2))
	}
	return arr[place], place
}

// MedianOfMedians searches for the median of the array using median-of-median algorithm
// It returns median and it's place in the array
func MedianOfMedians(arr []int) (int, int) {
	if len(arr) == 0 {
		return 0, 0
	}
	return medianOfMedians(append([]int{}, arr...))
}

func medianOfMedians(arr []int) (int, int) {
	segments := make([][5]int, int(math.Ceil(float64(len(arr))/divnLen)))

	for i := 0; i < len(arr); i++ {
		segments[i/divnLen][i%divnLen] = arr[i]
		if (i+1)%divnLen == 0 {
			sort.Slice(segments[i/divnLen][:], func(a, b int) bool {
				return segments[i/divnLen][a] < segments[i/divnLen][b]
			})
		}
	}

	if len(arr)%divnLen != 0 {
		for i := len(arr) % divnLen; i < 5; i++ {
			segments[len(segments)-1][i] = infinity
		}
	}

	medians := make([]int, len(segments))
	for i := 0; i < len(segments); i++ {
		medians[i] = segments[i][divnLen/2]
	}

	_, pos := quickSearch(medians, len(medians)/2, len(medians)-1)
	return quickSearch(arr, len(arr)/2, pos*divnLen+2)
}
