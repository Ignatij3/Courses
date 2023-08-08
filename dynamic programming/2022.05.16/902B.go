package main

import (
	"bufio"
	"fmt"
	"os"
)

type tree struct {
	id        int
	color     int
	cumDifAmt int
	children  []*tree
}

func (t *tree) insertAt(nodeId, color, newId int) {
	if t.id == nodeId {
		t.children = append(t.children, &tree{id: newId, color: color, cumDifAmt: 1, children: nil})
		return
	}

	t.insertInChild(nodeId, color, newId)
}

// Возвращает, всунул ли он узел в ребёнка, сделано для эффективности
func (t *tree) insertInChild(nodeId, color, newId int) bool {
	for i := range t.children {
		if t.children[i].id == nodeId {
			t.children[i].children = append(t.children[i].children, &tree{id: newId, color: color, cumDifAmt: 1, children: nil})
			return true
		}
	}

	for _, child := range t.children {
		if child.insertInChild(nodeId, color, newId) {
			return true
		}
	}

	return false
}

func (t *tree) setCumulativeDifferentColorAmount() {
	if t.children == nil {
		return
	}

	for _, child := range t.children {
		child.setCumulativeDifferentColorAmount()
		t.cumDifAmt += child.cumDifAmt
		if t.color == child.color {
			t.cumDifAmt--
		}
	}

}

func getData() *tree {
	var (
		nodeAmt int
		nodes   *tree = &tree{id: 1, color: 0, cumDifAmt: 1, children: []*tree{}}
	)

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d\n", &nodeAmt)

	strands := make([]int, nodeAmt-1)
	for i := 0; i < nodeAmt-1; i++ {
		fmt.Fscanf(reader, "%d", &strands[i])
	}

	fmt.Fscanf(reader, "\n")

	colors := make([]int, nodeAmt)
	for i := 0; i < nodeAmt; i++ {
		fmt.Fscanf(reader, "%d", &colors[i])
	}

	nodes.color = colors[0]
	for pos, strand := range strands {
		nodes.insertAt(strand, colors[pos+1], pos+2)
	}

	return nodes
}

func main() {
	nodes := getData()
	res := calculate(nodes)
	fmt.Println(res)
}

func calculate(nodes *tree) int {
	nodes.setCumulativeDifferentColorAmount()
	return nodes.cumDifAmt
}
