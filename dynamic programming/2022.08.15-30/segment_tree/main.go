package main

import (
	"bufio"
	"fmt"
	"os"

	"retard/segtree"
)

func getData() {
	var n, m int

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d %d\n", &n, &m)

	tree := segtree.InitTree(n)

	var letter rune
	for i := m - 1; i >= 0; i-- {
		fmt.Fscanf(reader, "%c %d %d\n", &letter, &n, &m)
		if letter == 'P' {
			tree.Put(n, m)
		} else if letter == 'G' {
			fmt.Println(tree.Get(n, m))
		}
	}
}

func main() {
	getData()
}
