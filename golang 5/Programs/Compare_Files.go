package main

import (
	"strconv"
	"fmt"
	"os"
)

func CompareNth(n int) bool {
	var (
		a, b *os.File
		anum, bnum int
		err1, err2 error
	)
	
	if n < 10 {
		a, _ = os.Open("../Files/results/result0" + strconv.Itoa(n) + ".dat")
		b, _ = os.Open("../Files/tests/answer0" + strconv.Itoa(n) + ".txt")
	} else {
		a, _ = os.Open("../Files/results/result" + strconv.Itoa(n) + ".dat")
		b, _ = os.Open("../Files/tests/answer" + strconv.Itoa(n) + ".txt")
	}
	
	fmt.Fscanf(a, "%d ", &anum)
	fmt.Fscanf(b, "%d ", &bnum)
	if anum == 0 {
		if bnum != 0 {return false}
		fmt.Fscanf(a, "%d", &anum)
		fmt.Fscanf(b, "%d", &bnum)
		return anum == bnum
	}
	
	for {
		_, err1 = fmt.Fscanf(a, "%d\n", &anum)
		_, err2 = fmt.Fscanf(b, "%d\n", &bnum)
		if err1 != nil && err2 != nil {return true} else if (err1 != nil || err2 != nil) && !(err1 != nil && err2 != nil) {return false}
		if anum != bnum {return false}
	}
	return true
}

func main() {
	for n := 1; n <= 26; n++ {
		fmt.Printf("%d) %t\n", n, CompareNth(n))
	}
}
