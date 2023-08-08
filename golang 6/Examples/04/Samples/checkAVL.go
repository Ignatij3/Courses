package main

import (
	"fmt"
	"order"
	"order/avl"
)

type Integer int

func (a Integer) Before(b order.Ordered) bool {
	return a < b.(Integer)
}

func (a Integer) Show() string {
	return fmt.Sprintf("%*d", order.ImageWidth, a)
}

func main() {
	var data []Integer
	for _, x := range []int{34, 45, 36, 7, 24, 2, 40, 27, 5, 3} {
		data = append(data, Integer(x))
	}
	tree := avl.NewAVLtree()
	for _, val := range data {
		tree.Insert(Integer(val))
	}
	order.ImageWidth = 3
	tree.Traversal(order.ReverseInOrder, func(x order.Key) {
		fmt.Printf("%s ", x.(Integer).Show())
	})
	//  45  40  36  34  27  24   7   5   3   2
	fmt.Println()
	for _, s := range tree.Diagram() {
		fmt.Println(s)
	}
	/* 			   |
	         ┌─── 24────────┐
	┌───  5──┐           ┌ 36─────┐
	2──┐     7        ┌ 34     ┌ 45
	   3             27       40
	*/
	tree.Delete(Integer(24))
	for _, s := range tree.Diagram() {
		fmt.Println(s)
	}
	/*	        |
		      ┌─── 27─────┐
	    ┌───  5──┐     ┌ 36─────┐
	    2──┐     7    34     ┌ 45
	       3                40
	*/
	tree.Delete(Integer(34))
	for _, s := range tree.Diagram() {
		fmt.Println(s)
	}
	/*	        |
		      ┌─── 27─────┐
	    ┌───  5──┐     ┌ 40──┐
	    2──┐     7    36    45
	       3
	*/
	tree.Delete(Integer(36))
	for _, s := range tree.Diagram() {
		fmt.Println(s)
	}
	/*	    	    |
	    	  ┌─── 27──┐
		┌───  5──┐    40──┐
		2──┐     7       45
		   3
	*/
	tree.Insert(Integer(4))
	for _, s := range tree.Diagram() {
		fmt.Println(s)
	}
	/*		           |
	    	     ┌─── 27──┐
		   ┌───  5──┐    40──┐
		┌  3──┐     7       45
		2     4
	*/
}
