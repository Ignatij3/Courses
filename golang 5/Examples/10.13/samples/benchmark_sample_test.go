package main 

// Запуск из командной строки:
//         go test -bench . benchmark_sample_test.go 
// В данном случае benchmark_sample_test.go - это имя файла, 
// в котором находится данная программа. Имя тестируемого 
// файла обязательно должно заканиваться на _test

import "testing"

func lala(s []int) {
	for i, x := range(s) {
		s[i] = (x-1)*x*(x+1)
	}
}
// Названия тестируемых функций должны начинаться на Benchmark,
// за которым идёт название, начинающееся с большой буквы

// Оценивается функция BenchmarkLala. 
func BenchmarkLala(b *testing.B)  {
	var s [100000]int
    for i:= 0; i < b.N; i++  {
		lala(s[:])
    }
}

// Оценивается функция BenchmarkLalala.
func BenchmarkLalala(b *testing.B)  {
	var s [100000]int
    for k:= 0; k < b.N; k++  {
        for i, x := range(s) {
            s[i] = (x-1)*x*(x+1)
        }
    }
}

// Оценивается функция BenchmarkLalala2.
func BenchmarkLalala2(b *testing.B)  {
	s:= make([]int, 100000)
    for k:= 0; k < b.N; k++  {
        for i, x := range(s) {
            s[i] = (x-1)*x*(x+1)
        }
    }
}

// OUTPUT
// goos: windows
// goarch: amd64
// BenchmarkLala-4      	   12103	     99203 ns/op
// BenchmarkLalala-4    	    7063	    161336 ns/op
// BenchmarkLalala2-4   	   12094	     98955 ns/op
// PASS
// ok  	command-line-arguments	5.765s
