package main

import (
	"fmt"
)

type notAsceding struct {
	min uint32
	max uint32
}

func getData() (uint64, []notAsceding) {
	var (
		curr, prev uint32
		n, max     uint32
		asc        uint64
		ascfound   bool
	)
	nasc := make([]notAsceding, 0)

	fmt.Scanf("%d\n", &n)

	for i := n; i > 0; i-- {
		fmt.Scanf("%d", &n)

		fmt.Scanf("%d", &curr)
		max, prev = curr, curr

		for j := n - 1; j > 0; j-- {
			fmt.Scanf("%d", &curr)
			if prev < curr {
				ascfound = true
			}
			prev = curr
		}

		if ascfound {
			asc++
			ascfound = false
		} else if prev == curr {
			nasc = append(nasc, notAsceding{min: curr, max: max})
		}
		fmt.Scanf("\n")
	}

	return asc, nasc
}

func main() {
	asc, nasc := getData()
	res := calculate(nasc)

	lenn := uint64(len(nasc))
	res += (asc+lenn)*(asc+lenn) - lenn*lenn

	fmt.Println(res)
}

func calculate(nasc []notAsceding) (res uint64) {
	for i := 0; i < len(nasc); i++ {
		for j := 0; j < len(nasc); j++ {
			if nasc[i].min < nasc[j].max {
				res++
			}
		}
	}
	return
}
