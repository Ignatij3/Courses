package main

import (
	"bufio"
	"fmt"
	"os"
)

func getData() uint64 {
	var n uint64

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d\n", &n)

	return n
}

func main() {
	n := getData()
	res := calculate(n)
	fmt.Println(res)
}

func calculate(n uint64) uint64 {
	var res uint64

	return res
}