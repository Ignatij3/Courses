package main

import (
	"bufio"
	"fmt"
	"os"
)

var catTolerance int

type node struct {
	id        int
	cat       bool
	neighbors []*node
}

func getData() []*node {
	var nodeCount int
	buf := bufio.NewReader(os.Stdin)

	fmt.Fscanf(buf, "%d %d\n", &nodeCount, &catTolerance)
	nodes := make([]*node, nodeCount)

	var cat int
	for i := range nodes {
		nodes[i] = &node{id: i}
		fmt.Fscanf(buf, "%d", &cat)
		if cat == 1 {
			nodes[i].cat = true
		}
	}
	fmt.Fscanf(buf, "\n")

	var node1, node2 int
	for i := 0; i < nodeCount-1; i++ {
		fmt.Fscanf(buf, "%d %d\n", &node1, &node2)
		nodes[node1-1].neighbors = append(nodes[node1-1].neighbors, nodes[node2-1])
		nodes[node2-1].neighbors = append(nodes[node2-1].neighbors, nodes[node1-1])
	}

	return nodes
}

func main() {
	nodes := getData()
	res := calculate(nil, nodes[0], 0)
	fmt.Println(res)
}

func calculate(parent *node, vertice *node, cats int) (res int) {
	if vertice.cat {
		cats++
	} else {
		cats = 0
	}

	if cats > catTolerance {
		return 0
	}

	if len(vertice.neighbors) == 1 && vertice.neighbors[0] == parent {
		return 1
	}

	for _, neighbor := range vertice.neighbors {
		if neighbor != parent {
			res += calculate(vertice, neighbor, cats)
		}
	}

	return res
}
