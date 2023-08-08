package segtree

// TODO: add update for single existing tree (for performance reasons)

// SegTree is persistent segment tree
type SegTree struct {
	toproot *node
	roots   []*node
	rcount  int
}

type node struct {
	left, right int
	lson, rson  *node
	data        int
	oath        func(int) int
}

func InitTree(n int) *SegTree {
	if n < 1 {
		n = 1
	}

	tree := &SegTree{
		toproot: nil,
		roots:   make([]*node, 2),
		rcount:  1,
	}

	tree.roots[0] = &node{
		left:  1,
		right: n,
	}
	tree.roots[0].initSegment()

	tree.roots[1] = &node{
		left:  1,
		right: n,
	}
	tree.roots[1].initSegment()

	tree.toproot = tree.roots[1]
	return tree
}

func (n *node) initSegment() {
	if n.right > n.left {
		n.lson = &node{
			left:  n.left,
			right: (n.left + n.right) / 2,
		}
		n.rson = &node{
			left:  ((n.left + n.right) / 2) + 1,
			right: n.right,
		}
		n.lson.initSegment()
		n.rson.initSegment()
	}
}

func (t *SegTree) AddTree() {
	t.roots = append(t.roots, new(node))
	t.rcount++
	*t.roots[t.rcount] = *t.toproot
	t.toproot = t.roots[t.rcount]
}

func (t *SegTree) UpdateSingle(pos int, upd func(int) int) {
	t.toproot.update(t.roots[t.rcount-1], pos, pos, upd)
}

func (t *SegTree) Update(left, right int, upd func(int) int) {
	t.toproot.update(t.roots[t.rcount-1], left, right, upd)
}

func (n *node) update(old *node, left, right int, upd func(int) int) {
	n.fulfillOath()
	n.copyFrom(old)
	n.fulfillOath()

	if n.left == n.right {
		n.data = upd(n.data)
		return
	}

	if left <= n.left && n.right <= right {
		n.data = upd(n.data)
		n.oath = upd
		return
	}

	if n.lson.right >= left {
		n.lson.update(old, left, right, upd)
		(*n).rson = (*old).rson
	}
	if n.rson.left <= right {
		n.rson.update(old, left, right, upd)
		(*n).lson = (*old).lson
	}

	n.data = n.lson.data + n.rson.data
}

func (t SegTree) Get(left, right int) int {
	return t.toproot.get(left, right)
}

func (n node) get(left, right int) int {
	if n.left >= left && n.right <= right {
		return n.data
	}

	if n.right < left || n.left > right {
		return 0
	}

	return n.lson.get(left, right) + n.rson.get(left, right)
}

func (n *node) fulfillOath() {
	if n.oath != nil {
		n.data = n.oath(n.data)
		if n.left != n.right {
			n.lson.oath = n.oath
			n.rson.oath = n.oath
		}
		n.oath = nil
	}
}

func (n *node) copyFrom(old *node) {
	*n = *old
	if n.left != n.right {
		n.lson = &node{}
		n.rson = &node{}
	}
}
