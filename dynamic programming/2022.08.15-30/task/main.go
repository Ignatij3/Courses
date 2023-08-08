package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type (
	SegmentTree struct {
		root *node
	}
	node struct {
		a, b int
		l, r *node
		sum  int
	}
	runner struct {
		id    byte
		place int
	}
)

func InitTree(n int) *SegmentTree {
	tree := new(SegmentTree)
	tree.root = &node{a: 1, b: n}
	tree.root.initSegment()
	return tree
}

func (t *SegmentTree) Put(place, num int) {
	t.root.put(place, num)
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
	if n.a == n.b && n.a == place {
		n.sum = num
		return
	}

	if n.l.b >= place {
		n.l.put(place, num)
	} else {
		n.r.put(place, num)
	}

	n.sum = n.l.sum + n.r.sum
}

func (t *SegmentTree) PlaceRunner(r *runner) {
	var n *node = t.root

	for n.a < n.b {
		n.sum--
		if r.place <= n.l.sum {
			n = n.l
		} else {
			r.place -= n.l.sum
			n = n.r
		}
	}
	r.place = n.a
}

func getData() (*SegmentTree, []runner) {
	var n int

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d\n", &n)

	tree := InitTree(n)
	runners := make([]runner, n)

	var letter byte
	for i := n - 1; i >= 0; i-- {
		fmt.Fscanf(reader, "%c %d\n", &letter, &n)
		runners[i] = runner{id: letter, place: n} //Закидываются в обратном порядке
		tree.Put(i+1, 1)
	}

	return tree, runners
}

func main() {
	tree, runners := getData()
	calculate(tree, runners)
	sort.Slice(runners, func(i, j int) bool {
		return runners[i].place < runners[j].place
	})
	for i := range runners {
		fmt.Printf("%c ", runners[i].id)
	}
}

func calculate(tree *SegmentTree, runners []runner) {
	for i := range runners {
		tree.PlaceRunner(&runners[i])
	}
}
