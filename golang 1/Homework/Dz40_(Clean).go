 package main
import (
			"fmt"
			"time"
			  )
/*
 * 
 * 
 * 
 * 
 * */	
func main() {
	fmt.Println("Loading...")
	time.Sleep(5*time.Second)
	q:=[...]int{21, 34, 832}
	for v:=0; v<101; v++  {
		fmt.Println("Downloading", "  16bit_f2llw%.exe ","                                    ", v, "%")
		time.Sleep(45*time.Millisecond)
	}
	fmt.Print("Continue? (y(1)/n(0))")
	var w int
	fmt.Scan(&w)
		if w==1 {
			fmt.Println("Before opening this file in geany, please disable your antivirus")
			fmt.Println("If you have troubles, contact developer by e-mail: ()%(N#$Y *#(Y(*@Y%@gmail.com")
		} 
			if w==0 {
				fmt.Println(q)
				fmt.Println("Before opening this file in geany, please disable your antivirus")
				fmt.Println("If you have troubles, contact developer by e-mail: $(&$)@%)YFH)$*$)@gmail.com")
			}
				for ;w!=1 && w!=0; {
					if w!=1 && w!=0 {
						fmt.Println("Error")
						fmt.Print("Continue? (y(1)/n(0))")
						fmt.Scan(&w)
					} 
					if w==0 {
						fmt.Println(q)
						fmt.Println("Before opening this file in geany, please disable your antivirus")
						fmt.Println("If you have troubles, contact developer by e-mail: $(&$)@%)YFH)$*$)@gmail.com")
						break
					} 
					if w==1{
						fmt.Println("Before opening this file in geany, please disable your antivirus")
						fmt.Println("If you have troubles, contact developer by e-mail: ()%(N#$Y *#(Y(*@Y%@gmail.com")
						break
					}
				}
}
