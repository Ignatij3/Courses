package main

import (
	"container/binarytree"
	"fmt"
)

//main
func main() {
	less := func(x, y interface{}) bool {
		return x.(int) < y.(int)
	}

	tree := binarytree.NewAVLTree(less)
	for _, val := range []int{34, 45, 36, 24, 7, 2, 40, 27, 5, 3} {
		tree.Insert(val)
	}
	fmt.Println()

	binarytree.ImageWidth = 4
	rootImage := func(x interface{}) string {
		return fmt.Sprintf("%*d", binarytree.ImageWidth, x.(int))
	}

	fmt.Println(tree.RotatedDiagram(">>>", rootImage))
	fmt.Println()

	fmt.Println(tree.Search(15))
	fmt.Println(tree.Search(24))
	fmt.Println()

	tree.Traverse(binarytree.ReverseInOrder, func(t binarytree.AVLTree) {
		fmt.Print(rootImage(t.Value()))
	})
	fmt.Printf("\n\n")

	for _, s := range tree.Diagram(rootImage) {
		fmt.Println(s)
	}
	fmt.Println()

	tree.Delete(5)
	for _, s := range tree.Diagram(rootImage) {
		fmt.Println(s)
	}
	fmt.Println()

	tree.Delete(3)
	tree.Delete(7)
	for _, s := range tree.Diagram(rootImage) {
		fmt.Println(s)
	}
	fmt.Println()
}

/*

>>>        45
>>>           40
>>>     36
>>>        34
>>>           27
>>>  24
>>>         7
>>>      5
>>>            3
>>>         2


false
true

  45  40  36  34  27  24   7   5   3   2

                   |
           ┌────  24───────────┐
   ┌────   5───┐           ┌  36───────┐
   2───┐       7       ┌  34       ┌  45
       3              27          40

               |
       ┌────  24───────────┐
   ┌   3───┐           ┌  36───────┐
   2       7       ┌  34       ┌  45
                  27          40

                   |
       ┌────────  36───────┐
   ┌  24───────┐       ┌  45
   2       ┌  34      40
          27

*/
