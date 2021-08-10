package binheap

import (
	"math"
)
	
type (
	BinaryHeap []Tdata
	Tdata int32
)

const (
	MinData = math.MinInt32
)
		
func (b *BinaryHeap) Init (data []Tdata) {
	for _, x := range data {
		(*b) = append(*b, x)	
	}	
	(*b).Heapify()
}	

func (b BinaryHeap) pushUp(place int) {
	if place >= len(b) || place <= 0 { 
		return
	}	
	x := b[place]
	parent:= (place-1)/2
	for  place > 0 && b[parent] < x {
		b[place] = b[parent]
		place = parent
		parent = (place-1)/2
	}
	b[place] = x	
}	

func (b BinaryHeap) pushDown(place int) {
	if place >= len(b) || place < 0 { 
		return
	}	
	x := b[place]
	for  {
		if 2*place + 1 >= len(b) { // лист - сыновей нет
			break
		}
		maxson := 2*place + 1  // левый сын
		rson:= maxson + 1
		if rson < len(b) && b[rson] > b[maxson] { // правый сын больше левого
			maxson = rson
		}
		if b[maxson] <= x {
			break
		}
		b[place] = b[maxson]	 	
		place = maxson
	}
	b[place] = x	
}	

func (b *BinaryHeap) Add(value Tdata) {
	(*b) = append(*b, value)
	(*b).pushUp(len(*b)-1)
}

func (b BinaryHeap) GetMax() Tdata {
	if len(b) > 0 {
		return b[0]
	} else {
		return MinData
	}
}

func (b *BinaryHeap) ExtractMax() Tdata {
	if len(*b) > 0 {
		max:= (*b)[0]
		(*b)[0] = (*b)[len(*b)-1]
		*b = (*b)[:len(*b)-1]
		(*b).pushDown(0)
		return max
	} else {
		return MinData
	}
}

func (b *BinaryHeap) Delete(place int) {
	if place >= len(*b) {
		return
	}
	x:= (*b)[len(*b)-1]
	*b = (*b)[:len(*b)-1]
	(*b).Change(place, x)
}	

func (b BinaryHeap) Change(place int, value Tdata) {
	b[place] = value
	if place > 0 && value > b[(place - 1) / 2] {
		b.pushUp(place)
	} else {
		b.pushDown(place)
	}		 
}

func (b BinaryHeap) Heapify() {
	for k:= (len(b) / 2) - 1; k >= 0; k-- {
		b.pushDown(k)
	}	
}	

type (
	Lmnt struct {
		Index int
		Value Tdata
	}		
	LocatorBinaryHeap struct {
		heap []Lmnt
		locator []int
	}	
)
		
func (b *LocatorBinaryHeap) Init (data []Tdata) {
	for i, x := range data {
		(*b).heap = append((*b).heap, Lmnt{i, x})	
		(*b).locator = append((*b).locator, i)
	}	
	(*b).Heapify()
}	

func (b LocatorBinaryHeap) pushUp(place int) {
// поднять элемент, стоящий в b.heap на месте #place	
	if place >= len(b.heap) || place <= 0 { 
		return
	}	
	t := b.heap[place]
	parent:= (place-1)/2
	for  place > 0 && b.heap[parent].Value < t.Value {
		b.heap[place] = b.heap[parent]
		b.locator[b.heap[parent].Index] = place
		place = parent
		parent = (place-1)/2
	}
	b.heap[place] = t
	b.locator[t.Index] = place	
}	

func (b LocatorBinaryHeap) pushDown(place int) {
// спустить элемент, стоящий в b.heap на месте #place	
	if place >= len(b.heap) || place < 0 { 
		return
	}	
	t := b.heap[place]
	for  {
		if 2*place + 1 >= len(b.heap) { // лист - сыновей нет
			break
		}
		maxson := 2*place + 1  // левый сын
		rson:= maxson + 1
		if rson < len(b.heap) && b.heap[rson].Value > b.heap[maxson].Value { // правый сын больше левого
			maxson = rson
		}
		if b.heap[maxson].Value <= t.Value {
			break
		}
		b.heap[place] = b.heap[maxson]	 	
		b.locator[b.heap[maxson].Index] = place
		place = maxson
	}
	b.heap[place] = t	
	b.locator[t.Index] = place	
}	

func (b *LocatorBinaryHeap) Add(t Lmnt) {
	if t.Index < 0 || t.Index >= len((*b).locator) || (*b).locator[t.Index] >= 0 {
	// Ошибка - кривой индекс или добавляем в кучу элемент, который сейчас уже в куче	
		return
	}	
	(*b).heap = append((*b).heap, t)
	(*b).locator[t.Index] = len((*b).heap)-1
	(*b).pushUp(len((*b).heap)-1)
}

func (b LocatorBinaryHeap) GetMax() Lmnt {
	if len(b.heap) > 0 {
		return b.heap[0]
	} else {
		return Lmnt{-1, MinData}
	}
}

func (b *LocatorBinaryHeap) ExtractMax() Lmnt {
	if len((*b).heap) > 0 {
		max:= (*b).heap[0]
		(*b).locator[(*b).heap[0].Index] = -1
		(*b).heap[0] = (*b).heap[len((*b).heap)-1]
		(*b).heap = (*b).heap[:len((*b).heap)-1]
		(*b).pushDown(0)
		return max
	} else {
		return Lmnt{-1, MinData}
	}
}

func (b *LocatorBinaryHeap) Delete(index int) {
	if index < 0 || index >= len((*b).locator) || (*b).locator[index] == -1 {
	// Ошибка - кривой индекс или удаляем из кучи элемент, который сейчас не в куче	
		return
	}	
	place:= (*b).locator[index]
	if place >= len((*b).heap) {
		return
	}
	(*b).locator[index] = -1
	t:= (*b).heap[len((*b).heap)-1]
	(*b).heap[place] = t
	(*b).locator[t.Index] = place
	(*b).heap = (*b).heap[:len((*b).heap)-1]
	(*b).Change(t)
}	

func (b LocatorBinaryHeap) Change(t Lmnt) {
	if t.Index < 0 || t.Index >= len(b.locator) || b.locator[t.Index] == -1 {
	// Ошибка - кривой индекс или изменяем элемент, который сейчас не в куче	
		return
	}	
	place:= b.locator[t.Index]
	x:= b.heap[place].Value
	b.heap[place] = t
	if t.Value > x {
		b.pushUp(place)
	} else {
		b.pushDown(place)
	}		 
}

func (b LocatorBinaryHeap) Heapify() {
	for k:= (len(b.heap) / 2) - 1; k >= 0; k-- {
		b.pushDown(k)
	}	
}	
