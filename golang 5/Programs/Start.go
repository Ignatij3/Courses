package main

import (
	"fmt"
)

func main() {
	a := 5
	p := &a
	pp := &p
	ppp := &pp
	pppp := &ppp
	ppppp := &pppp
	pppppp := &ppppp
	ppppppp := &pppppp
	pppppppp := &ppppppp
	ppppppppp := &pppppppp
	pppppppppp := &ppppppppp
	ppppppppppp := &pppppppppp
	
	fmt.Printf("------------ - %T\n", p)
	fmt.Printf("------------- - %T\n", pp)
	fmt.Printf("-------------- - %T\n", ppp)
	fmt.Printf("--------------- - %T\n", pppp)
	fmt.Printf("---------------- - %T\n", ppppp)
	fmt.Printf("----------------- - %T\n", pppppp)
	fmt.Printf("------------------ - %T\n", ppppppp)
	fmt.Printf("------------------- - %T\n", pppppppp)
	fmt.Printf("-------------------- - %T\n", ppppppppp)
	fmt.Printf("--------------------- - %T\n", pppppppppp)
	fmt.Printf("---------------------- - %T\n", ppppppppppp)
	fmt.Printf("------------------------ - %T\n", ppppppppppp)
	fmt.Printf("---------------------- - %T\n", ppppppppppp)
	fmt.Printf("--------------------- - %T\n", pppppppppp)
	fmt.Printf("-------------------- - %T\n", ppppppppp)
	fmt.Printf("------------------- - %T\n", pppppppp)
	fmt.Printf("------------------ - %T\n", ppppppp)
	fmt.Printf("----------------- - %T\n", pppppp)
	fmt.Printf("---------------- - %T\n", ppppp)
	fmt.Printf("--------------- - %T\n", pppp)
	fmt.Printf("-------------- - %T\n", ppp)
	fmt.Printf("------------- - %T\n", pp)
	fmt.Printf("------------ - %T\n", p)
	
	
}
