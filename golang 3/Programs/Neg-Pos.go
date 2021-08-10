package main

import (
	"fmt"
	"os"
	"time"
	"strings"
	"math/rand"
)

const (
		fileEL = "../Files/NumELite.dat"
		fileESL = "../Files/NumESLite.dat"
		fileL = "../Files/NumLite.dat"
		fileA = "../Files/NumAvg.dat"
		fileB = "../Files/NumBig.dat"
		fileS = "../Files/NumSuper.dat"
		fileI = "../Files/NumInsane.dat"
		fileHN = "../Files/HugeNumbers.dat"
)

func BubbleSort(a []int) ([]int, int) {
	var (
		comp, i, ch int
	)
	i = 1
	for {
		if a[i - 1] > a[i] {
			a[i - 1], a[i] = a[i], a[i - 1]
			ch = 1
		}
		if i == len(a) - 1 && ch == 0 {break}
		if i == len(a) - 1 {
			i = 0
			ch = 0
		}
		i++
		comp++
	}
	return a, comp
}

func QuickSort(a []int, comp int) ([]int, int) {
	//fmt.Println(comp)
	if len(a) < 2 {return a, comp}
	left, right := 0, len(a) - 1
	pivotIndex := rand.Int() % len(a)
	a[pivotIndex], a[right] = a[right], a[pivotIndex]
	for i := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
		comp++
	}
	a[left], a[right] = a[right], a[left]
	QuickSort(a[:left], comp)
	QuickSort(a[left + 1:], comp)
	return a, comp
}

func MergeSort(a []int) ([]int, int) {
	var num = len(a)
	var comp int
	if num == 1 {
		return a, comp
	}
	mid := int(num / 2)
	var (
		low = make([]int, mid)
		high = make([]int, num-mid)
	)
	for i := 0; i < num; i++ {
		if i < mid {
			low[i] = a[i]
		} else {
			high[i-mid] = a[i]
		}
	}
	low, comp = MergeSort(low)
	high, comp = MergeSort(high)
	a, comp = merge(low, high)
	return a, comp
}
  
func merge(low, high []int) (a []int, comp int) {
	a = make([]int, len(low) + len(high))
	i, comp := 0, 0
	for len(low) > 0 && len(high) > 0 {
		if low[0] < high[0] {
			a[i] = low[0]
			low = low[1:]
		} else {
			a[i] = high[0]
			high = high[1:]
		}
		i++
		comp++
	}
	for j := 0; j < len(low); j++ {
		a[i] = low[j]
		i++
	}
	for j := 0; j < len(high); j++ {
		a[i] = high[j]
		i++
	}
	return
}

func main() {
	var (
		nArr, sArr []int
		temp, totalNum, neg, zero, pos, i, choose, choose1, comp int = 0, 0, 0, 0, 0, 0, 0, 0, 0
		pArr, fileCH, wrTFile, filePath, fileAdd string = "", "", "", "../Files/", ""
	)
	
	fmt.Println("1) ", strings.Split(fileEL, "../Files/")[len(strings.Split(fileEL, "../Files/")) - 1], "   --> 100 чисел       [-100 - 100)")
	fmt.Println("2) ", strings.Split(fileESL, "../Files/")[len(strings.Split(fileESL, "../Files/")) - 1], "  --> 500 чисел       [-1 000 - 1 000)")
	fmt.Println("3) ", strings.Split(fileL, "../Files/")[len(strings.Split(fileL, "../Files/")) - 1], "    --> 1 000 чисел     [-1 000 - 1 000)")
	fmt.Println("4) ", strings.Split(fileA, "../Files/")[len(strings.Split(fileA, "../Files/")) - 1], "     --> 10 000 чисел    [-10 000 - 10 000)")
	fmt.Println("5) ", strings.Split(fileB, "../Files/")[len(strings.Split(fileB, "../Files/")) - 1], "     --> 50 000 чисел    [-10 000 - 10 000)")
	fmt.Println("6) ", strings.Split(fileS, "../Files/")[len(strings.Split(fileS, "../Files/")) - 1], "   --> 100 000 чисел   [-10 000 - 10 000)")
	fmt.Println("7) ", strings.Split(fileI, "../Files/")[len(strings.Split(fileI, "../Files/")) - 1], "  --> 1 000 000 чисел [-100 000 - 100 000)")
	fmt.Println("8) ", strings.Split(fileHN, "../Files/")[len(strings.Split(fileHN, "../Files/")) - 1], "--> 56 207 чисел    [2 165 - 900 000 000 000)")
	
	fmt.Print("Выберите файл, который хотите отсортировать (1 - 8): ")
	fmt.Scan(&choose)
	for choose < 1 || choose > 8 {
		fmt.Print("Данные введены неправильно, попробуйте ещё раз:")
		fmt.Scan(&choose)
	}
	if choose == 1 {
		fileCH = fileEL
	} else if choose == 2 {
		fileCH = fileESL
	} else if choose == 3 {
		fileCH = fileL
	} else if choose == 4 {
		fileCH = fileA
	} else if choose == 5 {
		fileCH = fileB
	} else if choose == 6 {
		fileCH = fileS
	} else if choose == 7 {
		fileCH = fileI
	} else if choose == 8 {
		fileCH = fileHN
	}
	
	file, err := os.Open(fileCH)
	if err != nil {fmt.Println(err)}
	defer file.Close()
	for {
		if _, err := fmt.Fscanln(file, &temp); err == nil {
			nArr = append(nArr, temp)
			totalNum++
			if nArr[i] < 0 {
					neg++
				} else if nArr[i] == 0 {
					zero++
				} else if nArr[i] > 0 {
					pos++
				}
				i++
		} else {
			break
		}
	}
	
	fmt.Println("1) Bubble Sort")
	fmt.Println("2) Quick Sort")
	fmt.Println("3) Merge Sort")
	fmt.Println("4) Все")
	fmt.Print("Выберите сорт, которым хотите сортировать (1 - 4 (\"4\" - Для выбора всех видов сорта)): ")
	fmt.Scan(&choose1)
	
	for choose1 < 1 || choose1 > 4 {
		fmt.Print("Данные введены неправильно, попробуйте ещё раз:")
		fmt.Scan(&choose1)
	}
	if choose1 == 1 || choose1 == 4 {
		start := time.Now()
		sArr, comp = BubbleSort(nArr)
		end := time.Since(start)
		t2 := start.Add(end)
		diff := t2.Sub(start)
		fmt.Println("Всего было", comp, "сравнений.")
		fmt.Println("Всего прошло", diff)
		fmt.Println("Сортировка - Bubble sort\n")
	}
	if choose1 == 2 || choose1 == 4 {
		start := time.Now()
		sArr, comp = QuickSort(nArr, 0)
		end := time.Since(start)
		t2 := start.Add(end)
		diff := t2.Sub(start)
		fmt.Println("Всего было", comp, "сравнений.")
		fmt.Println("Всего прошло", diff)
		fmt.Println("Сортировка - Quick sort\n")
	}
	if choose1 == 3 || choose1 == 4 {
		start := time.Now()
		sArr, comp = MergeSort(nArr)
		end := time.Since(start)
		t2 := start.Add(end)
		diff := t2.Sub(start)
		fmt.Println("Всего было", comp, "сравнений.")
		fmt.Println("Всего прошло", diff)
		fmt.Println("Сортировка - Merge sort\n")
	}
	
	fmt.Print("Вывести массив? (y/n): ")
	fmt.Scan(&pArr)
	for pArr != "n" && pArr != "y" {
		fmt.Print("Данные введены неправильно, попробуйте ещё раз (y/n):")
		fmt.Scan(&pArr)
	}
	
	if pArr == "y" {fmt.Println("\n", sArr, "\n")}
	if neg > 0 {fmt.Println("Всего здесь", neg, "чисел меньше нуля;")}
	if zero > 0 {fmt.Println("Всего здесь", zero, "нулей;")}
	if pos > 0 {fmt.Println("Всего здесь", pos, "чисел больше нуля;")}
	fmt.Println("Всего было найдено", totalNum, "чисел.")
	
	fmt.Print("Хотите записать это в файл? (y/n)")
	fmt.Scan(&wrTFile)
	for wrTFile != "y" && wrTFile != "n" {
		fmt.Print("Данные введены неправильно, попробуйте ещё раз(y/n):")
		fmt.Scan(&wrTFile)
	}
	if wrTFile == "y" {
		fmt.Print("Введите название файла (без расширения!):")
		fmt.Scan(&fileAdd)
	}
	
	filePath += fileAdd + ".dat"
	fileN, err1 := os.Create(filePath)
	if err1 != nil {fmt.Println(err1)}
	for i := range sArr {
		_, err = fileN.WriteString(fmt.Sprintf("%v\n", sArr[i]))
		if err != nil {fmt.Println(err); return}
	}
	fmt.Println("Ваш файл записан по путю", filePath)
	fileN.Close()
}
