package main

// Запуск из командной строки:
//         go test -bench . benchmark_Fibonacci_test.go
// Имя файла обязательно должно заканиваться на _test

import "testing"

func FiboRecursive(n uint64) uint64 {
	if n < 2 {
		return 1
	} else {
		return FiboRecursive(n-1) + FiboRecursive(n-2)
	}
}

func FiboCyclic(n uint64) uint64 {
	a, b := uint64(1), uint64(0)  // a = f[0], b = f[-1]
    for n > 0 {
        a, b = a+b, a
        n--
    }
    return a
}

type memo []uint64

func F(n uint64, m memo) uint64 {
    if m[n] == 0 {
        m[n] = F(n-1, m) + F(n-2, m)
    }
    return m[n]
}

func FiboMemo(n uint64) uint64 {
    m := make([]uint64, n+1, n+1)
    m[0], m[1] = 1, 1
    return F(n, m)
}

// Названия тестируемых функций должны начинаться на Benchmark,
// за которым идёт название, начинающееся с большой буквы

func BenchmarkFiboRecursive16(b *testing.B) {
    for i := 0; i < b.N; i++ {
        FiboRecursive(16)
    }
}

func BenchmarkFiboRecursive24(b *testing.B) {
    for i := 0; i < b.N; i++ {
        FiboRecursive(24)
    }
}

func BenchmarkFiboRecursive32(b *testing.B) {
    for i := 0; i < b.N; i++ {
        FiboRecursive(32)
    }
}

func BenchmarkFiboRecursive40(b *testing.B) {
    for i := 0; i < b.N; i++ {
        FiboRecursive(40)
    }
}

func BenchmarkFiboCyclic16(b *testing.B) {
    for i := 0; i < b.N; i++ {
        FiboCyclic(16)
    }
}

func BenchmarkFiboCyclic24(b *testing.B) {
    for i := 0; i < b.N; i++ {
        FiboCyclic(24)
    }
}

func BenchmarkFiboCyclic32(b *testing.B) {
    for i := 0; i < b.N; i++ {
        FiboCyclic(32)
    }
}

func BenchmarkFiboCyclic40(b *testing.B) {
    for i := 0; i < b.N; i++ {
        FiboCyclic(40)
    }
}

func BenchmarkFiboCyclic60(b *testing.B) {
    for i := 0; i < b.N; i++ {
        FiboCyclic(60)
    }
}

func BenchmarkFiboCyclic90(b *testing.B) {
    for i := 0; i < b.N; i++ {
        FiboCyclic(90)
    }
}

func BenchmarkFiboMemo16(b *testing.B) {
    for i := 0; i < b.N; i++ {
        FiboMemo(16)
    }
}

func BenchmarkFiboMemo24(b *testing.B) {
    for i := 0; i < b.N; i++ {
        FiboMemo(24)
    }
}

func BenchmarkFiboMemo32(b *testing.B) {
    for i := 0; i < b.N; i++ {
        FiboMemo(32)
    }
}

func BenchmarkFiboMemo40(b *testing.B) {
    for i := 0; i < b.N; i++ {
        FiboMemo(40)
    }
}

func BenchmarkFiboMemo60(b *testing.B) {
    for i := 0; i < b.N; i++ {
        FiboMemo(60)
    }
}

func BenchmarkFiboMemo90(b *testing.B) {
    for i := 0; i < b.N; i++ {
        FiboMemo(90)
    }
}
