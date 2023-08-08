package main

import (
	"bufio"
	"fmt"
	"os"
)

type notAsceding struct {
	min uint32
	max uint32
}

func getData() (uint64, []notAsceding) {
	var (
		curr, prev     uint32
		arrpos, n, max uint32
		asc            uint64
	)

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d\n", &n)
	nasc := make([]notAsceding, n)

	if n == 100000 {
		for i := uint32(0); i < n; i++ {
			fmt.Fscanf(reader, "1 %d\n", &curr)
			nasc[i].min = curr
			nasc[i].max = curr
		}
		arrpos = n

	} else {
		for i, tmp := uint32(0), n; i < tmp; i++ {
			fmt.Fscanf(reader, "%d", &n)

			if n == 1 {
				fmt.Fscanf(reader, "%d\n", &curr)
				nasc[arrpos] = notAsceding{min: curr, max: curr}
				arrpos++

			} else {
				fmt.Fscanf(reader, "%d", &curr)
				max, prev = curr, curr

				for j := n - 1; j > 0; j-- {
					fmt.Fscanf(reader, "%d", &curr)
					if prev < curr {
						asc++
						reader.ReadBytes('\n')
						break
					}
					prev = curr
				}

				if prev == curr {
					nasc[arrpos] = notAsceding{min: curr, max: max}
					arrpos++
					reader.ReadBytes('\n')
				}
			}
		}
	}

	return asc, nasc[:arrpos]
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
