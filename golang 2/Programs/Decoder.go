package main
import (
		"os"
		"io"
		"fmt"
)

func Unscripting() bool {
	check := true
	original, err1 := os.Open("..\\Files\\Homework.txt")
	if err1 != nil {
		fmt.Println(err1)
	}
	crypted, err2 := os.Open("..\\Files\\encoded\\CryptedH.txt")
	if err2 != nil {
		fmt.Println(err2)
	}
	decrypted, _ := os.Create("..\\Files\\decoded\\DecryptedH.txt")
	defer original.Close()
	defer crypted.Close()
	defer decrypted.Close()
	var k, c byte = 17, 0
	r1, r2 := make ([]byte, 1), make ([]byte, 1)
	for {
		n1, ctrl1 := original.Read(r1)
		n2, ctrl2 := crypted.Read(r2)
		if ctrl1 != nil || n1 != 1 {
			if ctrl1 != io.EOF {
				fmt.Println(ctrl1)
				check = false
			}
			break
		}
		if ctrl2 != nil || n2 != 1 {
			if ctrl2 != io.EOF {
				fmt.Println(ctrl2)
				check = false
			}
			break
		}
		k = r1[0]
		c = r2[0]
		R := make ([]byte, 1)
		R[0] = k ^ c
		_, ctrl := decrypted.Write(R)
		if ctrl != nil {
			fmt.Println(ctrl)
			check = false
			break
		}
		k = R[0]
	}
	return check
}

func main() {
	if Unscripting() == true {
		fmt.Println("Файл успешно декодирован")
	} else {
		fmt.Println("Произошла непредвиденная ошибка")
	}
}
