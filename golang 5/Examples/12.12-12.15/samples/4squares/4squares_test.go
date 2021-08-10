package main

// Запуск из командной строки:
//         go test -bench . benchmark_4squares_test.go
// Имя файла обязательно должно заканиваться на _test

import (
	"testing"
	"math"
)	

func search2rec(sum int, amount int, result []int) {
	if amount == 0  {
		if sum == 0 {
//			fmt.Println(result)
		}	
		return
	}	
	var start int
	if len(result) == 0 {
		start = 0
	} else {
		start = result[len(result)-1]	
	}	
	for i:= start; i*i <= sum; i++ {
		search2rec(sum - i*i, amount - 1, append(result, i))
	}	
}	

func search2cycle(sum int) {
	for i1:= 0; i1*i1 <= sum; i1++ {
		for i2:= i1; i2*i2 <= sum; i2++ {
			for i3:= i2; i3*i3 <= sum; i3++ {
				for i4:= i3; i4*i4 <= sum; i4++ {
					if i1*i1 + i2*i2 + i3*i3 + i4*i4 == sum {
//						fmt.Println(i1, i2, i3, i4)
					}	
				}
			}
		}			
	}	
}	

func search3rec(sum int, amount int, result []int) {
	if amount == 0  {
		if sum == 0 {
//			fmt.Println(result)
		}	
		return
	}	
	var start int
	if len(result) == 0 {
		start = 0
	} else {
		start = result[len(result)-1]	
	}	
	for i:= start; i*i*amount <= sum; i++ {
		search3rec(sum - i*i, amount - 1, append(result, i))
	}	
}

func search3cycle(sum int) {
	for i1:= 0; i1*i1*4 <= sum; i1++ {
		for i2:= i1; i2*i2*3 <= sum - i1*i1; i2++ {
			for i3:= i2; i3*i3*2 <= sum - i1*i1 - i2*i2; i3++ {
				for i4:= i3; i4*i4 <= sum - i1*i1 - i2*i2 - i3*i3; i4++ {
					if i1*i1 + i2*i2 + i3*i3 + i4*i4 == sum {
//						fmt.Println(i1, i2, i3, i4)
					}	
				}
			}
		}			
	}	
}	

func PerfectSquare (n int) (sqrt int, is bool) {
	sqrt = int( math.Round( math.Sqrt( float64(n) ) ) )
	return sqrt, sqrt*sqrt==n
}
	
func search4rec(sum int, amount int, result []int) {
	if amount == 1  {
		if _, ok:= PerfectSquare(sum); ok {
//			fmt.Println(append(result, x))
		}
		return
	}	
	var start int
	if len(result) == 0 {
		start = 0
	} else {
		start = result[len(result)-1]	
	}	
	for x:= start; x*x*amount <= sum; x++ {
		search4rec(sum - x*x, amount - 1, append(result, x))
	}	
}	

func search4cycle(sum int) {
	for x1:= 0; x1*x1*4 <= sum; x1++ {
		for x2:= x1; x2*x2*3 <= sum - x1*x1; x2++ {
			for x3:= x2; x3*x3*2 <= sum - x1*x1 - x2*x2; x3++ {
				if _, ok:= PerfectSquare(sum - x1*x1 - x2*x2 - x3*x3); ok {
//						fmt.Println(x1, x2, x3, x4)
				}
			}
		}			
	}	
}	

// Названия тестируемых функций должны начинаться на Benchmark,
// за которым идёт название, начинающееся с большой буквы

func BenchmarkSearch2Recursive(b *testing.B) {
    for i := 0; i < b.N; i++ {
        search2rec(50, 4, make([]int, 0))
    }
}

func BenchmarkSearch2Cycle(b *testing.B) {
    for i := 0; i < b.N; i++ {
        search2cycle(50)
    }
}

func BenchmarkSearch3Recursive(b *testing.B) {
    for i := 0; i < b.N; i++ {
        search3rec(50, 4, make([]int, 0))
    }
}

func BenchmarkSearch3Cycle(b *testing.B) {
    for i := 0; i < b.N; i++ {
        search3cycle(50)
    }
}

func BenchmarkSearch4Recursive(b *testing.B) {
    for i := 0; i < b.N; i++ {
        search4rec(50, 4, make([]int, 0))
    }
}

func BenchmarkSearch4Cycle(b *testing.B) {
    for i := 0; i < b.N; i++ {
        search4cycle(50)
    }
}
