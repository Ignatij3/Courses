package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type elem struct {
	value int
	pos   int
}

type node struct {
	val     *elem
	conns   []*node
	visited bool
}

type descElem []*elem

func (d descElem) Len() int {
	return len(d)
}

func (d descElem) Less(i, j int) bool {
	return d[i].value > d[j].value
}

func (d descElem) Swap(i, j int) {
	d[i].value, d[j].value = d[j].value, d[i].value
	d[i].pos, d[j].pos = d[j].pos, d[i].pos
}

func getData() ([]*elem, []*node) {
	var n, m int

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d %d\n", &n, &m)

	elems := make([]*elem, n)
	for i := 0; i < n; i++ {
		elems[i] = &elem{pos: i}
		fmt.Fscanf(reader, "%d ", &(*elems[i]).value)
	}
	fmt.Fscanf(reader, "\n")

	nodes := make([]*node, n)
	for i := 0; i < n; i++ {
		nodes[i] = &node{
			val:     elems[i],
			conns:   make([]*node, 0),
			visited: false,
		}
	}

	var a, b int
	for i := 0; i < m; i++ {
		fmt.Fscanf(reader, "%d %d\n", &a, &b)
		nodes[a-1].conns = append(nodes[a-1].conns, nodes[b-1])
		nodes[b-1].conns = append(nodes[b-1].conns, nodes[a-1])
	}

	return elems, nodes
}

func main() {
	elems, nodes := getData()
	calculate(nodes)

	writer := bufio.NewWriter(os.Stdout)
	for i := range elems {
		fmt.Fprintf(writer, "%d ", elems[i].value)
	}
	writer.Flush()
}

func calculate(nodes []*node) {
	stack := make([]*elem, 0)

	for i := range nodes {
		if !nodes[i].visited {
			stack = []*elem{}
			DFS(nodes[i], &stack)
			sortElements(append(make([]*elem, 0), stack...))
		}
	}
}

func DFS(node *node, stack *[]*elem) {
	if node.visited {
		return
	}

	*stack = append(*stack, node.val)

	node.visited = true
	for _, con := range node.conns {
		if !con.visited {
			DFS(con, stack)
		}
	}
}

func sortElements(component []*elem) {
	sort.Slice(component, func(a, b int) bool {
		return component[a].pos < component[b].pos
	})
	sort.Sort(descElem(component))
}
