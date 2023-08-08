package main

import (
	"fmt"
)

func getData() ([]uint64, []uint64) {
	var amt uint64
	fmt.Scanf("%d\n", &amt)

	n := make([]uint64, amt)
	m := make([]uint64, amt)

	for i := uint64(0); i < amt; i++ {
		fmt.Scanf("%d %d\n", &n[i], &m[i])
	}

	return n, m
}

func main() {
	n, m := getData()
	for i := range n {
		fmt.Println(ayoubFunction(n[i], m[i]))
	}
}

func ayoubFunction(n, m uint64) uint64 {
	if m == 0 {
		return 0
	}

	var count uint64
	count += (n*n + n) / 2

	//Следующие 4 строчки подсмотрены с Codeforces, так как я не имел понятия, куда ещё сокращать код
	//Именно код, который считал кол-во нулей слева и справаdd
	z := n - m
	k := z / (m + 1)
	count -= (m + 1) * k * (k + 1) / 2
	count -= (z % (m + 1)) * (k + 1)

	return count
}
