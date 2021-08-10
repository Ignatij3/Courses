package main

import (
		"fmt"
		"bufio"
		"os"
		"io"
		"strings"
)

func scanString() string {
	in:=bufio.NewReader(os.Stdin)
	str,err:=in.ReadString('\n')
	if err!=nil && err!=io.EOF {
		fmt.Println("Ошибка ввода:", err)
	}
	return strings.TrimRight(str, "\n\r")
}

func main() {
	var str string
	str = scanString()
	fmt.Println("'" + str + "'")
	RuneSym := []rune(str)
	L := 0
	R := len(RuneSym) - 1
	for L < R && RuneSym[L] == RuneSym[R] {
		fmt.Println("L =", L, " R =", R, " ok")
		L++
		R--
	}
	fmt.Println("EXIT! L =", L, " R =", R)
	if L >= R {
		fmt.Println("Строка симметрична")
	} else {
		fmt.Println("Строка не симметрична")
	}
}
