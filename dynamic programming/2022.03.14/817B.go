package main

import (
	"bufio"
	"fmt"
	"os"
)

func getData() uint64 {
	var (
		a, b, c, n uint32 = 1e9 + 1, 1e9 + 1, 1e9 + 1, 0
		an, bn, cn uint64
	)

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d\n", &n)

	for i := n; i > 0; i-- {
		fmt.Fscanf(reader, "%d", &n)

		if n <= a {
			if n == a {
				an, bn, cn = an+1, 1, bn
			} else {
				an, bn, cn = 1, an, bn
			}
			a, b, c = n, a, b
		} else if n <= b { //a, b, c = a, n, b
			if n == b {
				bn, cn = bn+1, 1
			} else {
				bn, cn = 1, bn
			}
			b, c = n, b
		} else if n < c { // a, b, c = a, b, n
			c = n
			cn = 1
		} else if n == c {
			cn++
		}
	}

	if a == c {
		an = (an * (an - 1) * (an - 2)) / 6
		bn, cn = 1, 1
	} else if a == b {
		an, bn = 1, 1
	} else if b == c {
		bn = (bn * (bn - 1)) / 2
		cn = 1
	}

	return an * bn * cn
}

func main() {
	res := getData()
	fmt.Println(res)
}
