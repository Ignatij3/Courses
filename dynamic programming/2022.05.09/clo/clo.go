package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type strand struct {
	id   int
	ends []int
}

type sHeap []*strand

func (s sHeap) Len() int {
	return len(s)
}

func (s sHeap) Less(i, j int) bool {
	return len(s[i].ends) < len(s[j].ends)
}

func (s sHeap) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s *sHeap) Push(x any) {
	*s = append(*s, x.(*strand))
}

func (s *sHeap) Pop() any {
	old := *s
	n := len(old)
	x := old[n-1]
	*s = old[0 : n-1]
	return x
}

func getData() (*sHeap, []*strand) {
	var knots, roads int

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d %d\n", &knots, &roads)
	if roads < knots {
		return nil, nil
	}

	strands := make([]*strand, knots)
	for i := range strands {
		strands[i] = &strand{id: i + 1, ends: []int{}}
	}

	var a, b int
	for i := 0; i < roads; i++ {
		fmt.Fscanf(reader, "%d %d\n", &a, &b)
		strands[a-1].ends = append(strands[a-1].ends, b)
		strands[b-1].ends = append(strands[b-1].ends, a)
	}

	var data *sHeap = &sHeap{}
	heap.Init(data)
	for i := range strands {
		heap.Push(data, strands[i])
	}

	return data, strands
}

func main() {
	data, strands := getData()
	if strands == nil {
		fmt.Println("NIE")
		return
	}
	arrows := connectArrows(data, strands)

	if arrows == nil || len(arrows) == 0 {
		fmt.Println("NIE")
	} else {
		fmt.Println("TAK")
	}

	for i := range arrows {
		fmt.Println(arrows[i])
	}
}

func connectArrows(data *sHeap, strands []*strand) []int {
	arrows := make([]int, len(strands))

	for data.Len() > 0 {
		strand := (heap.Pop(data)).(*strand)
		if len(strand.ends) == 0 {
			return nil
		}

		end := strand.ends[0]
		arrows[strand.id-1] = end

		for i := range strands[end-1].ends {
			if strands[end-1].ends[i] == strand.id {
				strands[end-1].ends = append(strands[end-1].ends[:i], strands[end-1].ends[i+1:]...)
				break
			}
		}
		heap.Fix(data, findID(data, end))
	}

	return arrows
}

func findID(data *sHeap, id int) int {
	for i := range *data {
		if (*data)[i].id == id {
			return i
		}
	}
	return -1
}
