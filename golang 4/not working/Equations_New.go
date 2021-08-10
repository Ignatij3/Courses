package main

import (
	"bufio"
	"math"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	task00 = "../Files/task00.txt"
	task01 = "../Files/task01.txt"
	task02 = "../Files/task02.txt"
	task03 = "../Files/task03.txt"
	task04 = "../Files/task04.txt"
	task05 = "../Files/task05.txt"
	task06 = "../Files/task06.txt"
	task07 = "../Files/task07.txt"
	task08 = "../Files/task08.txt"
	task09 = "../Files/task09.txt"
	task10 = "../Files/task10.txt"
	task11 = "../Files/task11.txt"
	task12 = "../Files/task12.txt"
	task13 = "../Files/taskZeros.txt"
	task14 = "../Files/taskLess.txt"
	task15 = "../Files/taskLessZeros.txt"
	task16 = "../Files/taskMore.txt"
	task17 = "../Files/taskMoreZeros.txt"
)

func Preview(slice2D [][]float64, LastNum []float64) {	
	for i := 0; i < len(slice2D); i++ {
		for j := 0; j < len(slice2D[i]); j++ {
			if slice2D[i][j] == 1 || slice2D[i][j] == -1 || slice2D[i][j] == 0 {
				if j == 0 {
					switch slice2D[i][j] {
						case -1 :
							fmt.Printf("-x%v", j+1)
						case 0 :
							fmt.Print("0")
						case 1 :
							fmt.Printf("x%v", j+1)
					}
				} else {
					switch slice2D[i][j] {
						case -1 :
							fmt.Printf(" - x%v", j+1)
						case 0 :
							fmt.Print(" + 0")
						case 1 :
							fmt.Printf(" + x%v", j+1)
					}
				}
			} else if j == 0 {
				fmt.Printf("%v*x%v", slice2D[i][j], j+1)
			} else if slice2D[i][j] >= 0 {
				fmt.Printf(" + %v*x%v", slice2D[i][j], j+1)
			} else if slice2D[i][j] < 0 {
				fmt.Printf(" - %v*x%v", math.Abs(slice2D[i][j]), j+1)
			}
			if j == len(slice2D[i]) - 1 {
				fmt.Printf(" = %v\n", LastNum[i])
			}
			
		}
	}
}

func Calculate(slice2D [][]float64, LastNum []float64) {
	var (
		multiplier, combined, top, resCheck float64
		originCol []int
		row0, col0, stopRow, stopCol, back, numlines, numvar, min, max int
		infinite bool
	)
	
	numlines = len(slice2D)
	numvar = len(slice2D[0])
	if numlines <= numvar {
		min, back = numlines, numlines - 1
		max = numvar
	} else if numlines > numvar {
		min, back = numvar, numvar - 1
		max = numlines
	}
	showSlice := make([][]float64, numlines, numlines)
	showLastNum := make([]float64, numlines, numlines)
	sorted := make([]float64, max, max) //Если неизвестных больше, то ставить по numvar, если меньше - по numlines
	answers := make([]float64, max, max)
	//stop = numlines - 1 //ПРОБЛЕМА
	stopRow = numlines - 1
	stopCol = numvar - 1
	
	for line := 0; line < numlines; line++ {
		for unknown := 0; unknown < numvar; unknown++ {
			showSlice[line] = append(showSlice[line], slice2D[line][unknown])
		}
		showLastNum[line] = LastNum[line]
	}
	
	for i := 0; i < numvar; i++ {originCol = append(originCol, i)}
	
	for step := 0; step < min; step++ {
		//Поиск главного элемента
		top = 0
		for newRow := step; newRow < numlines; newRow++ {
			for newCol := step; newCol < len(slice2D[newRow]); newCol++ {
				if top < math.Abs(slice2D[newRow][newCol]) {
					top = math.Abs(slice2D[newRow][newCol])
					row0, col0 = newRow, newCol
				}
			}
		}
		
		if top == 0 {
			resCheck = 0
			fmt.Println()
			for newRes := step; newRes < numlines; newRes++ {
				if resCheck <= math.Abs(LastNum[newRes]) {
					resCheck = math.Abs(LastNum[newRes])
				}
			}
			if resCheck != 0 {
				fmt.Println("Система не имеет решений\n")
				return
			} else {
				fmt.Println("Система имеет бесконечное множество решений\n")
				infinite = true
			}
		} else {
			// Переставляем колонки step и col0
			for changeRow := 0; changeRow < numlines; changeRow++ {
				slice2D[changeRow][step], slice2D[changeRow][col0] = slice2D[changeRow][col0], slice2D[changeRow][step]
			}
			
			originCol[step], originCol[col0] = originCol[col0], originCol[step]
			
			// Переставляем ряды step и row0
			for changeCol := step; changeCol < len(slice2D[row0]); changeCol++ {
				slice2D[step][changeCol], slice2D[row0][changeCol] = slice2D[row0][changeCol], slice2D[step][changeCol]
			}
			LastNum[step], LastNum[row0] = LastNum[row0], LastNum[step]
			
			// Вычисления 	
			for row := step + 1; row < numlines; row++  {
				multiplier = -slice2D[row][step] / slice2D[step][step]
				slice2D[row][step] = 0
				for col := step + 1; col < len(slice2D[row]); col++  {
					slice2D[row][col] += slice2D[step][col] * multiplier
				}
				LastNum[row] += LastNum[step]*multiplier
			}
		}
		//fmt.Println(step, "step")
		//Preview(slice2D, LastNum)
		//fmt.Println("\n")
		if top == 0 && resCheck == 0 {
			stopRow = step - int(math.Abs(float64(numlines) - float64(numvar))) - 1
			stopCol = step - int(math.Abs(float64(numlines) - float64(numvar))) - 1
			///stopCol / stopRow = (step - 1) - abs(numlines - numvar)
			break
		}
	}
	
	fmt.Println(stopRow, "stopRow")
	fmt.Println(stopCol, "stopCol")
	
	if numlines < numvar && top != 0 {infinite = true}
	for row, col := stopRow, stopCol; row >= 0 && col >= 0; {
			answers[row] = LastNum[row]
			for col2 := back; col2 > row; col2-- {
				combined = slice2D[row][col2] * answers[col2]
				answers[row] -= combined
			}
			answers[row] = answers[row] / slice2D[row][col] //ERROR сделать аналог col
			row--
			col--
	}
	
	for i := 0; i < numvar; i++ { //len(answers)
		sorted[originCol[i]] = answers[i] //Сделать так, что-бы числа не превосходили кол-во ответов
	}
	
	fmt.Println("Last")
	fmt.Println(answers, "\n")
	
	Show(numlines, numvar, sorted, showSlice, showLastNum, infinite)
}

func Show(numlines, numvar int, sorted []float64, showSlice [][]float64, showLastNum []float64, infinite bool) {
	var (
		rounded string
		split1, split2 []string
	)
	
	fmt.Println("\nResult-----------------------------------------------------------\n")
	for r := 0; r < numvar; r++ {
		if (infinite == true && sorted[r] == 0) || (numlines < numvar && r >= len(sorted)) {
			fmt.Printf("%d-е неизвестное - Any\n", r+1)
		} else {
			fmt.Printf("%d-е неизвестное - %v\n", r+1, sorted[r])
		}
	}
	fmt.Println()
	for i := 0; i < numlines; i++ {
		for j := 0; j < numvar; j++ {
			if (infinite == true && sorted[j] == 0) || (numlines < numvar && j >= len(sorted)) {
				rounded = fmt.Sprintf("%f", showSlice[i][j])
				split1 = strings.Split(rounded, ".")
				split2 = []string{"(-inf;+inf)", ""}
			} else {
				rounded = fmt.Sprintf("%f", showSlice[i][j])
				split1 = strings.Split(rounded, ".")
				rounded = fmt.Sprintf("%f", sorted[j])
				split2 = strings.Split(rounded, ".")
				
				split2[1] = "." + split2[1]
				if split2[1] == ".000000" {split2[1] = ""}
			}
			
			split1[1] = "." + split1[1]
			if split1[1] == ".000000" {split1[1] = ""}
			
			if (split1[0] == "0" && split1[1] == "") || (split2[0] == "0" && split2[1] == "") {
				if j == 0 {fmt.Print("0")} else {fmt.Print(" + 0")}
			} else if (split1[0] == "-1" || split1[0] == "1") && split1[1] == "" {
				if j == 0 {
					fmt.Printf("%s%s", split2[0], split2[1])
				} else {
					if splitnum, _ := strconv.Atoi(split2[0]); splitnum < 0 {
						switch split1[0] {
							case "-1" :
								fmt.Printf(" + %s%s", strings.ReplaceAll(split2[0], "-", ""), split2[1])
							case "1" :
								fmt.Printf(" - %s%s", strings.ReplaceAll(split2[0], "-", ""), split2[1])
						}
					} else if splitnum, _ := strconv.Atoi(split2[0]); splitnum >= 0 {
						switch split1[0] {
							case "-1" :
								fmt.Printf(" - %s%s", split2[0], split2[1])
							case "1" :
								fmt.Printf(" + %s%s", split2[0], split2[1])
						}
					}
				}
			} else if (split2[0] == "-1" || split2[0] == "1") && split2[1] == "" {
				if j == 0 {
					fmt.Printf("%s%s", split1[0], split1[1])
				} else {
					if splitnum, _ := strconv.Atoi(split1[0]); splitnum < 0 {
						switch split2[0] {
							case "-1" :
								fmt.Printf(" + %s%s", strings.ReplaceAll(split1[0], "-", ""), split1[1])
							case "1" :
								fmt.Printf(" - %s%s", strings.ReplaceAll(split1[0], "-", ""), split1[1])
						}
					} else if splitnum, _ := strconv.Atoi(split1[0]); splitnum >= 0 {
						switch split2[0] {
							case "-1" :
								fmt.Printf(" - %s%s", split1[0], split1[1])
							case "1" :
								fmt.Printf(" + %s%s", split1[0], split1[1])
						}
					}
				}
			} else {
				if j == 0 {
					fmt.Printf("%s%s*%s%s", split1[0], split1[1], split2[0], split2[1])
				} else if showSlice[i][j] > 0 {
					fmt.Printf(" + %s%s*%s%s", split1[0], split1[1], split2[0], split2[1])
				} else if showSlice[i][j] < 0 {
					fmt.Printf(" - %s%s*%s%s", strings.ReplaceAll(split1[0], "-", ""), split1[1], split2[0], split2[1])
				}
			}
			if j == len(showSlice[i]) - 1 {
				fmt.Printf(" = %v\n", showLastNum[i])
			}
		}
	}
	fmt.Println("\nResult-----------------------------------------------------------")
}

func Input() ([][]float64, []float64) {
	var (
		sideLeft [][]float64
		sideRight, line, newLine []float64
		choose int
		lines int64
		filePath string
	)
	
	fmt.Println("0) ", strings.Split(task00, "../Files/")[len(strings.Split(task00, "../Files/")) - 1])
	fmt.Println("1) ", strings.Split(task01, "../Files/")[len(strings.Split(task01, "../Files/")) - 1])
	fmt.Println("2) ", strings.Split(task02, "../Files/")[len(strings.Split(task02, "../Files/")) - 1])
	fmt.Println("3) ", strings.Split(task03, "../Files/")[len(strings.Split(task03, "../Files/")) - 1])
	fmt.Println("4) ", strings.Split(task04, "../Files/")[len(strings.Split(task04, "../Files/")) - 1])
	fmt.Println("5) ", strings.Split(task05, "../Files/")[len(strings.Split(task05, "../Files/")) - 1])
	fmt.Println("6) ", strings.Split(task06, "../Files/")[len(strings.Split(task06, "../Files/")) - 1])
	fmt.Println("7) ", strings.Split(task07, "../Files/")[len(strings.Split(task07, "../Files/")) - 1])
	fmt.Println("8) ", strings.Split(task08, "../Files/")[len(strings.Split(task08, "../Files/")) - 1])
	fmt.Println("9) ", strings.Split(task09, "../Files/")[len(strings.Split(task09, "../Files/")) - 1])
	fmt.Println("10) ", strings.Split(task10, "../Files/")[len(strings.Split(task10, "../Files/")) - 1])
	fmt.Println("11) ", strings.Split(task11, "../Files/")[len(strings.Split(task11, "../Files/")) - 1])
	fmt.Println("12) ", strings.Split(task12, "../Files/")[len(strings.Split(task12, "../Files/")) - 1])
	fmt.Println("13) ", strings.Split(task13, "../Files/")[len(strings.Split(task13, "../Files/")) - 1])
	fmt.Println("14) ", strings.Split(task14, "../Files/")[len(strings.Split(task14, "../Files/")) - 1])
	fmt.Println("15) ", strings.Split(task15, "../Files/")[len(strings.Split(task15, "../Files/")) - 1])
	fmt.Println("16) ", strings.Split(task16, "../Files/")[len(strings.Split(task16, "../Files/")) - 1])
	fmt.Println("17) ", strings.Split(task17, "../Files/")[len(strings.Split(task17, "../Files/")) - 1])
	
	start:
	fmt.Print("Выберите уравнение, которое хотите решить (0 - 17): ")
	fmt.Scan(&choose)
	
	switch choose {
		case 0 :
			filePath = task00
		case 1 :
			filePath = task01
		case 2 :
			filePath = task02
		case 3 :
			filePath = task03
		case 4 :
			filePath = task04
		case 5 :
			filePath = task05
		case 6 :
			filePath = task06
		case 7 :
			filePath = task07
		case 8 :
			filePath = task08
		case 9 :
			filePath = task09
		case 10 :
			filePath = task10
		case 11 :
			filePath = task11
		case 12 :
			filePath = task12
		case 13 :
			filePath = task13
		case 14 :
			filePath = task14
		case 15 :
			filePath = task15
		case 16 :
			filePath = task16
		case 17 :
			filePath = task17
		default:
			fmt.Println("Данные введены неправильно")
			goto start
	}
	fmt.Println()
	
	fin, _ := os.Open(filePath)
	defer fin.Close()
	
	scanner := bufio.NewScanner(fin)
	scanner.Scan()
	num := strings.Fields(scanner.Text())
	lines, _ = strconv.ParseInt(num[0], 10, 64)
	
	for scanner.Scan() {
		line = make([]float64, 0)
		newLine = make([]float64, 0)
		for _, snum := range strings.Fields(scanner.Text()) {
			if x, err := strconv.ParseFloat(snum, 64); err == nil {
				line = append(line, x)
				newLine = append(newLine, x)
			}
		}
		if len(line) == 0 {break}
		fmt.Println(len(line), "len(line)", lines)
		sideLeft = append(sideLeft, line[:len(line)-1])
		sideRight = append(sideRight, line[len(line)-1])
	}
	
	return sideLeft, sideRight
}

func main() { //2lazy2fix
	var continues string
	
	start:
	sideLeft, sideRight := Input()
	Preview(sideLeft, sideRight)
	Calculate(sideLeft, sideRight)
	sideLeft, sideRight = nil, nil
	
	fmt.Print("\nХотите решить ещё одно уравнение? (y/n)")
	fmt.Scan(&continues)
	
	for continues != "y" && continues != "n" {
		fmt.Print("Данные введены неправильно, попробуйте ещё раз:")
		fmt.Scan(&continues)
	}
	if continues == "y" {goto start} else {fmt.Println("\nУдачи")}
}
