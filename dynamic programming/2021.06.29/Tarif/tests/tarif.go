package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type list []revenue
type revenue struct {
	sum        int
	components []int
}

func main() {
	tariffs, plans := get_list()
	best_revenue := get_best_payment(tariffs, plans)
	print_answer(best_revenue, plans)
}

func get_list() ([]int, int) {
	var (
		abonents int
		plans    int
	)

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d %d\n", &abonents, &plans)
	tariffs := make([]int, abonents)

	for i := 0; i < abonents; i++ {
		fmt.Fscanf(reader, "%d ", &tariffs[i])
	}

	return tariffs, plans
}

func get_best_payment(tariffs []int, plans int) revenue {
	sort.Slice(tariffs, func(i, j int) bool { return tariffs[i] < tariffs[j] })

	revenue_matrix := make([]list, plans)
	revenue_matrix = calculate_with_n_plans(tariffs, revenue_matrix, plans-1)

	max_revenue := find_max_revenue(revenue_matrix[plans-1])

	return max_revenue
}

func calculate_with_n_plans(tariffs []int, revenue_matrix []list, plans int) []list {
	if plans > 0 {
		revenue_matrix = calculate_with_n_plans(tariffs, revenue_matrix, plans-1)
		revenue_matrix[plans] = calculate_max_revenues(&revenue_matrix[plans], revenue_matrix[plans-1], tariffs, plans)
		return revenue_matrix
	}

	revenue_matrix[0] = calculate_max_revenues(&revenue_matrix[0], nil, tariffs, 0)
	return revenue_matrix
}

func calculate_max_revenues(rev *list, prev_rev list, tariffs []int, plans int) list {
	if plans > 0 {
		*rev = append(*rev, revenue{
			sum:        tariffs[len(tariffs)-1],
			components: []int{tariffs[len(tariffs)-1]}})

		for pos := len(tariffs) - 2; pos >= 0; pos-- {
			make_new_rev_list(rev, prev_rev, tariffs, pos)
		}

	} else {
		make_first_rev_list(rev, tariffs)
	}

	return *rev
}

func make_new_rev_list(rev *list, prev_rev list, tariffs []int, pos int) {
	var (
		max_rev    int = -1
		total      int
		components []int
	)

	total = tariffs[pos]
	for rev_pos := pos + 1; rev_pos <= len(tariffs); rev_pos++ {
		if rev_pos == len(tariffs) {
			if total > max_rev {
				max_rev = total
				components = []int{tariffs[pos]}
			}

		} else if total+prev_rev[rev_pos].sum > max_rev {
			max_rev = total + prev_rev[rev_pos].sum
			components = *new([]int)
			components = append(prev_rev[rev_pos].components, components...)
			components = append([]int{tariffs[pos]}, components...)

		}
		total += tariffs[pos]
	}

	*rev = append(list{revenue{sum: max_rev, components: components}}, *rev...)
}

func make_first_rev_list(rev *list, tariffs []int) {
	var (
		people_amount int = len(tariffs)
		total         int
	)

	for _, tariff := range tariffs {
		total = tariff * people_amount
		*rev = append(*rev, revenue{sum: total, components: []int{tariff}})
		people_amount--
	}
}

func find_max_revenue(rev_list list) revenue {
	var max_rev_pos int = 0
	for pos, rev := range rev_list {
		if rev.sum > rev_list[max_rev_pos].sum {
			max_rev_pos = pos
		}
	}

	return rev_list[max_rev_pos]
}

func print_answer(best revenue, plans int) {
	var last_output int
	delete_zeros(&best)

	for i := 0; i < len(best.components); i++ {
		if best.components[i] <= last_output {
			best.components[i] = last_output + 1
		}
		fmt.Printf("%d ", best.components[i])
		last_output = best.components[i]
	}

	if len(best.components) < plans {
		for i := len(best.components); i < plans; i++ {
			last_output++
			fmt.Printf("%d ", last_output)
		}
	}
}

func delete_zeros(best *revenue) {
	j := 0
	for _, rev := range (*best).components {
		if rev != 0 {
			(*best).components[j] = rev
			j++
		}
	}
	(*best).components = (*best).components[:j]
}
