package main
import  (
	"fmt"
	"time"
)	

func searchDivisor(n int64, from, to int64, done chan int64) {
	// Ищет первый делитель числа n на интервале [from; to)
	// и отправляет его в канал
	// Если таких нет, то посылает в канал 0.
    for i:= from; i < to; i++ {
        if n%i == 0 {
            done <- i
            return
        } 
    }
    done <- 0
}

func main() {
    var n, start, step, res  int64
    n = 1000000087*1120000093
	
	t0 := time.Now()
    step = 8388608  // 2^23
	done := make(chan int64)
	count:= 0
	for start = 2; start*start <= n; start += step {
		// Разбиваем интервал от 2 до Sqrt(n) на отрезки
		// длиной 2^23 и для каждого из них запускаем 
		// go-функцию, которая ищет делитель
		go searchDivisor(n, start, start + step, done)
		// count - количество запущенных go-подпрограмм
		count++
	}	
	loop:
	for {
		// Собираем результаты выполнения go-подпрограмм
		select {
		case res = <- done:
			if res > 0  {
				// Если go-подпрограмма нашла делитель, заканчиваем процесс
				fmt.Printf("%d = %d * %d\n", n, res, n/res)
				break loop
			}  else  {
				// иначе, отмечаем, что одна go-подпрограмма закончила своё выполнение
				count--
			}				
		default:	
			// Если в канале ничего нет,
			if count == 0  {  // и запущенных go-подпрограмм не осталось,
				// то у числа n нет собственных делителей, и оно простое.
				fmt.Println(n, "is Prime")
				break loop
			}
		}
	}
	t1 := time.Now()
	fmt.Printf("The call took %v to run.\n", t1.Sub(t0))
}
