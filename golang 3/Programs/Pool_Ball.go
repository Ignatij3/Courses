package main

import "fmt"

/*
	* dir = 0 - RU - Right UP
	* dir = 1 - RD - Right Down
	* dir = 2 - LU - Left UP
	* dir = 3 - LD - Left Down
	* cases 1 - Left Down
	* cases 2 - Left UP
	* cases 3 - Right UP
	* cases 4 - Right Down
*/
func Move(w, h, v, wS, hS, dir int) (float64, int, string, bool, int) {
	var (
		i, wAdd, hAdd, seconds, bounce, cases int
		timeM float64
		smhdm string = "секунд"
		result bool = true
	)
	i = 1
	wAdd = wS
	hAdd = hS
	if w == h {i = -1; result = false}
	for i >= 0 {
		if ((hAdd >= h && dir == 0) || (wAdd <= 0 && dir == 3) || (dir == 1)) && i >= 0 {
			dir = 1
			for (0 < hAdd) && (w > wAdd) {
				hAdd -= v
				wAdd += v
				seconds++
				if hAdd <= 0 && wAdd >= w {cases = 4; i = -1}
				if hAdd < 0 {hAdd = 0; break}
				if wAdd > w {wAdd = w; break}
				if (wAdd == wS && hAdd == hS) && i > 0 {result = false; break; i = -1}
			}
			bounce++
		}
		if ((hAdd >= h && dir == 2) || (wAdd >= w && dir == 1) || (dir == 3)) && i >= 0 {
			dir = 3
			for (0 < hAdd) && (0 < wAdd) {
				hAdd -= v
				wAdd -= v
				seconds++
				if hAdd <= 0 && wAdd <= 0 {cases = 1; i = -1}
				if hAdd < 0 {hAdd = 0; break}
				if wAdd < 0 {wAdd = 0; break}
				if (wAdd == wS && hAdd == hS) && i > 0 {result = false; break; i = -1}
			}
			bounce++
		}
		if ((hAdd <= 0 && dir == 3) || (wAdd >= w && dir == 0) || (dir == 2)) && i >= 0 {
			dir = 2
			for (h > hAdd) && (0 < wAdd) {
				hAdd += v
				wAdd -= v
				seconds++
				if hAdd >= h && wAdd <= 0 {cases = 2; i = -1}
				if hAdd > h {hAdd = h; break}
				if wAdd < 0 {wAdd = 0; break}
				if (wAdd == wS && hAdd == hS) && i > 0 {result = false; break; i = -1}
			}
			bounce++
		}
		if ((hAdd <= 0 && dir == 1) || (wAdd <= 0 && dir == 2) || (dir == 0)) && i >= 0 {
			dir = 0
			for (h > hAdd) && (w > wAdd) {
				hAdd += v
				wAdd += v
				seconds++
				if hAdd >= h && wAdd >= w {cases = 3; i = -1}
				if hAdd > h {hAdd = h; break}
				if wAdd > w {wAdd = w; break}
				if (wAdd == wS && hAdd == hS) && i > 0 {result = false; break; i = -1}
			}
			bounce++
		}
		if i == -1 {bounce -= 1; seconds -= 1}
	}
	timeM = float64(seconds)
	if timeM >= 60 {
		timeM /= 60
		smhdm = "минуты"
	}
	if timeM >= 60 {
		timeM /= 60
		smhdm = "часов"
	}
	if timeM >= 24 {
		timeM /= 24
		smhdm = "день"
	}
	if timeM >= 30 {
		timeM /= 30
		smhdm = "месяцев"
	}
	return timeM, bounce, smhdm, result, cases
}

func main() {
	var (
		tableH, tableW, cTime, speed, wPos, hPos, direction int
	)
	fmt.Print("Введите ширину стола(x): ")
	fmt.Scan(&tableW)
	fmt.Print("Введите высоту стола(y): ")
	fmt.Scan(&tableH)
	for tableW <= 0 || tableH <= 0 {
		fmt.Println("Данные введены неправильно, попробуйте ещё раз")
		fmt.Println("Ширина стола(x): ")
		fmt.Scan(&tableW)
		fmt.Println("Высота стола(y): ")
		fmt.Scan(&tableH)
	}
	fmt.Print("Введите точку старта(x): ")
	fmt.Scan(&wPos)
	fmt.Print("Введите точку старта(y): ")
	fmt.Scan(&hPos)
	for wPos < 0 || wPos > tableW || hPos < 0 || hPos > tableH {
		fmt.Println("Данные введены неправильно, попробуйте ещё раз")
		fmt.Print("Точка старта(x): ")
		fmt.Scan(&wPos)
		fmt.Print("Точка старта(y): ")
		fmt.Scan(&hPos)
	}
	fmt.Println("Введите направление")
	fmt.Println("Направление 0 - Направо вверх")
	fmt.Println("Направление 1 - Направо вниз")
	fmt.Println("Направление 2 - Налево вверх")
	fmt.Println("Направление 3 - Налево вниз")
	fmt.Scan(&direction)
	for direction < 0 || direction > 3 {
		fmt.Println("Данные введены неправильно, попробуйте ещё раз")
		fmt.Print("Введите направление: ")
		fmt.Scan(&direction)
	}
	fmt.Print("Введите скорость шара: ")
	fmt.Scan(&speed)
	for speed <= 0 {
		fmt.Println("Данные введены неправильно, попробуйте ещё раз")
		fmt.Print("Введите скорость шара: ")
		fmt.Scan(&direction)
	}
	time, bounce, smhdm, result, cases := Move(tableW, tableH, speed, wPos, hPos, direction)
	cTime = int(time)
	if result {
		if cases == 1 {
			fmt.Println("\nВремя, за которое шар попал в левый нижний угол: ~", cTime, smhdm)
			fmt.Println("Кол-во отскоков: ", bounce)
		}
		if cases == 2 {
			fmt.Println("\nВремя, за которое шар попал в левый верхний угол: ~", cTime, smhdm)
			fmt.Println("Кол-во отскоков: ", bounce)
		}
		if cases == 3 {
			fmt.Println("\nВремя, за которое шар попал в правый верхний угол: ~", cTime, smhdm)
			fmt.Println("Кол-во отскоков: ", bounce)
		}
		if cases == 4 {
			fmt.Println("\nВремя, за которое шар попал в правый нижний угол: ~", cTime, smhdm)
			fmt.Println("Кол-во отскоков: ", bounce)
		}
	} else if result == false {
		fmt.Println("\nВ данной конфигурации стола шар никогда не попадёт ни в один угол")
		fmt.Println("Всего было", bounce, "отскоков")
	}
}
//fmt.Println(wAdd, "wAdd", hAdd, "hAdd", dir, "dir") - расставить в каждом for, при надобности
//Одна секунда равняется одному "шагу"
