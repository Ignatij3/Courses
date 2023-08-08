package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

type event struct {
	position int32
	time     int32
}
type border struct {
	left  int64
	right int64
}

func main() {
	events, max_speed := get_data()

	borders := init_lists(events, max_speed)
	max_arb := find_max_for_arbitrary(borders)
	max_zero := find_max_for_zero(borders)

	fmt.Println(max_zero, max_arb)
}

func get_data() ([]event, uint16) {
	var (
		event_amount uint32
		max_speed    uint16
	)

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d\n", &event_amount)
	events := make([]event, event_amount)

	for i := range events {
		fmt.Fscanf(reader, "%d %d\n", &events[i].position, &events[i].time)
	}

	fmt.Fscanf(reader, "%d\n", &max_speed)
	return events, max_speed
}

func init_lists(events []event, max_speed uint16) []border {
	borders := make([]border, len(events))

	for i, event := range events {
		borders[i].left = int64(event.position - (event.time * int32(max_speed)))
		borders[i].right = int64(event.position + (event.time * int32(max_speed)))
	}

	sort.Slice(borders, func(i, j int) bool { return borders[i].left >= borders[j].left })
	sort.Slice(borders, func(i, j int) bool { return borders[i].left == borders[j].left && borders[i].right <= borders[j].right })

	return borders
}

func find_max_for_arbitrary(borders []border) int {
	return find_LIS(borders)
}

func find_max_for_zero(borders []border) int {
	var neg_num int
	for borders[neg_num].left > 0 {
		neg_num++
	}

	borders = append(borders, border{})
	copy(borders[neg_num+1:], borders[neg_num:])
	borders[neg_num] = border{0, 0}

	return find_LIS(borders) - 1
}

func find_LIS(borders []border) int {
	lis_with_len := make([]int64, 1)
	lis_with_len[0] = math.MinInt64

	for _, border := range borders {
		pos := get_pos_of_smaller_num(lis_with_len, border.right)

		if pos == len(lis_with_len)-1 {
			lis_with_len = append(lis_with_len, border.right)
		} else {
			lis_with_len[pos+1] = border.right
		}
	}

	return len(lis_with_len) - 1
}

func get_pos_of_smaller_num(lis_with_len []int64, num int64) int {
	var (
		pos         int
		left, right int = 0, len(lis_with_len) - 1
		center      int
	)

	if num > lis_with_len[right] {
		pos = right
	} else {
		for right-left > 1 {
			center = (left + right) / 2

			if lis_with_len[center] < num {
				left = center
			} else {
				right = center
			}
		}

		pos = left
	}

	return pos
}
