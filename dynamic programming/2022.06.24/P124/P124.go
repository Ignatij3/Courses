package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type graph []*node
type sortBytes []byte

type node struct {
	letter  byte
	relPos  int
	next    []*node
	visited bool
	taken   bool
}

func (s sortBytes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortBytes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortBytes) Len() int {
	return len(s)
}

func getData() []graph {
	vars := make([]byte, 0)
	orderSets := make([]graph, 0)
	reader := bufio.NewReader(os.Stdin)

	var (
		str []byte
		err error
	)
	for err == nil {
		str, _, _ = reader.ReadLine()
		vars = parseVars(str)

		str, _, err = reader.ReadLine()
		orderSets = append(orderSets, parseRestrictions(vars, str))
		if str[len(str)-1] == '\x04' {
			break
		}
	}

	return orderSets
}

func parseVars(str []byte) []byte {
	vars := make([]byte, (len(str)/2)+1)
	for i, j := 0, 0; i < len(str); i, j = i+2, j+1 {
		vars[j] = str[i]
	}
	return vars
}

func parseRestrictions(vars []byte, str []byte) graph {
	orders := initOrders(vars)

	for i := 0; i < len(str); i += 4 {
		lhs := orders[findLetter(vars, str[i])]
		rhs := orders[findLetter(vars, str[i+2])]

		if !lhs.taken && !rhs.taken {
			rhs.relPos = 1
			rhs.taken = true

			lhs.relPos = 0
			lhs.taken = true
			lhs.next = []*node{rhs}

		} else if !lhs.taken {
			lhs.relPos = rhs.relPos - 1
			lhs.taken = true
			lhs.next = []*node{rhs}

			if rhs.relPos == 0 {
				lhs.relPos = 0
				elevatePositions(lhs)
			}

		} else if !rhs.taken {
			rhs.relPos = lhs.relPos + 1
			rhs.taken = true
			rhs.next = []*node{}

			lhs.next = append(lhs.next, rhs)

		} else {
			lhs.next = append(lhs.next, rhs)
			rhs.relPos = lhs.relPos + 1
			elevatePositions(rhs)
		}
	}

	return orders
}

func elevatePositions(letter *node) {
	for _, nextWord := range letter.next {
		nextWord.relPos = letter.relPos + 1
		elevatePositions(nextWord)
	}
}

func initOrders(vars []byte) graph {
	orders := make(graph, len(vars))
	for i, letter := range vars {
		orders[i] = &node{
			letter:  letter,
			relPos:  0,
			next:    nil,
			visited: false,
			taken:   false,
		}
	}
	return orders
}

func findLetter(vars []byte, lttr byte) int {
	for i, word := range vars {
		if word == lttr {
			return i
		}
	}
	return -1
}

func main() {
	orderSets := getData()
	for i := range orderSets {
		calculate(orderSets[i])
	}
}

func calculate(orders graph) {
	orderedLetters, freeLetters := topologicalSort(orders)
	letterSets := splitIntoSets(orderedLetters)
	sortLetters(letterSets, freeLetters)

	createPermutations(letterSets, freeLetters)
}

func createPermutations(letterSets []string, freeLetters []byte) {
	cache := make([][]string, len(letterSets))
	for i := range letterSets {
		cache[i] = permutations(letterSets[i])
	}

	writer := bufio.NewWriter(os.Stdout)
	positions := make([]int, len(cache))

	var (
		str string
		i   int
	)
	for positions[0] < len(cache[0]) {
		str = ""
		for i := range cache {
			str += cache[i][positions[i]]
		}
		fmt.Fprintf(writer, "%s\n", str)

		for i = len(positions) - 1; i >= 0; i-- {
			positions[i]++
			if positions[i] < len(cache[i]) {
				break
			}
			positions[i] = 0
		}
		if i == -1 {
			break
		}
	}

	writer.Flush()
	fmt.Println()
}

func topologicalSort(grph graph) ([]*node, []byte) {
	res := make([]*node, 0)
	free := make([]byte, 0)
	stack := make([]*node, 0)

	for i := range grph {
		if !grph[i].taken {
			free = append(free, grph[i].letter)
			continue
		}
		if !grph[i].visited {
			dfs(grph[i], &stack)
		}
	}

	for i := len(stack) - 1; i >= 0; i-- {
		res = append(res, stack[i])
	}

	return res, free
}

func dfs(elem *node, stack *[]*node) {
	elem.visited = true
	for i := range elem.next {
		if !elem.next[i].visited {
			dfs(elem.next[i], stack)
		}
	}
	*stack = append(*stack, elem)
}

func splitIntoSets(grph graph) []string {
	letterSets := make([]string, 0)

	var (
		lset            string
		currPos, oldPos int
	)
	for _, currnode := range grph {
		currPos = currnode.relPos

		if currPos != oldPos {
			letterSets = append(letterSets, lset)
			lset = ""
		}

		lset += string(currnode.letter)
		oldPos = currPos
	}

	return append(letterSets, lset)
}

func sortLetters(letterSets []string, freeLetters []byte) {
	for i := range letterSets {
		letterSets[i] = sortString(letterSets[i])
	}
	sort.Sort(sortBytes(freeLetters))
}

func sortString(s string) string {
	r := []byte(s)
	sort.Sort(sortBytes(r))
	return string(r)
}

func permutations(testStr string) []string {
	var n func(testStr []byte, p []string) []string
	n = func(testStr []byte, p []string) []string {
		if len(testStr) == 0 {
			return p
		} else {
			result := []string{}
			for _, e := range p {
				result = append(result, join([]byte(e), testStr[0])...)
			}
			return n(testStr[1:], result)
		}
	}

	output := []byte(testStr)
	return n(output[1:], []string{string(output[0])})
}

func join(ins []byte, c byte) (result []string) {
	for i := 0; i <= len(ins); i++ {
		result = append(result, string(ins[:i])+string(c)+string(ins[i:]))
	}
	return
}
