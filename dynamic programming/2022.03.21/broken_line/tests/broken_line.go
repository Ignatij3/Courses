package main

import "fmt"

//N stands for n-word
var N, start, stop uint16

func main() {
	fmt.Scanf("%d %d %d\n", &N, &start, &stop)
	res := calculate()
	fmt.Println(res)
}

func calculate() (res uint64) {
	for end := uint16(0); end < N; end++ {
		if end < start && end != stop {
			res += f(end, start)
		} else if end > start && end != stop {
			res += f(start, end)
		}
	}
	return
}

func f(left, right uint16) uint64 {
	return uint64(left) * uint64(right)
}
