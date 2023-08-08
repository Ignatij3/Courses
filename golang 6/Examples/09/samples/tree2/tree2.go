package main

import (
	"container/binarytree"
	"fmt"
)

//main
func main() {
	toTheLeft := func(x, y interface{}) bool {
		return x.(int)%5 < y.(int)%5
	}

	binarytree.ImageWidth = 4
	rootImage := func(x interface{}) string {
		return fmt.Sprintf("%*d", binarytree.ImageWidth, x.(int))
	}

	tree := binarytree.NewBinTree()
	for _, val := range []int{34, 45, 36, 24, 3, 7, 2, 40, 27, 5} {
		tree.Insert(val, toTheLeft)
	}
	fmt.Println()

	tree.Traverse(binarytree.ReversePreOrder, func(t binarytree.BinTree) {
		fmt.Print(rootImage(t.Value()))
	})
	fmt.Printf("\n\n")

	for _, s := range tree.Diagram(rootImage) {
		fmt.Println(s)
	}
	fmt.Println()

	fmt.Println(tree.RotatedDiagram(">>>", rootImage))
}

/*

  34  24  45  36   3   7   2  27  40   5

                                   |
   ┌────────────────────────────  34───┐
  45───────────┐                      24
       ┌────  36───────────────┐
      40───┐       ┌────────   3
           5       7───┐
                       2───┐
                          27

>>>     24
>>>  34
>>>            3
>>>                    27
>>>                  2
>>>               7
>>>        36
>>>               5
>>>           40
>>>     45

*/
