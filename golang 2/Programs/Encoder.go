package main
import (
		"os"
		"io"
		"fmt"
)

func Scripting() bool {
	check := true
	original, err1 := os.Open("..\\Files\\Homework.txt")
	if err1 != nil {
		fmt.Println(err1)
	}
	crypted, _ := os.Create("..\\Files\\encoded\\CryptedH.txt")
	defer original.Close()
	defer crypted.Close()	
	var k, c byte = 17, 0
	r := make([]byte, 1)
	for {
		n, ctrl := original.Read(r)
		if ctrl != nil || n != 1 {
			if ctrl != io.EOF {
				fmt.Println(ctrl)
				check = false
			}
			break
		}
		c = r[0]
		r[0] = k ^ c
		_, ctrl = crypted.Write(r)
		if ctrl != nil {
			fmt.Println(ctrl)
			check = false
			break
		}
		k = c
	}
	return check
}

func main() {
	if Scripting() == true {
		fmt.Println("Файл успешно закодирован")
	} else {
		fmt.Println("Произошла непредвиденная ошибка")
	}
}
