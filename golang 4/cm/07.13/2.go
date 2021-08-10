package main

import  "fmt"
	
func main()  {
	var  (
		a float64 = 1.0E-323
		b float64 = 2.0E-323
		d float64 = 8.0E-323
		y float64 = 2.3
		res1, res2 float64
	)
	res1 = (a*y+b)/(a*y+d);
	res2 = (a+b/y)/(a+d/y);
	fmt.Printf("%24.16e\n%24.16e\n%24.16e\n", res1, res2, res1-res2)
}
/*
  4.2857142857142855e-01
  4.4444444444444442e-01
 -1.5873015873015872e-02
*/
