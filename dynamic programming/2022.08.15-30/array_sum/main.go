package main

import (
	"bufio"
	"fmt"
	"os"

	"cock/segtree"
)

type query struct {
	xa, xb int
	ya, yb int
}

func getData() (*segtree.SegTree[segtree.SegTree[int]], []query) {
	var n, m int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanf(reader, "%d %d\n", &n, &m)

	set := func(data, value segtree.SegTree[int]) segtree.SegTree[int] {
		data.SetSegment(1, data.Size(), value.GetSeg(1, value.Size()))
		return data
	}
	add := func(a, b segtree.SegTree[int]) segtree.SegTree[int] {
		a.SetSegment(1, a.Size(), a.GetSeg(1, a.Size())+b.GetSeg(1, b.Size()))
		return a
	}
	initializer := func() segtree.SegTree[int] {
		return *segtree.InitTree[int](m, func(data, value int) int { return data + value }, func(a, b int) int { return a + b }, func() int { return 0 })
	}

	tree := segtree.InitTree[segtree.SegTree[int]](n, set, add, initializer)

	var (
		num  int
		temp segtree.SegTree[int] = initializer()
	)

	for i := 0; i < n; i++ {
		temp = tree.Get(i + 1)
		for j := 0; j < m; j++ {
			fmt.Fscanf(reader, "%d", &num)
			temp.Set(j, num)
		}

		tree.Set(i, temp)
		fmt.Fscanf(reader, "\n")
	}

	fmt.Fscanf(reader, "%d\n", &n)
	queries := make([]query, n)
	for i := range queries {
		fmt.Fscanf(reader, "%d %d %d %d\n", &queries[i].xa, &queries[i].xb, &queries[i].ya, &queries[i].yb)
	}

	for i := 0; i < tree.Size(); i++ {
		subt := tree.Get(i + 1)
		fmt.Printf("tree %d\n", i)
		subt.Traverse()
	}

	return tree, queries
}

func main() {
	tree, queries := getData()
	for _, qry := range queries {
		fmt.Println(calculate(tree, qry))
	}
}

func calculate(tree *segtree.SegTree[segtree.SegTree[int]], qry query) int {
	return tree.GetSeg(qry.xa, qry.xb).GetSeg(qry.ya, qry.yb)
}
