package median_search_test

import (
	"math/rand"
	"testing"

	"../median"
)

const (
	thousand = 1e3
	million  = 1e6
	billion  = 1e9
)

var data []int = make([]int, billion)

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

func benchmarkQuickSearch(b *testing.B, data []int) {
	for i := 0; i < b.N; i++ {
		median.QuickSearch(data)
	}
}

func benchmarkMedianOfMedians(b *testing.B, data []int) {
	for i := 0; i < b.N; i++ {
		median.MedianOfMedians(data)
	}
}

func BenchmarkQuickSearch(b *testing.B) {
	benchmarks := []struct {
		size int
		name string
	}{
		{thousand, "10^3"},
		{million, "10^6"},
		{billion, "10^9"},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(bb *testing.B) { benchmarkQuickSearch(bb, data[:bm.size]) })
	}
}
func BenchmarkMedianOfMedians(b *testing.B) {
	benchmarks := []struct {
		size int
		name string
	}{
		{thousand, "10^3"},
		{million, "10^6"},
		{billion, "10^9"},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(bb *testing.B) { benchmarkMedianOfMedians(bb, data[:bm.size]) })
	}
}
