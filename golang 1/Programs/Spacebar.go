package main

import "fmt"

func spcBytes(spc string) string {
	t:=[]rune(spc)
	var t2 []rune
	var z int
		for z=0; z<len(t); z++{
			if t[z]!='.' {
				t2= append(t2, t[z])
			}
		}
	spc=string(t2)
	return string(spc)
}

func main() {
	var spc string
	fmt.Scan(&spc)
	s:=spcBytes(spc)
	fmt.Println(s)
}
