package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	N    = int(1e3)
	SIZE = 10000
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

type (
	testcase struct {
		data     []int
		requests []request
		zeroAmt  int
	}
	request struct {
		letter  byte
		a, b, c int
	}
)

func initData() ([]int, int) {
	rand.Seed(time.Now().UnixNano())
	var size int = rand.Intn(SIZE) + 1

	arr := make([]int, size)
	zeros := 0
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(3)
		if arr[i] == 0 {
			zeros++
		}
	}

	return arr, zeros
}

func generateTestcase() testcase {
	arr, zeros := initData()
	test := testcase{
		data:     arr,
		requests: make([]request, rand.Intn(SIZE)+1),
		zeroAmt:  zeros,
	}

	for i := range test.requests {
		if rand.Intn(2) == 0 {
			test.requests[i] = request{
				letter: 'U',
				a:      rand.Intn(len(test.data)) + 1,
				b:      rand.Intn(5),
				c:      0,
			}
		} else {
			l := rand.Intn(len(test.data)) + 1
			r := l + rand.Intn(len(test.data)-(l-1))
			test.requests[i] = request{
				letter: 'S',
				a:      l,
				b:      r,
				c:      l + rand.Intn(test.zeroAmt+1),
			}
		}
	}

	return test
}

func smartTest(test testcase) []int {
	tree := InitTree(len(test.data))
	for i, num := range test.data {
		tree.Put(i+1, num)
	}
	res := make([]int, 0)

	for _, req := range test.requests {
		if req.letter == 'U' {
			tree.Put(req.a, req.b)
		} else if req.letter == 'S' {
			res = append(res, tree.Get(req.a, req.b, req.c))
		}
	}

	return res
}

func stupidTest(test testcase) []int {
	res := make([]int, 0)
	test.data = append([]int{1}, test.data...)

	for _, req := range test.requests {
		if req.letter == 'U' {
			test.data[req.a] = req.b
		} else if req.letter == 'S' {
			var zeroCount, place int

			for place = req.a; place <= req.b && zeroCount < req.c; place++ {
				if test.data[place] == 0 {
					zeroCount++
				}
			}

			if zeroCount < req.c {
				res = append(res, -1)
			} else {
				res = append(res, place-1)
			}
		}
	}
	return res
}

func compare(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	var (
		correct             bool
		resSmart, resStupid []int
	)

	for i := 0; i < N; i++ {
		test := generateTestcase()
		resSmart = smartTest(test)
		resStupid = stupidTest(test)
		correct = compare(resSmart, resStupid)

		if !correct {
			fmt.Printf("Error on test %d\n", i)
			fmt.Printf("test data: %v\nanswer: %v\nresult: %v\n\n", test.data, resStupid, resSmart)
			fmt.Println("requests:")
			resCounter := 0
			for _, req := range test.requests {
				if req.letter == 'U' {
					fmt.Printf("U %d %d\n", req.a, req.b)
				} else {
					fmt.Printf("S %d %d %d (ans: %d, res: %d)\n", req.a, req.b, req.c, resStupid[resCounter], resSmart[resCounter])
					resCounter++
				}
			}
			return
		}
	}
	fmt.Printf("All %d tests passed!", N)
}
