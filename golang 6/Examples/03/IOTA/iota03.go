package main

import "fmt"


const  (
	stub = -10000
	level1 = iota
	level2
	level3
)
/*
level1: 1 — int
level2: 2 — int
level3: 3 — int
*/

const  (
	alpha1, beta1 = iota, iota*3
	alpha2, beta2
	alpha3, beta3
)	
/*
alpha1: 0 — int;  beta1: 0 — int
alpha2: 1 — int;  beta2: 3 — int
alpha3: 2 — int;  beta3: 6 — int
*/
	
func main() {
   fmt.Printf("level1: %v — %T\n", level1, level1)
   fmt.Printf("level2: %v — %T\n", level2, level2)
   fmt.Printf("level3: %v — %T\n", level3, level3)
   fmt.Printf("alpha1: %v — %T;  ", alpha1, alpha1)
   fmt.Printf("beta1: %v — %T\n", beta1, beta1)
   fmt.Printf("alpha2: %v — %T;  ", alpha2, alpha2)
   fmt.Printf("beta2: %v — %T\n", beta2, beta2)
   fmt.Printf("alpha3: %v — %T;  ", alpha3, alpha3)
   fmt.Printf("beta3: %v — %T\n", beta3, beta3)
}

