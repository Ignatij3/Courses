package main

import (
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
	conns := make([]connection, len(*nodes))
	for i := range conns {
		conns[i].length = infinity - 1
	}

	sst[0].id = (*nodes)[0].id
	conns[0].length = infinity
	updateSmallestConns(*nodes, 0, conns)
	smallest := findSmallestConns(conns)

	for smallest.length != infinity {
		sst[smallest.a.id-1].conns = append(sst[smallest.a.id-1].conns, connection{smallest.length, sst[smallest.a.id-1], sst[smallest.b.id-1]})
		sst[smallest.b.id-1].conns = append(sst[smallest.b.id-1].conns, connection{smallest.length, sst[smallest.b.id-1], sst[smallest.a.id-1]})
		conns[smallest.b.id-1].length = infinity
		updateSmallestConns(*nodes, smallest.b.id-1, conns)
		smallest = findSmallestConns(conns)
	}
	*nodes = sst
}

func updateSmallestConns(nodes graph, updatedPos int, conns []connection) {
	for _, connNode := range nodes[updatedPos].conns {
		if conns[connNode.b.id-1].length != infinity && conns[connNode.b.id-1].length > connNode.length {
			conns[connNode.b.id-1] = connNode
		}
	}
}

func findSmallestConns(conns []connection) connection {
	smallest := connection{length: infinity}
	for _, conn := range conns {
		if conn.length < smallest.length {
			smallest = conn
		}
	}
	return smallest
}

func outputConns(nodes graph) {
	stack := make([]struct{ a, b int }, 0, len(nodes)-1)
	visited := make([]bool, len(nodes))
	for i := range nodes {
		if !visited[i] {
			dfs(nodes, nodes[i], visited, &stack)
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

func dfs(nodes graph, node *node, visited []bool, stack *[]struct{ a, b int }) {
	if visited[node.id-1] {
		return
	}

	visited[node.id-1] = true
	for _, connNode := range node.conns {
		if !visited[connNode.b.id-1] {
			*stack = append(*stack, struct{ a, b int }{connNode.a.id, connNode.b.id})
			dfs(nodes, connNode.b, visited, stack)
		}
	}
}
