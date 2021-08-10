package main

import "fmt"

func subseq (s1, s2 string) bool {
	r1:= []rune(s1)
    i1:= 0
    for _, c2 := range []rune(s2)   {
		if c2 == r1[i1]  {
			i1++
			if i1 >= len(r1)  {
				return true
			}	
		}	
	}	
	return false
}

func main()  {
	var pattern, source string
	fmt.Print("Enter pattern: ")
	fmt.Scan(&pattern)
	fmt.Print("Enter source: ")
	fmt.Scan(&source)
	fmt.Println(subseq(pattern, source))	
}	
