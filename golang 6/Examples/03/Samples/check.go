package main

import (
	"fmt"
	"order"
	"order/bst"
	"order/sortable"
)

type Integer int

func (a Integer) Before(b order.Ordered) bool {
	return a < b.(Integer)
}

func (a Integer) Image() string {
	return fmt.Sprintf("%*d", order.ImageWidth, a)
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

	fmt.Println("SortableCollection")
	ac := sortable.NewSortableCollection()
	addCollection(&ac, []Integer{2, 5, 7, 2, 4, 9, 1, 6})
	ac.HeapSort()
	ac.Print() // 1 2 2 4 5 6 7 9
	fmt.Println()

	fmt.Println("SortableList")
	al := sortable.NewSortableList()
	addList(&al, []Integer{2, 5, 7, 2, 4, 9, 1, 6})
	al.Print() // 6 1 9 4 2 7 5 2
	fmt.Println()
	al.HeapSort()
	al.Print() // 1 2 2 4 5 6 7 9
	fmt.Println()

	tree := bst.NewTree()
	data := []Integer{34, 45, 36, 7, 24, 2, 40, 27, 5, 3}
	for _, val := range data {
		tree.Insert(val)
	}
	order.ImageWidth = 4
	fmt.Println(tree.Find(Integer(24))) // true
	fmt.Println(tree.Find(Integer(18))) // false
	fmt.Println(tree.Find(Integer(3)))  // true
	tree.Traversal(order.InOrder, func(x order.Key) {
		fmt.Print(x.Image())
	})
	fmt.Println()
	fmt.Println(tree.DraftDisplay(">>"))
	//>>      45
	//>>            40
	//>>         36
	//>>   34
	//>>            27
	//>>         24
	//>>       7
	//>>             5
	//>>                3
	//>>          2
	fmt.Println()
	for _, s := range tree.Diagram() {
		fmt.Println(s)
	}
	fmt.Println()
	/*
	                                |
	                  ┌──────────   34──────────────┐
	   ┌──────────    7────┐              ┌─────   45
	   2─────────┐        24────┐        36────┐
	        ┌    5             27             40
	        3
	*/
	for _, val := range data {
		tree.Delete(val)
		tree.Traversal(order.InOrder, func(x order.Key) {
			fmt.Print(x.Image())
		})
		fmt.Printf(".    deleted:%s\n", val.Image())
		// 2   3   5   7  24  27  36  40  45.    deleted:  34
		// 2   3   5   7  24  27  34  36  40.    deleted:  45
		// 2   3   5   7  24  27  34  40  45.    deleted:  36
		// 2   3   5  24  27  34  36  40  45.    deleted:   7
		// 2   3   5   7  27  34  36  40  45.    deleted:  24
		// 3   5   7  24  27  34  36  40  45.    deleted:   2
		// 2   3   5   7  24  27  34  36  45.    deleted:  40
		// 2   3   5   7  24  34  36  40  45.    deleted:  27
		// 2   3   7  24  27  34  36  40  45.    deleted:   5
		// 2   5   7  24  27  34  36  40  45.    deleted:   3
		tree.Insert(val)
	}
	fmt.Println()
	tree.Traversal(0, func(x order.Key) {
		x = x.(Integer) + 1
	})
	tree.Traversal(order.ReverseInOrder, func(x order.Key) {
		x.Image()
	}) // 46  41  37  35  28  25   8   6   4   3
	fmt.Println()
}
