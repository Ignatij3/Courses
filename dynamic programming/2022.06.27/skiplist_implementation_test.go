package skiplist_test

import (
	"math/rand"
	"testing"

	. "../skiplist"
)

const (
	hundred  = 1e2
	thousand = 1e3
	million  = 1e6
	billion  = 1e9
)

var data [hundred * thousand]int

type boolgen struct {
	src       rand.Source
	cache     int64
	remaining int
}

func (b *boolgen) Bool() bool {
	if b.remaining == 0 {
		b.cache, b.remaining = b.src.Int63(), 63
	}

	result := b.cache&0x01 == 1
	b.cache >>= 1
	b.remaining--

	return result
}

func init() {
	var boolsrc boolgen = boolgen{src: rand.NewSource(69)}
	rand.Seed(69)

	for i := range data {
		if boolsrc.Bool() {
			data[i] = rand.Intn(billion / 2)
		} else {
			data[i] = -rand.Intn(billion / 2)
		}
	}
}

// refresh returns copy of data of required size.
// While doing so, it stops and on exit resumes timer
func refresh(b *testing.B, size int) []int {
	b.StopTimer()
	newdata := make([]int, size)
	copy(newdata, data[:size])
	b.StartTimer()
	return newdata
}

func benchmarkGeneral(b *testing.B, insert func(data []int)) {
	var cases = []struct {
		name string
		size int
	}{
		{"100", hundred},
		{"1k", thousand},
		{"10k", 10 * thousand},
		{"100k", hundred * thousand},
	}

	for _, c := range cases {
		test := func(b *testing.B) {
			data := refresh(b, c.size)
			for i := 0; i < b.N; i++ {
				insert(data)
				data = refresh(b, c.size)
			}
		}
		b.Run(c.name, test)
	}
}

func BenchmarkAdapter(b *testing.B) {
	benchmarkGeneral(b, func(data []int) {
		sl := NewSkipListInt(10, 0.5, func(a, b int) bool { return a < b })
		for _, v := range data {
			sl.Insert(v)
		}
	})
}

func BenchmarkGeneric(b *testing.B) {
	benchmarkGeneral(b, func(data []int) {
		sl := NewSkipListGeneric(10, 0.5, func(a, b int) bool { return a < b })
		for _, v := range data {
			sl.Insert(v)
		}
	})
}

func BenchmarkInterface(b *testing.B) {
	benchmarkGeneral(b, func(data []int) {
		sl := NewSkipListIface(10, 0.5, func(a, b interface{}) bool { return a.(int) < b.(int) })
		for _, v := range data {
			sl.Insert(v)
		}
	})
}
