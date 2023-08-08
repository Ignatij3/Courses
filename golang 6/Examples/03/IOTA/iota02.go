package main

import "fmt"

// It’s possible to make operation when using iotas
// Operation's expansion
const (
    n         = iota
    k     
    l int8	  = iota + 10
    m
    x float32 = iota * 0.3
    y
    z 
    u         = iota * 3.1
    v
    w
)
/*
n: 0 — int
k: 1 — int
l: 12 — int8
m: 13 — int8
x: 1.2 — float32
y: 1.5 — float32
z: 1.8 — float32
u: 21.7 — float64
v: 24.8 — float64
w: 27.9 — float64
*/

// Start from one
const  (
	level1 = iota + 1
	level2
	level3
)
/*
level1: 1 — int
level2: 2 — int
level3: 3 — int
*/

// Skip iota's value
const (
	a 		  = iota * 3 - 1
	b
	_
	c uint8   = 8
	_
	d		
	e		  = iota
	
)
/*
a: -1 — int    // iota = 0: 0*3 - 1
b: 2 — int     // iota = 1: 1*3 - 1
               // skip iota, iota = 2 
c: 8 — uint8   // iota = 3
               // skip iota, iota = 4 
d: 8 — uint8   // iota = 5
e: 6 — int     // iota = 6 
*/ 
 	
	
func main() {
   fmt.Printf("n: %v — %T\n", n, n)
   fmt.Printf("k: %v — %T\n", k, k)
   fmt.Printf("l: %v — %T\n", l, l)
   fmt.Printf("m: %v — %T\n", m, m)
   fmt.Printf("x: %v — %T\n", x, x)
   fmt.Printf("y: %v — %T\n", y, y)
   fmt.Printf("z: %v — %T\n", z, z)
   fmt.Printf("u: %v — %T\n", u, u)
   fmt.Printf("v: %v — %T\n", v, v)
   fmt.Printf("w: %v — %T\n", w, w)
   fmt.Printf("level1: %v — %T\n", level1, level1)
   fmt.Printf("level2: %v — %T\n", level2, level2)
   fmt.Printf("level3: %v — %T\n", level3, level3)
   fmt.Printf("a: %v — %T\n", a, a)
   fmt.Printf("b: %v — %T\n", b, b)
   fmt.Printf("c: %v — %T\n", c, c)
   fmt.Printf("d: %v — %T\n", d, d)
   fmt.Printf("e: %v — %T\n", e, e)
}

