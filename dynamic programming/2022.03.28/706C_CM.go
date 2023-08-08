package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const INFINITY = math.MaxInt64

func calculate() int64 {
	var (
		min1, min2  int64
		s1, s2      string
		str, strRev string
		w, wr       int64
		n           int
	)

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d\n", &n)

	cost := make([]int64, n, n)

	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d", &cost[i])
	}
	fmt.Fscan(reader, "\n")

	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%s\n", &str)

		w = INFINITY
		if str >= s1 {
			w = min1
		}
		if str >= s2 && min2 < w {
			w = min2
		}

		strRev = reverse(str)

		wr = INFINITY
		if strRev >= s1 {
			wr = min1 + cost[i]
		}
		if strRev >= s2 && min2+cost[i] < wr {
			wr = min2 + cost[i]
		}

		min1, s1 = w, str
		min2, s2 = wr, strRev
	}

	if min2 < min1 {
		min1 = min2
	}

	if min1 == INFINITY {
		return -1
	}

	return min1
}

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func main() {
	res := calculate()
	fmt.Println(res)
}
