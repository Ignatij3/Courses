package binheap

// Max-heap
import (
	"order"
)

type BinaryHeap []order.Ordered

func (b *BinaryHeap) Add(Element order.Ordered) {
	(*b) = append(*b, Element)
	(*b).pushUp(len(*b) - 1)
}

func (b BinaryHeap) Heapify() {
	for k := (len(b) / 2) - 1; k >= 0; k-- {
		b.pushDown(k)
	}
}

func (b BinaryHeap) GetMax() order.Ordered {
	if len(b) > 0 {
		return b[0]
	} else {
		return b[0]
	}
}

func (b *BinaryHeap) ExtractMax() order.Ordered {
	if len(*b) > 0 {
		max := (*b)[0]
		(*b)[0] = (*b)[len(*b)-1]
		*b = (*b)[:len(*b)-1]
		b.pushDown(0)
		return max
	} else {
		return (*b)[0]
	}
}

func (b *BinaryHeap) Delete(place int) {
	if place >= len(*b) || place < 0 {
		return
	}
	x := (*b)[len(*b)-1]
	*b = (*b)[:len(*b)-1]
	(*b).Change(place, x)
}

func (b BinaryHeap) Change(place int, Key order.Ordered) {
	if place >= len(b) || place < 0 {
		return
	}
	b[place] = Key
	if place > 0 && b[(place-1)/2].Before(Key) {
		b.pushUp(place)
	} else {
		b.pushDown(place)
	}
}

func (b BinaryHeap) pushUp(place int) {
	if place >= len(b) || place <= 0 {
		return
	}
	x := b[place]
	parent := (place - 1) / 2
	for place > 0 && b[parent].Before(x) {
		b[place] = b[parent]
		place = parent
		parent = (place - 1) / 2
	}
	b[place] = x
}

func (b BinaryHeap) pushDown(place int) {
	if place >= len(b) || place < 0 {
		return
	}
	x := b[place]
	for {
		if 2*place+1 >= len(b) { // лист - сыновей нет
			break
		}
		maxson := 2*place + 1 // левый сын
		rson := maxson + 1
		if rson < len(b) && b[maxson].Before(b[rson]) { // правый сын больше левого
			maxson = rson
		}
		if !x.Before(b[maxson]) {
			break
		}
		b[place] = b[maxson]
		place = maxson
	}
	b[place] = x
}
