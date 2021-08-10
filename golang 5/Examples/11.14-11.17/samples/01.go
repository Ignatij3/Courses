package main

import  (
    "fmt"
)    

type  (
    BinaryNode struct {
        key int
        lson *BinaryNode
        rson *BinaryNode
    }
    BinaryTree struct {
        root *BinaryNode
    }
)

func InitTree() BinaryTree {
    return BinaryTree{root: nil}
}

func (s BinaryTree) Empty() bool {
    return s.root == nil
}	

func (s *BinaryTree) Insert(n int) {
    if (*s).Empty() {
        (*s).root = &BinaryNode {key: n, lson: nil, rson: nil}
        return
    }
    current:= (*s).root
    for  { 
        if n < (*current).key  {
            if (*current).lson == nil  {
               (*current).lson = &BinaryNode{key: n, lson: nil, rson: nil}
               return
            }  else  {
               current = (*current).lson
            }
        }  else  {
        // n >= (*current).key  {
            if (*current).rson == nil  {
               (*current).rson = &BinaryNode{key: n, lson: nil, rson: nil}
               return
            }  else  {
               current = (*current).rson
            }
        }        
    }
}	

func (s BinaryTree) Search(n int) bool {
    current:= s.root
    for  {
        switch {
        case current == nil:
		    return false
		case n == (*current).key:	
		    return true
		case n < (*current).key:	
		    current = (*current).lson
		case n > (*current).key:	
		    current = (*current).rson
		}    
    }
}	

func main() {  
    data:= []int{34, 45, 36, 7, 24, 2, 40, 27, 5, 3, 52, 16, 32, 12, 15}
    tree:= InitTree()
    for _, key := range data {
        tree.Insert(key)
    }
    fmt.Println(tree.Search(12))   // true
    fmt.Println(tree.Search(18))   // false
    fmt.Println(tree.Search(3))    // true
}
