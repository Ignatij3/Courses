package main

import  (
	"fmt"
	"os"
	"bufio"	
	"strings"
	"strconv"	
)	

func main()  {
	var  (
		a, b []int
	)
	
	fin, _ := os.Open("cross.dat")
	defer fin.Close()

	scanner := bufio.NewScanner(fin)
	_ = scanner.Scan()   
 	for _, snum := range strings.Fields(scanner.Text()) {
		if c, err := strconv.Atoi(snum); err == nil {
			a = append(a, c)
		}
	}
	fmt.Println(a)
	
	_ = scanner.Scan()   
 	for _, snum := range strings.Fields(scanner.Text()) {
		if c, err := strconv.Atoi(snum); err == nil {
			b = append(b, c)
		}
	}
	fmt.Println(b)
	fmt.Print("Cross: ")
	for ia, ib := 0,0; ia<len(a) && ib < len(b);  {
		if a[ia] < b[ib]  {
			ia++
		}  else 
		if a[ia] > b[ib]  {
			ib++
		}  else
		// if a[ia] == b[ib]
		{	fmt.Print(a[ia], " ")
			ia++
			ib++
		}				
	}	 
	fmt.Println()
}	
