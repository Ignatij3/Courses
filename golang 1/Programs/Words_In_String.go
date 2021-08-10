package main

import (
	"fmt"
	"bufio"	
	"os"
	"io"
)

func WordsScan() string {
	in:=bufio.NewReader(os.Stdin)
	spacebar, err := in.ReadString('\n')
	if err!=nil && err!=io.EOF {
		fmt.Println("Ошибка  ввода:", err)
	}
	return spacebar
}

func WordsCount(spacebar string) int {
	rSpc:=[]rune(spacebar)
	var count int
	if rSpc[0]==' ' {
		fmt.Println("Ошибка! Введите текст начиная не с пробела")
		fmt.Println(rSpc)
	}
		for count=0; count<=len(rSpc); count++ {}
	return count
}

func main() {
	spacebar := WordsScan()
	count := WordsCount(spacebar)
	fmt.Println(count)
}
