package main

import (
	"fmt"
	"container/binarytree"
)


//main
func main() {

	less:= func (x, y interface{}) bool {
		return x.(int) < y.(int)	
	}	

	tree:= binarytree.NewBSTree(less) 
	for _, val := range []int {34, 45, 36, 24, 7, 2, 40, 27, 5, 3} {
		tree.Insert(val)
	}
	fmt.Println()

	binarytree.ImageWidth = 4
	rootImage:= func (x interface{}) string {
		return fmt.Sprintf("%*d", binarytree.ImageWidth, x.(int))
	}

	fmt.Println(tree.RotatedDiagram( ">>>", rootImage))
	fmt.Println()

	fmt.Println(tree.Search(15))
	fmt.Println(tree.Search(24))
	fmt.Println()
	
	tree.Traverse(binarytree.ReverseInOrder, func(t binarytree. BSTree) {
		fmt.Print(rootImage(t.Value()))
	})
	fmt.Printf("\n\n")
	
	for _, s:= range tree.Diagram(rootImage) {
		fmt.Println(s)	
	}	
	fmt.Println()
	tree.Delete(45)
	for _, s:= range tree.Diagram(rootImage) {
		fmt.Println(s)	
	}	
	fmt.Println()
 
}

/*

>>>     45
>>>           40
>>>        36
>>>  34
>>>        27
>>>     24
>>>         7
>>>               5
>>>                  3
>>>            2


false
true

  45  40  36  34  27  24   7   5   3   2

                           |
                   ┌────  34───────────┐
               ┌  24───┐       ┌────  45
   ┌────────   7      27      36───┐
   2───────┐                      40
       ┌   5
       3

                           |
                   ┌────  34───┐
               ┌  24───┐      36───┐
   ┌────────   7      27          40
   2───────┐
       ┌   5
       3

*/
