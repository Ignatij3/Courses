package segtree

// SegTree[T] is persistent segment tree with data of type T.
//
// Set function updates data in node to whatever the function returns.
// Argument of function is node's value which is to be updated and passed update value.
//
// Merge must return result of merging left and right subtrees's value.
// The arguments are, subsequently, the values they hold.
type SegTree[T TreeAble[T]] struct {
	toproot      *node[T]
	roots        []*node[T]
	rcount, size int
	initData     func() T
}

type node[T TreeAble[T]] struct {
	left, right int
	lson, rson  *node[T]
	data        T
	oath        func(T) T
}

// TODO убрать лишний нулевой корень

func InitTree[T TreeAble[T]](a, b int, initData func() T) *SegTree[T] {
	if b < a {
		a, b = b, a
	}

	tree := &SegTree[T]{
		toproot:  nil,
		roots:    make([]*node[T], 2),
		rcount:   1,
		size:     b - a + 1,
		initData: initData,
	}

	tree.roots[0] = &node[T]{
		left:  a,
		right: b,
		data:  tree.initData(),
	}

	tree.roots[0].initSegment(tree.initData)
	tree.roots[1] = tree.roots[0]
	tree.toproot = tree.roots[1]

	return tree
}

func (n *node[T]) initSegment(initdata func() T) {
	n.data = initdata()

	if n.right > n.left {
		n.lson = &node[T]{
			left:  n.left,
			right: (n.left + n.right) / 2,
		}
		n.rson = &node[T]{
			left:  ((n.left + n.right) / 2) + 1,
			right: n.right,
		}

		n.lson.initSegment(initdata)
		n.rson.initSegment(initdata)
	}
}

func (t *SegTree[T]) Traverse(do func(T, int, int)) {
	t.toproot.traverse(do)
}

func (n *node[T]) traverse(do func(T, int, int)) {
	if n != nil {
		n.lson.traverse(do)
		do(n.data, n.left, n.right)
		n.rson.traverse(do)
	}
}

func (t SegTree[T]) Size() int {
	return t.size
}

func (t SegTree[T]) Left() int {
	return t.toproot.left
}

func (t SegTree[T]) Right() int {
	return t.toproot.right
}

func (t *SegTree[T]) MergeTree() {
	t.roots = append(t.roots, new(node[T]))
	t.rcount++
	*t.roots[t.rcount] = *t.toproot
	t.toproot = t.roots[t.rcount]
}

func (t *SegTree[T]) Set(pos int, value T) {
	restoreTree := func(n, old *node[T]) {}
	if t.rcount > 1 {
		restoreTree = func(n, old *node[T]) {
			n.fulfillOath()
			n.copyFrom(old)

			if n.lson.right >= pos {
				(*n).rson = (*old).rson
			}
			if n.rson.left <= pos {
				(*n).lson = (*old).lson
			}
		}
	}

	t.toproot.setSegment(restoreTree, t.roots[t.rcount-1], pos, pos, value)
}

func (t *SegTree[T]) SetSegment(left, right int, value T) {
	restoreTree := func(n, old *node[T]) {}
	if t.rcount > 1 {
		restoreTree = func(n, old *node[T]) {
			n.fulfillOath()
			n.copyFrom(old)

			if n.lson.right >= left {
				(*n).rson = (*old).rson
			}
			if n.rson.left <= right {
				(*n).lson = (*old).lson
			}
		}
	}

	t.toproot.setSegment(restoreTree, t.roots[t.rcount-1], left, right, value)
}

func (n *node[T]) setSegment(restoreTree func(*node[T], *node[T]), old *node[T], left, right int, value T) {
	restoreTree(n, old)

	if n.left == n.right {
		n.data = n.data.SetVal(value)
		return
	}

	if left <= n.left && n.right <= right {
		n.data = n.data.SetVal(value)
		n.oath = func(data T) T { return data.SetVal(value) }
		return
	}

	if n.lson.right >= left {
		n.lson.setSegment(restoreTree, old, left, right, value)
	}
	if n.rson.left <= right {
		n.rson.setSegment(restoreTree, old, left, right, value)
	}

	n.data = n.data.MergeVal(n.lson.data, n.rson.data)
}

func (t SegTree[T]) Get(position int) T {
	return t.toproot.getSegment(position, position, t.initData)
}

func (t SegTree[T]) GetSegment(left, right int) T {
	return t.toproot.getSegment(left, right, t.initData)
}

func (n node[T]) getSegment(left, right int, initdata func() T) T {
	if n.left >= left && n.right <= right {
		return n.data
	}

	if n.right < left || n.left > right {
		return initdata()
	}

	return n.data.MergeVal(n.lson.getSegment(left, right, initdata), n.rson.getSegment(left, right, initdata))
}

func (n *node[T]) fulfillOath() {
	if n.oath != nil {
		n.data = n.oath(n.data)
		if n.left != n.right {
			n.lson.oath = n.oath
			n.rson.oath = n.oath
		}
		n.oath = nil
	}
}

func (n *node[T]) copyFrom(old *node[T]) {
	old.fulfillOath()
	*n = *old
	if old.lson != nil {
		n.lson = &node[T]{}
		n.rson = &node[T]{}
	}
}
