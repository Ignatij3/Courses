package main 

import "testing"

var secret int = 68578719

func guess1(left, right int) {
	c:= (left + right) / 2
    switch {
		case secret < c: 
			guess1(left, c-1)   	
		case secret > c:
			guess1(c+1, right)
		case secret == c:
			return
    }
}

func guess2(left, right int) {
	for {
        c:= (left + right) / 2
        switch {
			case secret < c: 
				right = c - 1   	
			case secret > c:
				left = c + 1
			case secret == c:
				return
        }    
    }
}

func BenchmarkRecursion(b *testing.B) {
	for i:= 0; i < b.N; i++  {
        guess1(0, 100000000)
    }
}

func BenchmarkCyclic(b *testing.B) {
	for i:= 0; i < b.N; i++  {
        guess2(0, 100000000)
    }
}


