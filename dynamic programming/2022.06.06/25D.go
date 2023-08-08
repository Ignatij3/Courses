package main

import (
	"bufio"
	"fmt"
	"os"
)

type root struct {
	id   int
	in   []*node
	size int
}

type node struct {
	id   int
	root *root
}

type vertex interface {
	getRoot() *root
	getID() int
}

type change struct {
	closeA, closeB int
	openA, openB   int
}

var result []change = make([]change, 0)

func (r *root) getRoot() *root {
	return r
}

func (r *root) getID() int {
	return r.id
}

func (n node) getRoot() *root {
	return n.root
}

func (n node) getID() int {
	return n.id
}

func buildRoad(vertices []vertex, a, b int) {
	vertA := &vertices[a-1]
	vertB := &vertices[b-1]
	rootA := (*vertA).getRoot()
	rootB := (*vertB).getRoot()

	if rootA == nil && rootB == nil {
		vertices[a-1] = &root{id: (*vertA).getID(), in: []*node{(*vertB).(*node)}, size: 2}
		vertices[b-1].(*node).root = (*vertA).(*root)
	} else if rootA == nil {
		rootB.in = append(rootB.in, (*vertA).(*node))
		vertices[a-1].(*node).root = rootB
		rootB.size++
	} else if rootB == nil {
		rootA.in = append(rootA.in, (*vertB).(*node))
		vertices[b-1].(*node).root = rootA
		rootA.size++
	} else if rootA.id == rootB.id {
		result = append(result, change{
			closeA: (*vertA).getID(),
			closeB: (*vertB).getID(),
			openA:  0, openB: 0,
		})
	} else {
		if rootB.size < rootA.size {
			rootA, rootB = rootB, rootA
		}
		mergeComponents(rootA, rootB)

		rootA.in = append(rootA.in, &node{id: rootB.id, root: rootA})
		if (*vertB).getID() == rootB.id {
			vertices[b-1] = rootA.in[rootA.size-2]
		} else {
			vertices[rootB.id-1] = rootA.in[rootA.size-2]
		}
	}
}

func mergeComponents(a, b *root) {
	for _, node := range b.in {
		node.root = a
		a.in = append(a.in, node)
	}
	a.size += b.size
}

func getData() []vertex {
	var n int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d\n", &n)

	vertices := make([]vertex, n)
	for i := 0; i < n; i++ {
		vertices[i] = &node{
			id:   i + 1,
			root: nil,
		}
	}

	for i := 0; i < n-1; i++ {
		var a, b int
		fmt.Fscanf(reader, "%d %d\n", &a, &b)
		buildRoad(vertices, a, b)
	}

	return vertices
}

func main() {
	vertices := getData()
	calculate(vertices)

	fmt.Println(len(result))
	for _, change := range result {
		fmt.Printf("%d %d %d %d\n", change.closeA, change.closeB, change.openA, change.openB)
	}
}

func calculate(vertices []vertex) {
	components := make([]vertex, 0)
	for _, node := range vertices {
		if node.getRoot() == nil {
			components = append(components, node)
		} else if node.getRoot().id == node.getID() {
			components = append(components, node.getRoot())
		}
	}

	for i := range result { //компонент должно быть на 1 больше, чем лишних дорог
		result[i].openA = components[i].getID()
		result[i].openB = components[i+1].getID()
	}
}
