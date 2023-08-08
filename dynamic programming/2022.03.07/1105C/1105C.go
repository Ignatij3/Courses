package main

import "fmt"

const MODULE = 1000000007

func getData() (n, l, r uint64) {
	fmt.Scanf("%d %d %d\n", &n, &l, &r)
	return
}

func main() {
	n, l, r := getData()
	res := calculate(n, l, r)
	fmt.Println(res)
}

func calculate(n, l, r uint64) uint64 {
	var (
		a0, a1, a2          uint64
		rem0, rem1, rem2    uint64 = 1, 0, 0
		temp0, temp1, temp2 uint64
	)

	a0 = r/3 - l/3
	a1 = (r+2)/3 - (l-1)/3
	a2 = (r - l + 1) - (a0 + a1)

	for i := uint64(0); i < n; i++ {
		temp0 = (a0*rem0 + a2*rem1 + a1*rem2) % MODULE
		temp1 = (a1*rem0 + a0*rem1 + a2*rem2) % MODULE
		temp2 = (a2*rem0 + a1*rem1 + a0*rem2) % MODULE

		rem0, rem1, rem2 = temp0, temp1, temp2
	}

	return rem0
}
