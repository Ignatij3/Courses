package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type pipe struct {
	vtx      *vertex
	capacity int
	flow     int
}

type network []*vertex

type vertex struct {
	conns   []pipe
	visited bool
}

func getData() network {
	var (
		vertices int
		pipes    int
	)

	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanf(reader, "%d %d\n", &vertices, &pipes)

	net := make(network, vertices+1)
	net[0] = &vertex{
		conns:   make([]pipe, 1),
		visited: false,
	}

	for i := 1; i < vertices+1; i++ {
		net[i] = &vertex{
			conns:   make([]pipe, 0, vertices),
			visited: false,
		}
	}

	net[0].conns[0] = pipe{
		vtx:      net[1],
		capacity: math.MaxInt,
		flow:     math.MaxInt,
	}
	net[1].conns = append(net[1].conns, pipe{
		vtx:      net[0],
		capacity: math.MaxInt,
		flow:     math.MaxInt,
	})

	var u, v, c int
	for i := 0; i < pipes; i++ {
		fmt.Fscanf(reader, "%d %d %d\n", &u, &v, &c)

		if pipeExists(net[u], net[v]) {
			posv := findVertice(net[u].conns, net[v])
			net[u].conns[posv].capacity += c

			posu := findVertice(net[v].conns, net[u])
			net[v].conns[posu].capacity += c

		} else {
			net[u].conns = append(net[u].conns, pipe{
				vtx:      net[v],
				capacity: c,
				flow:     0,
			})

			net[v].conns = append(net[v].conns, pipe{
				vtx:      net[u],
				capacity: c,
				flow:     0,
			})
		}
	}

	return net
}

func main() {
	net := getData()
	res := calculate(net)
	fmt.Println(res)
}

func calculate(net network) int {
	bfs(net[0], len(net))
	var res int
	for _, p := range net[len(net)-1].conns {
		res += p.flow
	}
	return res
}

func bfs(start *vertex, vertices int) {
	queue := make([]*vertex, 1, vertices)
	queue[0] = start

	for len(queue) > 0 {
		next := queue[0]
		queue = queue[1:]
		dobfs(next, &queue)
	}
}

func dobfs(node *vertex, queue *[]*vertex) {
	node.visited = true

	flow := 0
	for _, pipe := range node.conns {
		flow += pipe.flow
	}

	for i := range node.conns {
		if !node.conns[i].vtx.visited {
			*queue = append(*queue, node.conns[i].vtx)
		}

		if node.conns[i].flow == 0 {
			if node.conns[i].capacity <= flow {
				node.conns[i].flow = node.conns[i].capacity
			} else {
				node.conns[i].flow = flow
			}

			pos := findVertice(node.conns[i].vtx.conns, node)
			node.conns[i].vtx.conns[pos].flow = node.conns[i].flow
		}
	}
}

func pipeExists(a, b *vertex) bool {
	for i := range a.conns {
		if a.conns[i].vtx == b {
			return true
		}
	}
	return false
}

func findVertice(conns []pipe, v *vertex) int {
	for i := range conns {
		if conns[i].vtx == v {
			return i
		}
	}
	return -1
}
