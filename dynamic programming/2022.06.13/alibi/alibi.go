package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type graph [][]int64

//all objections are inclusive
type timeline struct {
	objections []trange
	left       int64
	right      int64
}

type murder struct {
	square   int64
	earliest int64
	latest   int64
}

type evidence struct {
	square int64
	time   int64
}

type trange struct {
	a, b int64
}

const infinity = math.MaxInt64

func NewTimeline(left, right int64) *timeline {
	return &timeline{
		objections: make([]trange, 0, 1000),
		left:       left,
		right:      right,
	}
}

func (t *timeline) AddObjectedRange(a, b int64) {
	if a > t.right || b < t.left {
		return
	}

	if t.left > a {
		a = t.left
	}
	if t.right < b {
		b = t.right
	}

	for i := 0; i < len(t.objections); i++ {
		if t.objections[i].a-1 > b { // 1
			t.objections = append(t.objections[:i+1], t.objections[i:]...)
			t.objections[i] = trange{a, b}
			return

		} else if (t.objections[i].a >= a && t.objections[i].a-1 <= b) && t.objections[i].b > b { // 2
			t.objections[i].a = a
			return

		} else if t.objections[i].a >= a && t.objections[i].b <= b { // 6
			t.objections = append(t.objections[:i], t.objections[i+1:]...)
			i--

		} else if t.objections[i].a < a && t.objections[i].b > b { // 3
			return

		} else if t.objections[i].a < a && (t.objections[i].b+1 >= a && t.objections[i].b <= b) { // 4
			t.objections[i].b = b
			for j := i + 1; j < len(t.objections); j++ {
				if t.objections[j].a-1 <= b && t.objections[j].b <= b { // 4-6
					t.objections = append(t.objections[:j], t.objections[j+1:]...)
					j--
				} else if t.objections[j].a-1 <= b && t.objections[j].b > b { // 4-2
					t.objections[i].b = t.objections[j].b
					t.objections = append(t.objections[:j], t.objections[j+1:]...)
					break
				}
			}
			return
		}
	}

	t.objections = append(t.objections, trange{a, b})
}

func (t timeline) IsAllowedInRange(a, b int64) bool {
	if len(t.objections) != 0 {
		for pos := 0; pos < len(t.objections) && t.objections[pos].a <= b; pos++ {
			if t.objections[pos].a <= a && t.objections[pos].b >= b {
				return false
			}
		}
	}
	return a >= t.left && b <= t.right
}

func getData() (graph, murder, []evidence, bool) {
	var n, m int

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d %d\n", &n, &m)

	squares := make(graph, n)
	for i := range squares {
		squares[i] = make([]int64, n)
	}

	var a, b, roadlen int64
	for i := 0; i < m; i++ {
		fmt.Fscanf(reader, "%d %d %d\n", &a, &b, &roadlen)
		squares[a-1][b-1] = roadlen
		squares[b-1][a-1] = roadlen
	}

	var homicide murder
	fmt.Fscanf(reader, "%d %d %d\n", &homicide.square, &homicide.earliest, &homicide.latest)

	fmt.Fscanf(reader, "%d\n", &n)
	evidenceList := make([]evidence, n)

	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d %d\n", &evidenceList[i].square, &evidenceList[i].time)
		if evidenceList[i].square == homicide.square && ((homicide.earliest <= evidenceList[i].time && evidenceList[i].time <= homicide.latest) || m == 0) {
			fmt.Println("YES")
			return nil, murder{}, nil, true
		}
	}

	if m == 0 {
		fmt.Println("NO")
	}
	return squares, homicide, evidenceList, m == 0
}

func main() {
	squares, homicide, evidenceList, prematureDiscontinuation := getData()
	if prematureDiscontinuation {
		return
	}

	res := calculate(squares, homicide, evidenceList)
	if res {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func calculate(squares graph, homicide murder, evidenceList []evidence) bool {
	distances := dijkstra(squares, homicide.square-1)
	killtmline := NewTimeline(homicide.earliest, homicide.latest)

	for _, ev := range evidenceList {
		if distances[ev.square-1] == infinity {
			return false
		}
		killtmline.AddObjectedRange(ev.time-int64(distances[ev.square-1])+1, ev.time+int64(distances[ev.square-1])-1)
	}

	return killtmline.IsAllowedInRange(homicide.earliest, homicide.latest)
}

func dijkstra(squares graph, src int64) []int64 {
	nodeamt := len(squares)
	dist := make([]int64, nodeamt)
	sptSet := make([]bool, nodeamt)

	for i := 0; i < nodeamt; i++ {
		dist[i] = infinity
		sptSet[i] = false
	}

	dist[src] = 0

	for count := 0; count < nodeamt-1; count++ {
		u := minDistance(dist, sptSet, nodeamt)
		sptSet[u] = true
		for v := 0; v < nodeamt; v++ {
			if !sptSet[v] && squares[u][v] != 0 && dist[u] != infinity && dist[u]+squares[u][v] < dist[v] {
				dist[v] = dist[u] + squares[u][v]
			}
		}
	}

	return dist
}

func minDistance(dist []int64, sptSet []bool, nodeamt int) int64 {
	var (
		min       int64 = infinity
		min_index int64
	)

	for v := 0; v < nodeamt; v++ {
		if !sptSet[v] && dist[v] <= min {
			min = dist[v]
			min_index = int64(v)
		}
	}

	return min_index
}
