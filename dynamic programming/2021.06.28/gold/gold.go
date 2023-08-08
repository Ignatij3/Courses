package main

import (
	"fmt"
	"time"
)

const (
	SAFE_CAPACITY = 22
	TOTAL_WEIGHT  = SAFE_CAPACITY + 1
)

type kg struct {
	is_valid    bool
	gold_bricks []int
}

func main() {
	/*gold_pile := make([]int, 2000)
	for i := range gold_pile {
		gold_pile[i] = rand.Intn(20) + 1
	}*/
	var gold_pile []int = []int{10, 7, 7, 3, 10, 7}

	start := time.Now()
	safe_fill_kg, gold_bricks := fill_safe(gold_pile)
	fmt.Println(time.Since(start))
	fmt.Printf("Safe (with capasity of %dkg) was filled with %dkg of gold, using %v\n", SAFE_CAPACITY, safe_fill_kg, gold_bricks)
}

func fill_safe(gold_pile []int) (int, []int) {
	weight_list := init_weight_list()

	for _, gold_brick := range gold_pile {
		add_weight(&weight_list, gold_brick)
	}

	return largest_kg_accessible(weight_list)
}

func init_weight_list() [TOTAL_WEIGHT]kg {
	var weight_list [TOTAL_WEIGHT]kg
	weight_list[0].is_valid = true
	return weight_list
}

func add_weight(weight_list *[TOTAL_WEIGHT]kg, gold_brick int) {
	for kg := SAFE_CAPACITY; kg >= gold_brick; kg-- {
		if (*weight_list)[kg-gold_brick].is_valid {
			var small_gold_pile []int
			small_gold_pile = append(small_gold_pile, (*weight_list)[kg-gold_brick].gold_bricks...)
			small_gold_pile = append(small_gold_pile, gold_brick)
			(*weight_list)[kg].is_valid = true
			(*weight_list)[kg].gold_bricks = small_gold_pile
		}
	}
}

func largest_kg_accessible(weight_list [TOTAL_WEIGHT]kg) (int, []int) {
	for kg := SAFE_CAPACITY; kg >= 0; kg-- {
		if weight_list[kg].is_valid {
			return kg, weight_list[kg].gold_bricks
		}
	}
	return 0, nil
}
