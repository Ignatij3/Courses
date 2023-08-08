package main

import (
	"bufio"
	"fmt"
	"os"
)

type query struct {
	l, r, k int
}

type PersisSegTree struct {
	root *node
}

type node struct {
	a, b int
	l, r *node
	sum  int
}

func InitTree(n int) *PersisSegTree {
	if n < 1 {
		n = 1
	}

	tree := new(PersisSegTree)
	tree.root = &node{a: 1, b: n}
	tree.root.initSegment()
	return tree
}

func (t *PersisSegTree) NewLevel(place, num int) *PersisSegTree {
	newtree := InitTree(1)
	t.root.putToLevel(newtree.root, place, num)
	return newtree
}

func (t *PersisSegTree) Put(place, num int) {
	t.root.put(place, num)
}

func (t PersisSegTree) Get(a, b int) int {
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

func (n *node) putToLevel(newNode *node, place, num int) {
	if n.a == n.b {
		*newNode = node{
			a:   n.a,
			b:   n.b,
			l:   nil,
			r:   nil,
			sum: n.sum + 1,
		}
		return
	}

	*newNode = node{
		a:   n.a,
		b:   n.b,
		l:   &node{},
		r:   &node{},
		sum: n.sum + 1,
	}

	if n.l.b >= place {
		n.l.putToLevel(newNode.l, place, num)
		(*newNode).r = (*n).r
	} else {
		(*newNode).l = (*n).l
		n.r.putToLevel(newNode.r, place, num)
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

func getData() ([]int, []query) {
	var n int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanf(reader, "%d\n", &n)
	arr := make([]int, n)

	for i := range arr {
		fmt.Fscanf(reader, "%d", &arr[i])
	}
	fmt.Fscanf(reader, "\n")

	fmt.Fscanf(reader, "%d\n", &n)
	queries := make([]query, n)

	for i := range queries {
		fmt.Fscanf(reader, "%d %d %d\n", &queries[i].l, &queries[i].r, &queries[i].k)
	}

	return arr, queries
}

func main() {
	arr, queries := getData()
	for _, n := range calculate(arr, queries) {
		fmt.Println(n)
	}
}

func calculate(arr []int, queries []query) []int {
	tree := InitTree(len(arr))
	for i, num := range arr {
		if num == 0 {
			tree.Put(i+1, 1)
		} else {
			tree.Put(i+1, 0)
		}
	}

	level := 1
	res := make([]int, len(queries))

	trees := make([]*PersisSegTree, len(arr)+1)
	trees[level] = tree

	for i, qry := range queries {
		if level < qry.k {
			for level < qry.k {
				level++
				increaseLevel(trees, arr, qry, level)
			}
		}
		res[i] = trees[qry.k].Get(qry.l, qry.r)
	}

	return res
}

func increaseLevel(trees []*PersisSegTree, arr []int, qry query, level int) {
	tempOld := trees[level-1]
	for i, num := range arr {
		if num == level-1 {
			tempOld = tempOld.NewLevel(i+1, 1)
		}
	}
	trees[level] = tempOld
}
