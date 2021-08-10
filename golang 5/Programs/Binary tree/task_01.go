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
		x int
		a []int
	)
	rand.Seed(time.Now().UnixNano())
	h := rand.Intn(1001)
	for i := 0; i < h; i++ {
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

func (s *BinaryTree) SearchHeight() int {
	if (*s).Empty() {return 0}
	var a, b, x int
	n := (*s).root
	if (*n).lson != nil {a = (*n).lson.Height()}
	if (*n).rson != nil {b = (*n).rson.Height()}
	x += a + b
	return x + 1
}

func (n *BinaryNode) Height() int {
	x := 1
	if n != nil {
		a := (*n).lson.Height()
		b := (*n).rson.Height()
		x += a + b
	}
	return x
}

func main() {  
    //data:= []int{34, 45, 36, 7, 24, 2, 40, 27, 5, 3, 52, 16, 32, 12, 15}
    data := FillSlice()
    tree := InitTree()
    for _, key := range data {
        tree.Insert(key)
    }
    x := tree.SearchHeight()
    fmt.Printf("Высота дерева - %d", x)
}
