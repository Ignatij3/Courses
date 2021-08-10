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
	for i := 0; i < 200; i++ {
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

func (s *BinaryTree) CountX(x int) int {
	var n int
	if (*s).Empty() {return 0}
	m := (*s).root
	
	if (*m).lson != nil && x < (*m).key {
		n = (*m).lson.Search(x)
	} else if (*m).rson != nil && x >= (*m).key {
		n = (*m).rson.Search(x)
	}
	if (*m).key == x {n++}
	return n
}

func (m *BinaryNode) Search(x int) int {
    var a, b, n int
    if m != nil {
        if (*m).key == x {n = 1}
        a = (*m).lson.Search(x)
        b = (*m).rson.Search(x)
	}
	return a + b + n
}

func main() {  
    var x, n int
    //data:= []int{34, 45, 36, 7, 24, 2, 40, 27, 5, 3, 52, 16, 32, 12, 15}
    data := FillSlice()
    tree := InitTree()
    for _, key := range data {
        tree.Insert(key)
    }
    
    for {
		fmt.Print("Введите число, которое хотите найти:")
		fmt.Scan(&x)
		n = tree.CountX(x)
		fmt.Printf("Ключ %d находится в дереве %d раз\n\n", x, n)
		n = 0
	}
}
