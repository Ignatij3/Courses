package main 

// Запуск из командной строки:
//         go test -bench . benchmark_sample_test.go 
// В данном случае benchmark_sample_test.go - это имя файла, 
// в котором находится данная программа. Имя тестируемого 
// файла обязательно должно заканиваться на _test

import 	(
    "container/list"
    "math"
    "testing"
	"math/rand"
)

const
    Infinity = math.MaxInt64
    
var d [100000000]int
func init() {
	for i, _ := range d {
	    d[i] = rand.Intn(1000000000)
	}
}	

func Sort(data [] int) {
    l := list.New()
    l.PushFront(Infinity)
    for _, x:= range data  {
        Insert(l, x)
    }
    l.Remove(l.Back())	
}

func Insert(l *list.List, num int)  {
    for e := (*l).Front(); ; e = e.Next() {
        if e.Value.(int) >= num  {
            (*l).InsertBefore(num, e)
            break
        }    
    }
}

// Названия тестируемых функций должны начинаться на Benchmark,
// за которым идёт название, начинающееся с большой буквы

func BenchmarkSort256(b *testing.B)  {
    for i:= 0; i < b.N; i++  {
        Sort(d[:256])
    }
}

func BenchmarkSort512(b *testing.B)  {
    for i:= 0; i < b.N; i++  {
        Sort(d[:512])
    }
}

func BenchmarkSort1024(b *testing.B)  {
    for i:= 0; i < b.N; i++  {
        Sort(d[:1024])
    }
}

func BenchmarkSort2048(b *testing.B)  {
    for i:= 0; i < b.N; i++  {
        Sort(d[:2048])
    }
}

func BenchmarkSort5096(b *testing.B)  {
    for i:= 0; i < b.N; i++  {
        Sort(d[:5096])
    }
}

func BenchmarkSort100000(b *testing.B)  {
    for i:= 0; i < b.N; i++  {
        Sort(d[:100000])
    }
}
