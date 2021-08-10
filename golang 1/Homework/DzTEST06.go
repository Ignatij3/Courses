package main
import (
			"fmt"
			)
func symmetry(sym string) []rune {
	fmt.Scan(&sym)
	rsym:=[]rune(sym)
	var i, I int
	for i==0; i<len(rsym)/2; i++ {
		I==len(rsym)-i-1		
	}
	if i, I != I, i { 
			fmt.Println("Они не симметричны")
	   }
	   return symmetry
}

func main() {
	symmetry=string(sym)
	fmt.Println(sym)
}
/*По идее, они должны проверяться на противоположность, 
 * но 12-я строка какая то странная*/
