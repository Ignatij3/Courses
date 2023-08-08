package main

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
)

type graph []*node

const infinity = math.MaxInt32

type node struct {
	id    int
	conns []connection
}

type connection struct {
	length int
	a, b   *node
}

type ConnHeap []connection

func (h ConnHeap) Len() int           { return len(h) }
func (h ConnHeap) Less(i, j int) bool { return h[i].length < h[j].length }
func (h ConnHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *ConnHeap) Push(x any) {
	*h = append(*h, x.(connection))
}

func (h *ConnHeap) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func newGraph(size int) graph {
	gph := make(graph, size)
	for i := range gph {
		gph[i] = &node{
			id:    i + 1,
			conns: make([]connection, 0),
		}
	}
	return gph
}

func getData() graph {
	var n, m int

	fmt.Scanf("%d %d\n", &n, &m)
	nodes := newGraph(n)

	var a, b, weight int
	for i := 0; i < m; i++ {
		fmt.Scanf("%d %d %d\n", &a, &b, &weight)
		nodes[a-1].conns = append(nodes[a-1].conns, connection{weight, nodes[a-1], nodes[b-1]})
		nodes[b-1].conns = append(nodes[b-1].conns, connection{weight, nodes[b-1], nodes[a-1]})
	}

	return nodes
}

func main() {
	nodes := getData()
	convertToSST(&nodes)
	outputConns(nodes)
}

func convertToSST(nodes *graph) {
	sst := newGraph(len(*nodes))
	conns := make(ConnHeap, 0)
	heap.Init(&conns)

	sst[0].id = (*nodes)[0].id
	updateSmallestConns(*nodes, sst, 0, &conns)
	smallest := heap.Pop(&conns).(connection)

	for smallest.length != infinity {
		sst[smallest.a.id-1].conns = append(sst[smallest.a.id-1].conns, connection{smallest.length, sst[smallest.a.id-1], sst[smallest.b.id-1]})
		sst[smallest.b.id-1].conns = append(sst[smallest.b.id-1].conns, connection{smallest.length, sst[smallest.b.id-1], sst[smallest.a.id-1]})
		fixHeapLengths(&conns, sst)
		updateSmallestConns(*nodes, sst, smallest.b.id-1, &conns)
		smallest = heap.Pop(&conns).(connection)
	}
	*nodes = sst
}

func fixHeapLengths(conns *ConnHeap, sst graph) {
	for i := len(*conns) - 1; i >= 0; i-- {
		if (*conns)[i].length != infinity && sameComponent(sst, sst[0], (*conns)[i].b) {
			(*conns)[i].length = infinity
			heap.Fix(conns, i)
		}
	}
}

func updateSmallestConns(nodes, sst graph, updatedPos int, conns *ConnHeap) {
	for _, connNode := range nodes[updatedPos].conns {
		if !sameComponent(sst, connNode.a, connNode.b) {
			heap.Push(conns, connNode)
		}
	}
}

func outputConns(nodes graph) {
	stack := make([]struct{ a, b int }, 0, len(nodes)-1)
	visited := make([]bool, len(nodes))
	for i := range nodes {
		if !visited[i] {
			dfs(nodes, nodes[i], visited, func(connNode connection) {
				stack = append(stack, struct{ a, b int }{connNode.a.id, connNode.b.id})
			})
		}
	}

	for i := range stack {
		if stack[i].b < stack[i].a {
			stack[i].a, stack[i].b = stack[i].b, stack[i].a
		}
	}

	sort.Slice(stack, func(i, j int) bool {
		return stack[i].a < stack[j].a || (stack[i].a == stack[j].a && stack[i].b < stack[j].b)
	})

	for _, conn := range stack {
		fmt.Printf("%d %d\n", conn.a, conn.b)
	}
}

func sameComponent(nodes graph, a, b *node) bool {
	if (len(a.conns) == 0 || len(b.conns) == 0) && a != b {
		return false
	}

	visited := make([]bool, len(nodes))
	for i := range nodes {
		dfs(nodes, nodes[i], visited, func(connNode connection) {})
		if visited[a.id-1] || visited[b.id-1] {
			return visited[a.id-1] && visited[b.id-1]
		}
	}

	return false
}

func dfs(nodes graph, node *node, visited []bool, process func(connNode connection)) {
	if visited[node.id-1] {
		return
	}

	visited[node.id-1] = true
	for _, connNode := range node.conns {
		if !visited[connNode.b.id-1] {
			process(connNode)
			dfs(nodes, connNode.b, visited, process)
		}
	}
}
