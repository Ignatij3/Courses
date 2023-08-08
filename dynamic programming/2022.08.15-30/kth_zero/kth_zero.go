package main

import (
	"bufio"
	"fmt"
	"os"
)

type segmentTree struct {
	root *node
}

type node struct {
	a, b int
	l, r *node
	sum  int
}

func InitTree(n int) *segmentTree {
	tree := new(segmentTree)
	tree.root = &node{a: 1, b: n}
	tree.root.initSegment()
	return tree
}

func (t *segmentTree) Put(place, num int) {
	t.root.put(place, num)
}

func (t segmentTree) Get(a, b, k int) int {
	var n *node = t.root

	k += t.root.count(1, a-1)
	if n.sum < k {
		return -1
	}

	for n.a < n.b {
		if k <= n.l.sum {
			n = n.l
		} else {
			k -= n.l.sum
			n = n.r
		}
	}

	if n.a > b || n.a < a {
		return -1
	}
	return n.a
}

func (n node) count(a, b int) int {
	if n.a >= a && n.b <= b {
		return n.sum
	}

	if n.b < a || n.a > b {
		return 0
	}

	return n.l.count(a, b) + n.r.count(a, b)
}

func (n *node) initSegment() {
	if n.b > n.a {
		n.l = &node{a: n.a, b: (n.a + n.b) / 2}
		n.r = &node{a: ((n.a + n.b) / 2) + 1, b: n.b}
		n.l.initSegment()
		n.r.initSegment()
	}
}

func (n *node) put(place, num int) {
	if n.a == n.b {
		n.sum = 0
		if num == 0 {
			n.sum++
		}
		return
	}

	if n.l.b >= place {
		n.l.put(place, num)
	} else {
		n.r.put(place, num)
	}

	n.sum = n.l.sum + n.r.sum
}

func getData() {
	var (
		n, num      int
		letter      byte
		place, l, r int
	)

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d\n", &n)

	tree := InitTree(n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d\n", &num)
		tree.Put(i+1, num)
	}

	fmt.Fscanf(reader, "%d\n", &n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%c ", &letter)

		if letter == 'U' {
			fmt.Fscanf(reader, "%d %d\n", &place, &num)
			tree.Put(place, num)
		} else if letter == 'S' {
			fmt.Fscanf(reader, "%d %d %d\n", &l, &r, &num)
			fmt.Println(tree.Get(l, r, num))
		}
	}
}

func main() {
	getData()
}
