package segtree

type SegmentTree struct {
	root *node
}

type node struct {
	a, b int
	l, r *node
	sum  int
}

func InitTree(n int) *SegmentTree {
	tree := new(SegmentTree)
	tree.root = &node{a: 1, b: n}
	tree.root.initSegment()
	return tree
}

func (t *SegmentTree) Put(place, num int) {
	t.root.put(place, num)
}

func (t SegmentTree) Get(a, b int) int {
	return t.root.get(a, b)
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

func (n node) get(a, b int) int {
	if n.a >= a && n.b <= b {
		return n.sum
	}

	if n.b < a || n.a > b {
		return 0
	}

	return n.l.get(a, b) + n.r.get(a, b)
}
