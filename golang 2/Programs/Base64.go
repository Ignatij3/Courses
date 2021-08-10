package main

import (
	"fmt"
	"os"
)

func toLetter(x byte) byte {
	return byte('A' + x)
}

func Scripting() bool {
	ready := true
	original, err1 := os.Open("..\\Files\\Homework.txt")
	if err1 != nil {
		fmt.Println(err1)
	}
	crypted, _ := os.Create("..\\Files\\encoded\\Base64C.txt")
	defer original.Close()
	defer crypted.Close()
	Byte := make([]byte, 3)
	Byte2 := make([]byte, 4)
	for {
		_, err2 := original.Read(Byte)
		if err2 != nil {
			break
		}
		o0 := Byte[0]
		o1 := Byte[1]
		o2 := Byte[2]
		c0 := o0 >> 2
		c1 := (o0&0x03)<<4 | (o1 >> 4)
		c2 := (o1&0x0f)<<2 | (o2 >> 6)
		c3 := (o2 & 0x3f)
		Byte2[0] = toLetter(c0)
		Byte2[1] = toLetter(c1)
		Byte2[2] = toLetter(c2)
		Byte2[3] = toLetter(c3)
		_, ctrl := crypted.Write(Byte2)
		if ctrl != nil {
			fmt.Println(ctrl)
			ready = false
			break
		}
	}
	return ready
}

func main() {
	if Scripting() == true {
		fmt.Println("Файл успешно закодирован")
	} else {
		fmt.Println("Произошла непредвиденная ошибка")
	}
}
