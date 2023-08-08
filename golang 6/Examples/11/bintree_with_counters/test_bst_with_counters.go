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

	binarytree.ImageWidth = 4
	rootImage:= func (x interface{}) string {
		return fmt.Sprintf("%*d", binarytree.ImageWidth, x.(int))
	}
	
	tree:= binarytree.NewBSTree(less) 
	for _, val := range []int {34, 45, 36, 24, 7, 2, 40, 27, 5, 3} {
		tree.Insert(val)
	}

	for _, s:= range tree.Diagram(rootImage) {
		fmt.Println(s)	
	}	
	tree.Traverse(binarytree.ReverseInOrder, func(t binarytree. BSTree) {
		fmt.Print(rootImage(t.Value()), "/", t.Size())
	})
	fmt.Printf("\n\n")
	
	for i:= 0; i<tree.Size(); i++ {
		fmt.Printf("%4d: %d", i, tree.FindKth(i).(int))
	}
	fmt.Println()
	
	fmt.Println()
	tree.Delete(45)
	tree.Delete(11)
	tree.Delete(24)
	tree.Delete(3)
	for _, s:= range tree.Diagram(rootImage) {
		fmt.Println(s)	
	}	
	fmt.Println()

	tree.Traverse(binarytree.InOrder, func(t binarytree. BSTree) {
		fmt.Print(rootImage(t.Value()), "/", t.Size())
	})
	fmt.Printf("\n\n")
	
	for i:= 0; i<tree.Size(); i++ {
		fmt.Printf("%4d: %d", i, tree.FindKth(i).(int))
	}
	fmt.Println()
}

/*                                                    |
                           |
                   ┌────  34───────────┐
               ┌  24───┐       ┌────  45
   ┌────────   7      27      36───┐
   2───────┐                      40
       ┌   5
       3
  45/3  40/1  36/2  34/10  27/1  24/6   7/4   5/2   3/1   2/3

   0: 2   1: 3   2: 5   3: 7   4: 24   5: 27   6: 34   7: 36   8: 40   9: 45

                   |
               ┌  34───┐
           ┌  27      36───┐
   ┌────   7              40
   2───┐
       5

   2/2   5/1   7/3  27/4  34/7  36/2  40/1

   0: 2   1: 5   2: 7   3: 27   4: 34   5: 36   6: 40

*/
