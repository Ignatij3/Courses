package main

import (
	"fmt"
	"os"
	"time"
	"strings"
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

type (
	lmnt struct {
		x int
		next *lmnt
		prev *lmnt
	}
	list struct {
		head *lmnt
	}
)

func (t list) Print() {
	for runner:= t.head; runner != nil; runner = (*runner).next {
		fmt.Print((*runner).x, " ")
	}
	fmt.Println("\n")
	/*
	end := t.head
	for ; (*end).next != nil; {end = (*end).next}
	for runner:= end; runner != nil; runner = (*runner).prev {
		fmt.Print((*runner).x, " ")
	}
	fmt.Println("\n")*/
}

func (t *list) Add(k int, adr *lmnt) *lmnt {
	(*t).head = &lmnt{k, (*t).head, nil}
	if adr != nil {(*adr).prev = (*t).head}
	return (*t).head
}

func BubbleSortSlice(a []int) ([]int, int) {
	var (
		changes int
	)
	
	for border := 0; border != len(a); border++ {
		for k := 1; k < len(a) - border; k++ {
			if a[k - 1] > a[k] {a[k - 1], a[k] = a[k], a[k - 1]; changes++}
		}
	}
	return a, changes
}

func (t *list) BubbleSortList1() int {
	var (
		changes, steps, i int
		runner1, runner2, runner3, stop, end *lmnt
	)
	for end = (*t).head; end != nil; {steps++; end = (*end).next}
	
	for border := (*t).head; border != end; border = (*border).next {
		for stop, i = (*t).head, 0; i < steps; i++ {stop = (*stop).next}
		for runner1, runner2, runner3 = (*(*t).head).next, (*t).head, nil; runner1 != stop && runner2 != stop; runner3, runner2, runner1 = runner2, runner1, (*runner1).next {
			if runner3 == nil && (*runner2).x > (*runner1).x {
				(*runner1).next, (*runner2).next, (*t).head = (*t).head, (*runner1).next, runner1
			} else if (*runner2).x > (*runner1).x {
				 (*runner1).next, (*runner2).next, (*runner3).next = runner2, (*runner1).next, runner1
				 if runner2 == end {for end = (*t).head; end != nil; {end = (*end).next}}
			}
			changes++
		}
		steps -= 1
	}
	return changes
}

func ShakerSortSlice(a []int) ([]int, int) {
	var (
		changes int
	)
	
	for border := 0; border != len(a); border++ {
		for k := border + 1; k < len(a) - border; k++ {
			if a[k - 1] > a[k] {a[k - 1], a[k] = a[k], a[k - 1]; changes++}
		}
		for k := len(a) - border - 1; k > border; k-- {
			if a[k - 1] > a[k] {a[k - 1], a[k] = a[k], a[k - 1]; changes++}
		}
	}
	return a, changes
}

func (t *list) ShakerSortList2() int {
	var (
		changes int
		runner1, runner2, start, end *lmnt
	)
	
	start, end = (*t).head, (*t).head
	for ; (*end).next != nil; {end = (*end).next}
	
	for ; (*start).next != end && start != end; {
		for runner2, runner1 = start, (*start).next; runner2 != end; {
			if (*runner2).x > (*runner1).x {
				if (*runner2).prev == nil {
					  (*runner1).next, (*runner2).next, (*(*runner1).next).prev, (*runner1).prev, (*runner2).prev = runner2, (*runner1).next, runner2, nil, runner1
					  (*t).head = runner1
				} else if (*runner1).next == nil {
					(*runner1).next, (*runner2).next, (*(*runner2).prev).next, (*runner1).prev, (*runner2).prev = runner2, nil, runner1, (*runner2).prev, runner1
				} else  {
					(*runner1).next, (*runner2).next, (*(*runner2).prev).next, (*(*runner1).next).prev, (*runner1).prev, (*runner2).prev = runner2, (*runner1).next, runner1, runner2, (*runner2).prev, runner1
				}
				runner2, runner1 = runner1, runner2
				if runner2 == end {end = runner1}
				changes++
			}
			
			if (*runner1).next != nil {
				runner2, runner1= runner1, (*runner1).next
			} else {
				break
			}
		}
		
		for runner1, runner2 = end, (*end).prev; runner1 != (*start).prev; {
			if (*runner2).x > (*runner1).x {
				if (*runner2).prev == nil {
					  (*runner1).next, (*runner2).next, (*(*runner1).next).prev, (*runner1).prev, (*runner2).prev = runner2, (*runner1).next, runner2, nil, runner1
					  (*t).head = runner1
				} else if (*runner1).next == nil {
					(*runner1).next, (*runner2).next, (*(*runner2).prev).next, (*runner1).prev, (*runner2).prev = runner2, nil, runner1, (*runner2).prev, runner1
				} else  {
					(*runner1).next, (*runner2).next, (*(*runner2).prev).next, (*(*runner1).next).prev, (*runner1).prev, (*runner2).prev = runner2, (*runner1).next, runner1, runner2, (*runner2).prev, runner1
				}
				runner2, runner1 = runner1, runner2
				if runner1 == start {start = runner2}
				changes++
			}
			
			if (*runner2).prev != nil {
				runner1, runner2 = runner2, (*runner2).prev
			} else {break}
		}
		
		start = (*start).next
		end = (*end).prev
	}
	return changes
}

func main() {
	var (
		nArr, sArr []int
		totalNum, neg, zero, pos, choose, choose1, changes, x int
		pArr, fileCH string
		t list
		address *lmnt
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
	switch choose {
		case 1:
			fileCH = fileEL
		case 2:
			fileCH = fileESL
		case 3:
			fileCH = fileL
		case 4:
			fileCH = fileA
		case 5:
			fileCH = fileB
		case 6:
			fileCH = fileS
		case 7:
			fileCH = fileI
		case 8:
			fileCH = fileHN
	}
	
	file, err := os.Open(fileCH)
	if err != nil {fmt.Println(err)}
	defer file.Close()
	
	fmt.Println("\n1) Bubble Sort Slice")
	fmt.Println("2) Bubble Sort List (1 way)")
	fmt.Println("3) Shaker Sort Slice")
	fmt.Println("4) Shaker Sort List (2 way)")
	fmt.Print("Выберите сортировку, которой хотите сортировать (1 - 4): ")
	fmt.Scan(&choose1)
	
	for {
		_, err := fmt.Fscanf(file, "%d\n", &x)
		if err != nil {break}
		if choose1 == 2 || choose1 == 4 {
			address = t.Add(x, address)
		} else {
			nArr = append(nArr, x)
		}
		
		totalNum++
		if x < 0 {
			neg++
		} else if x == 0 {
			zero++
		} else if x > 0 {
			pos++
		}
	}
	
	for choose1 < 1 || choose1 > 4 {
		fmt.Print("Данные введены неправильно, попробуйте ещё раз:")
		fmt.Scan(&choose1)
	}
	
	start := time.Now()
	switch choose1 {
		case 1:
			sArr, changes = BubbleSortSlice(nArr)
		case 2:
			changes = t.BubbleSortList1()
		case 3:
			sArr, changes = ShakerSortSlice(nArr)
		case 4:
			changes = t.ShakerSortList2()
	}
	end := time.Since(start)
	t2 := start.Add(end)
	diff := t2.Sub(start)
	
	fmt.Println("Всего было", changes, "перестановок")
	fmt.Println("Всего прошло", diff)
	fmt.Print("Вывести массив? (y/n): ")
	fmt.Scan(&pArr)
	
	for pArr != "n" && pArr != "y" {
		fmt.Print("Данные введены неправильно, попробуйте ещё раз (y/n):")
		fmt.Scan(&pArr)
	}
	
	if pArr == "y" && (choose1 == 2 || choose1 == 4) {
		t.Print()
	} else {
		fmt.Println("\n", sArr, "\n"); 
	}
	if neg > 0 {fmt.Println("Всего здесь", neg, "чисел меньше нуля;")}
	if zero > 0 {fmt.Println("Всего здесь", zero, "нулей;")}
	if pos > 0 {fmt.Println("Всего здесь", pos, "чисел больше нуля;")}
	fmt.Println("Всего было найдено", totalNum, "чисел.")
}
