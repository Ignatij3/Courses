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
	a, b  int
	l, r  *node
	left  int
	all   int
	right int
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

func (t segmentTree) Get(a, b int) int {
	_, res, _ := t.root.get(a, b)
	return res
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
		n.left, n.all, n.right = 1, 1, 1
		return
	}

	if n.l.b >= place {
		n.l.put(place, num)
	} else {
		n.r.put(place, num)
	}

	n.left = n.l.left
	if n.l.b-n.l.a+1 == n.l.left {
		n.left += n.r.left
	}

	n.all = max(n.l.all, n.r.all, n.l.right+n.r.left)

	n.right = n.r.right
	if n.r.b-n.r.a+1 == n.r.right {
		n.right += n.l.right
	}
}

func (n node) get(a, b int) (int, int, int) {
	if n.a >= a && n.b <= b {
		return n.left, n.all, n.right
	}

	if n.b < a || n.a > b {
		return 0, 0, 0
	}

	// return n.l.get(a, b) + n.r.get(a, b)
}

func max[T int | int64](elems ...T) T {
	var max T
	for _, n := range elems {
		if n > max {
			max = n
		}
	}
	return max
}

func getData() {
	var (
		n, a, b int
		command string
	)

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d\n", &n)

	tree := InitTree(n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d", &a)
		tree.Put(i+1, a)
	}

	fmt.Fscanf(reader, "%d\n", &n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%s %d %d\n", &command, &a, &b)

		if command == "UPDATE" {
			tree.Put(a, b)
		} else if command == "QUERY" {
			fmt.Println(tree.Get(a, b))
		}
	}
}

func main() {
	getData()
}
