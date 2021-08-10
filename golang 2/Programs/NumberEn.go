package main
import (
	"fmt"
	"encoding/binary"
	"os"
)

var Z uint32

func main() {
	OrFile, err := os.Open("..\\Files\\file.dat")
	if err != nil {
		fmt.Println(err)
		return
	}
	Trans := make([]byte, 4)
	var Store uint32
	var Min, Max uint32 = 100000000, 0
	var Count1, Count2, Count3, OddPow uint32
	for {
		_, err1 := OrFile.Read(Trans)
		if err1 != nil {
			break
		}
		Store = binary.LittleEndian.Uint32(Trans)
		if Store <= Min {
			Min = Store
		}
		if Store >= Max {
			Max = Store
		}
		if Store % 2 == 0 {
			Count1++
		}
		if Store % 2 != 0 {
			Count2++
		}
		if Store & 3 == 3 {
			Count3++
		}
		if Store % 2 != 0 && issqr(Store) {
			fmt.Println("Вот оно:", Store, "(", Z, ")")
			OddPow++
		}
	}
	fmt.Println("Вот максимальное число:", Max)
	fmt.Println("Вот минимальное число:", Min)
	fmt.Println("Вот количество чётных чисел:", Count1)
	fmt.Println("Вот количество нечётных чисел:", Count2)
	fmt.Println("Вот количество чисел, которые при делении на 2 дают отрицательное число:", Count3)
	fmt.Println("Вот количество чисел с нечётным квадратом:", OddPow)
}

func issqr(x uint32) bool {
	var i uint32
	for i = 0; i * i < x; i++ { }
	Z = i
	return i * i == x
}
