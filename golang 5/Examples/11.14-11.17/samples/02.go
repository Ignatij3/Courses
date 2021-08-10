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

func (t *BinaryTree) Insert(data int) {
    if (*t).Empty() {
        (*t).root = &BinaryNode{key: data, lson: nil, rson: nil}
    } else {
        (*t).root.Insert(data)
    }
}

func (node *BinaryNode) Insert(data int) {
    if data < (*node).key {
        if (*node).lson == nil {
            (*node).lson = &BinaryNode{key: data, lson: nil, rson: nil}
        } else {
            (*node).lson.Insert(data)
        }
    } else {
    //  data >= (*node).key		
        if (*node).rson == nil {
            (*node).rson = &BinaryNode{key: data, lson: nil, rson: nil}
        } else {
            (*node).rson.Insert(data)
        }
    }   
}

/*   Вариант вставки ключа:
func (t *BinaryTree) Insert(data int) {
    insertnode (&((*t).root), data)
}

func insertnode(node **BinaryNode, data int) {
    if *node == nil {
        *node= &BinaryNode{key: data, lson: nil, rson: nil}
    } else {
        if data < (*node).key {
            insertnode (&((*node).lson), data)
        } else {
        //  data >= (*node).key		
            insertnode (&((*node).rson), data)
        }   
    }
}
Конец варианта */

func (s BinaryTree) Search(n int) bool {
    if s.Empty() { return false }
    switch {
    case n == (*s.root).key:
        return true 
    case n < (*s.root).key:
        return BinaryTree{root: s.root.lson}.Search(n)
    case n > (*s.root).key:
        return BinaryTree{root: s.root.rson}.Search(n)
    }
    return true // или false - всё равно до этой строки не доходит
}	


func (node *BinaryNode) traceUp()  {
    if node != nil  {
        (*node).lson.traceUp()
        fmt.Print((*node).key, " ")
        (*node).rson.traceUp()
	}	
}

func (node *BinaryNode) traceDown()  {
    if node != nil  {
        (*node).rson.traceDown()
        fmt.Print((*node).key, " ")
        (*node).lson.traceDown()
	}	
}

func (s BinaryTree) Trace(dir int)  {
	// dir > 0 - выводит в порядке возрастания
	// dir < 0 - выводит в порядке убывания
    switch  {
    case dir > 0:
        s.root.traceUp()
    case dir < 0:
        s.root.traceDown()
    }    
    fmt.Println()
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
	tree.Trace(1)
	tree.Trace(-1)
}
