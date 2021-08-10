package main

import  (
	"fmt"
	"time"
)

func Working(k, n, a int) (int, int, int, int, int, int, int, int) {
	var (
		kC, nC, kT, nT, aT, kNn, nNk, NnNk, kn, kH, nH, check int
	)
	kH = k
	nH = n
	for i := 0; i >= 0 ; i++ {
		//fmt.Println(kC, "kC", nC, "nC", kT, "kT", nT, "nT", aT, "aT", k, "k", n, "n", kNn, "kNn", nNk, "nNk", NnNk, "NnNk", kn, "kn")
		if aT != k && aT != n {kn++} //coloured, marked
		if aT == k {
			if aT == k && aT == n {
				NnNk++ //not coloured, not marked
				check = 1
			}
			if aT == k && aT != n {nNk++} //not coloured, marked
			k += kH
			kC++ //not coloured
			kT--
		}
		if aT == n {
			if aT != k && aT == n && check == 0 {kNn++} //coloured, not marked
			n += nH
			nC++ //not marked
			nT--
			check = 0
		}
		kT++ //total count
		nT++ //total count
		aT++ //total count
		if aT == a {break}
	}
	return kC, nC, kT, nT, kNn, nNk, NnNk, kn
}

func main() {
	var (
		k, n, a int
	)
	fmt.Print("Введите кол-во времени до пропуска (k) для красящей машины: ")
	fmt.Scan(&k)
	fmt.Print("Введите кол-во времени до пропуска (n) для маркирующей машины: ")
	fmt.Scan(&n)
	fmt.Print("Введите кол-во времени (a), через которое рабочие заметили неисправность: ")
	fmt.Scan(&a)
	kC, nC, kT, nT, kNn, nNk, NnNk, kn := Working(k, n, a)
	fmt.Println("Всего покрашено: ", kT)
	time.Sleep(2 * time.Second)
	fmt.Println("Всего маркировано: ", nT)
	time.Sleep(2 * time.Second)
	fmt.Println("Всего не покрашено: ", kC)
	time.Sleep(2 * time.Second)
	fmt.Println("Всего не маркировано: ", nC)
	time.Sleep(2 * time.Second)
	fmt.Println("Всего не маркировано, но покрашено: ", kNn)
	time.Sleep(2 * time.Second)
	fmt.Println("Всего маркировано, но не покрашено: ", nNk)
	time.Sleep(2 * time.Second)
	fmt.Println("Всего не маркировано и не покрашено: ", NnNk)
	time.Sleep(2 * time.Second)
	fmt.Println("Всего маркировано и покрашено: ", kn)
}
