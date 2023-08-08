package segtree

// TreeAble is interface which must be implemented by the type for the tree to be able to contain it.
type TreeAble[T any] interface {
	// SetVal receives argument of Set() function, the method is called on data.
	// It must return the expected data to be in the node.
	SetVal(T) T
	// MergeVal receives lhs and rhs, the method is called on data.
	// It must return the expected result when merging data from left and right node in tree.
	MergeVal(T, T) T
}
