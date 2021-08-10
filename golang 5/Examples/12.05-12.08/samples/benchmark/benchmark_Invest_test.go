package main

// Запуск из командной строки:
//         go test -bench . benchmark_Invest_test.go
// Имя файла обязательно должно заканиваться на _test

import (
	"fmt"
    "os"
    "testing"
)

var (
    inc       [][]int
    n, maxsum int
)

func init() {
    f, _ := os.Open("income.dat")
    fmt.Fscan(f, &n, &maxsum)
    // n - количество вариантов инвестиций (производств)
    // sum - максимально возможная суммарная сумма инвестиций
    // inc[i][s] - доход, который приносит инвестиция
    //             s денег в производство #i
    inc = append(inc, make([]int, maxsum+1, maxsum+1))
    for i := 1; i <= n; i++ {
        inc = append(inc, make([]int, maxsum+1, maxsum+1))
        inc[i][0] = 0
        for s := 1; s <= maxsum; s++ {
            fmt.Fscan(f, &inc[i][s])
        }
    }
}

func MaxIncomeRecursive(n int, sum int) (result int) {
    // инвестируем sum денег в проекты 1..n
    var w int
    if n > 0 {
        result = inc[n][sum]
        for reminder := 1; reminder <= sum; reminder++ {
            w = inc[n][sum-reminder] + MaxIncomeRecursive(n-1, reminder)
            if w > result {
                result = w
            }
        }
        return result
    } else {
        return 0
    }
}

func MaxIncomeCyclic(n int, sum int) int {
    // инвестируем sum денег в проекты 1..n
    var max, w int
    var res [][]int
    res = append(res, make([]int, sum+1, sum+1))
    for k := 1; k <= n; k++ {
        res = append(res, make([]int, sum+1, sum+1))
        max = inc[k][sum]
        for s := 1; s <= sum; s++ {
            for reminder := 1; reminder <= s; reminder++ {
                w = inc[n][s-reminder] + res[k-1][reminder]
                if w > max {
                    max = w
                }
            }
            res[k][s] = max
        }
    }
    return res[n][sum]
}

type memo [][]int

func MaxIncome(n int, sum int, m memo) int {
    // Рекурсивная функция с мемоизацией
    // инвестируем sum денег в проекты 1..n
    var w, result int
    if n > 0 && m[n][sum] == 0 {
        result = inc[n][sum]
        for reminder := 1; reminder <= sum; reminder++ {
            w = inc[n][sum-reminder] + MaxIncome(n-1, reminder, m)
            if w > result {
                result = w
            }
        }
        m[n][sum] = result
    }
    return m[n][sum]
}

func MaxIncomeMemo(n int, sum int) (result int) {
    // Слайс res используется для мемоизации
    var res [][]int
    // Подготавливаем слайс res, он заполняется нулями.
    for i := 0; i <= n; i++ {
        res = append(res, make([]int, sum+1, sum+1))
    }
    // Рекурсивно (с мемоизацией) вычисляем максимальный доход,
    // получаемый при инвестировании sum денег в проекты 1..n.
    return MaxIncome(n, sum, res)
}

// Названия тестируемых функций должны начинаться на Benchmark,
// за которым идёт название, начинающееся с большой буквы

func BenchmarkInvestRecursive_5_10(b *testing.B) {
    for i := 0; i < b.N; i++ {
        MaxIncomeRecursive(5, 10)
    }
}

func BenchmarkInvestRecursive_4_20(b *testing.B) {
    for i := 0; i < b.N; i++ {
        MaxIncomeRecursive(4, 20)
    }
}

func BenchmarkInvestRecursive_5_20(b *testing.B) {
    for i := 0; i < b.N; i++ {
        MaxIncomeRecursive(5, 20)
    }
}

func BenchmarkInvestCyclic_5_10(b *testing.B) {
    for i := 0; i < b.N; i++ {
        MaxIncomeCyclic(5, 10)
    }
}

func BenchmarkInvestCyclic_4_20(b *testing.B) {
    for i := 0; i < b.N; i++ {
        MaxIncomeCyclic(4, 20)
    }
}

func BenchmarkInvestCyclic_5_20(b *testing.B) {
    for i := 0; i < b.N; i++ {
        MaxIncomeCyclic(5, 20)
    }
}

func BenchmarkInvestMemo_5_10(b *testing.B) {
    for i := 0; i < b.N; i++ {
        MaxIncomeMemo(5, 10)
    }
}

func BenchmarkInvestMemo_4_20(b *testing.B) {
    for i := 0; i < b.N; i++ {
        MaxIncomeMemo(4, 20)
    }
}

func BenchmarkInvestMemo_5_20(b *testing.B) {
    for i := 0; i < b.N; i++ {
        MaxIncomeMemo(5, 20)
    }
}
