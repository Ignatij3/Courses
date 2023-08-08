package main

import (
	"bufio"
	"fmt"
	"os"
)

func getData() ([]uint32, []string) {
	var n int

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d\n", &n)

	cost := make([]uint32, n)
	strs := make([]string, n)

	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d", &cost[i])
	}
	fmt.Fscan(reader, "\n")

	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%s\n", &strs[i])
	}

	return cost, strs
}

func main() {
	mana, strs := getData()
	res := calculate(mana, strs)
	fmt.Println(res)
}

func calculate(mana []uint32, strs []string) int64 {
	var res int64

	return res
}
