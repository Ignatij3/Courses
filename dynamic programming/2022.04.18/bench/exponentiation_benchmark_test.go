package bench_test

import (
	"math"
	"strconv"
	"testing"
)

const (
	thousand    = 1e3
	billion     = 1e9
	quintillion = 1e18
	modulo      = 1e9 + 7
)

func linear_exp(exp uint64) uint64 {
	var result uint64 = 2
	for i := uint64(0); i < exp; i++ {
		result = (result * 2) % modulo
	}
	return result
}

func logarithmic_exp(exp uint64) uint64 {
	binRepr := strconv.FormatUint(exp, 2) //little endian

	cache := make([]uint64, len(binRepr))
	cache[0] = 2

	i := 0
	cachelen := len(cache) - 1
	for i <= cachelen && binRepr[cachelen-i] == '0' {
		i++
		cache[i] = (cache[i-1] * cache[i-1]) % modulo
	}

	result := cache[i]
	for i++; i <= cachelen; i++ {
		cache[i] = (cache[i-1] * cache[i-1]) % modulo
		if binRepr[cachelen-i] == '1' {
			result = (result * cache[i-1]) % modulo
		}
	}

	return result
}

func logarithmic_exp_revision(base, exp uint64) uint64 {
	var res uint64 = 1

	for exp > 0 {
		if exp&1 == 1 {
			res = (res * base) % modulo
		}
		base = (base * base) % modulo
		exp >>= 1
	}

	return res
}

func Benchmark_Linear_1e3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		linear_exp(thousand)
	}
}

func Benchmark_Linear_1e9(b *testing.B) {
	for i := 0; i < b.N; i++ {
		linear_exp(billion)
	}
}

func Benchmark_MathPow_1e3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		math.Pow(2, thousand)
	}
}

func Benchmark_MathPow_1e9(b *testing.B) {
	for i := 0; i < b.N; i++ {
		math.Pow(2, billion)
	}
}

func Benchmark_MathPow_1e18(b *testing.B) {
	for i := 0; i < b.N; i++ {
		math.Pow(2, quintillion)
	}
}

func Benchmark_Log2_1e3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		logarithmic_exp(thousand)
	}
}

func Benchmark_Log2_1e9(b *testing.B) {
	for i := 0; i < b.N; i++ {
		logarithmic_exp(billion)
	}
}

func Benchmark_Log2_1e18(b *testing.B) {
	for i := 0; i < b.N; i++ {
		logarithmic_exp(quintillion)
	}
}

func Benchmark_Log2_Ver2_1e3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		logarithmic_exp_revision(2, thousand)
	}
}

func Benchmark_Log2_Ver2_1e9(b *testing.B) {
	for i := 0; i < b.N; i++ {
		logarithmic_exp_revision(2, billion)
	}
}

func Benchmark_Log2_Ver2_1e18(b *testing.B) {
	for i := 0; i < b.N; i++ {
		logarithmic_exp_revision(2, quintillion)
	}
}
