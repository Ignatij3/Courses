package main_test

import (
	"math"
	"math/rand"
	"testing"
)

const (
	hundred  = 1e2
	thousand = 1e3
	million  = 1e6
	billion  = 1e9
	infinity = math.MaxInt32
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

func BenchmarkSSTKruskal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SSTKruskal()
	}
}
