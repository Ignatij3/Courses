package main

import (
	"bufio"
	"fmt"
	"os"

	"./segtree"
)

type paint struct {
	isBlack     bool
	distinctAmt int
	lenSum      int
	segLen      int
	initialSeg  bool
}

func (a paint) SetVal(b paint) paint {
	a.isBlack = b.isBlack
	a.initialSeg = b.initialSeg
	if a.isBlack {
		a.distinctAmt = 1
		a.lenSum = a.segLen
	} else {
		a.distinctAmt = 0
		a.lenSum = 0
	}
	return a
}

func (data paint) MergeVal(a, b paint) paint {
	if data.initialSeg {
		data.segLen = a.segLen + b.segLen
	}

	data.isBlack = a.isBlack && b.isBlack
	if data.isBlack {
		data.distinctAmt = 1
		data.lenSum = data.segLen
	} else { //надо обрабатывать края
		data.distinctAmt = a.distinctAmt + b.distinctAmt
		data.lenSum = a.lenSum + b.lenSum
	}
	return data
}

type segment struct {
	color bool // black = true, white = false
	l, r  int
}

func getData() []*segment {
	var n int

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d\n", &n)

	var (
		color byte
		l, r  int
	)

	requests := make([]*segment, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%c %d %d\n", &color, &l, &r)
		if r < l {
			l, r = r, l
		}

		requests[i] = &segment{
			color: color == 'B',
			l:     l,
			r:     r,
		}
	}

	return requests
}

func main() {
	requests := getData()
	res := calculate(requests)
	for _, bsegs := range res {
		fmt.Println(bsegs)
	}
}

func calculate(requests []*segment) []string {
	newPaint := func() paint { return paint{false, 0, 0, 1, true} }

	l, r := getBorders(requests)
	tree := segtree.InitTree[paint](l, r, newPaint)
	for i := l; i < r; i++ {
		tree.Set(i, newPaint())
	}

	res := make([]string, len(requests))
	for i, req := range requests {
		tree.SetSegment(req.l, req.r, paint{req.color, 1, req.r - req.l + 1, req.r - req.l + 1, false})
		boolres := tree.GetSegment(tree.Left(), tree.Right())
		res[i] = fmt.Sprintf("%d %d", boolres.distinctAmt, boolres.lenSum)
	}

	return res
}

func getBorders(requests []*segment) (int, int) {
	min, max := int(1e9), 0
	for _, seg := range requests {
		if seg.l < min {
			min = seg.l
		}
		if seg.r > max {
			max = seg.r
		}
	}

	return min, max
}
