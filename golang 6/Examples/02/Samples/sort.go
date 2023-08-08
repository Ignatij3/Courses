package main

import (
	"fmt"
	"mypacks/order"
	"mypacks/order/sortable"
	"mypacks/order/sorted"
)

type Integer int

func (a Integer) Before(b order.Ordered) bool {
	return a < b.(Integer)
}

func (a Integer) Show() string {
	return fmt.Sprintf("%d ", a)
}

func addCollection(c *sortable.SortableCollection, data []Integer) {
	for _, x := range data {
		*c = append(*c, x)
	}
	return
}

func addList(l *sortable.SortableList, data []Integer) {
	for _, x := range data {
		(*l).Add(sortable.Node{Value: x, Next: sortable.NewSortableList()})
	}
	return
}

func main() {
	fmt.Println("SortedCollection")
	var sc sorted.SortedCollection
	for _, x := range []Integer{2, 5, 7, 2, 4, 9, 1, 6} {
		sc.Insert(x)
	}
	sc.Print() // 1 2 2 4 5 6 7 9
	fmt.Println()

	fmt.Println("SortedList")
	sl := sorted.NewSortedList()
	for _, x := range []Integer{2, 5, 7, 2, 4, 9, 1, 6} {
		sl.Insert(x)
	}
	sl.Print() // 1 2 2 4 5 6 7 9
	fmt.Println()

	fmt.Println("SortableCollection")
	ac := sortable.NewSortableCollection()
	addCollection(&ac, []Integer{2, 5, 7, 2, 4, 9, 1, 6})
	ac.QSort()
	ac.Print() // 1 2 2 4 5 6 7 9
	fmt.Println()
	ac = sortable.NewSortableCollection()
	addCollection(&ac, []Integer{2, 5, 7, 2, 4, 9, 1, 6})
	ac.BinaryMergeSort()
	ac.Print() // 1 2 2 4 5 6 7 9
	fmt.Println()
	ac = sortable.NewSortableCollection()
	addCollection(&ac, []Integer{2, 5, 7, 2, 4, 9, 1, 6})
	ac.NaturalMergeSort()
	ac.Print() // 1 2 2 4 5 6 7 9
	fmt.Println()

	fmt.Println("SortableList")
	al := sortable.NewSortableList()
	addList(&al, []Integer{2, 5, 7, 2, 4, 9, 1, 6})
	al.QSort()
	al.Print() // 1 2 2 4 5 6 7 9
	fmt.Println()
	al = sortable.NewSortableList()
	addList(&al, []Integer{2, 5, 7, 2, 4, 9, 1, 6})
	al.MergeSort()
	al.Print() // 1 2 2 4 5 6 7 9
	fmt.Println()
}
