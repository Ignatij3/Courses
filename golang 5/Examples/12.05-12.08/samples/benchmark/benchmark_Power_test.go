package main

// Запуск из командной строки:
//         go test -bench . benchmark_Power_test.go
// Имя файла обязательно должно заканиваться на _test

import "testing"

const mod = 1000000007

func Mult(a, b uint32) uint32 {
	return uint32(uint64(a) * uint64(b) % mod)
}

func PowerCyclic0(a uint32, n uint32) uint32 {
	var res uint32 = 1
	for ; n > 0; n-- {
		res = Mult(res, a)
	}
	return res
}

func PowerRecursive0(a uint32, n uint32) uint32 {
	if n == 0 {
		return 1
	}
	return Mult(a, PowerRecursive0(a, n-1))
}

func PowerRecursive1(a uint32, n uint32) uint32 {
	var b uint32
	if n == 0 {
		return 1
	} else {
		b = PowerRecursive1(a, n/2) // b = a^(n/2)
		b = Mult(b, b)
		if n%2 == 0 {
			return b
		} else {
			return Mult(b, a)
		}
	}
}


func PowerRecursive1a(a uint32, n uint32) uint32 {
	var b uint32
	if n == 0 {
		return 1
	} else {
		b = Mult (PowerRecursive1a(a, n/2), PowerRecursive1a(a, n/2)) 
		if n%2 == 0 {
			return b
		} else {
			return Mult(b, a)
		}
	}
}

func PowerRecursive2(a uint32, n uint32) uint32 {
	var b uint32
	if n == 0 {
		return 1
	} else {
		b = Mult(a, a) // b = a^2
		b = PowerRecursive2(b, n/2)
		if n%2 == 0 {
			return b
		} else {
			return Mult(b, a)
		}
	}
}

func PowerCyclic2(a uint32, n uint32) uint32 {
	var res uint32 = 1
	for n > 0 {
		if n%2 == 1 {
			res = Mult(res, a)
		}
		n /= 2
		a = Mult(a, a)
	}
	return res
}

// Названия тестируемых функций должны начинаться на Benchmark,
// за которым идёт название, начинающееся с большой буквы

func BenchmarkPowerCyclic0_6(b *testing.B) {
    for i := 0; i < b.N; i++ {
        PowerCyclic0(2020, 1000000)
    }
}

func BenchmarkPowerRecursive0_6(b *testing.B) {
    for i := 0; i < b.N; i++ {
        PowerRecursive0(2020, 1000000)
    }
}

func BenchmarkPowerRecursive1_6(b *testing.B) {
    for i := 0; i < b.N; i++ {
        PowerRecursive1(2020, 1000000)
    }
}

func BenchmarkPowerRecursive1_9(b *testing.B) {
    for i := 0; i < b.N; i++ {
        PowerRecursive1(2020, 1000000000)
    }
}

func BenchmarkPowerRecursive1a_6(b *testing.B) {
    for i := 0; i < b.N; i++ {
        PowerRecursive1a(2020, 1000000)
    }
}

func BenchmarkPowerRecursive2_6(b *testing.B) {
    for i := 0; i < b.N; i++ {
        PowerRecursive2(2020, 1000000)
    }
}

func BenchmarkPowerRecursive2_9(b *testing.B) {
    for i := 0; i < b.N; i++ {
        PowerRecursive2(2020, 1000000000)
    }
}

func BenchmarkPowerCyclic2_6(b *testing.B) {
    for i := 0; i < b.N; i++ {
        PowerCyclic2(2020, 1000000)
    }
}

func BenchmarkPowerCyclic2_9(b *testing.B) {
    for i := 0; i < b.N; i++ {
        PowerCyclic2(2020, 1000000000)
    }
}
