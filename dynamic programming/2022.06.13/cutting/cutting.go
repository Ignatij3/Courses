package main

import (
	"bufio"
	"fmt"
	"os"
)

type node struct {
	id          int
	connections []*node
	visited     bool
}

type graph []*node

func getData() {
	var nodes, connections, operations int

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d %d %d\n", &nodes, &connections, &operations)
	graph := make(graph, nodes)

	for i := range graph {
		graph[i] = &node{
			id:          i + 1,
			connections: make([]*node, 0),
			visited:     false,
		}
	}

	var a, b int
	for i := 0; i < connections; i++ {
		fmt.Fscanf(reader, "%d %d\n", &a, &b)
		graph[a-1].connections = append(graph[a-1].connections, graph[b-1])
		graph[b-1].connections = append(graph[b-1].connections, graph[a-1])
	}

	var op string
	for i := 0; i < operations; i++ {
		fmt.Fscanf(reader, "%s %d %d\n", &op, &a, &b)
		if op == "ask" {
			if a == b || DFS(graph[a-1], graph[b-1]) {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}
			refresh(graph)

		} else if op == "cut" {
			cut(graph, a, b)
		}
	}
}

func DFS(node *node, target *node) bool {
	if node.visited {
		return false
	}

	node.visited = true
	for _, n := range node.connections {
		if n == target || DFS(n, target) {
			return true
		}
	}

	return false
}

func refresh(graph graph) {
	for _, n := range graph {
		n.visited = false
	}
}

func cut(graph graph, a, b int) {
	for i, conn := range graph[a-1].connections {
		if conn == graph[b-1] {
			graph[a-1].connections = append(graph[a-1].connections[:i], graph[a-1].connections[i+1:]...)
			break
		}
	}

	for i, conn := range graph[b-1].connections {
		if conn == graph[a-1] {
			graph[b-1].connections = append(graph[b-1].connections[:i], graph[b-1].connections[i+1:]...)
			break
		}
	}
}

func main() {
	getData()
}
