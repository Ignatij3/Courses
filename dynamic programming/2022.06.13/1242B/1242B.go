package main

import (
	"bufio"
	"fmt"
	"os"
)

type path struct {
	a, b int
}

func getData() (int, map[path]bool) {
	var n, m int

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d %d\n", &n, &m)

	graph := make(map[path]bool, n)

	var u, v int
	for i := 0; i < m; i++ {
		fmt.Fscanf(reader, "%d %d\n", &u, &v)
		graph[path{u, v}] = true
		graph[path{v, u}] = true
	}

	return n, graph
}

func main() {
	n, graph := getData()
	res := calculate(n, graph)
	fmt.Println(res)
}

func calculate(qlen int, graph map[path]bool) int {
	queue := make([]int, qlen)
	compt := make([]int, 0, qlen)
	for i := range queue {
		queue[i] = i + 1
	}

	comptCount := 0
	for qlen > 0 {
		compt = compt[:0]
		compt = append(compt, queue[qlen-1])
		comptCount++
		qlen--

		for j := 0; j < len(compt); j++ {
			for k := 0; k < qlen; {
				if _, ok := graph[path{compt[j], queue[k]}]; !ok {
					compt = append(compt, queue[k])
					queue[k] = queue[qlen-1]
					qlen--
				} else {
					k++
				}
			}
		}
	}

	return comptCount - 1
}
