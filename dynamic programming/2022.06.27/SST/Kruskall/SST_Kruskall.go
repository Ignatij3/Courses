package main

import (
	"fmt"
	"sort"
)

type graph []*node

type node struct {
	id    int
	conns []*node
}

type connection struct {
	weight int
	a, b   int
}

func newGraph(size int) graph {
	gph := make(graph, size)
	for i := range gph {
		gph[i] = &node{
			id:    i + 1,
			conns: make([]*node, 0),
		}
	}
	return gph
}

func getData() (graph, []connection) {
	var n, m, a, b int

	fmt.Scanf("%d %d\n", &n, &m)
	nodes := newGraph(n)
	conns := make([]connection, m)

	for i := range conns {
		fmt.Scanf("%d %d %d\n", &conns[i].a, &conns[i].b, &conns[i].weight)
		a, b = conns[i].a, conns[i].b
		nodes[a-1].conns = append(nodes[a-1].conns, nodes[b-1])
		nodes[b-1].conns = append(nodes[b-1].conns, nodes[a-1])
	}

	return nodes, conns
}

func main() {
	nodes, conns := getData()
	sst := findSST(&nodes, conns)
	sort.Slice(sst, func(i, j int) bool { return sst[i].a < sst[j].a || (sst[i].a == sst[j].a && sst[i].b < sst[j].b) })
	for _, conn := range sst {
		fmt.Printf("%d %d\n", conn.a, conn.b)
	}
}

func findSST(nodes *graph, conns []connection) []connection {
	sort.Slice(conns, func(i, j int) bool { return conns[i].weight < conns[j].weight })
	*nodes = newGraph(len(*nodes))
	sst := make([]connection, 0, len(*nodes)-1)

	for i := range conns {
		if !sameComponent(*nodes, (*nodes)[conns[i].a-1], (*nodes)[conns[i].b-1]) {
			(*nodes)[conns[i].a-1].conns = append((*nodes)[conns[i].a-1].conns, (*nodes)[conns[i].b-1])
			(*nodes)[conns[i].b-1].conns = append((*nodes)[conns[i].b-1].conns, (*nodes)[conns[i].a-1])
			sst = append(sst, conns[i])
		}
		if len(sst) == cap(sst) {
			break
		}
	}

	return sst
}

func sameComponent(nodes graph, a, b *node) bool {
	if (len(a.conns) == 0 || len(b.conns) == 0) && a != b {
		return false
	}

	visited := make([]bool, len(nodes))
	for i := range nodes {
		dfs(nodes, nodes[i], visited)
		if visited[a.id-1] || visited[b.id-1] {
			return visited[a.id-1] && visited[b.id-1]
		}
	}

	return false
}

func dfs(nodes graph, node *node, visited []bool) {
	if visited[node.id-1] {
		return
	}

	visited[node.id-1] = true
	for _, connNode := range node.conns {
		dfs(nodes, connNode, visited)
	}
}
