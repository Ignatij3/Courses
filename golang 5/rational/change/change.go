package change

type Fraction struct {
	A int
	B int
}

func abs(n int) int {
	if n < 0 {return -n}
	return n
}

func (n Fraction) lcm(m Fraction) int {
	return (abs(n.B) * abs(m.B)) / gcd(n.B, m.B)
}

func gcd(a, b int) int {
	for ; b != 0; {a, b = b, a % b}
	return a
}

func (x *Fraction) Compare(y Fraction) int {
	if (*x).A == y.A && (*x).B == y.B {return 0}
	xA := (*x).A / (*x).B
	yA := y.A / y.B
	if xA < yA {return -1} else if xA > yA {return 1}
	return 2
}

func (x *Fraction) Reduce() {
	c := gcd((*x).A, (*x).B)
	(*x).A /= c
	(*x).B /= c
}

func (x *Fraction) Round() int {
	var k int
	if (*x).A > (*x).B {
		for ; (*x).A >= (*x).B; k++ {
			(*x).A -= (*x).B
		}
	}
	
	c := (*x).A / (*x).B
	if float64(c) >= 0.5 {
		return k + 1
	} else if float64(c) < 0.5 {
		return k
	}
	return -1
}

func (x *Fraction) Floor() int {
	var k int
	if (*x).A > (*x).B {
		for ; (*x).A >= (*x).B; k++ {
			(*x).A -= (*x).B
		}
	}
	return k
}

func (x *Fraction) Ceil() int {
	var k int
	if (*x).A > (*x).B {
		for ; (*x).A >= (*x).B; k++ {
			(*x).A -= (*x).B
		}
	}
	return k + 1
}

func (x *Fraction) Frac() Fraction {
	if (*x).A > (*x).B {
		for (*x).A >= (*x).B {
			(*x).A -= (*x).B
		}
	}
	return (*x)
}
