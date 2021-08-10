package main
import "fmt"
func main() {
	var a, b, rp, rm, rmm, rM, rd, rdd float64
	fmt.Println("отношения двух чисел")
	fmt.Print("А=")
	fmt.Scan(&a)
	fmt.Print("B=")
	fmt.Scan(&b)
	rp = a + b
	rm = a - b
	rmm = b - a
	rM = a * b
	if a == 0 || b == 0 {
		rd=0
		rdd=0
	} else {
		rd = a / b
		rdd = b / a
	}
	fmt.Println(a,"+",b,"=",rp)
	fmt.Println(a,"-",b,"=",rm)
	fmt.Println(b,"-",a,"=",rmm)
	fmt.Println(a,"*",b,"=",rM)
	fmt.Println(a,"/",b,"=",rd)
	fmt.Println(b,"/",a,"=",rdd)
	
	
	
	
	
	
}
/*
 * a - число
 * b - число
 * rp - a+b
 * rm - a-b
 * rmm - b-a
 * rM - a*b
 * rd - a/b
 * rdd - b/a
 */
