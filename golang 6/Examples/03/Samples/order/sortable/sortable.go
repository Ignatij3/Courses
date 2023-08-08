package sortable

import (
	"fmt"
	"math/rand"
	"order"
	"order/binheap"
)

type (
	SortableCollection []order.Key
)

func NewSortableCollection() SortableCollection {
	return []order.Key{}
}

func (a SortableCollection) HeapSort() {
	var bheap binheap.BinaryHeap
	for _, x := range a {
		bheap.Add(x)
	}
	for k := len(a) - 1; k >= 0; k-- {
		a[k] = bheap.ExtractMax().(order.Key)
	}
}

const cutoff = 16

func (a SortableCollection) qSort() {
	if len(a) <= cutoff {
		return
	}

	k := rand.Intn(len(a))
	a[0], a[k] = a[k], a[0]

	pivot := a[0]
	small, large := 1, len(a)-1
	for {
		for small < len(a) && a[small].Before(pivot) {
			small++
		}
		for pivot.Before(a[large]) {
			large--
		}
		if small >= large {
			break
		}
		a[small], a[large] = a[large], a[small]
		small++
		large--
	}
	a[0], a[large] = a[large], a[0]
	a[:large].qSort()
	a[large+1:].qSort()
}

func (a SortableCollection) QSort() {
	a.qSort()
	for i, x := range a {
		j := i
		for j > 0 && x.Before(a[j-1]) {
			a[j] = a[j-1]
			j--
		}
		a[j] = x
	}
}

func (a SortableCollection) merge(start2 int) {
	res := make([]order.Key, len(a))
	i1, i2, ires := 0, start2, 0
	for i1 < start2 && i2 < len(a) {
		if a[i1].Before(a[i2]) {
			res[ires] = a[i1]
			i1++
		} else {
			res[ires] = a[i2]
			i2++
		}
		ires++
	}
	if i2 == len(a) {
		copy(a[ires:], a[i1:])
	}
	copy(a, res[:ires])
}

func (a SortableCollection) BinaryMergeSort() {
	if len(a) <= 1 {
		return
	}
	a[:len(a)/2].BinaryMergeSort()
	a[len(a)/2:].BinaryMergeSort()
	a.merge(len(a) / 2)
}

func (a SortableCollection) NaturalMergeSort() {
	// инициализация - заполняем слайс start
	start := []int{0}
	for i := 1; i < len(a); i++ {
		if a[i].Before(a[i-1]) {
			start = append(start, i)
		}
	}
	start = append(start, len(a))
	// сортировка
	for len(start) > 2 {
		// проходим по всему массиву, склеивая пары соседних серий
		for k := 0; k < len(start)-2; k += 2 {
			a[start[k]:start[k+2]].merge(start[k+1] - start[k])
		}
		// преобразуем слайс start: start[2] -> start[1],
		// start[4] -> start[2], start[6] -> start[3] и т.д.
		k := 0
		for {
			k += 2
			if k >= len(start) {
				break
			}
			start[k/2] = start[k]
		}
		start = start[:k/2]
		// если перед этим было нечётное количество серий, то
		// надо добавить конец последней серии - len(l)
		if start[len(start)-1] < len(a) {
			start = append(start, len(a))
		}
	}
}

func (a SortableCollection) Print() {
	for _, x := range a {
		fmt.Print(x.Image())
	}
}

type (
	Node struct {
		Value order.Key
		Next  SortableList
	}
	SortableList struct {
		First *Node
	}
)

func (a Node) Before(b order.Ordered) bool {
	return a.Value.Before(b.(Node).Value)
}

func (a Node) Image() string {
	return a.Value.Image()
}

func NewSortableList() SortableList {
	return SortableList{} // <==> SortableList{nil} - empty list
}

func (b SortableList) Empty() bool {
	return b.First == nil
}

func (b *SortableList) Add(x Node) {
	x.Next = *b
	(*b).First = &x
}

func (a *SortableList) HeapSort() {
	var bheap binheap.BinaryHeap
	for !(*a).Empty() {
		bheap.Add(*(*a).First)
		*a = (*(*a).First).Next
	}
	*a = NewSortableList()
	for k := len(bheap) - 1; k >= 0; k-- {
		(*a).Add(bheap.ExtractMax().(Node))
	}
}

func (a *SortableList) QSort() {
	if (*a).Empty() || (*(*a).First).Next.Empty() {
		return
	}
	pivotNode := *(*a).First // var pivotNode Node
	a1, a2 := (*(*a).First).Next.separateByPivot(pivotNode)
	a1.QSort()
	a2.QSort()
	pivotNode.Next = a2
	a2.First = &pivotNode
	var runner SortableList
	if a1.Empty() {
		*a = a2
	} else {
		runner = a1
		for !(*runner.First).Next.Empty() {
			runner = (*runner.First).Next
		}
		(*runner.First).Next = a2
		*a = a1
	}
}

func (a SortableList) separateByPivot(pivot order.Ordered) (SortableList, SortableList) {
	runner := a // var runner SortableList
	var next SortableList
	a1, a2 := NewSortableList(), NewSortableList()
	for !runner.Empty() {
		next = (*runner.First).Next // <==>  next = (*(runner.First)).Next
		if (*runner.First).Before(pivot) {
			(*runner.First).Next = a1
			a1 = runner
		} else {
			(*runner.First).Next = a2
			a2 = runner
		}
		runner = next
	}
	return a1, a2
}

func (a *SortableList) MergeSort() {
	if (*a).Empty() || (*(*a).First).Next.Empty() {
		return
	}
	a1, a2 := *a, (*a.First).Next
	for !a2.Empty() && !(*a2.First).Next.Empty() {
		a1 = (*a1.First).Next
		a2 = (*(*a2.First).Next.First).Next
	}
	a2, (*a1.First).Next, a1 = (*a1.First).Next, SortableList{nil}, *a
	a1.MergeSort()
	a2.MergeSort()
	*a = merge(a1, a2)
}

func merge(a SortableList, b SortableList) SortableList {
	if a.Empty() {
		return b
	}
	if b.Empty() {
		return a
	}
	var res SortableList
	if (*a.First).Value.Before((*b.First).Value) {
		res, a = a, (*a.First).Next
	} else {
		res, b = b, (*b.First).Next
	}
	runner := res.First
	for !a.Empty() && !b.Empty() {
		if (*a.First).Value.Before((*b.First).Value) {
			(*runner).Next, a = a, (*a.First).Next
		} else {
			(*runner).Next, b = b, (*b.First).Next
		}
		runner = (*runner).Next.First
	}
	if a.Empty() {
		(*runner).Next = b
	} else {
		// b.Empty()
		(*runner).Next = a
	}
	return res
}

func (a SortableList) Print() {
	if a.Empty() {
		return
	}
	fmt.Print((*a.First).Image())
	(*a.First).Next.Print()
}
