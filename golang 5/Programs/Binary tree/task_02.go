package main

import  (
    "fmt"
    "time"
    "math/rand"
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

func FillSlice() []int {
	var (
		x, stop int
		a []int
	)
	
	rand.Seed(time.Now().UnixNano())
	stop = rand.Intn(501)
	
	for i := 0; i < stop; i++ {
		x = rand.Intn(101)
		switch rand.Intn(2) {
			case 0:
				a = append(a, x)
			case 1:
				a = append(a, -x)
		}
	}
	return a
}

func InitTree() BinaryTree {
    return BinaryTree{nil}
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

func (s *BinaryTree) CountNodes() (int, int, int, int) {
	if (*s).Empty() {return 1, 0, 0, 1}
	n := (*s).root
	var nodes, oneSon, twoSons, leafs, a1, a2, a3, a4, b1, b2, b3, b4 int
	nodes++
	
	if (*n).lson == nil && (*n).rson == nil {
		leafs++
	} else if (*n).lson != nil && (*n).rson != nil {
		twoSons++
	} else if (*n).lson == nil || (*n).rson == nil {
		oneSon++
	}
	if (*n).lson != nil {a1, a2, a3, a4 = (*n).lson.CNodes()}
	if (*n).rson != nil {b1, b2, b3, b4 = (*n).rson.CNodes()}
	
	leafs += a4 + b4
	nodes += a1 + b1
	oneSon += a2 + b2
	twoSons += a3 + b3
	return nodes, oneSon, twoSons, leafs
}

func (n *BinaryNode) CNodes() (int, int, int, int) {
	var nodes, oneSon, twoSons, leafs int
	if n != nil {
		nodes = 1
		if (*n).lson == nil && (*n).rson == nil {
			leafs = 1
		} else if (*n).lson != nil && (*n).rson != nil {
			twoSons = 1
		} else if (*n).lson == nil || (*n).rson == nil {
			oneSon = 1
		}
		a1, a2, a3, a4 := (*n).lson.CNodes()
		b1, b2, b3, b4 := (*n).rson.CNodes()
		
		leafs += a4 + b4
		nodes += a1 + b1
		oneSon += a2 + b2
		twoSons += a3 + b3
	}
	return nodes, oneSon, twoSons, leafs
}

func main() {  
    //data:= []int{34, 45, 36, 7, 24, 2, 40, 27, 5, 3, 52, 16, 32, 12, 15}
    data := FillSlice()
    tree := InitTree()
    for _, key := range data {
        tree.Insert(key)
    }
	nodes, oneSon, twoSons, leafs := tree.CountNodes()
    fmt.Printf("Количество вершин - %d\nКоличество вершин с одним потомком - %d\nКоличество вершин с двумя потомками - %d\nКоличество листьев - %d\n", nodes, oneSon, twoSons, leafs)
}
