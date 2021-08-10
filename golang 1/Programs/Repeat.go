package main
import (
		"fmt"
		"bufio"
		"os"
		"strconv"
		)

func exists (name string) bool {
	f, err:=os.Open(name)
	if err==nil {
		f.Close()
		return true
	} else {
		return !os.IsNotExist(err)
	}
}

func NumbersRepeat() map[int]int {
	Num := make(map[int]int)
	fin, _ := os.Open("..\\Files\\Repeating numbers.txt")
	scan := bufio.NewScanner(fin)
	for scan.Scan() {
		n, _ := strconv.Atoi(scan.Text())
		Num[n]++
	}
	fin.Close()
	return Num
}

func main() {
	Num := NumbersRepeat();
	for x, _ := range Num{
		fmt.Println("[",x, "]", ":", "[", Num[x], "]")
		fmt.Printf("\n")
		x++
	}
}
